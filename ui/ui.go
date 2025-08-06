package ui

import (
	"errors"
	"fmt"
	"github.com/amiralitaherkhany/number-guessing-game/game"
	"strconv"
	"strings"
)

const (
	welcomeMessage               = "Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.\n\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n\nEnter your choice:"
	selectedDifficultyMessage    = "\nGreat! You have selected the %s difficulty level.\nLet's start the game!\n\n"
	enterGuessMessage            = "Enter your guess:"
	correctGuessMessage          = "Congratulations! You guessed the correct number in %d attempts.\n"
	incorrectGuessGreaterMessage = "Incorrect! The number is greater than %d.\n\n"
	incorrectGuessLesserMessage  = "Incorrect! The number is less than %d.\n\n"
	ranOutOfChancesMessage       = "Game Over: You ran out of chances, The random number was %d!\n"
)

func SelectedDifficultyAndStartGame(level string) {
	fmt.Printf(selectedDifficultyMessage, level)
}

func ShowWelcome() {
	fmt.Print(welcomeMessage)
}

func PromptDifficulty() (int, error) {
	var difficultyLevel string
	_, err := fmt.Scan(&difficultyLevel)
	if err != nil {
		return 0, fmt.Errorf("input scanning error: %v", err)
	}
	selectedDifficultyLevel, err := strconv.Atoi(strings.TrimSpace(difficultyLevel))
	if err != nil {
		return 0, errors.New("invalid input. please enter a number!")
	}
	return selectedDifficultyLevel, nil
}

func PromptGuess() (int, error) {
	var inputGuess string
	fmt.Print(enterGuessMessage)
	_, err := fmt.Scan(&inputGuess)
	if err != nil {
		return 0, fmt.Errorf("input scanning error: %v", err)
	}
	guess, err := strconv.Atoi(strings.TrimSpace(inputGuess))
	if err != nil {
		return 0, errors.New("invalid input. please enter a number!")
	}
	return guess, nil
}

func ShowResult(result game.GuessResult, guess int) {
	switch result {
	case game.Lesser:
		fmt.Printf(incorrectGuessLesserMessage, guess)
	case game.Greater:
		fmt.Printf(incorrectGuessGreaterMessage, guess)
	default:
		panic("unhandled result case")
	}
}

func ShowWin(attempts uint8) {
	fmt.Printf(correctGuessMessage, attempts)
}

func ShowGameOver(correct int) {
	fmt.Printf(ranOutOfChancesMessage, correct)
}
