package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 21)
	aoc.Local(part1, "part1", "input.aoc", 1675)
	aoc.Local(part2, "part2", "sample.aoc", 40)
	aoc.Local(part2, "part2", "input.aoc", nil)
}

func DoStep(area *aoc.Area, row int) *aoc.Area {
	// if 'S', make one directly below into '|' beam
	// if '|':
	// 		1. If below is '.', change to beam '|'
	//		2. If below is splitter '^', change '.' to sides to beams '|'
	newArea := area.Clone()
	for col := 0; col < area.Width; col++ {
		ch := area.Get(row, col)
		if ch == 'S' || ch == '|' {
			below := area.Get(row+1, col)
			switch below {
			case '.':
				newArea.Set(row+1, col, '|')
			case '^':
				newArea.Set(row+1, col-1, '|')
				newArea.Set(row+1, col+1, '|')
			}
		}
	}
	return newArea
}

func CountSplits(area *aoc.Area) int {
	count := 0
	for row := 1; row < area.Height; row++ {
		for col := 0; col < area.Width; col++ {
			if area.Is(row, col, '^') && area.Is(row-1, col, '|') {
				count++
			}
		}
	}
	return count
}

func part1(content string) interface{} {
	area := aoc.ParseArea(content)
	// fmt.Printf("Initial Area:\n%v\n", area)
	for row := 0; row < area.Height-1; row++ {
		area = DoStep(area, row)
		// fmt.Printf("After Row %d:\n%v\n", row, area)
	}
	return CountSplits(area)
}

func part2(content string) interface{} {
	area := aoc.ParseArea(content)

	// counts will be how many paths lead to each position
	counts := make([]int, area.Width)
	// first row starts with a '1' for the 'S'
	for i := 0; i < area.Width; i++ {
		if area.Is(0, i, 'S') {
			counts[i] = 1
			break
		}
	}
	for row := 1; row < area.Height; row++ {
		newCounts := make([]int, area.Width)
		for col := 0; col < area.Width; col++ {
			// for empty '.', add value above
			if area.Is(row, col, '.') {
				newCounts[col] += counts[col]
			}
			if area.Is(row, col, '^') {
				newCounts[col-1] += counts[col]
				newCounts[col+1] += counts[col]
			}
		}
		counts = newCounts
		// fmt.Printf("After Row %d:\n%v\n", row, area)
	}
	sum := 0
	for _, count := range counts {
		sum += count
	}
	return sum
}
