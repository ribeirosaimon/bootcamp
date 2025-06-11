package product

import (
	"errors"
	"github.com/ribeirosaimon/bootcamp/aula06/ex01/internal/domain/dto"
	"github.com/ribeirosaimon/bootcamp/aula06/ex01/internal/domain/entity"
	"strconv"
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
	ConsumerPrice(ids []string) (dto.ConsumerPrice, error)
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

func (s *service) ConsumerPrice(ids []string) (dto.ConsumerPrice, error) {
	newIds := make([]int, 0, len(ids))
	for _, id := range ids {
		atoi, err := strconv.Atoi(id)
		if err != nil {
			return dto.ConsumerPrice{}, err
		}
		newIds = append(newIds, atoi)
	}

	products, err := s.productRepository.GetByIdsAndCount(newIds)
	if err != nil {
		return dto.ConsumerPrice{}, err
	}

	var total float64
	productEntities := make([]entity.Product, 0, len(products))
	for _, product := range products {
		total += float64(product.ConsumerQuantity) * product.Product.Price
		productEntities = append(productEntities, *product.Product)
	}

	switch {
	case len(newIds) > 10:
		total *= 1.21
	case len(newIds) > 10 && len(newIds) > 20:
		total *= 1.17
	case len(newIds) > 30:
		total *= 1.15
	default:
		return dto.ConsumerPrice{}, errors.New("no product found")
	}

	return dto.ConsumerPrice{
		Products:   productEntities,
		TotalPrice: total,
	}, nil
}
