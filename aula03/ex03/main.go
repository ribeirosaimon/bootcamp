package main

import (
	"fmt"
)

func main() {

	for _, v := range []struct {
		salario   uint
		categoria Categoria
	}{
		{130, CatA},
		{130, CatB},
		{130, CatC},
	} {
		salario := calcSalario(v.salario, v.categoria)
		fmt.Printf("Categoria %d e Salário %d\n", v.categoria, salario)
	}

}

type Categoria uint

const (
	CatA Categoria = 3000
	CatB Categoria = 1500
	CatC Categoria = 1000
)

func calcSalario(tempTrabalhado uint, cat Categoria) (salario uint) {
	tempEmHoras := tempTrabalhado / 60
	salario = tempEmHoras * uint(cat)
	switch cat {
	case CatA:
		salario += (salario * 50) / 100
	case CatB:
		salario += (salario * 15) / 100
	default:
		return salario
	}

	return
}
