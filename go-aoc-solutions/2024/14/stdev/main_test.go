package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestComboSecond(t *testing.T) {
	// I found width of 101 has vertical pattern at 14
	// Height of 103 has horizontal pattern at 76
	// answer was when they combined at 7286
	c := Combo{101, 14, 103, 76}
	result := c.Second()
	aoc.ExpectJson(t, 7286, result)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}
