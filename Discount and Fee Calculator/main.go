package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Strategy type defines the function signature for our strategies
type Strategy func(float64) float64

// Calculator struct holds the strategy
type Calculator struct {
	strategy Strategy
}

// NewCalculator creates a new calculator with the given strategy
func NewCalculator(strategy Strategy) *Calculator {
	return &Calculator{strategy: strategy}
}

// Calculate applies the strategy to the amount
func (c *Calculator) Calculate(amount float64) float64 {
	return c.strategy(amount)
}

// Strategy implementations
func PercentageDiscount(percent float64) Strategy {
	return func(amount float64) float64 {
		return amount - (amount * percent / 100)
	}
}

func FixedDiscount(fixedAmount float64) Strategy {
	return func(amount float64) float64 {
		return amount - fixedAmount
	}
}

func AddFee(feeAmount float64) Strategy {
	return func(amount float64) float64 {
		return amount + feeAmount
	}
}

// askQuestion helper function for CLI input
func askQuestion(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	return reader.ReadString('\n')
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get original amount
	amountInput, err := askQuestion("Enter the original amount: ", reader)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	amountInput = strings.TrimSpace(amountInput)
	originalAmount, err := strconv.ParseFloat(amountInput, 64)
	if err != nil || originalAmount < 0 {
		fmt.Println("Please enter a valid positive number for amount.")
		return
	}

	// Get strategy type
	strategyType, err := askQuestion("Choose strategy (percentage, fixed, fee): ", reader)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	strategyType = strings.TrimSpace(strings.ToLower(strategyType))

	var calculator *Calculator
	var valueInput string
	var value float64

	switch strategyType {
	case "percentage":
		valueInput, err = askQuestion("Enter discount percentage (e.g., 10 for 10%): ", reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		valueInput = strings.TrimSpace(valueInput)
		value, err = strconv.ParseFloat(valueInput, 64)
		if err != nil || value < 0 || value > 100 {
			fmt.Println("Enter a valid percentage between 0 and 100.")
			return
		}
		calculator = NewCalculator(PercentageDiscount(value))

	case "fixed":
		valueInput, err = askQuestion("Enter fixed discount amount: ", reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		valueInput = strings.TrimSpace(valueInput)
		value, err = strconv.ParseFloat(valueInput, 64)
		if err != nil || value < 0 {
			fmt.Println("Enter a valid positive fixed amount.")
			return
		}
		calculator = NewCalculator(FixedDiscount(value))

	case "fee":
		valueInput, err = askQuestion("Enter fee amount to add: ", reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		valueInput = strings.TrimSpace(valueInput)
		value, err = strconv.ParseFloat(valueInput, 64)
		if err != nil || value < 0 {
			fmt.Println("Enter a valid positive fee amount.")
			return
		}
		calculator = NewCalculator(AddFee(value))

	default:
		fmt.Println("Invalid strategy type.")
		return
	}

	if calculator != nil {
		finalAmount := calculator.Calculate(originalAmount)
		fmt.Printf("Final amount after applying %s strategy: $%.2f\n", strategyType, finalAmount)
	}
}
