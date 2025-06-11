package dto

import "github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/domain/entity"

type ConsumerPrice struct {
	Products   []entity.Product `json:"products"`
	TotalPrice float64          `json:"total_price"`
}

type ConsumerPriceProductsWithCount struct {
	ConsumerQuantity int
	Product          *entity.Product
}
