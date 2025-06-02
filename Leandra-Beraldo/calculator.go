package main

import (
	"errors"
	"fmt"
)

func sum(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero is not possible.")
	}
	return a / b, nil
}

func main() {
	var operation int
	var a, b float64

	fmt.Println("Calculator")
	fmt.Println("Which operation you would like to perform?",
		"\n1. Sum",
		"\n2. Subtract",
		"\n3. Multiply",
		"\n4. Divide")
	fmt.Print("Enter the operation number: ")
	fmt.Scan(&operation)

	fmt.Print("Enter the first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&b)
	var result float64

	switch operation {
	case 1:
		result = sum(a, b)
		fmt.Println("Result:", result)
	case 2:
		result = subtract(a, b)
		fmt.Println("Result:", result)
	case 3:
		result = multiply(a, b)
		fmt.Println("Result:", result)
	case 4:
		var err error
		result, err = divide(a, b)
		if err != nil {
			fmt.Println("Division by zero is not possible.")
			return
		}
	default:
		fmt.Println("Invalid operation selected.")
		return
	}
}
