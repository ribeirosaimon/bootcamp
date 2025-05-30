package main

import "fmt"

func main() {
	var (
		lastName string = "Smith"
		age      int    = 35
	)

	boolean := false

	var (
		salary    string = "45857.90"
		firstName string = "Mary"
	)

	fmt.Printf("%s, %d, %t, %s, %s", lastName, age, boolean, salary, firstName)
}
