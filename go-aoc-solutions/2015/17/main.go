package main

import (
	"cmp"
	"slices"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/XX
	aoc.LocalArgs(part1, "part1", "sample.aoc", 4, 25)
	aoc.LocalArgs(part1, "part1", "input.aoc", 1638, 150) // 1317636 is too high
	aoc.LocalArgs(part2, "part2", "sample.aoc", 3, 25)
	aoc.LocalArgs(part2, "part2", "input.aoc", 17, 150)
}

var sizes []int

func countContainers(index int, remaining int) int {
	if index >= len(sizes) {
		return 0
	}
	count := 0
	r := remaining - sizes[index]
	if r > 0 {
		count += countContainers(index+1, r)
	}
	if r == 0 {
		count++ // we finished 150 exactly with THIS container
	}
	// try to not use this container
	return count + countContainers(index+1, remaining)
}

var usedWays = map[int]int{} // counts of ways for each # of containers used

func countContainers2(index, remaining, used int) {
	if index >= len(sizes) {
		return
	}

	// ccount using this container, passing on if used
	r := remaining - sizes[index]
	if r > 0 {
		countContainers2(index+1, r, used+1)
	}

	// count if THIS container finishes it off
	if r == 0 {
		usedWays[used+1]++
	}

	// try to not use this container
	countContainers2(index+1, remaining, used)
}

func part1(contents string, args ...interface{}) interface{} {
	sizes = aoc.ParseInts(contents)
	target := args[0].(int)
	// start by sorting largest first, only thing that makes sense
	slices.SortFunc(sizes, func(a, b int) int { return -cmp.Compare(a, b) })

	return countContainers(0, target)
}

func part2(contents string, args ...interface{}) interface{} {
	usedWays = map[int]int{} // need to reset
	sizes = aoc.ParseInts(contents)
	target := args[0].(int)
	// start by sorting largest first, only thing that makes sense
	slices.SortFunc(sizes, func(a, b int) int { return -cmp.Compare(a, b) })

	countContainers2(0, target, 0)
	minContainers := 0
	for k := range usedWays {
		if minContainers == 0 || k < minContainers {
			minContainers = k
		}
	}
	return usedWays[minContainers]
}
