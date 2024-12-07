package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type tea1model struct {
}

func (m tea1model) Init() tea.Cmd {
	return nil
}

func (m tea1model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m tea1model) View() string {
	return "Hello, world!"
}

func tea1() {
	// p := tea.NewProgram(
	// 	model{},
	// 	tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
	// 	tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	// )

	p := tea.NewProgram(tea1model{})
	m, err := p.Run()
	if err != nil {
		panic("tea did not run!")
	}
	fmt.Printf("tea did run!  model is: %v\n", m)
}
