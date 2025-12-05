package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func checkNeighbors(t *testing.T, area *aoc.Area, row, col, expected int) {
	count := CountNeighbors(area, row, col)
	if count != expected {
		t.Errorf("CountNeighbors(%d, %d) = %d; want %d", row, col, count, expected)
	}
}

func TestCountNeighbors(t *testing.T) {
	contents := `@.@
@@@
..@`
	area := aoc.ParseArea(contents)
	checkNeighbors(t, area, 0, 0, 2)
	checkNeighbors(t, area, 0, 1, -1)
	checkNeighbors(t, area, 0, 2, 2)
	checkNeighbors(t, area, 1, 0, 2)
	checkNeighbors(t, area, 1, 1, 5)
	checkNeighbors(t, area, 1, 2, 3)
	checkNeighbors(t, area, 2, 0, -1)
	checkNeighbors(t, area, 2, 1, -1)
	checkNeighbors(t, area, 2, 2, 2)
}

func TestPart1(t *testing.T) {
	contents := `@.@
@@@
..@`
	result := part1(contents)
	if result != 5 {
		t.Errorf("part1(...) = %d; want %d", result, 5)
	}
}
