package main

import (
	"strconv"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 357)
	aoc.Local(part1, "part1", "input.aoc", 17535)
	aoc.Local(part2, "part2", "sample.aoc", 3121910778619)
	aoc.Local(part2, "part2", "input.aoc", nil)
}

func CalculateJoltage(line string) int {
	start := 0
	for i := start + 1; i < len(line)-1; i++ {
		if line[i] > line[start] {
			start = i
		}
	}
	end := start + 1
	for i := end + 1; i < len(line); i++ {
		if line[i] > line[end] {
			end = i
		}
	}
	value, err := strconv.Atoi(line[start:start+1] + line[end:end+1])
	if err != nil {
		panic(err)
	}
	return value
}

func part1(contents string) interface{} {
	lines := aoc.ParseLines(contents)
	sum := 0
	for _, line := range lines {
		value := CalculateJoltage(line)
		sum += value
	}
	return sum
}

func CalculateJoltage2(line string) int {
	// pick the highest 12 digits from the start, first one counts and
	// we must leave room to fill out the full 12 digits
	positions := make([]int, 12)
	last_pos := -1
	for i := 0; i < 12; i++ {
		pos := last_pos + 1
		for j := pos + 1; j <= len(line)-12+i; j++ {
			if line[j] > line[pos] {
				pos = j
			}
		}
		positions[i] = pos
		last_pos = pos
	}

	// make large number from the digits
	value := 0
	for i := 0; i < len(positions); i++ {
		value = value*10 + int(line[positions[i]]-'0')
	}
	return value
}

func part2(contents string) interface{} {
	lines := aoc.ParseLines(contents)
	sum := 0
	for _, line := range lines {
		value := CalculateJoltage2(line)
		sum += value
	}
	return sum
}
