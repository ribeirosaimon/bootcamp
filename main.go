package main

import (
	"fmt"
	"log"
)

func main() {

	log.Printf("Hello World")
	fmt.Println("Hello Word")

	for i := range []string{"1", "2"} {
		fmt.Println(i)
	}
}
