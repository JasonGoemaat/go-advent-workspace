package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type tea2model struct {
	x, y int
}

func (m tea2model) Init() tea.Cmd {
	return nil
}

func (m tea2model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up":
			m.y = max(0, m.y-1)
		case "down":
			m.y = m.y + 1
		case "left":
			m.x = max(0, m.x-1)
		case "right":
			m.x = m.x + 1
		}
	}
	return m, nil
}

func (m tea2model) View() string {
	sb := strings.Builder{}
	spaces := strings.Repeat(" ", m.x+6)
	for _ = range m.y {
		sb.WriteString(spaces)
		sb.WriteRune('|')
		sb.WriteRune('\n')
	}
	for _ = range m.x {
		sb.WriteRune('-')
	}
	sb.WriteString("Hello, world!")
	return sb.String()
}

func tea2() {
	// p := tea.NewProgram(
	// 	model{},
	// 	tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
	// 	tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	// )

	p := tea.NewProgram(tea2model{})
	m, err := p.Run()
	if err != nil {
		panic("tea did not run!")
	}
	fmt.Printf("tea did run!  model is: %v\n", m)
}
