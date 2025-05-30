package main

import (
	"fmt"
)

func main() {

	var (
		temp  = float32(11.2)
		umi   = uint(98)
		press = float32(12.4)
	)

	fmt.Printf("%f, %d, %f", temp, umi, press)

}
