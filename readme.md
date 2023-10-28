# Number Guessing Game in Go

Welcome to the Number Guessing Game implemented in Go! This is a simple command-line game where you attempt to guess a randomly generated secret number. Have fun trying to crack the code!

# Code Structure

Here's a brief overview of the code structure:

-main.go: This is the main program that controls the game loop.

- getRandomNumber(): Generates a random secret number using Go's math/rand package.
- getUserInput(): Collects user input and handles any errors during input using Go's fmt.Scan() function.
- isMatching(secret, guess): Compares the secret number and the player's guess, providing feedback based on the comparison.
  Feel free to explore the code and customize the game to suit your preferences. Enhance the game by adding new features or making improvements as you see fit.

Enjoy the Number Guessing Game in Go! If you have any questions, issues, or suggestions, please don't hesitate to reach out or open an issue on this repository.

## How to Play

1. **Prerequisites**:

   - Make sure you have Go installed on your system. If you don't have Go installed, you can download it from [the official Go website](https://golang.org/dl/).

2. **Clone the Repository**:

   - Open your terminal and clone this repository to your local machine by running the following command:

   ` git clone https://github.com/yourusername/number-guessing-game-go.git`

3. **Navigate to the Project Directory**:

- Navigate into the project directory using the `cd` (change directory) command followed by the path to the cloned folder, for example

4. **Run the Game**:

- `go run main.go`

5. **Guess the number**:

- The game will generate a random secret number between 0 and 10.
- You will be prompted to enter your guess. Simply type in your guess and press Enter.
- If your guess is too low, it will say "Too low!"
- If your guess is too high, it will say "Too large!"
- If your guess is correct, it will celebrate with "You guessed right!".
