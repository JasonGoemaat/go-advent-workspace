package main

import (
	"fmt"
	"regexp"
	"strconv"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/7
	aoc.Local(part1, "part1", "sample.aoc", 3749)
	aoc.Local(part1, "part1", "input.aoc", 1289579105366)
	aoc.Local(part2, "part2", "sample.aoc", 11387)
	aoc.Local(part2, "part2", "input.aoc", 92148721834692)
}

type puzzleLine struct {
	Goal   int64
	Values []int
}

func parseContent(content string) []puzzleLine {
	rxLine := regexp.MustCompile("(\\d+): ([^\n\r]+)")
	rxSpace := regexp.MustCompile("[ ]+")
	// [][]string{{"190: 10 19", "190", "10 19"}, {"3267: 81 40 27", "3267", "81 40 27"}}
	matches := rxLine.FindAllStringSubmatch(content, -1)
	parsed := make([]puzzleLine, len(matches))
	for i, match := range matches {
		goal, _ := strconv.ParseInt(match[1], 10, 64)
		parts := rxSpace.Split(match[2], -1)
		values := make([]int, len(parts))
		for j, s := range parts {
			values[j], _ = strconv.Atoi(s)
		}
		parsed[i] = puzzleLine{goal, values}
	}
	return parsed
}

func (pl puzzleLine) doesWork(sum int64, index int) bool {
	// if index has passed the end, return true only if matched exactly
	if index >= len(pl.Values) {
		return sum == pl.Goal
	}

	// if sum is too large, return false, no way to go smaller
	// WAIT:   Now that I think about it, there could be a `0`
	// in the middle somewhere while the rest of the puzzle would
	// add and multiple to get a valid result.   Premature
	// optimization is the cause of me messing up first a lot of
	// times with these :)
	// if sum > pl.Goal {
	// 	return false
	// }

	// only add first number, but try multiply with others
	if index > 0 {
		multiplied := sum * int64(pl.Values[index])
		if pl.doesWork(multiplied, index+1) {
			return true
		}
	}

	// try to add and call recursively
	added := sum + int64(pl.Values[index])
	return pl.doesWork(added, index+1)
}

func part1(content string) interface{} {
	total := int64(0)
	parsed := parseContent(content)
	for _, line := range parsed {
		if line.doesWork(0, 0) {
			total += line.Goal
		}
	}
	return total
}

func (pl puzzleLine) doesWork2(sum int64, index int) bool {
	if index >= len(pl.Values) {
		return sum == pl.Goal
	}

	// only add first number, but try multiply and concatenation with others
	if index > 0 {
		multiplied := sum * int64(pl.Values[index])
		if pl.doesWork2(multiplied, index+1) {
			return true
		}

		concatenated, _ := strconv.ParseInt(fmt.Sprintf("%d%d", sum, pl.Values[index]), 10, 64)
		if pl.doesWork2(concatenated, index+1) {
			return true
		}
	}

	// try to add and call recursively
	added := sum + int64(pl.Values[index])
	return pl.doesWork2(added, index+1)
}

func part2(content string) interface{} {
	total := int64(0)
	parsed := parseContent(content)
	for _, line := range parsed {
		if line.doesWork2(0, 0) {
			total += line.Goal
		}
	}
	return total
}
