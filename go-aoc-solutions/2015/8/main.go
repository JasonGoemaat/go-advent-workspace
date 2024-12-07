package main

import (
	"fmt"
	"regexp"
	"strconv"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/8
	aoc.Local(part1, "part1", "sample.aoc", 12) // sample is (2-0)+(5-3)+(10-7)+(6-1) = 2 + 2 + 3 + 5 = 12
	aoc.Local(part1, "part1", "sample.aoc", 28) // sample is (2-0)+(5-3)+(10-7)+(6-1) = 2 + 2 + 3 + 5 = 12
	aoc.Local(part1, "part1", "input.aoc", 1371)
	aoc.Local(part2, "part2", "sample.aoc", 19)
	aoc.Local(part2, "part2", "input.aoc", 2117)
}

var rx = regexp.MustCompile(`\\\\|\\"|\\x[a-f0-9][a-f0-9]`)
var rxWhitespace = regexp.MustCompile(`\s`)
var rxWhitespaceTricky = regexp.MustCompile(`\\x[a-f0-9][a-f0-9]`)

func countWhitespace(line string) int {
	ws := rxWhitespace.FindAllString(line, -1)
	return len(ws)
}

func countWhitespaceTricky(line string) int {
	ws := rxWhitespaceTricky.FindAllString(line, -1)
	total := 0
	for _, x := range ws {
		bv, err := strconv.ParseInt(x[2:], 16, 0)
		if err != nil {
			panic("ERROR PARSING \\xVV")
		}
		b := byte(bv)
		if rxWhitespace.Match([]byte{b}) {
			total += 1
			// panic(fmt.Sprintf("ERROR WHITESPACE IN \\xVV - '%s' : %s", x, line))
		}
	}
	return total
}

func countRemoved(line string) int {
	total := 0
	all := rx.FindAllString(line, -1)
	for _, v := range all {
		total += len(v) - 1
	}
	return total + 2 // for beginning and ending double quotes
}

func part1(contents string) interface{} {
	lines := aoc.ParseLines(contents)
	sum := 0
	for _, line := range lines {
		removed := countRemoved(line)
		// count := len(line) - removed
		if countWhitespace(line) > 0 {
			panic("WHITESPACE!")
		}
		// tricky := countWhitespaceTricky(line)
		// if tricky > 0 {
		// 	// panic("WHITESPACE!")
		// 	removed += tricky
		// }
		// count := len(line) - special
		// value := len(line) - count
		sum += removed
		// fmt.Printf("total: %d, count: %d, removed: %d -  %s\n", len(line), count, removed, line)
	}
	return sum
}

func part2(contents string) interface{} {
	total := 0
	for _, line := range aoc.ParseLines(contents) {
		// fmt.Printf("%d - %s\n", len(line), line)
		newLine := fmt.Sprintf("%q", line) // not including new quotes
		newLen := len(newLine)
		// fmt.Printf("%d - %s\n", newLen, newLine)
		diff := newLen - len(line)
		// fmt.Printf("  Line %d adding %d\n", i, diff)
		total += diff
	}
	return total
}
