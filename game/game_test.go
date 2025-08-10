package game

import (
	"fmt"
	"testing"
)

type mockGame struct {
	*Game
}

func newMockGame(targetNumber int, chances uint8) *mockGame {
	g := &mockGame{
		Game: &Game{
			RandomNumber: targetNumber,
			Chances:      chances,
			Attempts:     0,
			IsGameOver:   false,
		},
	}
	return g
}

func TestNewGame(t *testing.T) {
	chances := uint8(5)
	g := New(chances)

	if g.Chances != chances {
		t.Errorf("Expected chances %d, got %d", chances, g.Chances)
	}

	if g.Attempts != 0 {
		t.Errorf("Expected attempts to be 0 at start, got %d", g.Attempts)
	}

	if g.RandomNumber < 1 || g.RandomNumber > 100 {
		t.Errorf("RandomNumber out of range: %d", g.RandomNumber)
	}

	if g.IsGameOver != false {
		t.Errorf("Expected IsGameOver to be false at start, got %v", g.IsGameOver)
	}
}

func TestGuess(t *testing.T) {
	targetNumber := 42
	chances := uint8(3)
	g := newMockGame(targetNumber, chances)
	tests := []struct {
		input    int
		expected GuessResult
	}{
		{input: 42, expected: Correct},
		{input: 30, expected: Greater},
		{input: 60, expected: Lesser},
		{input: -11231, expected: Greater},
		{input: 2139874, expected: Lesser},
		{input: 0, expected: Greater},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("input:%d", test.input), func(t *testing.T) {
			result := g.Guess(test.input)

			if uint8(i+1) >= chances {
				if !g.IsGameOver {
					t.Errorf("expexted IsGameOver to be true")
				}
			} else if g.IsGameOver {
				t.Errorf("expexted IsGameOver to be false")
			}

			if result != test.expected {
				t.Errorf("test failed,input: %v, expected: %v, result: %v", test.input, test.expected, result)
			}
		})
	}

}
