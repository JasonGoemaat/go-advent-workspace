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

// this should return two paths
func TestPart2(t *testing.T) {
	contents := `#####
#...#
#S#E#
#...#
#####`
	// aoc.ExpectJson(t, 3004, part2(contents))
	aoc.ExpectJson(t, -1, part2(contents))
	// aoc.ExpectJson(t, 1, 1)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}
