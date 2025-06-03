package main

import "fmt"

func main() {
	var (
		salario   uint
		idade     uint8
		empregado bool
		err       error
	)

	fmt.Print("Digite sua idade numero inteiro: ")
	if _, err = fmt.Scanln(&idade); err != nil {
		panic(err)
	}
	if idade < 22 {
		fmt.Printf("Você não tem idade para a operação")
		return
	}

	// bloco para GC
	{
		var entrada string
		fmt.Print("Está empregado? digite `s` ou `n`:")

		if _, err = fmt.Scanln(&entrada); err != nil {
			panic(err)
		}
		empregado = entrada == "s"
	}
	// fim do bloco para GC

	if !empregado {
		fmt.Printf("Grato!")
		return
	} else {
		fmt.Printf("Quanto tempo em seu emprego em anos?")
		var tempoEmprego uint8
		if _, err = fmt.Scanln(&tempoEmprego); err != nil {
			panic(err)
		}
		if tempoEmprego == 0 {
			fmt.Printf("Você não tem tempo de trabalho")
			return
		}
	}

	fmt.Print("Digite seu salário anual em numero inteiro: ")
	if _, err = fmt.Scanln(&salario); err != nil {
		panic(err)
	}

	if salario < 100 {
		fmt.Printf("Você pagara juros")
		return
	}
	fmt.Printf("Você não pagara juros")
	return
}
