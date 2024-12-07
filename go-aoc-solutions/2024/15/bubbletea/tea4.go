package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
	tea "github.com/charmbracelet/bubbletea"
)

type tea4model struct {
	area           *aoc.Area
	moves          []rune
	previousRender string
	currentRender  string
	step           int
	pos            int
	rendering      bool
	auto           bool
	delay          int
}

type msgRender struct{}
type msgStep struct{}

func CmdRender() tea.Msg {
	return msgRender{}
}

func (m tea4model) Init() tea.Cmd {
	return tea.Cmd(func() tea.Msg { return msgRender{} })
}

func (m tea4model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up":
			if m.delay < 1000 {
				m.delay = m.delay * 10
			}
		case "down":
			if m.delay > 1 {
				m.delay = m.delay / 10
			}
		case " ":
			// space toggles auto-mode, and beings setipping if set
			m.auto = !m.auto
			if m.auto {
				return m, tea.Cmd(func() tea.Msg { return msgStep{} })
			}
		case "s":
			// s disabled auto-mode, does one step if already disabled
			if m.auto {
				m.auto = false
				return m, nil
			}
			return m, tea.Cmd(func() tea.Msg { return msgStep{} })
		}
	case msgRender:
		m.previousRender = m.currentRender
		m.currentRender = m.Render()
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
			time.Sleep(time.Millisecond * time.Duration(m.delay))
			cmds = append(cmds, tea.Cmd(func() tea.Msg { return msgStep{} }))
		}
	}
	if len(cmds) > 0 {
		return m, tea.Batch(cmds...)
	}
	return m, nil
}

func (m tea4model) Render() string {
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

func (m tea4model) View() string {
	return m.Render()
}

func (m *tea4model) Step() {
	m.pos = move2(m.area, byte(m.moves[m.step]), m.pos)
	m.step = m.step + 1
}

func tea4() {
	contents, err := os.ReadFile(aoc.GetSubPath("sample2.aoc"))
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}
	tea4part2(string(contents))
}

func tea4part2(contents string) interface{} {
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

	model := tea4model{area: area, moves: []rune(s), currentRender: "INITIALIZED", pos: pos, rendering: true, delay: 10}
	p := tea.NewProgram(model)
	// init should send the msgRender message
	m, err := p.Run()
	if err != nil {
		panic(fmt.Sprintf("Error running tea program!\n%v", err))
	}

	// it's done, print the current state which should be final
	m2 := m.(tea4model)
	fmt.Printf("%s", m2.currentRender)

	return calculateValue2(area)
}
