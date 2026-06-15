package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")
	baseURL := os.Getenv("BASE_URL")

	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	if dbName == "" {
		dbName = "comic_db"
	}
	if baseURL == "" {
		baseURL = "https://komiku.id/"
	}

	// Initialize database
	db, err := NewDatabase(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to MongoDB successfully")

	// Initialize scraper
	scraper := NewScraper(baseURL, db)

	// Run incremental scraping
	log.Println("Starting incremental comic scraping...")
	err = runIncrementalScraping(scraper, db)
	if err != nil {
		log.Fatalf("Scraping failed: %v", err)
	}

	log.Println("Scraping completed successfully")
}

func runIncrementalScraping(scraper *Scraper, db *Database) error {
	// Strategy: Scrape newest comics first, then get detailed info for each
	page := 1
	totalComics := 0
	maxPages := 100 // Adjust based on your needs

	for page <= maxPages {
		log.Printf("Scraping page %d...", page)

		comics, err := scraper.GetNewestComics(page)
		if err != nil {
			log.Printf("Error scraping page %d: %v", page, err)
			page++
			continue
		}

		if len(comics) == 0 {
			log.Println("No more comics found, stopping")
			break
		}

		log.Printf("Found %d comics on page %d", len(comics), page)

		// Process each comic
		for i, comic := range comics {
			log.Printf("Processing comic %d/%d: %s", i+1, len(comics), comic.Title)

			// Check if comic already exists in database
			existing, err := db.GetComicByEndpoint(comic.Endpoint)
			if err == nil && existing != nil {
				log.Printf("Comic already exists, skipping: %s", comic.Title)
				continue
			}

			// Get detailed info
			detailedInfo, err := scraper.GetComicInfo(comic.Endpoint)
			if err != nil {
				log.Printf("Error getting comic info for %s: %v", comic.Title, err)
				continue
			}

			// Merge basic info with detailed info
			comic.Author = detailedInfo.Author
			comic.Status = detailedInfo.Status
			comic.Rating = detailedInfo.Rating
			comic.Genre = detailedInfo.Genre
			comic.Thumbnail = detailedInfo.Thumbnail
			if comic.Desc == "" {
				comic.Desc = detailedInfo.Desc
			}
			comic.CreatedAt = time.Now()
			comic.UpdatedAt = time.Now()

			// Get chapter list
			chapters, err := scraper.GetChapterList(comic.Endpoint)
			if err != nil {
				log.Printf("Error getting chapter list for %s: %v", comic.Title, err)
			} else {
				comic.ChapterList = chapters
				log.Printf("Found %d chapters for %s", len(chapters), comic.Title)
			}

			// Save to database
			err = db.SaveComic(comic)
			if err != nil {
				log.Printf("Error saving comic %s: %v", comic.Title, err)
				continue
			}

			totalComics++
			log.Printf("Saved comic %d: %s", totalComics, comic.Title)

			// Add delay to avoid overwhelming the server
			time.Sleep(2 * time.Second)
		}

		page++
		time.Sleep(3 * time.Second)
	}

	log.Printf("Scraping completed. Total comics saved: %d", totalComics)
	return nil
}
