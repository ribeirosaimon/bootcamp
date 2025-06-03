package main

import (
	"fmt"
)

type Person struct {
	ID          uint8
	name        string
	DateOfBirth string
}

type Employee struct {
	ID       uint8
	Position string
	Person   Person
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("Employee %s with ID: %d\n", e.Person.name, e.ID)
}

func NewEmployee(id uint8, name, dateOfBirth string) *Employee {
	return &Employee{
		ID:       id,
		Position: name,
		Person:   Person{id, name, dateOfBirth},
	}
}
func main() {
	employee := NewEmployee(1, "Teste", "27/10")
	employee.PrintEmployee()
}
