package main

import "fmt"

func main() {
	var (
		salario uint
		err     error
	)

	fmt.Print("Digite seu salário anual em numero inteiro: ")
	if _, err = fmt.Scanln(&salario); err != nil {
		panic(err)
	}

	percentage := getPercentage(salario)
	fmt.Printf("Seu salário com a dedução de %d%% é de %d \n", percentage, salario-(salario*percentage)/100)
}

func getPercentage(salario uint) uint {
	var percentage uint = 0

	if salario > 50 {
		percentage = 17
	}

	if salario > 150 {
		percentage += 10
	}
	return percentage
}
