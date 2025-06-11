package main

import (
	"bookstore/config"
	"bookstore/models" // Import models for AutoMigrate
	"bookstore/router"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv" // Optional: For loading .env file
)

func main() {
	// Load environment variables (optional)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, assuming environment variables are set.")
	}
	// Connect to database
	config.ConnectDatabase()

	// Auto-migrate the database schema
	err = config.DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	log.Println("Database auto-migration complete.")


	// Initialize router
	r := router.InitRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}