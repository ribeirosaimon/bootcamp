package main

import (
	"fmt"
)

type Students struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Date    string `json:"date"`
}

func main() {
	st := Students{
		ID:      1,
		Name:    "Aula",
		Surname: "Aula",
		Date:    "2022-01-01",
	}
	fmt.Printf("%+v\n", st)
}
