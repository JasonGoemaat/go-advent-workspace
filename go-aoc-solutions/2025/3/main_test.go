package main

import (
	"testing"
)

func TestCalculateJoltage(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1234", 34},
		{"1111", 11},
		{"98769", 99},
		{"3725", 75},
	}
	for _, test := range tests {
		result := CalculateJoltage(test.input)
		if result != test.expected {
			t.Errorf("CalculateJoltage(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestCalculateJoltage2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}
	for _, test := range tests {
		result := CalculateJoltage2(test.input)
		if result != test.expected {
			t.Errorf("CalculateJoltage(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
