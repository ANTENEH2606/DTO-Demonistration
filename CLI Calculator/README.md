# Simple CLI Calculator in Go

This is a basic command-line calculator written in Go that performs simple arithmetic operations: addition, subtraction, multiplication, and division.

## Features

- Accepts two numbers as input
- Supports four operators: `+`, `-`, `*`, `/`
- Handles division by zero error
- Validates operator input

## How It Works

- The program collects user input through `getInput()`.
- It then calls `calculate()` to perform the arithmetic operation based on the operator.
- If there is an error (invalid operator or division by zero), the program prints the error and terminates.
- Otherwise, it prints the calculated result.


## Error Handling

- If you enter an invalid operator, the program prints an error message and exits.
- If you attempt division by zero, the program prints an error message and exits.


## Usage

1. To run the program in terminal write "go run main.go".
2. Enter the first number when prompted.
3. Enter the operator (`+`, `-`, `*`, or `/`).
4. Enter the second number.
5. The program will display the result or an error if input is invalid.



