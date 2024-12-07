package main

import (
	"regexp"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/XX
	aoc.Local(part1, "part1", "sample.aoc", 7)
	aoc.Local(part1, "part1", "input.aoc", 0)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	// aoc.Local(part2, "part2", "input.aoc", 0)
}

type Link struct {
	From string
	To   string
}

type Model struct {
	Links     map[Link]bool
	Nodes     map[string]bool
	NodeLinks map[string][]Link
	Reachable map[string]map[string]bool
}

func (m *Model) AddLink(n1, n2 string) {
	l := Link{n1, n2}
	m.Links[l] = true

	m.Nodes[n1] = true

	list, found := m.NodeLinks[n1]
	if !found {
		list = []Link{}
	}
	list = append(list, l)
	m.NodeLinks[n1] = list

	r, found := m.Reachable[n1]
	if !found {
		r = map[string]bool{}
	}
	r[n2] = true
	m.Reachable[n1] = r
}

func part1(contents string) interface{} {
	m := &Model{
		Links:     map[Link]bool{},
		Nodes:     map[string]bool{},
		NodeLinks: map[string][]Link{},
		Reachable: map[string]map[string]bool{},
	}
	reCodes := regexp.MustCompile(`([a-z][a-z])\-([a-z][a-z])`)
	found := reCodes.FindAllStringSubmatch(contents, -1)
	for _, line := range found {
		m.AddLink(line[1], line[2])
		m.AddLink(line[2], line[1])
	}
	tNodes := []string{}
	for k, _ := range m.Nodes {
		if k[0] == 't' {
			tNodes = append(tNodes, k)
		}
	}

	// use this as a Set where key is sorted list of nodes in triplet
	resultMap := map[string]bool{}
	for _, tn := range tNodes {
		local := m.NodeLinks[tn]
		for i := 0; i < len(local)-1; i++ {
			n1 := local[i].To
			for j := i + 1; j < len(local); j++ {
				n2 := local[j].To
				if m.Links[Link{n1, n2}] {
					a, b, c := tn, n1, n2
					if a > b {
						a, b = b, a // [1] is largest of [0] or [1]
					}
					if b > c {
						b, c = c, b // [2] is largest of [1] or [2]
					}
					if a > b {
						a, b = b, a // [1] is largest of [0] or [1]
					}
					resultMap[a+"-"+b+"-"+c] = true
				}
			}
		}
	}
	// for k, _ := range resultMap {
	// 	fmt.Printf("%s\n", k)
	// }
	return len(resultMap)
}

func part2(contents string) interface{} {
	return 0
}
