package main

import (
	"slices"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/1
	aoc.Local(part1, "part1", "sample.aoc", 11)
	aoc.Local(part1, "part1", "input.aoc", 2904518)
	aoc.Local(part2, "part2", "sample.aoc", 31)
	aoc.Local(part2, "part2", "input.aoc", 18650129)
}

func part1(contents string) interface{} {
	var ints = aoc.ParseLinesToInts(aoc.ParseLines(contents))
	var a, b = takeElement(ints, 0), takeElement(ints, 1)
	slices.Sort(a)
	slices.Sort(b)
	total := 0
	for i, v := range a {
		total += distance(v, b[i])
	}
	return total
}

func part2(contents string) interface{} {
	var ints = aoc.ParseLinesToInts(aoc.ParseLines(contents))
	var a, b = takeElement(ints, 0), takeElement(ints, 1)
	m := make(map[int]int)
	for _, v := range b {
		m[v]++
	}

	total := 0
	for _, v := range a {
		total += v * m[v]
	}

	return total
}

func takeElement(numbers [][]int, index int) []int {
	var a = make([]int, len(numbers))
	for i, n := range numbers {
		a[i] = n[index]
	}
	return a
}

func distance(i, j int) int {
	if i < j {
		return j - i
	}
	return i - j
}
