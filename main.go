package main

import (
	"fmt"
	"github.com/amiralitaherkhany/number-guessing-game/difficulty"
	"github.com/amiralitaherkhany/number-guessing-game/game"
	"github.com/amiralitaherkhany/number-guessing-game/timer"
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

	gameTimer := timer.New()

StartOfGame:

	myGame := game.New(chances)

	gameTimer.StartTimer()

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
			gameTimer.StopTimer()
			ui.ShowWin(myGame.Attempts, gameTimer.GetTime())
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
