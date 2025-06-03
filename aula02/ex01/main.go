package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		entrada      string
		countEntrada int
		err          error
	)

	fmt.Print("Digite algo: ")
	if countEntrada, err = fmt.Scanln(&entrada); err != nil {
		panic(err)
	}
	fmt.Printf("Você digitou: %s, que tem %d letras \n", entrada, countEntrada)
	mostrarLetras(entrada)
	os.Exit(0)
}

func mostrarLetras(v string) {
	for index := range v {
		log.Println(string(v[index]))
	}
}
