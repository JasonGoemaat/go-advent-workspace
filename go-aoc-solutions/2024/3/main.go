package main

import (
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/3
	// fmt.Println("Expecting 161, 48 for sample")
	// NOTE: part 2 has a different sample
	// fmt.Println("Expecting 174336360, 88802350 for input")
	aoc.SolveALocal("sample.txt", part1, 161)
	aoc.SolveALocal("input.txt", part1, 174336360)
	aoc.SolveALocal("sample2.txt", part2, 48)
	aoc.SolveALocal("input.txt", part2, 88802350)
	// aoc.SolveLocal(part1, part2)
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
