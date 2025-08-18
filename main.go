package main

import (
	"log"

	"vizon-ai-backend/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables")
	}

	// Initialize database connection
	err = database.InitDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Ensure database connection is closed when the application exits
	defer func() {
		if err := database.CloseDatabase(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Your application logic here
	log.Println("Application started successfully")
}
