package main

import (
	"log"
	"net/http"
	"os"

	"vizon-ai-backend/internal/database"
	"vizon-ai-backend/internal/routes"

	"github.com/go-chi/chi/v5"
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

	r := chi.NewRouter()
	routes.SetupRoutes(r)
	port := os.Getenv("PORT")

	// Listen on port
	http.ListenAndServe(port, r)
}
