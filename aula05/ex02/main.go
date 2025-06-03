package main

import "fmt"

func main() {
	product := factory(Medium, 120).Price()
	fmt.Printf("%+v\n", product)
}

type ProductType string

const (
	Small  ProductType = "small"
	Medium ProductType = "medium"
	Large  ProductType = "large"
)

type Product interface {
	Price() uint
}

func factory(productType ProductType, value uint) Product {
	switch productType {
	case Small:
		return &SmallProduct{
			value: value,
		}
	case Medium:
		return &MediumProduct{
			value: value,
		}
	case Large:
		return &LargeProduct{
			value: value,
		}
	}
	return nil
}

type SmallProduct struct {
	value uint
}

func (s *SmallProduct) Price() uint {
	return s.value
}

type MediumProduct struct {
	value uint
}

func (s *MediumProduct) Price() uint {
	return uint(float64(s.value) * 1.6)
}

type LargeProduct struct {
	value uint
}

func (s *LargeProduct) Price() uint {
	return uint(float64(s.value)*1.6) + 2500
}
