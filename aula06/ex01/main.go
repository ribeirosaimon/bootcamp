package main

import "fmt"

type Salary struct {
	value uint
}

type SalaryError struct {
}

func (s SalaryError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}

func main() {
	if err := highSalary(1200); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Must pay tax.")
}

func highSalary(salary uint) error {
	v := Salary{
		value: salary,
	}
	if v.value >= 150 {
		return SalaryError{}
	}
	return nil
}
