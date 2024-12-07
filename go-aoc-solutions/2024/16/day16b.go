package main

import (
	"fmt"
	"strings"

	"github.com/JasonGoemaat/go-aoc/aoc"
)

type day16b struct {
	area         *aoc.Area
	start, end   day16bposition
	queue        Queue[day16bposition]
	scores       map[day16bposition]int
	previous     map[day16bposition][]day16bposition
	endpositions map[day16bposition]bool
	score        int
	step         int
}

type day16bposition struct {
	r, c   int
	facing int
}

func NewDay16b(contents string) *day16b {
	state := day16b{}
	state.Init(contents)
	return &state
}

func (state *day16b) Render() string {
	sb := strings.Builder{}
	for r := range state.area.Height {
		for c := range state.area.Width {
			sb.WriteRune(rune(state.area.Data[r*state.area.Width+c]))
		}
		sb.WriteRune('\n')
	}
	answer := fmt.Sprintf("Step %d, score %d, queue has %d", state.step, state.score, state.queue.Count())
	return fmt.Sprintf("%s\n%s", sb.String(), answer)
}

func (state *day16b) Init(contents string) {
	state.area = aoc.ParseArea(contents)
	state.step = 0
	state.queue = Queue[day16bposition]{backing: make([]day16bposition, 4096), head: 0, tail: 0}
	state.scores = map[day16bposition]int{}
	state.previous = map[day16bposition][]day16bposition{} // how we got here with the lowest score, could be multiple
	state.endpositions = map[day16bposition]bool{}         // positions leading to the best end score

	for i, b := range state.area.Data {
		if b == 'S' {
			r, c := state.area.IndexToRowCol(i)
			state.start = day16bposition{r, c, 0}
		} else if b == 'E' {
			r, c := state.area.IndexToRowCol(i)
			state.end = day16bposition{r, c, 0}
		}
	}
	state.queue.Enqueue(state.start)
}

func (state *day16b) backtrace(pos day16bposition, handled map[day16bposition]bool) {
	if handled[pos] {
		return
	}
	handled[pos] = true
	index := state.area.RowColToIndex(pos.r, pos.c)
	state.area.Data[index] = 'O'
	previous, exists := state.previous[pos]
	if !exists {
		return
	}
	for _, p := range previous {
		state.backtrace(p, handled)
	}
}

func (state *day16b) calculateScore() int {
	// find all cheapest ways to get to the end, may be multiple facings
	ends := []day16bposition{}
	bestScore := 0
	for facing := range 4 {
		searchPos := day16bposition{state.end.r, state.end.c, facing}
		score, exists := state.scores[searchPos]
		if !exists || (score > bestScore && bestScore > 0) {
			continue
		}
		if bestScore == 0 || score < bestScore {
			bestScore = score
			// new list containing our searchPos
			ends = []day16bposition{searchPos}
		} else {
			// add to list
			ends = append(ends, searchPos)
		}
	}
	handled := map[day16bposition]bool{}
	for _, end := range ends {
		state.backtrace(end, handled)
	}
	count := 0
	for _, b := range state.area.Data {
		if b == 'O' {
			count++
		}
	}
	// fmt.Printf("%s\n", state.Render())
	return count
}

func (state *day16b) try(pos, previous day16bposition, score int) {
	// if (state.end.r == pos.r) && (state.end.c == pos.c) {
	// 	if (state.score == 0) || (score < state.score) {
	// 		state.score = score

	// 	}
	// 	if state.score == score {

	// 	}
	// 	return
	// }

	existingScore, exists := state.scores[pos]
	if exists {
		// we already have this position/facing in the score list
		if existingScore < score {
			// if there is a better way to get to this position/facing, just return
			return
		}
		if existingScore == score {
			// there is the same position/facing/score, add previous as another way if it isn't already there
			for _, prev := range state.previous[pos] {
				if prev.r == previous.r && prev.c == previous.c && prev.facing == previous.facing {
					// previous already exists and score is the same, return
					return
				}
			}
			// previous position does not exist in list yet, add it
			state.previous[pos] = append(state.previous[pos], previous)
			return
		}
	}

	// doesn't exist or new score is less, replace existing
	state.scores[pos] = score
	state.previous[pos] = []day16bposition{previous}

	// queue if we are not at end
	if (pos.r != state.end.r) || (pos.c != state.end.c) {
		state.queue.Enqueue(pos)
	}
}

func (state *day16b) Step() bool {
	if state.queue.IsEmpty() {
		return false
	}
	state.step = state.step + 1
	pos := state.queue.Dequeue()
	r, c, facing := pos.r, pos.c, pos.facing
	score := state.scores[pos]

	// check if turning can be used to find cheaper or same cost x,y,facing if not already at end
	// NOTE: end will never be added to queue so we don't worry
	state.try(day16bposition{r, c, (facing + 1) & 3}, pos, score+1000)
	state.try(day16bposition{r, c, (facing + 2) & 3}, pos, score+2000)
	state.try(day16bposition{r, c, (facing + 3) & 3}, pos, score+1000)

	// try moving forwards
	switch pos.facing {
	case 0:
		c = c + 1
	case 1:
		r = r - 1
	case 2:
		c = c - 1
	case 3:
		r = r + 1
	}
	if r < 0 || r >= state.area.Height || c < 0 || c >= state.area.Width {
		return true
	}
	b := state.area.Get(r, c)
	if b == '.' || b == 'E' { // don't worry about returning through 'S', that shouldn't occur
		state.try(day16bposition{r, c, facing}, pos, score+1)
	}
	return true
}
