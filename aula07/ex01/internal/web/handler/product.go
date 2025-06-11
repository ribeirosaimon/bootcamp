package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/domain/entity"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/product"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/web/response"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/web/response/apperror"
	"net/http"
	"strconv"
	"strings"
)

type productHandler struct {
	productService product.Service
}

type Product interface {
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	SearchProducts(w http.ResponseWriter, r *http.Request)
	SaveProduct(w http.ResponseWriter, r *http.Request)
	ConsumerPrice(w http.ResponseWriter, r *http.Request)
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
		apperror.NewError(err).Build(w)
		return
	}

	pdt, err := p.productService.GetById(id)
	if err != nil {
		apperror.NewError(err).Build(w)
		return
	}
	response.Success(
		response.WithData(pdt),
		response.WithStatus(http.StatusOK),
	).Build(w)
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
			apperror.NewError(err).Build(w)
			return
		}
	}
	products, err := p.productService.Get(quantityOfItems)
	if err != nil {
		apperror.NewError(err).Build(w)
		return
	}

	response.Success(
		response.WithData(products),
		response.WithStatus(http.StatusOK),
	).Build(w)
}

func (p *productHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	priceGtStr := r.URL.Query().Get("priceGt")
	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		apperror.NewError(err).Build(w)
		return
	}

	products := p.productService.GetByPriceGt(priceGt)
	response.Success(
		response.WithData(products),
		response.WithStatus(http.StatusOK),
	).Build(w)
}

func (p *productHandler) SaveProduct(w http.ResponseWriter, r *http.Request) {
	var pdt entity.Product
	if err := json.NewDecoder(r.Body).Decode(&pdt); err != nil {
		apperror.NewError(err).Build(w)
		return
	}
	if err := pdt.IsValid(); err != nil {
		apperror.NewError(err).Build(w)
		return
	}

	if err := p.productService.Save(&pdt); err != nil {
		apperror.NewError(err).Build(w)
		return
	}
	response.Success(
		response.WithData(pdt),
		response.WithStatus(http.StatusCreated),
	).Build(w)
}

func (p *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var pdt entity.Product
	if err := json.NewDecoder(r.Body).Decode(&pdt); err != nil {
		apperror.NewError(err).Build(w)
		return
	}
	if err := pdt.IsValid(); err != nil {
		apperror.NewError(err).Build(w)
		return
	}
	if err := p.productService.Update(&pdt); err != nil {
		apperror.NewError(err).Build(w)
		return
	}
	response.Success(
		response.WithData(pdt),
		response.WithStatus(http.StatusOK),
	).Build(w)
}

func (p *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		apperror.NewError(err).Build(w)
		return
	}

	if err = p.productService.Delete(id); err != nil {
		apperror.NewError(err).Build(w)
		return
	}

	response.Success(
		response.WithStatus(http.StatusNoContent),
	).Build(w)
}

func (p *productHandler) ConsumerPrice(w http.ResponseWriter, r *http.Request) {
	consumerPrice := r.URL.Query().Get("consumer_price")
	if consumerPrice == "" {
		apperror.NewError(fmt.Errorf("consumer_price is required")).Build(w)
		return
	}

	consumerPriceIds := strings.Split(consumerPrice, ",")
	price, err := p.productService.ConsumerPrice(consumerPriceIds)
	if err != nil {
		apperror.NewError(err).Build(w)
	}

	response.Success(
		response.WithData(price),
		response.WithStatus(http.StatusNoContent),
	).Build(w)
}

func (p *productHandler) IsPublished(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err = p.productService.IsPublished(id); err != nil {
		apperror.NewError(err).Build(w)
	}

	response.Success(
		response.WithStatus(http.StatusOK),
	).Build(w)
}
