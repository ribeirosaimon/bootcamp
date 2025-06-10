package product

import (
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/domain/entity"
	"strings"
)

type service struct {
	productRepository Repository
}

type Service interface {
	IsPublished(id int) error
	GetById(id int) (entity.Product, error)
	Get(qtd int) ([]entity.Product, error)
	GetByPriceGt(qtd float64) []entity.Product
	Save(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id int) error
}

func NewService(productRepository Repository) *service {
	return &service{productRepository: productRepository}
}

func (s *service) IsPublished(id int) error {

	product, err := s.GetById(id)
	if err != nil {
		return err
	}
	product.IsPublished = true

	if err = s.productRepository.Update(&product); err != nil {
		return err
	}
	return nil
}

func (s *service) GetById(id int) (entity.Product, error) {
	return s.productRepository.GetById(id)
}

func (s *service) Get(qtd int) ([]entity.Product, error) {
	return s.productRepository.Get(qtd)
}

func (s *service) GetByPriceGt(qtd float64) []entity.Product {
	return s.productRepository.GetByPriceGt(qtd)
}

func (s *service) Save(product *entity.Product) error {
	return s.productRepository.Save(product)
}

func (s *service) Update(product *entity.Product) error {
	return s.productRepository.Update(product)
}

func (s *service) Delete(id int) error {
	return s.productRepository.Delete(id)
}

func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
