package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/6
	aoc.Local(part1, "part1", "sample.aoc", 41)
	aoc.Local(part1, "part1", "input.aoc", 4890)
	aoc.Local(part2, "part2", "sample.aoc", 6)
	aoc.Local(part2, "part2", "input.aoc", 1995)
}

var up = []int{-1, 0}
var down = []int{1, 0}
var left = []int{0, -1}
var right = []int{0, 1}

// directions in order, start by going up
var directions = [][]int{up, right, down, left}

var rowCount = 0
var colCount = 0
var theMap []byte = nil

func at(r, c int) byte {
	pos, outside := getPos(r, c)
	if outside {
		return 0
	}
	return theMap[pos]
}

func getPos(r, c int) (int, bool) {
	if r < 0 || r >= rowCount || c < 0 || c >= colCount {
		return -1, true
	}
	return r*colCount + c, false
}

func move(r, c, dir int) (int, int, bool) {
	newR := r + directions[dir][0]
	newC := c + directions[dir][1]
	_, outside := getPos(newR, newC)
	return newR, newC, outside
}

func part1(content string) interface{} {
	var lines = aoc.ParseLines(content)
	rowCount = len(lines)
	colCount = len(lines[0])
	theMap = make([]byte, rowCount*colCount)
	startRow := 0
	startCol := 0
	r := 0
	c := 0
	for r = 0; r < rowCount; r++ {
		for c = 0; c < colCount; c++ {
			pos, _ := getPos(r, c)
			theMap[pos] = lines[r][c]
			if theMap[pos] == '^' {
				startRow = r
				startCol = c
			}
		}
	}

	visitedCount := 0
	currentDirection := 0
	r = startRow
	c = startCol
	for r >= 0 && r < rowCount && c >= 0 && c < colCount {
		pos, _ := getPos(r, c)
		if theMap[pos] != 'X' {
			visitedCount++
			theMap[pos] = 'X'
		}
		nextR, nextC, outside := move(r, c, currentDirection)
		if outside {
			break
		}
		if at(nextR, nextC) == '#' {
			currentDirection = (currentDirection + 1) % 4
		} else {
			r = nextR
			c = nextC
		}
	}
	return visitedCount
}

///--------------------------------------------------------------------------------
/// Part 2
///--------------------------------------------------------------------------------

// We need to store a copy for after we place the test obstacle, will
// be copied from the main map every time.
// for each location, we need to store if we've been there travelling in each
// direction.  For this I'll use bitmasks, 1 == direction 0 (up), 2 = right,
// 3 == down, 4 ==  left.   This will also be copied every time
var mapCopy []byte = nil
var beenMoving []byte = nil
var beenMovingCopy []byte = nil

// We also need to keep track of blockers we have already placed that have
// led to loops so we don't try to place them again.
var blockers map[int]bool = nil

// / This is called to place a block in a map square and see if there is a
// / loop.   If there is a loop, it updates the map of positions where we
// / can place a block to cause a loop and restores the state, then returns
// / true.
func hasLoop(r, c, dir int) bool {
	// save var theMap []byte = nil
	copy(mapCopy, theMap)
	copy(beenMovingCopy, beenMoving)
	nr, nc, _ := move(r, c, dir)
	np, _ := getPos(nr, nc)
	mapCopy[np] = 'O'

	for r >= 0 && r < rowCount && c >= 0 && c < colCount {
		pos, isOutside := getPos(r, c)
		if isOutside {
			break
		}
		if (beenMovingCopy[pos] & (1 << dir)) > 0 {
			// found our loop!
			return true
		}

		beenMovingCopy[pos] |= (1 << dir)
		if mapCopy[pos] != 'X' {
			mapCopy[pos] = 'X'
		}
		nextR, nextC, outside := move(r, c, dir)
		if outside {
			break
		}

		// turn right instead if moving ahead would hit a blocker
		nextPos, _ := getPos(nextR, nextC)
		if mapCopy[nextPos] == '#' || mapCopy[nextPos] == 'O' {
			dir = (dir + 1) % 4
			continue
		}
		r = nextR
		c = nextC
	}

	return false
}

func part2(content string) interface{} {
	var lines = aoc.ParseLines(content)
	rowCount = len(lines)
	colCount = len(lines[0])
	theMap = make([]byte, rowCount*colCount)
	mapCopy = make([]byte, len(theMap))
	beenMoving = make([]byte, len(theMap))
	beenMovingCopy = make([]byte, len(theMap))
	blockers = make(map[int]bool)

	startRow := 0
	startCol := 0
	r := 0
	c := 0
	for r = 0; r < rowCount; r++ {
		for c = 0; c < colCount; c++ {
			pos, _ := getPos(r, c)
			theMap[pos] = lines[r][c]
			beenMoving[pos] = 0
			if theMap[pos] == '^' {
				startRow = r
				startCol = c
			}
		}
	}

	visitedCount := 0 // don't need this
	blockerCount := 0
	currentDirection := 0
	r = startRow
	c = startCol
	for r >= 0 && r < rowCount && c >= 0 && c < colCount {
		pos, isOutside := getPos(r, c)
		if isOutside {
			break
		}
		if theMap[pos] != 'X' {
			visitedCount++
			theMap[pos] = 'X'
		}
		nextR, nextC, outside := move(r, c, currentDirection)
		if outside {
			break
		}

		// turn right instead if moving ahead would hit a blocker
		nextPos, _ := getPos(nextR, nextC)
		nextItem := theMap[nextPos]
		if nextItem == '#' {
			beenMoving[pos] |= (1 << currentDirection)
			currentDirection = (currentDirection + 1) % 4
			continue
		}

		// next spot is open and would not cause use to move outside
		// so try to put a blocker there and check for a loop if we haven't
		// found one on that spot already
		// NOTE: Solution might be to check if path has already been used, it
		// wouldn't make sense to add a blocker if the path was already taken
		// if theMap[nextPos] == '.' && !blockers[nextPos] && (nextR != startRow || nextC != startCol) {
		if theMap[nextPos] == '.' && !blockers[nextPos] {
			if hasLoop(r, c, currentDirection) {
				blockers[nextPos] = true
				blockerCount++
			}
		}

		r = nextR
		c = nextC
	}
	return blockerCount
}
