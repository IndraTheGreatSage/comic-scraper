package main

import (
	"log"
	"os"

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
	port := os.Getenv("PORT")

	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	if dbName == "" {
		dbName = "comic_db"
	}
	if port == "" {
		port = "8080"
	}

	// Initialize database
	db, err := NewDatabase(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to MongoDB successfully")

	// Initialize API
	api := NewAPI(db)
	router := api.SetupRoutes()

	// Start server
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
