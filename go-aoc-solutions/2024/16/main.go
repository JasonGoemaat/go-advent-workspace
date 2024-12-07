package main

import (
	"fmt"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/16
	aoc.Local(part1, "part1", "sample.aoc", 7036)
	aoc.Local(part1, "part1", "sample2.aoc", 11048)
	aoc.Local(part1, "part1", "input.aoc", 73432)
	aoc.Local(part2, "part2", "sample.aoc", 45)
	aoc.Local(part2, "part2", "sample2.aoc", 64)
	aoc.Local(part2, "part2", "input.aoc", 0)
}

func part1(contents string) interface{} {
	solver := NewDay16(contents)
	for !solver.queue.IsEmpty() {
		// fmt.Printf("%s\n", solver.Render())
		solver.Step()
	}

	fmt.Printf("%s\n", solver.Render())
	return solver.score
}

func part2(contents string) interface{} {
	// for super-simple test
	// 	contents2 := `#####
	// #...#
	// #.#.#
	// #...#
	// #S#E#
	// #...#
	// #####`
	// 	solver := NewDay16b(contents2)

	solver := NewDay16b(contents)
	for solver.Step() {
	}
	return solver.calculateScore()
}
