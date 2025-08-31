package routes

import (
	"net/http"
	"vizon-ai-backend/internal/handlers/health"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router) {
	r.Get("/health", http.HandlerFunc(health.HandleHealth))
}
