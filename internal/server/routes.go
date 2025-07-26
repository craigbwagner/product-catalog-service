package server

import (
	"github.com/craigbwagner/product-catalog-service-go/internal/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", product.List)
	})


	return r
}
