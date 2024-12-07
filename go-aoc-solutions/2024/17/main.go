package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/17
	aoc.Local(part1, "part1", "sample.aoc", "4,6,3,5,6,3,5,2,1,0")
	aoc.Local(part1, "part1", "input.aoc", "5,1,4,0,5,1,0,2,6")

	// used for testing my fastLoop() which is go code representing the
	// 'assembly' instructions input
	// aoc.Local(part2FastLoopSample, "fastLoop", "sample.aoc", "4,6,3,5,6,3,5,2,1,0")
	// aoc.Local(part2FastLoopInput, "fastLoop", "input.aoc", "5,1,4,0,5,1,0,2,6")

	// aoc.Local(part2, "part2", "sample.aoc", 117440) // programmed for part 2, won't work
	aoc.Local(part2, "part2", "input.aoc", 202322936867370)
}
