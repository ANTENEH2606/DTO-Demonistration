package main

import (
	"fmt"
	"os"
)

func main() {
	num1, operator, num2 := getInput()
	result, err := calculate(num1, operator, num2)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Result:", result)
}

// getInput handles all user input
func getInput() (float64, string, float64) {
	var operator string
	var num1, num2 float64

	fmt.Println("Enter first number:")
	fmt.Scan(&num1)

	fmt.Println("Enter operator (+, -, *, /):")
	fmt.Scan(&operator)

	fmt.Println("Enter second number:")
	fmt.Scan(&num2)

	return num1, operator, num2
}

// calculate performs the math operation and handles errors
func calculate(num1 float64, operator string, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}
