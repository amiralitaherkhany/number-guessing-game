package ngg

import (
	"github.com/amiralitaherkhany/number-guessing-game/random"
)

type GuessResult int

const (
	Correct GuessResult = iota
	Lesser
	Greater
)

type Game struct {
	RandomNumber int
	Chances      uint8
	Attempts     uint8
}

func (g *Game) NewAttempt() (isGameOver bool) {
	g.Attempts++
	if g.Attempts > g.Chances {
		isGameOver = true
	}
	return
}

func New(chances uint8) *Game {
	return &Game{
		Chances:      chances,
		RandomNumber: random.GenerateNumber(1, 100),
		Attempts:     0,
	}
}

func (g *Game) Guess(guess int) (result GuessResult) {
	switch {
	case guess == g.RandomNumber:
		result = Correct
	case guess < g.RandomNumber:
		result = Greater
	case guess > g.RandomNumber:
		result = Lesser
	}
	return
}
