package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ribeirosaimon/bootcamp/aula03/ex01/domain"
	"os"
)

type Product interface {
	GetProductById(id int) (*domain.Product, error)
	GetProducts(qtd int) ([]domain.Product, error)
	GetProductsPriceGt(qtd float64) []domain.Product
	SaveProduct(product domain.Product) (int, error)
}
type repository struct {
	data map[int]*domain.Product
}

func NewProduct() *repository {

	file, err := os.ReadFile("aula03/ex01/products.json")
	if err != nil {
		return nil
	}
	buf := bytes.NewBuffer(file)
	products := make([]domain.Product, 0)
	if err = json.NewDecoder(buf).Decode(&products); err != nil {
		panic(err)
	}

	data := make(map[int]*domain.Product)
	for _, product := range products {
		data[product.Id] = &product
	}

	return &repository{
		data: data,
	}
}

func (r *repository) GetProducts(qtd int) ([]domain.Product, error) {

	res := make([]domain.Product, 0, qtd)
	for i := 0; i < qtd; i++ {
		res = append(res, *r.data[i])
	}
	return res, nil
}

func (r *repository) GetProductById(id int) (*domain.Product, error) {
	if res, ok := r.data[id]; ok {
		return res, nil
	}

	return nil, errors.New("product not found")
}

func (r *repository) GetProductsPriceGt(value float64) []domain.Product {
	res := make([]domain.Product, 0)
	for _, product := range r.data {
		if product.Price > value {
			res = append(res, *product)
		}
	}
	return res
}

func (r *repository) SaveProduct(product domain.Product) (int, error) {
	lastIndex := len(r.data) + 1
	product.Id = lastIndex
	r.data[lastIndex] = &product
	return lastIndex, nil
}
