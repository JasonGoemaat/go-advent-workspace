package main

import (
	"fmt"
	"strconv"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", nil)
	aoc.Local(part1, "part1", "input.aoc", nil)
	aoc.Local(part2, "part2", "sample2.aoc", nil)
	aoc.Local(part2, "part2", "input.aoc", nil)
}

type Piece struct {
	Index int
	Grid  []byte
	Used  int
}

type Tree struct {
	Width, Height int
	Counts        []int
}

type Puzzle struct {
	Pieces []Piece
	Trees  []Tree
}

func ParseTrees(content string) []Tree {
	lines := aoc.ParseLines(content)
	intsArrays := aoc.ParseLinesToInts(lines)
	result := make([]Tree, len(lines))
	for i, ints := range intsArrays {
		result[i] = Tree{ints[0], ints[1], ints[2:]}
	}
	return result
}

func ParsePiece(pieceString string) Piece {
	lines := aoc.ParseLines(pieceString)
	index, _ := strconv.Atoi(lines[0])
	grid := make([]byte, 9)
	copy(grid[0:3], []byte(lines[1]))
	copy(grid[3:6], []byte(lines[2]))
	copy(grid[6:9], []byte(lines[3]))
	used := 0
	for _, r := range grid {
		if r == '#' {
			used++
		}
	}
	return Piece{index, grid, used}
}

func ParsePuzzle(content string) Puzzle {
	groups := aoc.ParseGroups(content)
	trees := ParseTrees(groups[len(groups)-1])
	pieces := make([]Piece, len(groups)-1)
	for i, pieceString := range groups[:len(groups)-1] {
		pieces[i] = ParsePiece(pieceString)
	}
	return Puzzle{pieces, trees}
}

func part1(content string) any {
	// empties := []int{4, 2, 3, 2, 2, 2}
	puzzle := ParsePuzzle(content)
	negative := 0
	some := 0
	alot := 0
	for i, tree := range puzzle.Trees {
		totalSize := tree.Width * tree.Height
		totalUsed := 0
		for j, value := range tree.Counts {
			totalUsed += value * puzzle.Pieces[j].Used
		}
		totalEmpty := totalSize - totalUsed
		if totalEmpty < 0 {
			negative++
		} else if totalEmpty < 100 {
			some++
		} else {
			alot++
		}
		fmt.Printf("%d: %d empty and %d used\n", i, totalSize-totalUsed, totalUsed)
	}
	fmt.Printf("Negative: %d\n", negative)
	fmt.Printf("Some: %d\n", some)
	fmt.Printf("Alot: %d\n", alot)
	return nil
}

func part2(content string) any {
	return nil
}
