package main

import (
	"fmt"

	tui "github.com/JasonGoemaat/go-aoc/aoc/tui"
)

type MyState struct {
	MyTitle    string
	MyContents string
	StepCount  int
}

func (state *MyState) Render() string {
	return fmt.Sprintf("Title: %s\nContents: %s\nStepCount: %d", state.MyTitle, state.MyContents, state.StepCount)
}

func (state *MyState) IsDone() bool {
	return state.StepCount > 0
}

func (state *MyState) Step() {
	state.StepCount++
}

func (state *MyState) GetSolution() interface{} {
	if state.StepCount <= 0 {
		return 0
	}
	return 42
}

func NewTestModel(contents string) *MyState {
	var tm MyState
	tm = MyState{MyTitle: "MyTitle", MyContents: contents}
	return &tm
}

func main() {
	tm := NewTestModel("Hello, world!")
	tui.RunProgram(tm)
}
