package main

import (
	"fmt"
	"os"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
	tea "github.com/charmbracelet/bubbletea"
)

var x model

func main() {
	// showcolors()
	// showcolors2()
	// showascii()
	tea5()
	return

	// https://adventofcode.com/2015/day/X
	// aoc.Local(part1, "part1", "sample.aoc", 2028)
	// aoc.Local(part1, "part1", "sample2.aoc", 10092)
	// aoc.Local(part1, "part1", "input.aoc", 1505963)
	// aoc.Local(part2, "part2", "sample3.aoc", 618) // simple example in puzzle
	// aoc.Local(part2, "part2", "sample2.aoc", 9021)
	// aoc.Local(part2, "part2", "input.aoc", 1543141)

	// testing console window opening debug, set launch.json option:
	//				"console": "externalTerminal"
	// fmt.Println("Hello World")
	// var input int
	// fmt.Scanf("%d", &input)
	// fmt.Printf("Hello %v", input)
	// os.Exit(1)

	// Load some text for our viewport
	content, err := os.ReadFile(aoc.GetSubPath("sample2.aoc"))
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(
		model{content: string(content)},
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}

	// p := tui.NewViewportProgram(tui.NewModel("Day 15"))
}

func move(area *aoc.Area, m byte, pos int) int {
	// find space or wall in direction, if wall just return
	change := 0
	// vis := visualizeArea(area)

	// NOTE: problem input is surrounded by walls, so no need to
	// check bounds since we'll always hit a wall, therefore we
	// can just track index change in grid
	if m == '<' {
		change = -1
	} else if m == '>' {
		change = 1
	} else if m == '^' {
		change = -area.Width
	} else if m == 'v' {
		change = area.Width
	} else {
		panic(fmt.Sprintf("Invalid move: %v", m))
	}

	// move foward until hitting empty space or wall
	newPos := pos + change
	for (area.Data[newPos] != '.') && (area.Data[newPos] != '#') {
		newPos += change
	}

	// if the final position is a wall, we can't move
	if area.Data[newPos] == '#' {
		return pos
	}

	// must have hit '.' or empty space, move everything forward 1
	for newPos != pos {
		aoc.LogF("Moving %c at %d replacing %c at %d\n",
			area.Data[newPos-change], newPos-change,
			area.Data[newPos], newPos)
		outputArea(area)
		area.Data[newPos] = area.Data[newPos-change]
		outputArea(area)
		aoc.LogF("\n")
		newPos -= change
	}

	// we moved over 1
	area.Data[pos] = '.'
	return pos + change
}

func outputArea(area *aoc.Area) {
	if !aoc.LoggingEnabled {
		return
	}
	for row := range area.Height {
		fmt.Println(string(area.Data[row*area.Width : (row+1)*area.Width]))
	}
}

func visualizeArea(area *aoc.Area) []string {
	lines := make([]string, area.Height)
	for row := range area.Height {
		lines[row] = string(area.Data[row*area.Width : (row+1)*area.Width])
	}
	return lines
}

func calculateValue(area *aoc.Area) int {
	total := 0
	for r := range area.Height {
		for c := range area.Width {
			if area.Get(r, c) == 'O' {
				total += (100 * r) + c
			}
		}
	}
	return total
}

func part1(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	s := strings.Join(aoc.ParseLines(groups[1]), "")
	area := aoc.ParseArea(groups[0])
	moves := []byte(s)

	aoc.LogF("\n\n\n----------------------------------------\nInitial state:")
	outputArea(area)

	// find start
	pos := 0
	for i, b := range area.Data {
		if b == '@' {
			pos = i
			break
		}
	}

	for _, b := range moves {
		pos = move(area, b, pos)
		aoc.LogF("\n")
		aoc.LogF("Move %c:\n", b)
		outputArea(area)
	}

	return calculateValue(area)
}
