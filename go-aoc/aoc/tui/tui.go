package tui

// Code to handle tui elements and a program that allows stepping through or
// running solvers in a standard way.

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	TuiModel struct {
		// info for current run controlled by this package
		CurrentStep int
		Rendering   bool
		Auto        bool
		Delay       int
		State       TuiActions
	}
	TuiRenderMessage struct{}
	TuiStepMessage   struct{}
	TuiActions       interface {
		Render() string
		IsDone() bool
		Step()
		GetSolution() interface{}
	}
)

func (model TuiModel) Init() tea.Cmd {
	return tea.Cmd(func() tea.Msg { return TuiRenderMessage{} })
}

func (model TuiModel) View() string {
	if model.Rendering {
		return model.State.Render()
	} else {
		return "rendering paused..."
	}
}

func (model TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return model, tea.Quit
		case "up":
			if model.Delay < 1 {
				model.Delay = 1
			} else if model.Delay < 1000 {
				model.Delay = model.Delay * 10
			}
		case "down":
			if model.Delay > 1 {
				model.Delay = model.Delay / 10
			} else {
				model.Delay = 0
			}
		case " ":
			// space toggles auto-mode, and beings setipping if set
			model.Auto = !model.Auto
			if model.Auto {
				return model, tea.Cmd(func() tea.Msg { return TuiStepMessage{} })
			}
		case "r":
			// toggle rendering
			model.Rendering = !model.Rendering
		case "s":
			// s disabled auto-mode, does one step if already disabled
			if model.Auto {
				model.Auto = false
				return model, nil
			}
			return model, tea.Cmd(func() tea.Msg { return TuiStepMessage{} })
		}
	case TuiStepMessage:
		if !model.State.IsDone() {
			model.State.Step()
		} else {
			return model, tea.Quit
		}
		if model.Auto {
			if model.Delay > 0 {
				time.Sleep(time.Millisecond * time.Duration(model.Delay))
			}
			cmds = append(cmds, tea.Cmd(func() tea.Msg { return TuiStepMessage{} }))
		}
	}
	if len(cmds) > 0 {
		return model, tea.Batch(cmds...)
	}
	return model, nil
}

func RunProgram(state TuiActions) {
	model := TuiModel{0, true, false, 100, state}
	program := tea.NewProgram(model)
	m, err := program.Run()
	if err != nil {
		panic(fmt.Sprintf("Error running tea program!\n%v", err))
	}

	// it's done, print the current state which should be final
	resultModel := m.(TuiModel)
	fmt.Printf("%s\n", resultModel.State.Render())
	fmt.Printf("Solution: %v\n", resultModel.State.GetSolution())
}
