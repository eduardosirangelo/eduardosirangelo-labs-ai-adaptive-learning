package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/health", healthHandler)
	r.Get("/ready",  readinessHandler)

	// Business routes
	r.Route("/content", func(r chi.Router) {
		// r.Get("/{studentId}", getContentHandler)
		// r.Post(...), etc.
	})
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return 
	}
}

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("READY"))
	if err != nil {
		return 
	}
}
