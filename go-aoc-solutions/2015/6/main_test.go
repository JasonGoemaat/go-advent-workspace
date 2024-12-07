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

func TestLine(t *testing.T) {
	aoc.ExpectJson(t, 1000000, part1("turn on 0,0 through 999,999"))
	aoc.ExpectJson(t, 1000000, part1("toggle 0,0 through 999,999"))
	aoc.ExpectJson(t, 4000, part1("turn on 0,0 through 999,999\nturn off 10,10 through 11,11\ntoggle 1,1 through 998,998"))
}
