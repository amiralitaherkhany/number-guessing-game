package random

import (
	"math/rand"
	"time"
)

func GenerateNumber(min, max int) int {
	//[min,max] are inclusive
	randSource := rand.NewSource(time.Now().UnixNano())
	randomObj := rand.New(randSource)
	randomNumber := randomObj.Intn((max+1)-min) + 1
	return randomNumber
}
