package main

import (
	"testing"
)

func TestStringLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Kristina", 8},
		{"shrestha", 8},
		{"Data", 4},
		{"", 0}, // Edge case: empty string
	}

	for _, test := range tests {
		result := StringLength(test.input)
		if result != test.expected {
			t.Errorf("StringLength(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
