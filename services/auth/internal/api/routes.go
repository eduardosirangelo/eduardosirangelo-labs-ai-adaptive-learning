package api

import (
	"github.com/eduardosirangelo/go-angular-ai-adaptive-learning/services/auth/internal/handler"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	// Health-check endpoint
	r.Get("/health", handler.Health)
	r.Get("/ready",  handler.Ready)

	// Authentication endpoints
	r.Post("/login", handler.Login)
	r.Post("/refresh", handler.Refresh)
	r.Post("/logout", handler.Logout)
}
