package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/domain/entity"
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/product"
	"net/http"
	"strconv"
)

type productHandler struct {
	productService product.Service
}

type Product interface {
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	SearchProducts(w http.ResponseWriter, r *http.Request)
	SaveProduct(w http.ResponseWriter, r *http.Request)
	IsPublished(w http.ResponseWriter, r *http.Request)
}

func NewProduct(productService product.Service) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (p *productHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	pdt, err := p.productService.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pdt)
}

func (p *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
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
	products, err := p.productService.Get(quantityOfItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (p *productHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	priceGtStr := r.URL.Query().Get("priceGt")
	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	products := p.productService.GetByPriceGt(priceGt)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (p *productHandler) SaveProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := product.IsValid(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := p.productService.Save(&product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (p *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := product.IsValid(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := p.productService.Update(&product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (p *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err = p.productService.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (p *productHandler) IsPublished(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	p.productService.IsPublished(id)
}
