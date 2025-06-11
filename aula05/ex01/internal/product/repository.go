package product

import (
	"bytes"
	"encoding/json"
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/internal/apperror"
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/internal/domain/entity"

	"os"
)

type Repository interface {
	GetById(id int) (entity.Product, error)
	Get(qtd int) ([]entity.Product, error)
	GetByPriceGt(qtd float64) []entity.Product
	Save(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id int) error
}

type repository struct {
	data map[int]*entity.Product
}

func NewRepository() *repository {
	file, err := os.ReadFile("aula05/ex01/products.json")
	if err != nil {
		return nil
	}
	buf := bytes.NewBuffer(file)
	products := make([]entity.Product, 0)
	if err = json.NewDecoder(buf).Decode(&products); err != nil {
		panic(err)
	}

	data := make(map[int]*entity.Product)
	for _, product := range products {
		data[product.Id] = &product
	}

	return &repository{
		data: data,
	}
}

func (r *repository) Get(qtd int) ([]entity.Product, error) {
	res := make([]entity.Product, 0, qtd)
	for i := 1; i <= qtd; i++ {
		res = append(res, *r.data[i])
	}
	return res, nil
}

func (r *repository) GetById(id int) (entity.Product, error) {
	if res, ok := r.data[id]; ok {
		return *res, nil
	}

	return entity.Product{}, apperror.NewAppErrorNotFound()
}

func (r *repository) GetByPriceGt(value float64) []entity.Product {
	res := make([]entity.Product, 0)
	for _, product := range r.data {
		if product.Price > value {
			res = append(res, *product)
		}
	}
	return res
}

func (r *repository) Save(product *entity.Product) error {
	lastIndex := len(r.data) + 1
	product.Id = lastIndex
	r.data[lastIndex] = product
	return nil
}

func (r *repository) Update(product *entity.Product) error {
	if _, ok := r.data[product.Id]; !ok {
		return apperror.NewAppErrorNotFound()
	}
	r.data[product.Id] = product
	return nil
}

func (r *repository) Delete(id int) error {
	if _, ok := r.data[id]; ok {
		delete(r.data, id)
	}
	return nil
}
