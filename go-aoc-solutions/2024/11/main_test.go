package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestMain(t *testing.T) {
	aoc.ExpectJson(t, 1, 1)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}

func TestPart1Easy(t *testing.T) {
	aoc.ExpectJson(t, 3, part1Depth("125 17", 1))
	aoc.ExpectJson(t, 4, part1Depth("125 17", 2))
	aoc.ExpectJson(t, 5, part1Depth("125 17", 3))
	aoc.ExpectJson(t, 9, part1Depth("125 17", 4))
	aoc.ExpectJson(t, 13, part1Depth("125 17", 5))
	aoc.ExpectJson(t, 22, part1Depth("125 17", 6))
}
