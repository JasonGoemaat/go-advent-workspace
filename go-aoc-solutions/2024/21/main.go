package main

import (
	part2 "github.com/JasonGoemaat/go-aoc-solutions/2024/21/part2"
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/21
	aoc.Local(part2.Part1, "part1", "sample.aoc", 126384)
	aoc.Local(part2.Part1, "part1", "input.aoc", 138764)
	aoc.Local(part2.Part2, "part2", "sample.aoc", 154115708116294)
	aoc.Local(part2.Part2, "part2", "input.aoc", 169137886514152)
}
