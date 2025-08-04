# Discount and Fee Calculator in Go

This is a simple command-line calculator implemented in Go that applies different pricing strategies—percentage discount, fixed discount, or adding a fee—to an input amount using the Strategy design pattern.

## Features

- Supports three strategies:
  - **Percentage Discount**: applies a percentage off the original amount
  - **Fixed Discount**: subtracts a fixed amount from the original amount
  - **Add Fee**: adds a fixed fee to the original amount
- Uses the Strategy pattern to allow flexible swapping of calculation logic
- Command-line interaction prompting user for input values and strategy type
- Validates user input to enforce correct values (e.g., positive numbers, valid percentages)

## Code Structure

- `Strategy` is a function type that takes an amount (`float64`) and returns the modified amount.
- `Calculator` holds a strategy and applies it with its `Calculate` method.
- Three strategy factory functions:
  - `PercentageDiscount(percent float64)`
  - `FixedDiscount(fixedAmount float64)`
  - `AddFee(feeAmount float64)`
- User input is handled via the console using buffered input and parsing.
- Input validation ensures valid amounts and strategy choices.
- The main loop applies the Strategy pattern to calculate the final amount based on user input.s

## Usage

1. Run the program by writing "go run main.go".
2. Enter the original amount when prompted.
3. Choose the strategy: type `percentage`, `fixed`, or `fee`.
4. Enter the corresponding value depending on the strategy:
   - Discount percentage (for `percentage`)
   - Fixed discount amount (for `fixed`)
   - Fee amount (for `fee`)
5. The program calculates and displays the final amount after applying the chosen strategy.



