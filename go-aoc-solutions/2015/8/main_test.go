package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestMain(t *testing.T) {
	aoc.ExpectJson(t, 1, 1)
	sample := `"tihfbgcszuej\"c\xfbvoqskkhbgpaddioo"`
	actual := countRemoved(sample)
	expected := 6
	aoc.ExpectJson(t, expected, actual)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}

func TestMain2(t *testing.T) {
	aoc.ExpectJson(t, 1, 1)
	sample := `"tihfbgcszuejc\xfbvoqskkhbgpaddioo"`
	actual := countRemoved(sample)
	expected := 6
	aoc.ExpectJson(t, expected, actual)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}
