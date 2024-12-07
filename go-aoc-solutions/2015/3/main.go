package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/3
	// aoc.Local(part1, "part1", "sample.aoc", 0)
	aoc.Local(part1, "part1", "input.aoc", 2572)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 0)
}

// However, the elf back at the north pole has had a little too much eggnog,
// and so his directions are a little off, and Santa ends up visiting some
// houses more than once. How many houses receive at least one present?

// ah, I see this for 2 keys:

type key struct {
	x int
	y int
}

func part1(contents string) interface{} {
	visited := map[key]bool{}
	x, y := 0, 0
	for _, ch := range contents {
		visited[key{x, y}] = true
		if ch == '^' {
			y -= 1
		} else if ch == 'v' {
			y += 1
		} else if ch == '<' {
			x -= 1
		} else if ch == '>' {
			x += 1
		}
	}
	return len(visited)
}

func part2(contents string) interface{} {
	visited := map[key]bool{}
	x, y := []int{0, 0}, []int{0, 0}
	for i, ch := range contents {
		j := i & 1
		visited[key{x[j], y[j]}] = true
		if ch == '^' {
			y[j] -= 1
		} else if ch == 'v' {
			y[j] += 1
		} else if ch == '<' {
			x[j] -= 1
		} else if ch == '>' {
			x[j] += 1
		}
	}
	return len(visited)
}
