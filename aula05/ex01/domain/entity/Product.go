package entity

import "errors"

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" required:"true"`
	Quantity    int     `json:"quantity"  required:"true"`
	CodeValue   string  `json:"code_value"  required:"true"`
	IsPublished bool    `json:"is_published"  required:"true"`
	Expiration  string  `json:"expiration"  required:"true"`
	Price       float64 `json:"price"  required:"true"`
}

func (e Product) IsValid() error {
	if e.Quantity < 1 {
		return errors.New("quantity must be greater than zero")
	}
	if e.Price < 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
