package main

import (
	"fmt"

	"github.com/JasonGoemaat/go-aoc/aoc"
)

type day16 struct {
	area       *aoc.Area
	start, end day16position
	queue      Queue[day16position]
	scores     map[day16position]int
	score      int
	step       int
}

type day16position struct {
	r, c   int
	facing int
}

func NewDay16(contents string) *day16 {
	state := day16{}
	state.Init(contents)
	return &state
}

func (state *day16) Render() string {
	return fmt.Sprintf("Step %d, score %d, queue has %d", state.step, state.score, state.queue.Count())
}

func (state *day16) Init(contents string) {
	state.area = aoc.ParseArea(contents)
	state.step = 0
	state.queue = Queue[day16position]{backing: make([]day16position, 4096), head: 0, tail: 0}
	state.scores = map[day16position]int{}

	for i, b := range state.area.Data {
		if b == 'S' {
			r, c := state.area.IndexToRowCol(i)
			state.start = day16position{r, c, 0}
		} else if b == 'E' {
			r, c := state.area.IndexToRowCol(i)
			state.end = day16position{r, c, 0}
		}
	}
	state.queue.Enqueue(state.start)
}

func (state *day16) try(pos day16position, score int) {
	// if (state.end.r == pos.r) && (state.end.c == pos.c) {
	// 	if (state.score == 0) || (score < state.score) {
	// 		state.score = score
	// 	}
	// 	return
	// }
	existingScore, exists := state.scores[pos]
	if exists && (existingScore <= score) {
		return
	}
	state.scores[pos] = score
	state.queue.Enqueue(pos)
}

func (state *day16) Step() {
	state.step = state.step + 1
	pos := state.queue.Dequeue()
	r, c, facing := pos.r, pos.c, pos.facing
	score := state.scores[pos]

	// check if turns exist or we can do faster
	state.try(day16position{r, c, (facing + 1) & 3}, score+1000)
	state.try(day16position{r, c, (facing + 2) & 3}, score+2000)
	state.try(day16position{r, c, (facing + 3) & 3}, score+1000)

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
	if r < 0 || r > state.area.Height || c < 0 || c > state.area.Width {
		return
	}
	b := state.area.Get(r, c)
	if b == 'E' {
		if (state.end.r == r) && (state.end.c == c) {
			if (state.score == 0) || (score < (state.score + 1)) {
				state.score = score + 1
			}
		}
		return
	}
	if b == '.' {
		state.try(day16position{r, c, facing}, score+1)
	}
}
