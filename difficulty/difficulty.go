package difficulty

import (
	"errors"
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
	case Medium:
		numberOfChances = uint8(5)
	case Hard:
		numberOfChances = uint8(3)
	default:
		return 0, errors.New("invalid choice!")
	}
	return numberOfChances, nil
}
