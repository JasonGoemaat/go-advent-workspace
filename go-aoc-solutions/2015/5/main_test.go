package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestMain(t *testing.T) {
	aoc.ExpectJson(t, 1, part1("ugknbfddgicrmopn")) // at least 3 vowels, dupe dd, and no bad
	aoc.ExpectJson(t, 1, part1("aaa"))              // at least 3 vowels and dupe which is fine
	aoc.ExpectJson(t, 0, part1("jchzalrnumimnmhp")) // no double letter
	aoc.ExpectJson(t, 0, part1("haegwjzuvuyypxyu")) // contains bad combo xy
	aoc.ExpectJson(t, 0, part1("dvszwmarrgswjxmb")) // only one vowel
}
