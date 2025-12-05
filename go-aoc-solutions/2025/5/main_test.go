package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestParseIntRange(t *testing.T) {
	r := aoc.ParseIntRange("123-999")
	if r.Start != 123 {
		t.Errorf("ParseIntRange expecting Start to be %d but got %d", 123, r.Start)
	}
	if r.End != 999 {
		t.Errorf("ParseIntRange expecting End to be %d but got %d", 999, r.End)
	}
}

func TestParseIntRanges(t *testing.T) {
	tests := struct {
		input    string
		expected []aoc.IntRange
	}{
		input: `123-999
45-9817
91883818-333828281`,
		expected: []aoc.IntRange{
			{Start: 123, End: 999},
			{Start: 45, End: 9817},
			{Start: 91883818, End: 333828281},
		},
	}
	ranges := aoc.ParseIntRanges(tests.input)
	for i, r := range ranges {
		expected := tests.expected[i]
		if r.Start != expected.Start {
			t.Errorf("ParseIntRange expecting Start to be %d but got %d", expected.Start, r.Start)
		}
		if r.End != expected.End {
			t.Errorf("ParseIntRange expecting End to be %d but got %d", expected.End, r.End)
		}
	}
}
