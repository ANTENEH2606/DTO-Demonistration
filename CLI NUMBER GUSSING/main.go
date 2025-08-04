package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	playGame()
}

func playGame() {
	secretNumber := generateSecretNumber(1, 100)
	attempts := 0

	displayWelcomeMessage(1, 100)

	for {
		guess, err := getPlayerGuess(1, 100)
		if err != nil {
			fmt.Println(err)
			continue
		}

		attempts++

		result := evaluateGuess(guess, secretNumber)
		if result == 0 { // Correct guess
			fmt.Printf("Congratulations! You guessed correctly in %d attempts.\n", attempts)
			if askToPlayAgain() {
				playGame() // Restart game
			}
			return
		}
	}
}

func generateSecretNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func displayWelcomeMessage(min, max int) {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Printf("I've picked a number between %d and %d.\n", min, max)
	fmt.Println("Can you guess it?")
}

func getPlayerGuess(min, max int) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your guess: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input. Please try again")
	}

	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid whole number")
	}

	if guess < min || guess > max {
		return 0, fmt.Errorf("please guess between %d and %d", min, max)
	}

	return guess, nil
}

func evaluateGuess(guess, secretNumber int) int {
	switch {
	case guess < secretNumber:
		fmt.Println("Too low! Try a higher number.")
		return -1
	case guess > secretNumber:
		fmt.Println("Too high! Try a lower number.")
		return 1
	default:
		return 0 // Correct guess
	}
}

func askToPlayAgain() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Play again? (y/n): ")
	playAgain, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(playAgain)) == "y"
}
