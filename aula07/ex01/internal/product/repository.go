package product

import (
	"bytes"
	"encoding/json"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/domain/dto"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/domain/entity"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/web/response/apperror"
	"os"
)

type Repository interface {
	GetById(id int) (entity.Product, error)
	Get(qtd int) ([]entity.Product, error)
	GetByPriceGt(qtd float64) []entity.Product
	Save(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id int) error
	GetByIds(ids []int) ([]entity.Product, error)
	GetByIdsAndCount(ids []int) (map[int]*dto.ConsumerPriceProductsWithCount, error)
}

type repository struct {
	data      map[int]*entity.Product
	path      string
	lastIndex int
}

type repositoryOption func(*repository)

func WithPath(path string) repositoryOption {
	return func(r *repository) {
		r.path = path
	}
}

func NewRepository(opt ...repositoryOption) *repository {
	repo := &repository{
		data: make(map[int]*entity.Product),
		path: "aula07/ex01/products.json",
	}
	for _, v := range opt {
		v(repo)
	}
	file, err := os.ReadFile(repo.path)
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
		data:      data,
		lastIndex: len(products),
	}
}

func (r *repository) Get(qtd int) ([]entity.Product, error) {
	index := r.lastIndex
	if r.lastIndex <= index {
		index = r.lastIndex
	}
	res := make([]entity.Product, 0, index)
	for i := 1; i <= index; i++ {
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

func (r *repository) GetByIds(ids []int) ([]entity.Product, error) {
	res := make([]entity.Product, 0, len(ids))
	conf := make(map[int]uint8)
	for _, id := range ids {
		if _, ok := conf[id]; ok {
			continue
		}
		if v, ok := r.data[id]; ok {
			res = append(res, *v)
			conf[id] = 0
		}
	}
	return res, nil
}

func (r *repository) GetByIdsAndCount(ids []int) (map[int]*dto.ConsumerPriceProductsWithCount, error) {
	res := make(map[int]*dto.ConsumerPriceProductsWithCount)
	for _, id := range ids {
		if resp, ok := res[id]; ok {
			resp.ConsumerQuantity++
			continue
		}
		product := r.data[id]
		res[id] = &dto.ConsumerPriceProductsWithCount{
			Product:          product,
			ConsumerQuantity: 1,
		}
	}
	return res, nil
}
