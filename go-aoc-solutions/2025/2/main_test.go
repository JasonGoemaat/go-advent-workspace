package main

import (
	"reflect"
	"testing"
)

func TestGetDigitCount(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{9, 1},
		{10, 2},
		{123, 3},
		{1234, 4},
		{999999999, 9},
		{10000000, 8},
	}

	for _, test := range tests {
		result := getDigitCount(test.input)
		if result != test.expected {
			t.Errorf("getDigitCount(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestPowerOfTen(t *testing.T) {
	tests := []struct {
		exp      int
		expected int
	}{
		{0, 1},
		{1, 10},
		{2, 100},
		{3, 1000},
		{4, 10000},
	}

	for _, test := range tests {
		result := powerOfTen(test.exp)
		if result != test.expected {
			t.Errorf("powerOfTen(%d) = %d; expected %d", test.exp, result, test.expected)
		}
	}
}
func TestPart1Test(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{22, true},
		{123, false},
		{4444, true},
		{10001, false},
	}

	for _, test := range tests {
		result := part1Validate(test.input)
		if result != test.expected {
			t.Errorf("part1Test(%d) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestParseIntRanges(t *testing.T) {
	sampleInput := "10-19,30-39,50-59"
	expectedOutput := []IntRange{{Start: 10, End: 19}, {Start: 30, End: 39}, {Start: 50, End: 59}}

	result := ParseIntRanges(sampleInput)
	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("ParseIntRanges(%q) = %v; expected %v", sampleInput, result, expectedOutput)
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1-999", 11 + 22 + 33 + 44 + 55 + 66 + 77 + 88 + 99}, // 9 pairs
	}

	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("part1(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestPart2Validate(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{22, true},
		{123, false},
		{4444, true},
		{10001, false},
		{59595959, true},
		{58595959, false},
	}

	for _, test := range tests {
		result := part2Validate(test.input)
		if result != test.expected {
			t.Errorf("part2Validate(%d) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
