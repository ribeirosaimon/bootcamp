package main

import (
	"errors"
	"fmt"
)

type Salary struct {
	value uint
}

type SalaryError struct {
}

func (s SalaryError) Error() string {
	return "Error: salary is less than 10000"
}

func main() {

	tax, err := calcSalaryTax(1020)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The tax is %d\n", tax)
}

func calcSalaryTax(salary uint) (uint, error) {
	v := Salary{
		value: salary,
	}
	if v.value >= 80 {
		return 10, nil
	}
	return 0, errors.New("the worker cannot have worked less than 80 hours per month")
}
