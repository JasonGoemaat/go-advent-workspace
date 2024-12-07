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

func TestBad1(t *testing.T) {
	// Difference 253 => 0 for group 39
	contents := `Button A: X+36, Y+17
Button B: X+43, Y+77
Prize: X=3253, Y=1933`
	s1 := solve(contents)
	s2 := solve2(contents, false) // should be A 82 times and B 7 times
	aoc.ExpectJson(t, s1, s2)
}

func TestBad2(t *testing.T) {
	// Difference 253 => 0 for group 39
	contents := `Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450`
	s1 := solve(contents)
	s2 := solve2(contents, false) // should be A 38 and B 86
	aoc.ExpectJson(t, s1, s2)
}

func TestBad3(t *testing.T) {
	// Difference 253 => 0 for group 39
	contents := `Button A: X+11, Y+14
Button B: X+80, Y+28
Prize: X=4187, Y=2450`
	s1 := solve(contents)
	s2 := solve2(contents, false) // should be A 97 and B 39
	aoc.ExpectJson(t, s1, s2)
}
