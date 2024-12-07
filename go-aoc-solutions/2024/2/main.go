package main

import (
	"slices"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/2
	aoc.Local(part1, "part1", "sample.aoc", 2)
	aoc.Local(part1, "part1", "input.aoc", 660)
	aoc.Local(part2, "part2", "sample.aoc", 4)
	aoc.Local(part2, "part2", "input.aoc", 689)
}

func part1(content string) interface{} {
	numbers := aoc.ParseLinesToInts(aoc.ParseLines(content))
	safeCount := 0
	for _, row := range numbers {
		if isSafeIncreasing(row) || isSafeDecreasing(row) {
			safeCount++
		}
	}
	return safeCount
}

func part2(content string) interface{} {
	numbers := aoc.ParseLinesToInts(aoc.ParseLines(content))
	safeCount := 0
	for _, row := range numbers {
		if isSafeIncreasingLenient(row) || isSafeDecreasingLenient(row) {
			safeCount++
		}
	}
	return safeCount
}

func isSafeIncreasing(values []int) bool {
	for i, value := range values {
		if i > 0 && (value <= values[i-1] || value > values[i-1]+3) {
			return false
		}
	}
	return true
}

func isSafeDecreasing(values []int) bool {
	for i, value := range values {
		if i > 0 && (value >= values[i-1] || value < values[i-1]-3) {
			return false
		}
	}
	return true
}

func remove(slice []int, i int) []int {
	a := slices.Clone(slice)
	return slices.Delete(a, i, i+1)
}

func testIncreasing(a, b int) bool {
	if a >= b {
		return false
	}
	if (b - a) > 3 {
		return false
	}
	return true
}

func isSafeIncreasingLenient(values []int) bool {
	for i, value := range values {
		if i > 0 && ((value <= values[i-1]) || (value > values[i-1]+3)) {
			if isSafeIncreasing(remove(values, i-1)) {
				return true
			}
			if isSafeIncreasing(remove(values, i)) {
				return true
			}
			return false
		}
	}
	return true
}

func isSafeDecreasingLenient(values []int) bool {
	for i, value := range values {
		if i > 0 && (value >= values[i-1] || value < values[i-1]-3) {
			if isSafeDecreasing(remove(values, i-1)) {
				return true
			}
			if isSafeDecreasing(remove(values, i)) {
				return true
			}
			return false
		}
	}
	return true
}
