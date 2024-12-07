package main

import (
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/3
	// NOTE: part 2 has a different sample
	aoc.Local(part1, "part1", "sample.aoc", 161)
	aoc.Local(part1, "part1", "input.aoc", 174336360)
	aoc.Local(part2, "part2", "sample2.aoc", 48)
	aoc.Local(part2, "part2", "input.aoc", 88802350)
}

func part1(contents string) interface{} {
	return calculate(contents)
}

func part2(contents string) interface{} {
	dos := strings.Split(contents, "do()")
	total := 0
	for _, s := range dos {
		split := strings.Split(s, "don't()")
		total += calculate(split[0])
	}
	return total
}

func calculate(text string) int {
	rx := regexp.MustCompile(`mul\((\d\d??\d??),(\d\d??\d??)\)`)
	results := rx.FindAllStringSubmatch(text, -1)
	total := 0
	for _, match := range results {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		total += a * b
	}
	return total
}
