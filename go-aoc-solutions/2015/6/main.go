package main

import (
	"regexp"
	"strconv"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

// This one is interesting.   It involves a lot of data, a 1000x1000 grid with
// a million cells.  There are 300 commands that will turn on, turn off, or toggle
// any lights in a rectangle.

// One way would be just to have a 1mb grid and do the commands.   Worst case that
// would be 300mb of data changes, so maybe I'll just do that.

// But I think there must be some better option.   I could just keep track of 'on'
// rectangles.   Then for each of the commands I adjust so there's rectangles of
// one type.   The problem is that each one could divide multiple areas into multiple
// more areas.   I could end up with a ton of rectangles.   As a small example the
// first 150 could create 150 lines across the area.   The next 1 could toggle column
// 2, making it 600 lines - splitting the existing 150 lines in two by turning one
// of them off and creating another 151 lines for the sections between and outside
// those lines that were toggled on.

func main() {
	// https://adventofcode.com/2015/day/X
	// aoc.Local(part1, "part1", "sample.aoc", 0)
	aoc.Local(part1, "part1", "input.aoc", 543903) // 56ms, not bad
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 14687245) // 93ms, not bad
}

var reNumbers = regexp.MustCompile(`\d+`)

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func toggle(value byte) byte {
	return 1 - value
}

func turnOn(value byte) byte {
	return 1
}

func turnOff(value byte) byte {
	return 0
}

func doCommand(area []byte, line string) {
	var cmd = toggle
	if line[:8] == "turn off" {
		cmd = turnOff
	} else if line[:7] == "turn on" {
		cmd = turnOn
	}
	numStrings := reNumbers.FindAllString(line, -1)
	x1, y1 := parseInt(numStrings[0]), parseInt(numStrings[1])
	x2, y2 := parseInt(numStrings[2]), parseInt(numStrings[3])
	index := y1*1000 + x1
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			area[index] = cmd(area[index])
			index++
		}
		index += (999 - (x2 - x1))
	}
}

func part1(contents string) interface{} {
	lines := aoc.ParseLines(contents)
	area := make([]byte, 1000000)
	for _, line := range lines {
		doCommand(area, line)
	}
	total := 0
	for i := range area {
		total += int(area[i])
	}
	return total
}

func toggle2(value int) int {
	return value + 2
}

func turnOn2(value int) int {
	return value + 1
}

func turnOff2(value int) int {
	if value < 1 {
		return 0
	}
	return value - 1
}

func doCommand2(area []int, line string) {
	var cmd = toggle2
	if line[:8] == "turn off" {
		cmd = turnOff2
	} else if line[:7] == "turn on" {
		cmd = turnOn2
	}
	numStrings := reNumbers.FindAllString(line, -1)
	x1, y1 := parseInt(numStrings[0]), parseInt(numStrings[1])
	x2, y2 := parseInt(numStrings[2]), parseInt(numStrings[3])
	index := y1*1000 + x1
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			area[index] = cmd(area[index])
			index++
		}
		index += (999 - (x2 - x1))
	}
}

func part2(contents string) interface{} {
	lines := aoc.ParseLines(contents)
	area := make([]int, 1000000)
	for _, line := range lines {
		doCommand2(area, line)
	}
	total := 0
	for i := range area {
		total += area[i]
	}
	return total
}
