# Comic Scraper with Database Integration

Standalone comic scraper based on komikku-api logic, with MongoDB integration for building an independent comic database.

## Features

- **Independent Database**: Stores all comic data in your own MongoDB instance
- **Incremental Scraping**: Gradually pulls 5000+ comics without overwhelming the source
- **Rate Limiting**: Built-in delays to respect the source server
- **Resume Capability**: Skips already-scraped comics
- **Complete Data**: Scrapes comic info, chapters, and metadata

## Prerequisites

- Go 1.21 or higher
- MongoDB (local or remote)
- Internet connection

## Installation

1. Clone or create the project:
```bash
cd comic-scraper
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables in `.env`:
```
MONGODB_URI=mongodb://localhost:27017
DB_NAME=comic_db
BASE_URL=https://komiku.id/
```

## Usage

### Start MongoDB

**Local MongoDB:**
```bash
# Using Docker
docker run -d -p 27017:27017 --name mongodb mongo:latest

# Or install MongoDB locally
# Windows: Download from https://www.mongodb.com/try/download/community
# macOS: brew install mongodb-community
# Linux: sudo apt-get install mongodb
```

**Remote MongoDB (MongoDB Atlas):**
1. Create free account at https://www.mongodb.com/cloud/atlas
2. Create a cluster
3. Get connection string and update `.env`

### Run the Scraper

```bash
go run main.go scraper.go database.go models.go
```

Or build and run:
```bash
go build -o comic-scraper
./comic-scraper
```

## Scraping Strategy

The scraper uses an incremental approach:

1. **Page-by-Page Scraping**: Starts from newest comics page
2. **Duplicate Detection**: Checks database before scraping
3. **Detailed Info**: Fetches full comic details for each entry
4. **Chapter Lists**: Retrieves all available chapters
5. **Rate Limiting**: 2-3 second delays between requests

## Database Schema

### Comics Collection
```json
{
  "_id": "ObjectId",
  "title": "Comic Title",
  "image": "https://...",
  "desc": "Description",
  "type": "Manga/Manhwa/Manhua",
  "endpoint": "/manga/slug/",
  "thumbnail": "https://...",
  "author": "Author Name",
  "status": "Ongoing/Completed",
  "rating": "15 Tahun",
  "genre": ["Action", "Adventure"],
  "chapter_list": [
    {
      "name": "Chapter 1",
      "endpoint": "/ch/chapter-1/"
    }
  ],
  "created_at": "ISODate",
  "updated_at": "ISODate"
}
```

### Chapters Collection
```json
{
  "_id": "ObjectId",
  "endpoint": "/ch/chapter-1/",
  "title": "Chapter Title",
  "images": ["https://...", "https://..."]
}
```

## Customization

### Adjust Scraping Limits

Edit `main.go`:
```go
maxPages := 100 // Change to desired number of pages
```

### Change Rate Limits

Edit `scraper.go`:
```go
c.Limit(&colly.LimitRule{
    Parallelism: 2,  // Increase for faster scraping
    Delay: 2 * time.Second,  // Decrease for faster scraping
})
```

## Next Steps

After building your database:

1. **Create Web API**: Build a REST API to serve comic data
2. **Frontend**: Create a web interface (React, Vue, or plain HTML)
3. **Deployment**: Deploy to VPS (DigitalOcean, AWS, etc.)
4. **Updates**: Schedule regular scraping to keep data fresh

## Deployment Options

### Laptop/Local Server
- Run scraper periodically using cron/Task Scheduler
- Host web API locally or use ngrok for testing

### VPS (Recommended)
- Rent a cheap VPS ($5-10/month)
- Install MongoDB and Go
- Run scraper as background service
- Host web API on same server

### Cloud Functions
- Deploy scraper to AWS Lambda/Google Cloud Functions
- Use cloud MongoDB (MongoDB Atlas)
- Trigger on schedule

## Troubleshooting

**MongoDB Connection Error:**
- Ensure MongoDB is running
- Check connection string in `.env`
- Verify firewall settings

**Scraping Too Slow:**
- Increase `Parallelism` in scraper.go
- Decrease `Delay` (but be respectful to source)
- Run multiple instances with different page ranges

**Missing Data:**
- Check if source website structure changed
- Update CSS selectors in `scraper.go`
- Add error logging for debugging

## Legal Considerations

- Respect robots.txt of source website
- Consider terms of service
- Use scraped data for personal/educational purposes
- Don't redistribute scraped content commercially

## License

This is a educational project. Use responsibly.
