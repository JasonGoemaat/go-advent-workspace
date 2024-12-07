package main

import (
	"fmt"
	"strings"

	"github.com/JasonGoemaat/go-aoc/aoc"
	"github.com/JasonGoemaat/go-aoc/aoc/astar"
	"github.com/JasonGoemaat/go-aoc/aoc/tui"
	"github.com/charmbracelet/lipgloss"
)

var visualize = false

func main() {
	// https://adventofcode.com/2024/day/18
	aoc.Local(func(contents string) interface{} { return part1(contents, 7, 12) }, "part1", "sample.aoc", 22)
	aoc.Local(func(contents string) interface{} { return part1(contents, 71, 1024) }, "part1", "input.aoc", 336)
	aoc.Local(func(contents string) interface{} { return part2(contents, 7, 0) }, "part2", "sample.aoc", "6,1")
	aoc.Local(func(contents string) interface{} { return part2(contents, 71, 0) }, "part2", "input.aoc", "24,30")
}

func part1(contents string, size, count int) interface{} {
	tm := NewMyState(contents, size, count)

	if visualize {
		tui.RunProgram(tm)
	} else {
		for !tm.IsDone() {
			tm.Step()
		}
	}

	return tm.GetSolution()
}

func part2(contents string, size, count int) interface{} {
	return part1(contents, size, count)
}

var (
	sNormal  = lipgloss.NewStyle()                                  // no style
	sOpen    = lipgloss.NewStyle().Background(lipgloss.Color("2"))  // green
	sClosed  = lipgloss.NewStyle().Background(lipgloss.Color("1"))  // red
	sPath    = lipgloss.NewStyle().Background(lipgloss.Color("4"))  // blue
	sCurrent = lipgloss.NewStyle().Background(lipgloss.Color("13")) // purple
	sValue   = lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Bold(true)
)

type MyState struct {
	Title      string
	Contents   string
	astar      *astar.AStar
	tiles      [][]int
	step       int
	noSolution bool
	path       map[astar.AStarPosition]*astar.AStarNode
	final      *astar.AStarNode
}

// Render function, takes state and returns stylized (using lipgloss)
// string representing the grid.   Calculates it's own path for
// part 1 since it
func (state *MyState) Render() string {
	// path to end if we have it, or most recent check if stepping
	as := state.astar
	pathLength := 0
	path := map[astar.AStarPosition]bool{}
	for node := as.Last; node != nil; node = node.Parent {
		path[node.Position] = true
		pathLength++
	}
	pathLength--

	// nifty render
	sb := strings.Builder{}
	for y := range state.astar.Area.Height {
		for x := range state.astar.Area.Width {
			pos := astar.AStarPosition{X: x, Y: y}
			style := sNormal
			if as.Last != nil && as.Last.Position == pos {
				sb.WriteString(sCurrent.Render("O"))
			} else {
				if path[pos] {
					style = sPath
				} else if as.Open[pos] != nil {
					style = sOpen
				} else if as.Closed[pos] != nil {
					style = sClosed
				}
				sb.WriteString(style.Render(string(rune(as.Area.Get(y, x)))))
			}
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	done := ""
	if state.IsDone() {
		done = " - " + rv("DONE!")
	}
	sb.WriteString(fmt.Sprintf("Path Length: %s, Step: %s/%s%s\n", rv(pathLength), rv(state.step), rv(len(state.tiles)), done))
	sb.WriteString(fmt.Sprintf("Solution: %v\n", rv(state.GetSolution())))
	return sb.String()
}

func rv(value interface{}) string {
	return sValue.Render(fmt.Sprintf("%v", value))
}

func (state *MyState) IsDone() bool {
	return state.noSolution || (state.step >= len(state.tiles))
}

var recalculateAlways = false

func (state *MyState) Step() {
	if state.noSolution || (state.step >= len(state.tiles)) {
		return
	}
	step := state.tiles[state.step]
	index := state.astar.Area.RowColToIndex(step[1], step[0])
	state.astar.Area.Data[index] = '#'

	// if we have a path and are using fast, we only need to recalculate our path
	// if the tile is in it
	if recalculateAlways || state.final == nil || (state.path[astar.AStarPosition{X: step[0], Y: step[1]}] != nil) {
		state.astar.Reset()
		state.final = state.astar.GetShortestPath()

		if state.final == nil {
			state.noSolution = true
			return
		}
		state.path = map[astar.AStarPosition]*astar.AStarNode{}
		for node := state.final; node != nil; node = node.Parent {
			state.path[node.Position] = node
		}
	}
	state.step++
}

// two possible values, integer path length, or coordinates of last tile that
// blocked our path if there is no solution
func (state *MyState) GetSolution() interface{} {
	if state.noSolution {
		return fmt.Sprintf("%d,%d", state.tiles[state.step][0], state.tiles[state.step][1])
	}
	final := state.astar.Closed[state.astar.End]
	if final == nil {
		return "NO SOLUTION"
	}
	return pathLength(final)
}

func pathLength(finalNode *astar.AStarNode) int {
	if finalNode == nil {
		return 0
	}
	length := 0
	for node := finalNode.Parent; node != nil; node = node.Parent {
		length++
	}
	return length
}

func NewMyState(contents string, size, count int) *MyState {
	var tm MyState
	line := strings.Repeat(".", size)
	sb := strings.Builder{}
	for r := 0; r < size; r++ {
		if r > 0 {
			sb.WriteRune('\n')
		}
		sb.WriteString(line)
	}
	area := aoc.ParseArea(sb.String())
	astar := astar.NewAStar(area, astar.AStarPosition{X: 0, Y: 0}, astar.AStarPosition{X: area.Width - 1, Y: area.Height - 1})
	tiles := aoc.ParseIntsPerLine(contents)
	if count > 0 {
		tiles = tiles[0:count]
	}
	tm = MyState{Title: "Day 18", Contents: contents, astar: astar, tiles: tiles}
	return &tm
}
