package main

import (
	"github.com/amiralitaherkhany/number-guessing-game/difficulty"
	"github.com/amiralitaherkhany/number-guessing-game/ngg"
	"github.com/amiralitaherkhany/number-guessing-game/ui"
	"log"
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

	game := ngg.New(chances)

	for {
		if isRanOutOfChances := game.NewAttempt(); isRanOutOfChances {
			ui.ShowGameOver(game.RandomNumber)
			break
		}
		//
		guess, err := ui.PromptGuess()
		if err != nil {
			log.Fatalln(err)
		}
		//
		guessResult := game.Guess(guess)
		if guessResult == ngg.Correct {
			ui.ShowWin(game.Attempts)
			break
		}
		//
		ui.ShowResult(guessResult, guess)
	}
}
