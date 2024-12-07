package part1

import (
	"github.com/JasonGoemaat/go-aoc/aoc"
	"github.com/JasonGoemaat/go-aoc/aoc/tui"
)

type (
	Keypad struct {
		Rows       [][]string
		CurrentPos KeypadPos
		InitialPos KeypadPos
		Preference []string
	}
	KeypadPos struct {
		X, Y int
	}
)

func NewNumericKeypad() Keypad {
	keypad := Keypad{
		// Layout:     []string{"789", "456", "123", " 0A"},
		Rows: [][]string{
			{"7", "8", "9"},
			{"4", "5", "6"},
			{"1", "2", "3"},
			{" ", "0", "A"},
		},
		CurrentPos: KeypadPos{X: 2, Y: 3},
		InitialPos: KeypadPos{X: 2, Y: 3},
		Preference: []string{"^", "<", ">", "v"},
	}
	return keypad
}

func NewControlKeypad() Keypad {
	keypad := Keypad{
		// Layout:     []string{" ^A", "<v>"},
		Rows: [][]string{
			{" ", "^", "A"},
			{"<", "v", ">"},
		},
		CurrentPos: KeypadPos{X: 2, Y: 0},
		InitialPos: KeypadPos{X: 2, Y: 0},
		Preference: []string{"v", "<", ">", "^"},
	}
	return keypad
}

func (keypad Keypad) FindKey(key string) *KeypadPos {
	for r, row := range keypad.Rows {
		for c, value := range row {
			if value == key {
				return &KeypadPos{c, r}
			}
		}
	}
	return nil
}

func (keypad Keypad) MovesTo(pos KeypadPos) {

}

type (
	ActionStartLine  struct{}
	ActionFinishLine struct{}
	ActionEnqueue    struct {
		Pad  int
		Keys string
	}
	ActionProcessQueue struct {
		Pad int
	}
	ActionEndQueue struct {
		Pad int
	}
	ActionContinueQueue struct {
		Pad int
	}
)

type MyState struct {
	Lines            []string
	Scores           []int
	CurrentLine      int
	IsProcessingLine bool
	CurrentPad       int
	Pads             [4]Keypad
	Queues           [4]string
	QueueIndex       [4]int
	History          [4][]string
	KeyPressed       [4]string
	LastAction       [4]interface{}
	NextAction       interface{}
	StepCount        int
	TotalScore       int
}

func NewMyState(contents string) MyState {
	lines := aoc.ParseLines(contents)
	state := MyState{
		Lines:            lines,
		Scores:           make([]int, len(lines)),
		CurrentLine:      0,
		IsProcessingLine: false,
		CurrentPad:       0,
		Pads: [4]Keypad{
			NewNumericKeypad(),
			NewControlKeypad(),
			NewControlKeypad(),
			NewControlKeypad(),
		},
		Queues:     [4]string{"", "", "", ""},
		QueueIndex: [4]int{0, 0, 0, 0},
		History:    [4][]string{{}, {}, {}, {}},
		KeyPressed: [4]string{"", "", "", ""},
		NextAction: ActionStartLine{},
		StepCount:  0,
	}
	return state
}

func (ms *MyState) IsDone() bool {
	return ms.CurrentLine >= len(ms.Lines)
}

func (ms *MyState) Step() {
	// count our steps
	ms.StepCount++

	// start each step by clearing pressed keys
	for i := range 4 {
		ms.KeyPressed[i] = ""
	}

	// perform action
	switch t := ms.NextAction.(type) {
	case ActionStartLine:
		ms.CurrentPad = 0
		ms.DoStartLine(t)
	case ActionFinishLine:
		ms.CurrentPad = 0
		ms.DoFinishLine(t)
	case ActionEnqueue:
		ms.CurrentPad = t.Pad
		ms.DoEnqueue(t)
	case ActionProcessQueue:
		ms.CurrentPad = t.Pad
		ms.DoProcessQueue(t)
	case ActionContinueQueue:
		ms.CurrentPad = t.Pad
		ms.DoContinueQueue(t)
	case ActionEndQueue:
		ms.CurrentPad = t.Pad
		ms.DoEndQueue(t)
	}
}

func (ms *MyState) GetSolution() interface{} {
	return ms.TotalScore
}

var visualize = true

func Part1(contents string) interface{} {
	ms := NewMyState(contents)

	if visualize {
		tui.RunProgram(&ms)
	} else {
		for !ms.IsDone() {
			ms.Step()
		}
	}

	return ms.GetSolution()
}
