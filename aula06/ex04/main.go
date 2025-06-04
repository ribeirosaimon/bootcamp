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
	var salaryErr SalaryError
	err := highSalary(180)
	if errors.Is(err, salaryErr) {
		fmt.Printf("salary is greater than 150")
	}
	fmt.Printf("Must pay tax.")
}

func highSalary(salary uint) error {
	v := Salary{
		value: salary,
	}
	if v.value >= 150 {
		return fmt.Errorf("salary is greater than 150")
	}
	return nil
}
