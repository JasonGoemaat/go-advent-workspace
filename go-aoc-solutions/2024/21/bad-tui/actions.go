package part1

import (
	"fmt"
	"strconv"
	"strings"
)

// Perform ActionStartLine. CurrentLine is the line to start.
// Need to reset all keypads and data and queue up the line.
func (ms *MyState) DoStartLine(action ActionStartLine) {
	for kpi := range 4 {
		ms.Pads[kpi].CurrentPos = ms.Pads[kpi].InitialPos
		ms.Queues[kpi] = ""
		ms.QueueIndex[kpi] = 0
		if kpi != 0 {
			// don't reset first history, so we show socres
			ms.History[kpi] = []string{}
		}
		ms.KeyPressed[kpi] = ""
	}
	ms.CurrentPad = 0
	ms.IsProcessingLine = true
	ms.NextAction = ActionEnqueue{0, ms.Lines[ms.CurrentLine]}
}

// Perform ActionFinishLine.  Calculate score from last Pad history
// and integer from digits in line (first 3 characters).
func (ms *MyState) DoFinishLine(action ActionFinishLine) {
	allMoves := strings.Join(ms.History[3], "")
	movesCount := len(allMoves)
	digitValue, _ := strconv.Atoi(ms.Lines[ms.CurrentLine][0:3])
	ms.Scores[ms.CurrentLine] = movesCount * digitValue
	ms.TotalScore += movesCount * digitValue

	// REPLACE last history which would be what was on the keypad with scores
	ms.History[0][len(ms.History[0])-1] = fmt.Sprintf("%s: %d * %d = %d", ms.Queues[0], movesCount, digitValue, movesCount*digitValue)
	ms.CurrentLine++
	ms.NextAction = ActionStartLine{}
}

// Setup action.Pad with action.Keys, queue up ProcessQueue next
func (ms *MyState) DoEnqueue(action ActionEnqueue) {
	ms.Queues[action.Pad] = action.Keys
	ms.QueueIndex[action.Pad] = 0
	ms.NextAction = ActionProcessQueue{Pad: action.Pad}
}

// Handle queue on specified pad.
//
// This will do one of two things:
// 1. On final Pad, we don't need to actually do anything, just submit ContinueQueue
// 2. On other Pads, calculate moves required for next key submit Enqueue(action.Pad + 1, keys)
func (ms *MyState) DoProcessQueue(action ActionProcessQueue) {
	if action.Pad == 3 {
		ms.NextAction = ActionContinueQueue(action)
		return
	}

	pad := action.Pad
	queueIndex := ms.QueueIndex[action.Pad]
	queue := ms.Queues[action.Pad]

	// TODO: calculate moves for child to do next move in queue
	// Example 1: Digital Keypad at 'A' (2,3) needs to move to '0' (1,3) and press it
	nextKey := queue[queueIndex : queueIndex+1]
	pos := ms.Pads[pad].CurrentPos
	target := ms.Pads[pad].FindKey(nextKey)
	sb := strings.Builder{}
	for _, p := range ms.Pads[pad].Preference {
		for p == "<" && target.X < pos.X {
			sb.WriteString(p)
			pos.X--
		}
		for p == ">" && target.X > pos.X {
			sb.WriteString(p)
			pos.X++
		}
		for p == "^" && target.Y < pos.Y {
			sb.WriteString(p)
			pos.Y--
		}
		for p == "v" && target.Y > pos.Y {
			sb.WriteString(p)
			pos.Y++
		}
	}
	sb.WriteString("A")

	// ok, submit to next pad in list
	ms.NextAction = ActionEnqueue{action.Pad + 1, sb.String()}
}

// submitted by DoEndQueue of lower queue
// this is continued when pressing a key, so mark it as pressed
// move queue index forward and submit EndQueue if done, ProcessQueue if not
func (ms *MyState) DoContinueQueue(action ActionContinueQueue) {
	currentKey := ms.Queues[action.Pad][ms.QueueIndex[action.Pad] : ms.QueueIndex[action.Pad]+1]
	ms.KeyPressed[action.Pad] = currentKey
	if action.Pad == 3 {
		ms.Pads[action.Pad].CurrentPos = *ms.Pads[action.Pad].FindKey(currentKey)
	}
	if action.Pad > 0 {
		pp := &ms.Pads[action.Pad-1]
		switch currentKey {
		case "<":
			pp.CurrentPos.X--
		case ">":
			pp.CurrentPos.X++
		case "^":
			pp.CurrentPos.Y--
		case "v":
			pp.CurrentPos.Y++
		case "A":
			// ms.KeyPressed[action.Pad-1] = ms.Queues[action.Pad-1][ms.QueueIndex[action.Pad-1] : ms.QueueIndex[action.Pad-1]+1]
		}
	}
	ms.QueueIndex[action.Pad]++
	if ms.QueueIndex[action.Pad] >= len(ms.Queues[action.Pad]) {
		// ms.NextAction = ActionEndQueue{action.Pad}
		ms.NextAction = ActionEndQueue(action)
	} else {
		// ms.NextAction = ActionProcessQueue{action.Pad}
		ms.NextAction = ActionProcessQueue(action)
	}
}

// submitted by ContinueQueue when it moves past the end instead of ProcessQueue
// submits ContinueQueue for the previous queue
func (ms *MyState) DoEndQueue(action ActionEndQueue) {
	// move queue to history
	ms.History[action.Pad] = append(ms.History[action.Pad], ms.Queues[action.Pad])
	if action.Pad == 0 {
		ms.NextAction = ActionFinishLine{}
	} else {
		ms.NextAction = ActionContinueQueue{action.Pad - 1}
	}
}
