package api

import (
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mikolabarkouski/calculator/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (h *Handler) Router() http.Handler {
	r := chi.NewRouter()

	// CORS middleware
	r.Use(corsMiddleware)

	r.Get("/health", h.HealthCheck)

	r.Post("/calculate", h.calculate)

	r.Post("/package", h.addPackage)
	r.Delete("/package/{id}", h.deletePackage)
	r.Get("/packages", h.getPackages)

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}

// corsMiddleware handles CORS headers and preflight requests
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")

		// Handle preflight requests (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
