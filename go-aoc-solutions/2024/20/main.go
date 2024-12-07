package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/JasonGoemaat/go-aoc/aoc"
	"github.com/JasonGoemaat/go-aoc/aoc/astar"
	"github.com/JasonGoemaat/go-aoc/aoc/tui"
	"github.com/charmbracelet/lipgloss"
)

var visualize = false

func main() {
	// https://adventofcode.com/2024/day/20
	aoc.Local(func(contents string) interface{} { return part1(contents, 8) }, "part1", "sample.aoc", 14) // 14 requiring 8 or more
	aoc.Local(func(contents string) interface{} { return part1(contents, 100) }, "part1", "input.aoc", 1323)
	aoc.Local(func(contents string) interface{} { return part2(contents, 50) }, "part2", "sample.aoc", 285)
	aoc.Local(func(contents string) interface{} { return part2(contents, 100) }, "part2", "input.aoc", 983905)
}

func part1(contents string, savingsRequired int) interface{} {
	tm := NewMyState(contents, savingsRequired)

	if visualize {
		tui.RunProgram(tm)
	} else {
		for !tm.IsDone() {
			tm.Step()
		}
	}

	return tm.GetSolution()
}

func part2(contents string, savingsRequired int) interface{} {
	tm := NewMyState(contents, savingsRequired)
	costRequired := savingsRequired * 10
	current := tm.final
	for current.Parent != nil {
		other := current.Parent.Parent // one step is stoopid, actually 2 and 3 also
		for other != nil && (current.G-other.G) < costRequired {
			other = other.Parent
		}
		for other != nil {
			dx, dy := current.Position.X-other.Position.X, current.Position.Y-other.Position.Y
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			d := dx + dy
			if d <= 20 {
				savings := ((current.G - other.G) / 10) - d // subtract cheat length from savings
				if savings >= savingsRequired {
					tm.savings[savings]++
				}
			}
			other = other.Parent
		}
		current = current.Parent
	}

	type sav struct {
		amount, count int
	}
	total := 0
	savs := make([]sav, 0, len(tm.savings))
	for k, v := range tm.savings {
		savs = append(savs, sav{k, v})
		total += v
	}
	slices.SortFunc(savs, func(a, b sav) int { return cmp.Compare(a.amount, b.amount) })
	return total
}

var (
	sNormal  = lipgloss.NewStyle()                                  // no style
	sOpen    = lipgloss.NewStyle().Background(lipgloss.Color("2"))  // green
	sClosed  = lipgloss.NewStyle().Background(lipgloss.Color("1"))  // red
	sPath    = lipgloss.NewStyle().Background(lipgloss.Color("4"))  // blue
	sCurrent = lipgloss.NewStyle().Background(lipgloss.Color("13")) // purple
	sValue   = lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Bold(true)
	sCheat   = lipgloss.NewStyle().Background(lipgloss.Color("3")) // yellow

)

type MyState struct {
	Title           string
	Contents        string
	astar           *astar.AStar
	step            int
	path            map[astar.AStarPosition]*astar.AStarNode
	final           *astar.AStarNode
	current         *astar.AStarNode
	direction       int
	visWall         *astar.AStarPosition
	visOther        *astar.AStarPosition
	savings         map[int]int // map of savings amount and counter
	visSavings      int
	savingsRequired int
}

// Return string to print for current state, only programmed really for
// part 1
func (state *MyState) Render() string {
	// path to end if we have it, or most recent check if stepping
	as := state.astar
	pathLength := 0

	pathLength--

	// nifty render
	sb := strings.Builder{}
	for y := range state.astar.Area.Height {
		for x := range state.astar.Area.Width {
			pos := astar.AStarPosition{X: x, Y: y}
			style := sNormal
			if state.visWall != nil && (*state.visWall == pos) {
				sb.WriteString(sCheat.Render("1"))
			} else if state.visOther != nil && (*state.visOther == pos) {
				sb.WriteString(sCheat.Render("2"))
			} else if state.current != nil && state.current.Position == pos {
				sb.WriteString(sCheat.Render("O"))
			} else {
				if state.path[pos] != nil {
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
	sb.WriteString(fmt.Sprintf("Path Length: %s, Step: %s/%s\n", rv(pathLength), rv(state.step), done))
	sb.WriteString(fmt.Sprintf("Solution: %v\n", rv(state.GetSolution())))
	return sb.String()
}

func rv(value interface{}) string {
	return sValue.Render(fmt.Sprintf("%v", value))
}

func (state *MyState) IsDone() bool {
	return state.current.Position == state.astar.Start
}

func (state *MyState) Step() {
	if state.IsDone() {
		return
	}
	current := state.current
	var wallPosition, otherPosition astar.AStarPosition
	var dx, dy = 0, 0
	switch state.direction {
	case 0:
		dx, dy = 1, 0
	case 1:
		dx, dy = 0, 1
	case 2:
		dx, dy = -1, 0
	case 3:
		dx, dy = 0, -1
	default:
		panic("INVALID DIRECTION")
	}
	wallPosition.X = current.Position.X + dx
	otherPosition.X = wallPosition.X + dx
	wallPosition.Y = current.Position.Y + dy
	otherPosition.Y = wallPosition.Y + dy
	state.direction = (state.direction + 1) & 3
	if state.direction == 0 {
		state.current = current.Parent
	}
	other := state.path[otherPosition]
	if other == nil || state.astar.Area.Get(wallPosition.Y, wallPosition.X) != '#' {
		state.visWall, state.visOther = nil, nil
		return
	}
	if other.G < current.G {
		// cheaper
		savings := ((current.G - other.G) / 10) - 2
		state.savings[savings]++
		state.visSavings = savings
	}
}

// two possible values, integer path length, or coordinates of last tile that
// blocked our path if there is no solution
func (state *MyState) GetSolution() interface{} {
	// part 1 - how many save at least 100 picoseconds
	count := 0
	for k, v := range state.savings {
		if k >= state.savingsRequired {
			count += v
		}
	}
	return count
}

func NewMyState(contents string, savingsRequired int) *MyState {
	var tm MyState
	area := aoc.ParseArea(contents)
	var start, end astar.AStarPosition
	for i, b := range area.Data {
		if b == 'S' {
			r, c := area.IndexToRowCol(i)
			start.X = c
			start.Y = r
		} else if b == 'E' {
			r, c := area.IndexToRowCol(i)
			end.X = c
			end.Y = r
		}
	}
	as := astar.NewAStar(area, start, end)
	tm = MyState{Title: "Day 20", Contents: contents, astar: as, savingsRequired: savingsRequired}
	tm.final = as.GetShortestPath()
	tm.current = tm.final
	path := map[astar.AStarPosition]*astar.AStarNode{}
	for node := tm.final; node != nil; node = node.Parent {
		path[node.Position] = node
	}
	tm.path = path
	tm.savings = map[int]int{}
	return &tm
}
