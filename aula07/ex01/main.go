package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("execução concluída")
	}()
	file, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	// sempre após err, pois file pode ser nil

	defer file.Close()
}
