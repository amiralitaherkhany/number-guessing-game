package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	welcomeMessage               = "Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.\n\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n\nEnter your choice:"
	selectedDifficultyMessage    = "\nGreat! You have selected the %s difficulty level.\nLet's start the game!\n\n"
	enterGuessMessage            = "Enter your guess:"
	correctGuessMessage          = "Congratulations! You guessed the correct number in %d attempts."
	incorrectGuessGreaterMessage = "Incorrect! The number is greater than %d.\n\n"
	incorrectGuessLesserMessage  = "Incorrect! The number is less than %d.\n\n"
	ranOutOfChancesMessage       = "Game Over: You ran out of chances, The random number was %d!\n"
)
const (
	EasyLevel   uint8 = 10
	MediumLevel uint8 = 5
	HardLevel   uint8 = 3
)

func main() {
	fmt.Print(welcomeMessage)
	var difficultyLevel string
	_, err := fmt.Scan(&difficultyLevel)
	if err != nil {
		log.Fatalln("input scanning error:", err)
	}
	selectedDifficultyLevel, err := strconv.Atoi(strings.TrimSpace(difficultyLevel))
	if err != nil {
		log.Fatalln("invalid input. please enter a number!")
	}
	var numberOfChances uint8
	switch selectedDifficultyLevel {
	case 1:
		numberOfChances = EasyLevel
		fmt.Printf(selectedDifficultyMessage, "Easy")
	case 2:
		numberOfChances = MediumLevel
		fmt.Printf(selectedDifficultyMessage, "Medium")
	case 3:
		numberOfChances = HardLevel
		fmt.Printf(selectedDifficultyMessage, "Hard")
	default:
		log.Fatalln("invalid choice!")
	}

	randomNum := generateRandomNumber()

	var inputGuess string
	var attempts uint8
StartOfGuessing:
	for {
		attempts++
		if attempts > numberOfChances {
			fmt.Printf(ranOutOfChancesMessage, randomNum)
			break StartOfGuessing
		}
		fmt.Print(enterGuessMessage)
		_, err := fmt.Scan(&inputGuess)
		if err != nil {
			log.Fatalln("input scanning error:", err)
		}
		guess, err := strconv.Atoi(strings.TrimSpace(inputGuess))
		if err != nil {
			log.Fatalln("invalid input. please enter a number!")
		}

		switch {
		case guess == randomNum:
			fmt.Printf(correctGuessMessage, attempts)
			break StartOfGuessing
		case guess < randomNum:
			fmt.Printf(incorrectGuessGreaterMessage, guess)
		case guess > randomNum:
			fmt.Printf(incorrectGuessLesserMessage, guess)
		}

	}
}

func generateRandomNumber() int {
	randSource := rand.NewSource(time.Now().UnixNano())
	randomObj := rand.New(randSource)
	randomNumber := randomObj.Intn(100) + 1
	return randomNumber
}
