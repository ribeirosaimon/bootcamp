package main

import "fmt"

type Animal string

const (
	Dog       Animal = "Dog"
	Cat       Animal = "Cat"
	Hamster   Animal = "Hamster"
	Tarantula Animal = "Tarantula"
)

func main() {
	animal, err := getAnimal(Dog)
	if err != nil {
		fmt.Println(err)
	}

	v := animal(10)
	fmt.Printf("Precisa de %.2f Kg\n", v)
}

func getAnimal(animal Animal) (func(int) float64, error) {
	switch animal {
	case Dog:
		return getDog, nil
	case Cat:
		return getCat, nil
	case Hamster:
		return getHamster, nil
	case Tarantula:
		return getTarantula, nil
	}
	return nil, fmt.Errorf("no animal found for %s", animal)
}

func getTarantula(qtd int) float64 {
	return float64(qtd*15) / 100
}

func getHamster(qtd int) float64 {
	return float64(qtd*25) / 100
}

func getCat(qtd int) float64 {
	return float64(qtd) * 5
}
func getDog(qtd int) float64 {
	return float64(qtd) * 10
}
