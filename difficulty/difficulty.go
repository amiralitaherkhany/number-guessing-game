package difficulty

import (
	"errors"
	"github.com/amiralitaherkhany/number-guessing-game/ui"
)

type Level = int

const (
	Easy Level = iota + 1
	Medium
	Hard
)

func GetChances(level Level) (uint8, error) {
	var numberOfChances uint8
	switch level {
	case Easy:
		numberOfChances = uint8(10)
		ui.SelectedDifficultyAndStartGame("Easy")
	case Medium:
		numberOfChances = uint8(5)
		ui.SelectedDifficultyAndStartGame("Medium")
	case Hard:
		numberOfChances = uint8(3)
		ui.SelectedDifficultyAndStartGame("Hard")
	default:
		return 0, errors.New("invalid choice!")
	}
	return numberOfChances, nil
}
