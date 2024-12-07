package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestPart1Sample(t *testing.T) {
	aoc.Local(part1, "Part1", "sample.aoc", 14)
}

func TestRangeForLoop(t *testing.T) {
	for i := range 10 {
		t.Log(i)
	}
	t.Fail() // so we can see output without '-v' option
}
