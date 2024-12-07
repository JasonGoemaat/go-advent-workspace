package main

import (
	"cmp"
	"regexp"
	"slices"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/XX
	aoc.Local(part1, "part1", "sample.aoc", 0)
	// aoc.Local(part1, "part1", "input.aoc", 0)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	// aoc.Local(part2, "part2", "input.aoc", 0)
}

type Node struct{}

type Model struct {
	Reachable map[string][]string
	// Nodes     map[string][]Node
}

func (m *Model) SetReachable(from, to string) {
	list, found := m.Reachable[from]
	if !found {
		list = []string{}
	}
	list = append(list, to)
	m.Reachable[from] = list
}

func part1(contents string) interface{} {
	m := &Model{map[string][]string{}}
	reCodes := regexp.MustCompile(`([a-z][a-z])\-([a-z][a-z])`)
	found := reCodes.FindAllStringSubmatch(contents, -1)
	for _, line := range found {
		m.SetReachable(line[1], line[2])
		m.SetReachable(line[2], line[1])
	}
	nodes := []string{}
	for k, _ := range m.Reachable {
		nodes = append(nodes, k)
	}
	// special sort, 't' first
	slices.SortFunc(nodes, func(a string, b string) int {
		if a[0] == 't' {
			if b[0] == 't' {
				return cmp.Compare(a[1], b[1])
			} else {
				return -1
			}
		}
		if b[0] == 't' {
			return 1
		}
		return cmp.Compare(a, b)
	})
	// for i := 0; i < len(nodes)-2 && nodes[i][0] == 't'; i++ {
	// 	n1 := nodes[i]
	// 	for j := i + 1; j < len(nodes)-1; j++ {
	// 		// n2 := nodes[j]
	// 		// if m.Reachable[n1, n2] {

	// 		// }
	// 	}
	// }
	return 0
}

func part2(contents string) interface{} {
	return 0
}
