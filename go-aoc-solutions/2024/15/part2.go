package main

import (
	"fmt"
	"strings"

	"github.com/JasonGoemaat/go-aoc/aoc"
)

// Decided to create a new file for part 2 since it is so different

// If the tile is #, the new map contains ## instead.
// If the tile is O, the new map contains [] instead.
// If the tile is ., the new map contains .. instead.
// If the tile is @, the new map contains @. instead.

// first I need to reprocess the area where width will be width * 2
func doubleArea(original *aoc.Area) *aoc.Area {
	area := aoc.Area{Width: original.Width * 2, Height: original.Height, Data: make([]byte, len(original.Data)*2)}
	index := 0
	for r := 0; r < original.Height; r++ {
		for c := 0; c < original.Width; c++ {
			och := original.Get(r, c)
			if och == '#' || och == '.' {
				area.Data[index] = och
				area.Data[index+1] = och
			} else if och == 'O' {
				area.Data[index] = '['
				area.Data[index+1] = ']'
			} else if och == '@' {
				area.Data[index] = '@'
				area.Data[index+1] = '.'
			} else {
				panic("Invalid character in area")
			}
			index += 2
		}
	}
	return &area
}

func part2(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	s := strings.Join(aoc.ParseLines(groups[1]), "")
	area := aoc.ParseArea(groups[0])
	area = doubleArea(area)
	moves := []byte(s)

	aoc.LoggingEnabled = false
	aoc.LogF("\n\n\n----------------------------------------\nInitial state:")
	outputArea(area)

	// find start
	pos := 0
	for i, b := range area.Data {
		if b == '@' {
			pos = i
			break
		}
	}

	for _, b := range moves {
		pos = move2(area, b, pos)
		aoc.LogF("\n")
		aoc.LogF("Move %c:\n", b)
		outputArea(area)
	}

	return calculateValue2(area)
}

// Here's where most changes are required.   I need to rethink it.  I think
// Horizontal moves can be pretty much the same I just allow moving all '['
// and ']' instead of 'O', but vertical moves can end up moving multiple
// things from the row above or below.   What I need I think then is to
// build a list of things that will move up.   If I find something that
// will hit a wall in the list, don't move anything.
func move2(area *aoc.Area, m byte, pos int) int {
	// find space or wall in direction, if wall just return
	change := 0
	// vis := visualizeArea(area)

	// NOTE: problem input is surrounded by walls, so no need to
	// check bounds since we'll always hit a wall, therefore we
	// can just track index change in grid
	if m == '<' {
		change = -1
	} else if m == '>' {
		change = 1
	} else if m == '^' {
		change = -area.Width
		return move2Vertical(area, m, pos, change)
	} else if m == 'v' {
		change = area.Width
		return move2Vertical(area, m, pos, change)
	} else {
		panic(fmt.Sprintf("Invalid move: %v", m))
	}

	// move foward until hitting empty space or wall
	newPos := pos + change
	for (area.Data[newPos] != '.') && (area.Data[newPos] != '#') {
		newPos += change
	}

	// if the final position is a wall, we can't move
	if area.Data[newPos] == '#' {
		return pos
	}

	// must have hit '.' or empty space, move everything forward 1
	for newPos != pos {
		aoc.LogF("Moving %c at %d replacing %c at %d\n",
			area.Data[newPos-change], newPos-change,
			area.Data[newPos], newPos)
		outputArea(area)
		area.Data[newPos] = area.Data[newPos-change]
		outputArea(area)
		aoc.LogF("\n")
		newPos -= change
	}

	// we moved over 1
	area.Data[pos] = '.'
	return pos + change
}

func move2Vertical(area *aoc.Area, m byte, pos int, change int) int {
	r, c := area.IndexToRowCol(pos)
	aoc.LogF("\n\nmove2Vertical(%d,%d) - '%c'\n", r, c, m)
	outputArea(area)

	isPushed := map[int]bool{} // if this position has been pushed, we need to take what's below, otherwise use empty
	added := map[int]bool{}

	// start with '@' at pos
	positions := []int{pos}
	added[pos] = true

	checking := 0
	for checking < len(positions) {
		next := positions[checking] + change
		if area.Data[next] == '#' {
			return pos // something hit a wall
		}
		if area.Data[next] == '.' {
			// nothing to add for this position, but flag as pushed so we can
			// move what is pushing up
			// TODO: I don't think there's a need for this the way I decided to do it
			isPushed[next] = true
			checking++
			continue
		}
		// Handle left side of box
		if area.Data[next] == '[' {
			// add only this block as pushed no matter what
			isPushed[next] = true

			// only append if not existing
			if !added[next] {
				positions = append(positions, next)
				added[next] = true
			}

			// append other 1/2 of box if not existing
			if !added[next+1] {
				positions = append(positions, next+1)
				added[next+1] = true
			}
			checking++
			continue
		}

		// Handle right side of box
		if area.Data[next] == ']' {
			// add only this block as pushed no matter what
			isPushed[next] = true

			// only append if not existing
			if !added[next] {
				positions = append(positions, next)
				added[next] = true
			}

			// append other 1/2 of box if not existing
			if !added[next-1] {
				positions = append(positions, next-1)
				added[next-1] = true
			}
			checking++
			continue
		}

		panic("Unknown data in map!")
	}

	// ok, everything in positions is moving up, we should be able to traverse
	// from the back down to the start and move those up, replacing with empties
	// that will be filled by other positions later
	for i := len(positions) - 1; i >= 0; i-- {
		pos := positions[i]
		area.Data[pos+change] = area.Data[pos]
		if isPushed[pos] {
			area.Data[pos] = area.Data[pos-change]
		} else {
			area.Data[pos] = '.'
		}
	}

	aoc.LogF("Result:\n")
	outputArea(area)
	return pos + change
}

func calculateValue2(area *aoc.Area) int {
	total := 0
	for r := range area.Height {
		for c := range area.Width {
			if area.Get(r, c) == '[' {
				total += (100 * r) + c
			}
		}
	}
	return total
}
