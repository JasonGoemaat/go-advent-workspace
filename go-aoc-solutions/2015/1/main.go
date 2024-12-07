package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/1
	// aoc.Local(part1, "part1", "sample.aoc", 14)
	aoc.Local(part1, "part1", "input.aoc", 74)
	// aoc.Local(part2, "part2", "sample.aoc", 34)
	aoc.Local(part2, "part2", "input.aoc", 1795)
}

func part1(contents string) interface{} {
	floor := 0
	for _, ch := range contents {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
	}
	return floor
}

func part2(contents string) interface{} {
	floor := 0
	for i, ch := range contents {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}
