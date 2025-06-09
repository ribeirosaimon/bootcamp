package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ribeirosaimon/bootcamp/aula03/ex01/handler"
	"github.com/ribeirosaimon/bootcamp/aula03/ex01/repository"
	"net/http"
)

type server struct {
	productHandler handler.Product
	healthHandler  handler.Health
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

	server := NewServer(handler.NewProduct(repository.NewProduct()), handler.NewHealth())

	api.Get("/ping", server.healthHandler.Ping)

	api.Route("/products", func(r chi.Router) {
		r.Get("/", server.productHandler.GetProductById)
		r.Get("/{id}", server.productHandler.GetProductById)
		r.Get("/search", server.productHandler.SearchProducts)
		r.Post("/", server.productHandler.SaveProduct)
	})

	http.ListenAndServe(":8080", api)
}
