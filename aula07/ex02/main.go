package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("execução concluída")
	}()
	file, err := os.Open("aula07/ex02/customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	fmt.Printf("%v\n", file.Name())
	// sempre após err, pois file pode ser nil
	defer file.Close()
}
