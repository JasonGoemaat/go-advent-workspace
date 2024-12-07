package part1

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func fg(v interface{}) lipgloss.Style {
	s := fmt.Sprintf("%v", v)
	return lipgloss.NewStyle().Foreground(lipgloss.Color(s))
}

var (
	sNormal         = lipgloss.NewStyle()
	sBlue           = fg(12)
	sGold           = fg(11)
	sCyan           = fg(14)
	sRed            = fg(1)
	sCentered       = lipgloss.NewStyle().Width(32).Align(lipgloss.Center)
	sCenteredKeypad = lipgloss.NewStyle().Width(24).Align(lipgloss.Center)
	sCenteredBorder = sCentered.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("13"))
	sTableCurrent   = sBlue.Reverse(true).Padding(1).PaddingLeft(2).PaddingRight(2)
	sTablePressed   = sNormal.Bold(true).Reverse(true).Padding(1).PaddingLeft(2).PaddingRight(2)
	sTableNormal    = sBlue.Padding(1).PaddingLeft(2).PaddingRight(2)
)

func (ms *MyState) GetActionString() string {
	switch t := ms.NextAction.(type) {
	case ActionStartLine:
		return "ActionStartLine"
	case ActionFinishLine:
		return "ActionFinishLine"
	case ActionEnqueue:
		return fmt.Sprintf("ActionEnqueue(%d, %q)", t.Pad, t.Keys)
	case ActionProcessQueue:
		return fmt.Sprintf("ActionProcessQueue(%d)", t.Pad)
	case ActionContinueQueue:
		return fmt.Sprintf("ActionContinueQueue(%d)", t.Pad)
	case ActionEndQueue:
		return fmt.Sprintf("ActionEndQueue(%d)", t.Pad)
	default:
		return "UNKNOWN ACTION"
	}
}

func rv(st lipgloss.Style, v interface{}) string {
	return st.Render(fmt.Sprintf("%v", v))
}

func rf(st lipgloss.Style, f string, args ...interface{}) string {
	s := fmt.Sprintf(f, args...)
	return st.Render(s)
}

func renderActionInfo(ms *MyState) string {
	line := rv(sRed, "NONE")
	if ms.CurrentLine < len(ms.Lines) {
		line = rv(sGold, ms.Lines[ms.CurrentLine])
	}
	lines := []string{
		"Action: " + sCyan.Render(ms.GetActionString()),
		"Steps: " + rv(sCyan, ms.StepCount),
		"Line: " + rv(sCyan, ms.CurrentLine) + " - " + rv(sGold, line),
	}
	// return sCenteredBorder.Render(strings.Join(lines, "\n"))
	return sCenteredBorder.Render(strings.Join(lines, "\n"))
}

var rows = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{" ", "0", "A"},
}

var tbl = table.New().
	Border(lipgloss.NormalBorder()).
	BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
	BorderRow(true).
	BorderColumn(true).
	StyleFunc(func(row, col int) lipgloss.Style {
		switch {
		case rows[row][col] == "2":
			return sTableCurrent
		case rows[row][col] == "9":
			return sTablePressed
		default:
			return sTableNormal
		}
	}).
	// Headers("LANGUAGE", "FORMAL", "INFORMAL").
	Rows(rows...)

func renderKeypad(ms *MyState, index int) string {
	var sb = strings.Builder{}

	colorString := "99"
	if index == ms.CurrentPad {
		colorString = "15"
	}
	var tbl = table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(colorString))).
		BorderRow(true).
		BorderColumn(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case ms.Pads[index].Rows[row][col] == ms.KeyPressed[index]:
				return sTablePressed
			case (col == ms.Pads[index].CurrentPos.X) && (row == ms.Pads[index].CurrentPos.Y):
				return sTableCurrent
			default:
				return sTableNormal
			}
		}).
		Rows(ms.Pads[index].Rows...)

	sb.WriteString(tbl.Render())
	return sb.String()
}

// var normalBorder = lipgloss.Border{
// 	Top:          "─",
// 	Bottom:       "─",
// 	Left:         "│",
// 	Right:        "│",
// 	TopLeft:      "┌",
// 	TopRight:     "┐",
// 	BottomLeft:   "└",
// 	BottomRight:  "┘",
// 	MiddleLeft:   "├",
// 	MiddleRight:  "┤",
// 	Middle:       "┼",
// 	MiddleTop:    "┬",
// 	MiddleBottom: "┴",
// }

func renderColumn(ms *MyState, index int) string {
	q := ""
	qi := ms.QueueIndex[index]
	if qi < len(ms.Queues[index]) {
		if qi > 0 {
			q = sBlue.Render(ms.Queues[index][0:qi])
		}
		if qi < len(ms.Queues[index]) {
			q = q + sGold.Render(ms.Queues[index][qi:qi+1])
		}
		if qi < (len(ms.Queues[index]) - 1) {
			q = q + sBlue.Render(ms.Queues[index][qi+1:len(ms.Queues[index])])
		}
	}

	kps := renderKeypad(ms, index)
	li := max(0, len(ms.History[index])-10)
	hi := min(li+10, len(ms.History[index]))
	history := strings.Join(ms.History[index][li:hi], "\n")
	extra := ""
	if (hi - li) < 10 {
		extra = strings.Repeat("\n", 10-(hi-li))
	}
	all := strings.Join([]string{
		kps,
		"",
		q,
		"",
		history,
	}, "\n") + extra
	return sCenteredKeypad.Render(all)
}

func (ms *MyState) Render() string {
	actionInfo := renderActionInfo(ms)

	paddedLeft := lipgloss.NewStyle().PaddingLeft(3)

	kp1 := renderColumn(ms, 0)
	kp2 := paddedLeft.Render(renderColumn(ms, 1))
	kp3 := paddedLeft.Render(renderColumn(ms, 2))
	kp4 := paddedLeft.Render(renderColumn(ms, 3))
	columns := lipgloss.JoinHorizontal(lipgloss.Top, kp1, kp2, kp3, kp4, actionInfo)
	return columns
}
