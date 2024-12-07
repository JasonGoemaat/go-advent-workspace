package main

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

var (
	reColors = regexp.MustCompile("[wubrg]+")
)

func main() {
	// https://adventofcode.com/2015/day/X
	// aoc.Local(part1, "part1", "sample.aoc", 6)
	// aoc.Local(part1, "part1", "input.aoc", 267)
	aoc.Local(part2, "part2", "sample.aoc", 16)
	aoc.Local(part2, "part2", "input.aoc", 0)
}

var allowDupes = true
var totalLength = 0

func matchBytes(source []byte, offset int, test []byte) bool {
	if (offset + len(test)) > len(source) {
		return false
	}
	for i := range len(test) {
		if source[offset+i] != test[i] {
			return false
		}
	}
	return true
}

func can(pattern []byte, start int, towels [][]byte) bool {
	if start >= len(pattern) {
		return true
	}

	for _, towel := range towels {
		if matchBytes(pattern, start, towel) {
			if can(pattern, start+len(towel), towels) {
				return true
			}
		}
	}
	return false
}

func part1(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	towels := reColors.FindAllString(groups[0], -1)
	slices.SortFunc(towels, func(a, b string) int {
		return cmp.Compare(len(b), len(a))
	})
	towelsBytes := make([][]byte, len(towels))
	for i, towel := range towels {
		towelsBytes[i] = []byte(towel)
	}
	patterns := aoc.ParseLines(groups[1])
	goodPatterns := 0

	for _, pattern := range patterns {
		patternBytes := []byte(pattern)
		if can(patternBytes, 0, towelsBytes) {
			goodPatterns++
			fmt.Printf("'%s': YES!\n", pattern)
		} else {
			fmt.Printf("'%s': NO!\n", pattern)
		}
	}
	return goodPatterns
}

func can2(pattern []byte, start int, towels [][]byte) int {
	if start >= len(pattern) {
		return 1
	}

	count := 0
	for _, towel := range towels {
		if matchBytes(pattern, start, towel) {
			count += can2(pattern, start+len(towel), towels)
		}
	}
	return count
}

// BOOST: memoize so if we get to the same spot another way, we already
// know how many options there are
var waysAtIndex map[int]int

func part2(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	towels := reColors.FindAllString(groups[0], -1)
	slices.SortFunc(towels, func(a, b string) int {
		return cmp.Compare(len(b), len(a))
	})

	patterns := aoc.ParseLines(groups[1])
	goodPatterns := 0

	for i, pattern := range patterns {
		towelsBytes := make([][]byte, 0, len(towels))
		for _, towel := range towels {
			if strings.Index(pattern, towel) >= 0 {
				towelsBytes = append(towelsBytes, []byte(towel))
			}
		}

		patternBytes := []byte(pattern)
		myPatterns := can2(patternBytes, 0, towelsBytes)
		fmt.Printf("%03d: %d\n", i, myPatterns)
		goodPatterns += myPatterns
	}
	return goodPatterns
}
