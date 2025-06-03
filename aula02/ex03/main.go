package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("%s tem %v anos \n", "Benjamin", employees["Benjamin"])

	for key, value := range employees {
		if value > 21 {
			fmt.Printf("%s tem %v anos \n", key, value)
		}
	}

	employees["Frederico"] = 25

	delete(employees, "Pedro")
}
