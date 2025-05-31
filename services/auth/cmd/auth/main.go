package main

import (
	"fmt"
	"github.com/eduardosirangelo/go-angular-ai-adaptive-learning/pkg/config"
	httpmw "github.com/eduardosirangelo/go-angular-ai-adaptive-learning/pkg/http"
	"github.com/eduardosirangelo/go-angular-ai-adaptive-learning/services/auth/internal/api"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	// 1) Load the configuration (port, timeouts, etc.)
	cfg, err := config.Load("auth")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 2) Create a new router
	r := chi.NewRouter()
	r.Use(httpmw.Chain(
		cfg.RequestTimeout,
		cfg.CORSOrigins,
		cfg.RateLimit,
		cfg.RateWindow,
		cfg.CompressionLevel,
	))

	api.RegisterRoutes(r)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Auth listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
