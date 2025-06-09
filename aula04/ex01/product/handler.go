package product

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ribeirosaimon/bootcamp/aula03/ex01/domain"
	"net/http"
	"strconv"
)

type handler struct {
	productRepository Repository
}

type Handler interface {
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	SearchProducts(w http.ResponseWriter, r *http.Request)
	SaveProduct(w http.ResponseWriter, r *http.Request)
}

func NewProduct(productRepository Repository) *handler {
	return &handler{
		productRepository: productRepository,
	}
}

func (p *handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	pdt, err := p.productRepository.GetProductById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pdt)
}

func (p *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
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
	products, err := p.productRepository.GetProducts(quantityOfItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (p *handler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	priceGtStr := r.URL.Query().Get("priceGt")
	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	products := p.productRepository.GetProductsPriceGt(priceGt)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (p *handler) SaveProduct(w http.ResponseWriter, r *http.Request) {
	var pdt domain.Product
	if err := json.NewDecoder(r.Body).Decode(&pdt); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	savedId, err := p.productRepository.SaveProduct(pdt)
	pdt.Id = savedId
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pdt)
}
