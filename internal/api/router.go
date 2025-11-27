package api

import (
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mikolabarkouski/calculator/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (h *Handler) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", h.HealthCheck)

	r.Post("/calculate", h.calculate)

	r.Post("/package", h.addPackage)
	r.Delete("/package", h.deletePackage)
	r.Get("/packages", h.getPackages)

	//UI
	r.Get("/app", h.RenderApp)

	//Swagger-UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
