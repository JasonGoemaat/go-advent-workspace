package main

import (
	"fmt"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/1
	fmt.Println("Expecting XX, YY for sample")
	fmt.Println("Expecting XX, YY for input")
	aoc.SolveLocal(part1, part2)
}

func part1(contents string) interface{} {
	var ints = aoc.ParseLinesToInts(aoc.ParseLines(contents))
	return len(ints)
}

func part2(contents string) interface{} {
	var ints = aoc.ParseLinesToInts(aoc.ParseLines(contents))
	return len(ints)
}
