package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minValue(numbers ...float64) float64 {
	minValue := numbers[0]
	for _, number := range numbers {
		minValue = math.Min(minValue, number)
	}

	return minValue
}

func maxValue(numbers ...float64) float64 {
	maxValue := numbers[0]
	for _, number := range numbers {
		maxValue = math.Max(maxValue, number)
	}

	return maxValue
}

func averageValue(numbers ...float64) float64 {
	var sum float64

	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}

func operation(operation string) (func(numbers ...float64) float64, error) {
	switch operation {
	case "minimum":
		return minValue, nil
	case "maximum":
		return maxValue, nil
	case "average":
		return averageValue, nil
	default:
		return nil, errors.New("operation invalid")
	}
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		panic(err)
	}
	averageFunc, err := operation(average)
	if err != nil {
		panic(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		panic(err)
	}

	minVal := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	avgVal := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxVal := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minVal)
	fmt.Println(avgVal)
	fmt.Println(maxVal)
}
