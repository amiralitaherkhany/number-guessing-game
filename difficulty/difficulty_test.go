package difficulty

import (
	"fmt"
	"testing"
)

func TestGetChances_Success(t *testing.T) {
	tests := []struct {
		input    int
		expected uint8
	}{
		{input: 1, expected: 10},
		{input: 2, expected: 5},
		{input: 3, expected: 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input:%d", test.input), func(t *testing.T) {
			result, err := GetChances(test.input)
			if err != nil {
				t.Errorf("GetChances returned unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("test failed,input: %v, expected: %v, result: %v", test.input, test.expected, result)
			}
		})
	}
}

func TestGetChances_Error(t *testing.T) {
	tests := []struct {
		input    int
		expected uint8
	}{
		{input: 0, expected: 0},
		{input: -1, expected: 0},
		{input: 4, expected: 0},
		{input: 5, expected: 0},
		{input: 10, expected: 0},
		{input: -10, expected: 0},
		{input: -1000008098, expected: 0},
		{input: 8759438790, expected: 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input:%d", test.input), func(t *testing.T) {
			result, err := GetChances(test.input)
			if err == nil {
				t.Error("GetChances did not return an error when expected")
			}
			if result != test.expected {
				t.Errorf("test failed,input: %v, expected: %v, result: %v", test.input, test.expected, result)
			}
		})
	}
}
