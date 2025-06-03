package main

import "fmt"

func main() {
	var nums = []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Printf("%f", calcMedia(nums...))
}

func calcMedia(nums ...float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}
