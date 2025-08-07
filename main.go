package main

import (
	"github.com/amiralitaherkhany/number-guessing-game/difficulty"
	"github.com/amiralitaherkhany/number-guessing-game/game"
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

	ui.SelectedDifficultyAndStartGame(chances)

	myGame := game.New(chances)

	for {
		if isRanOutOfChances := myGame.NewAttempt(); isRanOutOfChances {
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
}
