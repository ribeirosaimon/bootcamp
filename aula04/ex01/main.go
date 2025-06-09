package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ribeirosaimon/bootcamp/aula03/ex01/handler"
	"github.com/ribeirosaimon/bootcamp/aula04/ex01/health"
	"github.com/ribeirosaimon/bootcamp/aula04/ex01/product"
	"net/http"
)

type server struct {
	productHandler product.Handler
	healthHandler  health.Handler
}

func NewServer(productHandler handler.Product, healthHandler handler.Health) *server {
	return &server{
		productHandler: productHandler,
		healthHandler:  healthHandler,
	}
}

func main() {
	api := chi.NewRouter()
	api.Use(middleware.Logger)
	api.Use(middleware.Recoverer)

	srv := NewServer(handler.NewProduct(product.NewRepository()), handler.NewHealth())

	api.Get("/ping", srv.healthHandler.Ping)

	api.Route("/products", func(r chi.Router) {
		r.Get("/", srv.productHandler.GetProductById)
		r.Get("/{id}", srv.productHandler.GetProductById)
		r.Get("/search", srv.productHandler.SearchProducts)
		r.Post("/", srv.productHandler.SaveProduct)
	})

	http.ListenAndServe(":8080", api)
}
