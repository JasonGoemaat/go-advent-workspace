package main

import (
	"regexp"
	"sort"
	"strconv"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

type JunctionBox struct {
	X, Y, Z int
}

type Link struct {
	Distance       int
	Index1, Index2 int
}

type Puzzle struct {
	JunctionBoxes    []JunctionBox
	Connections      int // Number of connections to make
	Links            []Link
	JunctionCircuits map[int]int // index is junction box index, value is circuit index
	CircuitCount     int
}

var rxJunctionBox = regexp.MustCompile(`\d+`)

func ParseJunctionBox(line string) JunctionBox {
	matches := rxJunctionBox.FindAllString(line, -1)
	X, _ := strconv.Atoi(matches[0])
	Y, _ := strconv.Atoi(matches[1])
	Z, _ := strconv.Atoi(matches[2])
	return JunctionBox{X, Y, Z}
}

func CreatePuzzle(contents string, connections int) *Puzzle {
	lines := aoc.ParseLines(contents)
	junctionBoxes := make([]JunctionBox, len(lines))
	for i, line := range lines {
		junctionBoxes[i] = ParseJunctionBox((line))
	}
	links := make([]Link, (len(junctionBoxes)-1)*len(junctionBoxes)/2)
	linkCount := 0
	for i := 0; i < len(junctionBoxes)-1; i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			distance := CalculateDistance(&junctionBoxes[i], &junctionBoxes[j])
			links[linkCount] = Link{distance, i, j}
			linkCount++
		}
	}
	sort.Slice(links, func(i, j int) bool { return links[i].Distance < links[j].Distance })
	return &Puzzle{junctionBoxes, connections, links, make(map[int]int), 0}
}

func CalculateDistance(a, b *JunctionBox) int {
	dx := (a.X - b.X) * (a.X - b.X)
	dy := (a.Y - b.Y) * (a.Y - b.Y)
	dz := (a.Z - b.Z) * (a.Z - b.Z)
	return dx + dy + dz
}

func main() {
	aoc.Local(func(c string) interface{} { return part1(c, 10) }, "part1", "sample.aoc", 40)
	aoc.Local(func(c string) interface{} { return part1(c, 1000) }, "part1", "input.aoc", 112230)
	aoc.Local(func(c string) interface{} { return part2(c, 10) }, "part2", "sample.aoc", 25272)
	aoc.Local(func(c string) interface{} { return part2(c, 10) }, "part2", "input.aoc", 2573952864)
}

func part1(contents string, connections int) interface{} {
	puzzle := CreatePuzzle(contents, connections)
	for i := 0; i < connections; i++ {
		link := puzzle.Links[i]
		circuit1 := puzzle.JunctionCircuits[link.Index1]
		circuit2 := puzzle.JunctionCircuits[link.Index2]
		if circuit1 == 0 && circuit2 == 0 {
			puzzle.CircuitCount++
			puzzle.JunctionCircuits[link.Index1] = puzzle.CircuitCount
			puzzle.JunctionCircuits[link.Index2] = puzzle.CircuitCount
		} else if circuit1 == 0 {
			puzzle.JunctionCircuits[link.Index1] = circuit2
		} else if circuit2 == 0 {
			puzzle.JunctionCircuits[link.Index2] = circuit1
		} else if circuit1 != circuit2 {
			for key, value := range puzzle.JunctionCircuits {
				// replace circuit2 with circuit 1, merging them
				if value == circuit2 {
					puzzle.JunctionCircuits[key] = circuit1
				}
			}
		}
	}
	circuitSizes := make(map[int]int)
	for _, value := range puzzle.JunctionCircuits {
		circuitSizes[value] = circuitSizes[value] + 1
	}
	circuitSizeList := make([]int, 0)
	for _, value := range circuitSizes {
		circuitSizeList = append(circuitSizeList, value)
	}
	sort.Slice(circuitSizeList, func(i, j int) bool {
		return circuitSizeList[i] > circuitSizeList[j]
	})

	return circuitSizeList[0] * circuitSizeList[1] * circuitSizeList[2]
}

func part2(contents string, connections int) interface{} {
	puzzle := CreatePuzzle(contents, connections)
	lastLinkIndex := 0
	for i := 0; i < len(puzzle.Links); i++ {
		link := puzzle.Links[i]
		circuit1 := puzzle.JunctionCircuits[link.Index1]
		circuit2 := puzzle.JunctionCircuits[link.Index2]
		if circuit1 == 0 && circuit2 == 0 {
			puzzle.CircuitCount++
			puzzle.JunctionCircuits[link.Index1] = puzzle.CircuitCount
			puzzle.JunctionCircuits[link.Index2] = puzzle.CircuitCount
			lastLinkIndex = i
		} else if circuit1 == 0 {
			puzzle.JunctionCircuits[link.Index1] = circuit2
			lastLinkIndex = i
		} else if circuit2 == 0 {
			puzzle.JunctionCircuits[link.Index2] = circuit1
			lastLinkIndex = i
		} else if circuit1 != circuit2 {
			for key, value := range puzzle.JunctionCircuits {
				// replace circuit2 with circuit 1, merging them
				if value == circuit2 {
					puzzle.JunctionCircuits[key] = circuit1
				}
			}
			lastLinkIndex = i
		}
	}
	jb1 := puzzle.JunctionBoxes[puzzle.Links[lastLinkIndex].Index1]
	jb2 := puzzle.JunctionBoxes[puzzle.Links[lastLinkIndex].Index2]
	return jb1.X * jb2.X
}
