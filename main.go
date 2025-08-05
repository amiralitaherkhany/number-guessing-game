package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	welcomeMessage = "Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.\n\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n\nEnter your choice:"
)

func main() {
	fmt.Print(welcomeMessage)
	randomNum := generateRandomNumber()
	fmt.Print(randomNum)
}

func generateRandomNumber() int {
	randSource := rand.NewSource(time.Now().UnixNano())
	randomObj := rand.New(randSource)
	randomNumber := randomObj.Intn(100) + 1
	return randomNumber
}
