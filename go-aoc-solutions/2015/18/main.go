package main

import (
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/XX
	aoc.LocalArgs(part1, "part1", "sample.aoc", 4, 4)
	aoc.LocalArgs(part1, "part1", "input.aoc", 768, 100)
	aoc.LocalArgs(part2, "part2", "sample.aoc", 17, 5)
	aoc.LocalArgs(part2, "part2", "input.aoc", 0, 100)
}

func show(area *aoc.Area) string {
	sb := strings.Builder{}
	for row := range area.Height {
		for col := range area.Width {
			sb.WriteByte(area.Get(row, col))
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	return sb.String()
}

func part1(contents string, args ...interface{}) interface{} {
	onByte := byte('#')
	steps := args[0].(int)
	area := aoc.ParseArea(contents)
	surrounding := make([]int, len(area.Data))
	for i := 0; i < steps; i++ {
		// rendering := show(area)
		// fmt.Printf("\nStep %d\n%s\n", i, rendering)
		for i := 0; i < len(surrounding); i++ {
			surrounding[i] = 0
		}
		for row := 0; row < area.Height; row++ {
			for col := 0; col < area.Width; col++ {
				index := area.RowColToIndex(row, col)
				for i := range 9 {
					nr := row - 1 + i/3
					nc := col - 1 + (i % 3)
					if i != 4 && area.Is(nr, nc, onByte) {
						surrounding[index]++
					}
				}
			}
		}
		for index, b := range area.Data {
			r, c := area.IndexToRowCol(index)
			c, r = r, c

			// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
			// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
			if b == onByte && surrounding[index] != 2 && surrounding[index] != 3 {
				area.Data[index] = '.'
			}
			if b != onByte && surrounding[index] == 3 {
				area.Data[index] = '#'
			}
		}
	}
	count := 0
	for _, b := range area.Data {
		if b == onByte {
			count++
		}
	}
	return count
}

func stickCorners(area *aoc.Area) {
	lastRowIndex := (area.Height - 1) * area.Width
	lastColIndex := area.Width - 1
	area.Data[0] = '#'
	area.Data[lastColIndex] = '#'
	area.Data[lastRowIndex] = '#'
	area.Data[lastRowIndex+lastColIndex] = '#'
}

// Wow, super easy.   Was expecting something with tons of iterations where
// we'd have to look for repeaters or generators.  Just have to remember to
// stick the 4 corners on.   Before the first iteration and after every
// iteration seems to do the trick.
func part2(contents string, args ...interface{}) interface{} {
	onByte := byte('#')
	steps := args[0].(int)
	area := aoc.ParseArea(contents)
	stickCorners(area)
	surrounding := make([]int, len(area.Data))
	for i := 0; i < steps; i++ {
		for i := 0; i < len(surrounding); i++ {
			surrounding[i] = 0
		}
		for row := 0; row < area.Height; row++ {
			for col := 0; col < area.Width; col++ {
				index := area.RowColToIndex(row, col)
				for i := range 9 {
					nr := row - 1 + i/3
					nc := col - 1 + (i % 3)
					if i != 4 && area.Is(nr, nc, onByte) {
						surrounding[index]++
					}
				}
			}
		}
		for index, b := range area.Data {
			r, c := area.IndexToRowCol(index)
			c, r = r, c

			// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
			// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
			if b == onByte && surrounding[index] != 2 && surrounding[index] != 3 {
				area.Data[index] = '.'
			}
			if b != onByte && surrounding[index] == 3 {
				area.Data[index] = '#'
			}
		}
		stickCorners(area)
	}
	count := 0
	for _, b := range area.Data {
		if b == onByte {
			count++
		}
	}
	return count
}
