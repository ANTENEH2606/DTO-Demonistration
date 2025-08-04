# Number Guessing Game in Go

A simple interactive command-line number guessing game written in Go. The program picks a random number between 1 and 100, and the player tries to guess it.

## Features

- Generates a random secret number between 1 and 100
- Prompts the user to enter guesses and provides feedback whether the guess is too low, too high, or correct
- Counts and displays the number of attempts taken to guess the number
- Validates user input to ensure guesses are whole numbers within the allowed range
- Offers an option to play again after a successful guess


## Code Structure

- `main()` starts the game by calling `playGame()`.
- `playGame()` manages the game loop, guesses, and replay logic.
- `generateSecretNumber(min, max int)` creates a random secret number.
- `getPlayerGuess(min, max int)` prompts and validates user input.
- `evaluateGuess(guess, secretNumber int)` provides feedback on the guess.
- `askToPlayAgain()` asks the player if they want to restart the game.

## Error Handling

- The game detects non-integer inputs and prompts the user to try again.
- Guesses outside the allowed range (1 to 100) are rejected with an appropriate message.


## How to Play/Run the code

1. To run the program in terminal write "go run main.go".
2. The program will announce the game and the range of the secret number.
3. Enter your guesses when prompted.
4. The program will tell you if your guess is too low, too high, or correct.
5. When you guess correctly, the program displays the total attempts.
6. Choose whether to play another round or exit.



