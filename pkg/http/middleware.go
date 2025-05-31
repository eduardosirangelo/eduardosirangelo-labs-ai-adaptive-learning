package http

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"net/http"
	"time"
)

// Chain returns a Middlewares type from a slice of middleware handlers.
// Use: r.Use(http.Chain(cfg))
func Chain(
	requestTimeout time.Duration,
	corsOrigins []string,
	rateLimit int,
	rateWindow time.Duration,
	compressionLevel int,
) func(http.Handler) http.Handler {
	mws := chi.Chain(
		chiMiddleware.RequestID,               // X-Request-ID :contentReference[oaicite:5]{index=5}
		chiMiddleware.RealIP,                  // correct RemoteAddr via X-Forwarded-For :contentReference[oaicite:6]{index=6}
		chiMiddleware.Logger,                  // logs start/end and duration :contentReference[oaicite:7]{index=7}
		chiMiddleware.Recoverer,               // captures panics → 500 :contentReference[oaicite:8]{index=8}
		chiMiddleware.Timeout(requestTimeout), // per-request timeout :contentReference[oaicite:9]{index=9}
		Cors(corsOrigins),                     // CORS via go-chi/cors :contentReference[oaicite:10]{index=10}
		Compress(compressionLevel),            // gzip compression :contentReference[oaicite:11]{index=11}
		RateLimiter(rateLimit, rateWindow),    // rate limit via httprate :contentReference[oaicite:12]{index=12}
		SecureHeaders,                         // security headers (HSTS, XSS, etc.) :contentReference[oaicite:13]{index=13}
	)
	return mws.Handler
}

// Cors returns a CORS middleware with allowed origins
func Cors(origins []string) func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})
}

// Compress applies gzip to the body of the response
func Compress(level int) func(http.Handler) http.Handler {
	return chiMiddleware.Compress(level)
}

// RateLimiter limits N requests per time window per IP
func RateLimiter(max int, window time.Duration) func(http.Handler) http.Handler {
	return httprate.Limit(max, window)
}

// SecureHeaders adds HTTP security headers
func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}
