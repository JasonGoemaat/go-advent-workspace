package main

import (
	"fmt"
	"strings"

	"github.com/JasonGoemaat/go-aoc/aoc"
	astar "github.com/JasonGoemaat/go-aoc/aoc/astar"
	tui "github.com/JasonGoemaat/go-aoc/aoc/tui"
	"github.com/charmbracelet/lipgloss"
)

var (
	sNormal  = lipgloss.NewStyle()                                  // no style
	sOpen    = lipgloss.NewStyle().Background(lipgloss.Color("2"))  // green
	sClosed  = lipgloss.NewStyle().Background(lipgloss.Color("1"))  // red
	sPath    = lipgloss.NewStyle().Background(lipgloss.Color("4"))  // blue
	sCurrent = lipgloss.NewStyle().Background(lipgloss.Color("13")) // purple
)

type MyState struct {
	Title    string
	Contents string
	astar    *astar.AStar
}

func (state *MyState) Render() string {
	// path to end if we have it, or most recent check if stepping
	as := state.astar
	pathLength := 0
	path := map[astar.AStarPosition]bool{}
	for node := as.Last; node != nil; node = node.Parent {
		path[node.Position] = true
		pathLength++
	}

	// nifty render
	sb := strings.Builder{}
	for y := range state.astar.Area.Height {
		for x := range state.astar.Area.Width {
			pos := astar.AStarPosition{x, y}
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
	sb.WriteString(fmt.Sprintf("Path Cost: %v, Path Length: %d\n", state.GetSolution(), pathLength))
	return sb.String()
}

func (state *MyState) IsDone() bool {
	return (state.astar.Closed[state.astar.End] != nil) || (len(state.astar.Open) == 0)
}

func (state *MyState) Step() {
	state.astar.StepShortestPath()
}

func (state *MyState) GetSolution() interface{} {
	final := state.astar.Closed[state.astar.End]
	if final == nil {
		return "NO SOLUTION"
	}
	return fmt.Sprintf("%d", final.G)
}

func NewMyState(contents string) *MyState {
	var tm MyState
	area := aoc.ParseArea(contents)
	astar := astar.NewAStar(area, astar.AStarPosition{X: 0, Y: 0}, astar.AStarPosition{X: area.Width - 1, Y: area.Height - 1})
	tm = MyState{Title: "Test A*", Contents: contents, astar: astar}
	return &tm
}

var sample = `...#...
..#..#.
....#..
...#..#
..#..#.
.#..#..
#.#....`

func mainTui() {
	tm := NewMyState(sample)
	tui.RunProgram(tm)
	fmt.Printf("The solution is: %s\n", tm.GetSolution())
}

func main() {
	tm := NewMyState(sample)
	// tui.RunProgram(tm)
	for !tm.IsDone() {
		tm.Step()
	}
	fmt.Printf("The solution is: %s\n", tm.GetSolution())
}
