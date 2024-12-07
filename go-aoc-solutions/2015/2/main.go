package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/2
	aoc.Local(part1, "part1", "sample.aoc", 58+43)
	aoc.Local(part1, "part1", "input.aoc", 1598415)
	aoc.Local(part2, "part2", "sample.aoc", 38)
	aoc.Local(part2, "part2", "input.aoc", 3812909)
}

func part1(contents string) interface{} {
	dimensions := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	total := 0
	for _, line := range dimensions {
		// fmt.Printf("line: %v", line)
		a := line[0] * line[1]
		b := line[0] * line[2]
		c := line[1] * line[2]
		smallest := a
		if b < smallest {
			smallest = b
		}
		if c < smallest {
			smallest = c
		}
		thisPackage := (a+b+c)*2 + smallest
		// fmt.Printf("%d: %d, %d, %d - smallest %d, size: %d\n", i, a, b, c, smallest, thisPackage)
		total += thisPackage
	}
	return total
}

func part2(contents string) interface{} {
	dimensions := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	total := 0
	for _, line := range dimensions {
		// fmt.Printf("line: %v", line)
		l := line[0]
		w := line[1]
		h := line[2]

		// get smallest in a, b
		a, b, c := l, w, h
		if b < a {
			a, b = b, a
		}
		if c < b {
			b, c = c, b
		}
		_ = c // prevent warning

		ribbon := l * w * h
		thisPackage := a + a + b + b + ribbon
		// fmt.Printf("%d: %d, %d, %d - smallest %d, size: %d\n", i, a, b, c, smallest, thisPackage)
		total += thisPackage
	}
	return total
}
