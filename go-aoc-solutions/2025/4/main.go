package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 13)
	aoc.Local(part1, "part1", "input.aoc", 1460)
	aoc.Local(part2, "part2", "sample.aoc", 43)
	aoc.Local(part2, "part2", "input.aoc", nil)
}

func CountNeighbors(area *aoc.Area, row, col int) int {
	count := -1
	if !area.Is(row, col, '@') {
		return -1
	}
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if area.Is(row+r, col+c, '@') {
				count++
			}
		}
	}
	return count
}

func part1(contents string) interface{} {
	area := aoc.ParseArea(contents)
	count := 0
	for r := 0; r < area.Height; r++ {
		for c := 0; c < area.Width; c++ {
			neighbors := CountNeighbors(area, r, c)
			if neighbors != -1 && neighbors < 4 {
				count++
			}
		}
	}
	return count
}

func RemoveTP(area *aoc.Area) int {
	count := 0
	for r := 0; r < area.Height; r++ {
		for c := 0; c < area.Width; c++ {
			neighbors := CountNeighbors(area, r, c)
			if neighbors != -1 && neighbors < 4 {
				count++
				area.Set(r, c, '.')
			}
		}
	}
	return count
}

func part2(contents string) interface{} {
	area := aoc.ParseArea(contents)
	total := 0
	for removed := RemoveTP(area); removed > 0; {
		total += removed
		removed = RemoveTP(area)
	}
	return total
}
