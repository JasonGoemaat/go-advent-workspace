package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/X
	aoc.Local(part1, "part1", "sample.aoc", 605)
	aoc.Local(part1, "part1", "input.aoc", 251)
	aoc.Local(part2, "part2", "sample.aoc", 982)
	aoc.Local(part2, "part2", "input.aoc", 898)
}

type edge struct {
	from     string
	to       string
	distance int
}

type state struct {
	visited  map[string]bool
	edges    []edge
	distance int
	current  string
}

func parseLine(line string) (node1, node2 string, distance int) {
	a1 := strings.Split(line, " = ")
	aLeft := strings.Split(a1[0], " to ")
	distance, err := strconv.Atoi(a1[1])
	if err != nil {
		panic(fmt.Sprintf("Invalid distance '%s': %s", a1[1], line))
	}
	return aLeft[0], aLeft[1], distance
}

func parseState(contents string) state {
	lines := aoc.ParseLines(contents)
	edges := []edge{}
	visited := make(map[string]bool, len(lines)/2)
	for _, line := range lines {
		node1, node2, distance := parseLine(line)
		visited[node1] = false
		visited[node2] = false
		edges = append(edges, edge{node1, node2, distance})
		edges = append(edges, edge{node2, node1, distance})
	}
	return state{visited, edges, 0, ""}
}

func getDistance(st *state, ed edge, useMin bool) int {
	// if we've already visited, that shouldn't happen
	if st.visited[ed.to] {
		panic("ALREADY VISISTED!")
	}

	// new distance adds edge distance
	distance := st.distance + ed.distance

	// fmt.Printf("%d -> %d: %v\n", st.distance, distance, ed)

	// new location is to of edge
	location := ed.to

	// copy visited map and add the new location
	newVisited := map[string]bool{}
	for k, v := range st.visited {
		newVisited[k] = v
	}
	newVisited[location] = true

	// filter out remaining edges that are from or to our location
	newEdges := []edge{}
	currentPaths := []edge{}
	for _, e := range st.edges {
		if e.from == location {
			currentPaths = append(currentPaths, e)
		} else if e.to != location {
			newEdges = append(newEdges, e)
		}
	}

	// quit when we have no edges left and report distance
	if len(currentPaths) == 0 {
		return distance
	}

	// create new state from our data
	stNew := state{newVisited, newEdges, distance, location}

	// try each edge where current location is the from
	minDistance := 0xffffffff
	if !useMin {
		minDistance = 0
	}

	for _, e := range currentPaths {
		d := getDistance(&stNew, e, useMin)
		if useMin {
			if d < minDistance {
				minDistance = d
			}
		} else {
			if d > minDistance {
				minDistance = d
			}
		}
	}

	return minDistance
}

func part1(contents string) interface{} {
	st := parseState(contents)
	minDistance := 0x7fffffff
	// try each starting location
	for k, _ := range st.visited {
		distance := getDistance(&st, edge{"", k, 0}, true)
		if distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func part2(contents string) interface{} {
	st := parseState(contents)
	maxDistance := 0
	// try each starting location
	for k, _ := range st.visited {
		distance := getDistance(&st, edge{"", k, 0}, false)
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance
}
