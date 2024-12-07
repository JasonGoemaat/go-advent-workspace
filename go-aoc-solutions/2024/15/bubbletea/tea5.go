package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tea5model struct {
	area           *aoc.Area
	moves          []rune
	previousRender string
	currentRender  string
	step           int
	pos            int
	rendering      bool
	auto           bool
	delay          int
	moved          map[int]bool
	lastMove       rune
	elapsed        time.Duration
}

// NOTE: re-uses some types in tea4.go

// styles

var (
	sWall  = lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Bold(true)
	sRobot = lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Bold(true)
	sBox   = lipgloss.NewStyle().Foreground(lipgloss.Color("142")) // 15 is white, 142 is tan
	sMoved = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
)

func (m tea5model) Init() tea.Cmd {
	return tea.Cmd(func() tea.Msg { return msgRender{} })
}

func (m tea5model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "up":
			if m.delay < 1 {
				m.delay = 1
			} else if m.delay < 1000 {
				m.delay = m.delay * 10
			}
			m.currentRender = m.RenderFancy()
		case "down":
			if m.delay > 1 {
				m.delay = m.delay / 10
			} else {
				m.delay = 0
			}
			m.currentRender = m.RenderFancy()
		case " ":
			// space toggles auto-mode, and beings setipping if set
			m.auto = !m.auto
			if m.auto {
				return m, tea.Cmd(func() tea.Msg { return msgStep{} })
			}
		case "r":
			// toggle rendering
			m.rendering = !m.rendering
		case "s":
			// s disabled auto-mode, does one step if already disabled
			if m.auto {
				m.auto = false
				return m, nil
			}
			return m, tea.Cmd(func() tea.Msg { return msgStep{} })
		}
	case msgRender:
		if m.rendering {
			m.previousRender = m.currentRender
			m.currentRender = m.RenderFancy()
		}
	case msgStep:
		if m.step < len(m.moves) {
			m.Step()
		} else {
			return m, tea.Quit
		}
		if m.rendering {
			cmds = append(cmds, tea.Cmd(func() tea.Msg { return msgRender{} }))
		}
		if m.auto {
			if m.delay > 0 {
				time.Sleep(time.Millisecond * time.Duration(m.delay))
			}
			cmds = append(cmds, tea.Cmd(func() tea.Msg { return msgStep{} }))
		}
	}
	if len(cmds) > 0 {
		return m, tea.Batch(cmds...)
	}
	return m, nil
}

func (m tea5model) Render() string {
	var sb strings.Builder
	var r, c = m.area.IndexToRowCol(m.pos)
	solution := calculateValue2(m.area)
	sb.WriteString(fmt.Sprintf("Step %d value %d\n\n", m.step, solution))
	sb.WriteString(fmt.Sprintf("delay: %d\n\n", m.delay))
	for r = 0; r < m.area.Height; r++ {
		for c = 0; c < m.area.Width; c++ {
			sb.WriteRune(rune(m.area.Data[r*m.area.Width+c]))
		}
		sb.WriteRune('\n')
	}
	if m.step >= len(m.moves) {
		sb.WriteString(fmt.Sprintf("\nDONE!\nSolution: %d\n", solution))
	}
	return sb.String()
}

func chmapMove(move rune) rune {
	// order left,right,up,down
	// see: https://symbl.cc/en/unicode/blocks/geometric-shapes/#subblock-25A0

	// big white (hollow) triangles (too big for terminal char, shows color of next square on right)
	// chars := []rune{'\u25c1','\u25b7','\u25b3','\u25bd'}

	// small black ones (missing one for big triangle in black)
	chars := []rune{'\u25c2', '\u25b8', '\u25b4', '\u25be'}

	switch rune(move) {
	case '<':
		return chars[0]
	case '>':
		return chars[1]
	case '^':
		return chars[2]
	case 'v':
		return chars[3]
	default:
		return '@'
	}
}

func chmapBox(r rune) rune {
	// 228f, 2290 - square image of and square original of: ⊏⊐ ⊏⊐
	switch r {
	case '[':
		return '\u228f'
	case ']':
		return '\u2290'
	default:
		return r
	}
}

func (m *tea5model) RenderFancy() string {
	var sb strings.Builder
	var _r, c = m.area.IndexToRowCol(m.pos)
	solution := calculateValue2(m.area)
	sb.WriteString(fmt.Sprintf("Step %d value %d\n\n", m.step, solution))
	// ms := m.elapsed.Microseconds()
	// elapsed := fmt.Sprintf("%dµs", ms)
	ms := m.elapsed.Nanoseconds()
	elapsed := fmt.Sprintf("%dns", ms)
	sb.WriteString(fmt.Sprintf("delay: %d, last step took:%s\n\n", m.delay, elapsed))
	for _r = 0; _r < m.area.Height; _r++ {
		for c = 0; c < m.area.Width; c++ {
			var pos = m.area.RowColToIndex(_r, c)
			var r = rune(m.area.Data[pos])
			var rs = string(r)
			if r == '\u0123' {
				r = rune(rs[0]) // stupid go wants me to declare variables
			}
			if r == '@' {
				r = chmapMove(m.lastMove)
				if m.moved[pos] {
					sb.WriteString(sMoved.Render(string(r)))
				} else {
					sb.WriteString(sRobot.Render(string(r)))
				}
			} else if r == '#' {
				sb.WriteString(sWall.Render("#"))
				// sb.WriteString(sWall.Render(string(r)))
			} else if r == '[' || r == ']' {
				if m.moved[pos] {
					sb.WriteString(sMoved.Render(string(chmapBox(r))))
				} else {
					sb.WriteString(sBox.Render(string(chmapBox(r))))
				}
			} else if r == '.' {
				sb.WriteRune('\u00b7')
			} else {
				sb.WriteRune(r)
			}
		}
		sb.WriteRune('\n')
	}
	if m.step >= len(m.moves) {
		sb.WriteString(fmt.Sprintf("\nDONE!\nSolution: %d\n", solution))
	}
	return sb.String()
}

func (m tea5model) View() string {
	if m.rendering {
		return m.currentRender
	} else {
		return "rendering paused..."
	}
}

func (m *tea5model) Step() {
	start := time.Now()
	m.pos = m.move2(m.area, byte(m.moves[m.step]), m.pos)
	m.elapsed = time.Since(start)
	m.step = m.step + 1
}

func tea5() {
	contents, err := os.ReadFile(aoc.GetSubPath("input.aoc"))
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}
	tea5part2(string(contents))
}

func tea5part2(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	s := strings.Join(aoc.ParseLines(groups[1]), "")

	area := doubleArea(aoc.ParseArea(groups[0]))
	pos := 0
	for i, b := range area.Data {
		if b == '@' {
			pos = i
			break
		}
	}

	model := tea5model{area: area, moves: []rune(s), currentRender: "INITIALIZED", pos: pos, rendering: true, delay: 10}
	p := tea.NewProgram(model)
	// init should send the msgRender message
	m, err := p.Run()
	if err != nil {
		panic(fmt.Sprintf("Error running tea program!\n%v", err))
	}

	// it's done, print the current state which should be final
	m2 := m.(tea5model)
	fmt.Printf("%s", m2.Render())

	return calculateValue2(area)
}

// Here's where most changes are required.   I need to rethink it.  I think
// Horizontal moves can be pretty much the same I just allow moving all '['
// and ']' instead of 'O', but vertical moves can end up moving multiple
// things from the row above or below.   What I need I think then is to
// build a list of things that will move up.   If I find something that
// will hit a wall in the list, don't move anything.
func (model *tea5model) move2(area *aoc.Area, m byte, pos int) int {
	model.lastMove = rune(m)
	var moved = map[int]bool{}

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
		return model.move2Vertical(area, m, pos, change)
	} else if m == 'v' {
		change = area.Width
		return model.move2Vertical(area, m, pos, change)
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
		if area.Data[newPos-change] != '@' {
			// don't count player as moved
			moved[newPos] = true
		}
		area.Data[newPos] = area.Data[newPos-change]
		newPos -= change
	}

	// we moved over 1
	area.Data[pos] = '.'
	// moved[pos] = true // don't count empty space
	// don't set model if the only thing moved is the robot
	if len(moved) > 0 {
		model.moved = moved
	}
	return pos + change
}

func (model *tea5model) move2Vertical(area *aoc.Area, m byte, pos int, change int) int {
	var moved = map[int]bool{}
	isPushed := map[int]bool{} // if this position has been pushed, we need to take what's below, otherwise use empty
	added := map[int]bool{}

	// start with '@' at pos
	positions := []int{pos}
	added[pos] = true

	checking := 0
	for checking < len(positions) {
		next := positions[checking] + change
		if area.Data[next] == '#' {
			return pos // something hit a wall
		}
		if area.Data[next] == '.' {
			// nothing to add for this position, but flag as pushed so we can
			// move what is pushing up
			// TODO: I don't think there's a need for this the way I decided to do it
			isPushed[next] = true
			checking++
			continue
		}
		// Handle left side of box
		if area.Data[next] == '[' {
			// add only this block as pushed no matter what
			isPushed[next] = true

			// only append if not existing
			if !added[next] {
				positions = append(positions, next)
				added[next] = true
			}

			// append other 1/2 of box if not existing
			if !added[next+1] {
				positions = append(positions, next+1)
				added[next+1] = true
			}
			checking++
			continue
		}

		// Handle right side of box
		if area.Data[next] == ']' {
			// add only this block as pushed no matter what
			isPushed[next] = true

			// only append if not existing
			if !added[next] {
				positions = append(positions, next)
				added[next] = true
			}

			// append other 1/2 of box if not existing
			if !added[next-1] {
				positions = append(positions, next-1)
				added[next-1] = true
			}
			checking++
			continue
		}

		panic("Unknown data in map!")
	}

	// ok, everything in positions is moving up, we should be able to traverse
	// from the back down to the start and move those up, replacing with empties
	// that will be filled by other positions later
	for i := len(positions) - 1; i >= 0; i-- {
		pos := positions[i]
		area.Data[pos+change] = area.Data[pos]
		moved[pos+change] = true
		if isPushed[pos] {
			area.Data[pos] = area.Data[pos-change]
			moved[pos] = true
		} else {
			area.Data[pos] = '.'
			moved[pos] = true // for pretty render
		}
	}

	// remove moved that are player or empty spaces now
	for k := range moved {
		if area.Data[k] == '.' || area.Data[k] == '@' {
			delete(moved, k)
		}
	}
	if len(moved) > 0 {
		model.moved = moved
	}
	return pos + change
}
