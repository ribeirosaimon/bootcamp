package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ribeirosaimon/bootcamp/aula02/ex01/domain"
	"os"
)

type Repository interface {
	GetProductById(id int) (*domain.Product, error)
	GetProducts(qtd int) ([]domain.Product, error)
	GetProductsPriceGt(qtd float64) []domain.Product
}
type repository struct {
	data []domain.Product
}

func NewProduct() *repository {

	file, err := os.ReadFile("aula02/ex01/products.json")
	if err != nil {
		return nil
	}
	buf := bytes.NewBuffer(file)
	products := make([]domain.Product, 0)
	if err = json.NewDecoder(buf).Decode(&products); err != nil {
		panic(err)
	}

	return &repository{
		data: products,
	}
}

func (r *repository) GetProducts(qtd int) ([]domain.Product, error) {
	res := make([]domain.Product, 0, qtd)
	for i := 0; i < qtd; i++ {
		res = append(res, r.data[i])
	}
	return res, nil
}

func (r *repository) GetProductById(id int) (*domain.Product, error) {
	for _, product := range r.data {
		if product.Id == id {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func (r *repository) GetProductsPriceGt(value float64) []domain.Product {
	res := make([]domain.Product, 0)
	for _, product := range r.data {
		if product.Price > value {
			res = append(res, product)
		}
	}
	return res
}
