package main

import (
	"encoding/json"
	"github.com/ribeirosaimon/bootcamp/aula02/ex01/repository"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	api := chi.NewRouter()
	api.Use(middleware.Logger)
	api.Use(middleware.Recoverer)

	repository := repository.NewProduct()

	api.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("pong")
		return
	})

	api.Group(func(group chi.Router) {
		group.Get("/products", func(w http.ResponseWriter, r *http.Request) {
			var (
				quantityOfItems int
				err             error
			)
			qtdItemsParam := r.URL.Query().Get("items")
			if qtdItemsParam == "" {
				quantityOfItems = 10
			} else {
				quantityOfItems, err = strconv.Atoi(qtdItemsParam)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
			products, err := repository.GetProducts(quantityOfItems)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
		})

		group.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
			urlId := chi.URLParam(r, "id")
			id, err := strconv.Atoi(urlId)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			product, err := repository.GetProductById(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
		})
		group.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {
			priceGtStr := r.URL.Query().Get("priceGt")
			priceGt, err := strconv.ParseFloat(priceGtStr, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			products := repository.GetProductsPriceGt(priceGt)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
		})
	})

	http.ListenAndServe(":8080", api)
}
