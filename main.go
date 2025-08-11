package main

import (
	"fmt"
	"github.com/amiralitaherkhany/number-guessing-game/difficulty"
	"github.com/amiralitaherkhany/number-guessing-game/game"
	"github.com/amiralitaherkhany/number-guessing-game/ui"
	"log"
	"os"
)

func main() {
	ui.ShowWelcome()

	difficultyLevel, err := ui.PromptDifficulty()
	if err != nil {
		log.Fatalln(err)
	}

	chances, err := difficulty.GetChances(difficultyLevel)
	if err != nil {
		log.Fatalln(err)
	}

	ui.SelectedDifficultyAndStartGame(chances)

StartOfGame:

	myGame := game.New(chances)

	for {
		if myGame.IsGameOver {
			ui.ShowGameOver(myGame.RandomNumber)
			break
		}
		//
		guess, err := ui.PromptGuess()
		if err != nil {
			log.Fatalln(err)
		}
		//
		guessResult := myGame.Guess(guess)
		if guessResult == game.Correct {
			ui.ShowWin(myGame.Attempts)
			break
		}
		//
		ui.ShowResult(guessResult, guess)
	}
	isWant, err := ui.AskForNewGame()
	if err != nil {
		log.Fatalln(err)
	}
	if isWant {
		fmt.Println()
		goto StartOfGame
	} else {
		fmt.Println()
		os.Exit(0)
	}
}
