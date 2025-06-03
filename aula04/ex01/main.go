package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID           uint
	Price        uint
	Category     uint
	Name         string
	Descriptions string
}
type Products struct {
	product []Product
}

func (p *Products) AddProduct(product Product) {
	p.product = append(p.product, product)
}

func (p *Products) GetAllProducts() []Product {
	return p.product
}

func (p *Products) GetProductByID(id uint) (Product, error) {
	for _, product := range p.product {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("not found")
}

var products = Products{
	product: []Product{
		{
			ID:           1,
			Price:        100,
			Category:     2,
			Name:         "Livro",
			Descriptions: "Um livro interessante",
		},
		{
			ID:           2,
			Price:        50,
			Category:     3,
			Name:         "Caneta",
			Descriptions: "Para escrever",
		},
	},
}

func main() {
	product, err := products.GetProductByID(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(product)
}
