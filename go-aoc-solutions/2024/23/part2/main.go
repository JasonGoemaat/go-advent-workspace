package main

import (
	"cmp"
	"regexp"
	"slices"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/23
	aoc.Local(part2, "part2", "sample.aoc", "co,de,ka,ta")                           // 16 nodes
	aoc.Local(part2, "part2", "input.aoc", "az,ed,hz,it,ld,nh,pc,td,ty,ux,wc,yg,zz") // 520 nodes
}

type Graph struct {
	Nodes map[string]bool
	Links map[string]map[string]bool
}

var maxList map[string]bool

// find largest connected subroup
func (g *Graph) LargestConnectedSubgroup(depth int, list map[string]bool) int {
	if len(g.Nodes) == 0 {
		// end of connected nodes, if our path is larger than the current largest,
		// copy our current set of connected nodes to the largest
		if len(list) > len(maxList) {
			maxList = map[string]bool{}
			for k, _ := range list {
				maxList[k] = true
			}
		}
		return 0
	}
	maxConnected := 0
	names := make([]string, 0, len(g.Nodes))
	remaining := make(map[string]bool, len(g.Nodes))
	for name, _ := range g.Nodes {
		remaining[name] = true
		names = append(names, name)
	}

	// for depth 0 we start with the 't's, so sort names so 't' are first
	if depth == 0 {
		slices.SortFunc(names, func(a, b string) int {
			if a[0] == 't' {
				return -1
			}
			if b[0] == 't' {
				return 1
			}
			return cmp.Compare(a, b)
		})
	}

	for _, name := range names {
		// as we're setpping through the remaining, remove one we're checking and
		// treat the remaining as a sub-list
		delete(remaining, name)

		// for depth 0, care about names starting with 't' only
		if depth > 0 || name[0] == 't' {
			newLinks := map[string]map[string]bool{}
			newNodes := map[string]bool{}

			// create new graph with the remaining nodes connected to this one
			for k, _ := range g.Links[name] {
				if remaining[k] {
					// node connected to one we're checking and in remaining l ist
					newNodes[k] = true

					// copy links if they're in the remaining list
					newLinks[k] = map[string]bool{}

					for l, _ := range g.Links[k] {
						if remaining[l] {
							// one linked to 'name' and in remaining list, so link will be in child list
							newLinks[k][l] = true
						}
					}
				}
			}

			childGraph := Graph{Nodes: newNodes, Links: newLinks}
			list[name] = true // add to our list
			result := childGraph.LargestConnectedSubgroup(depth+1, list)
			if result > maxConnected {
				maxConnected = result
			}
			delete(list, name) // remove from list, we're moving on
		}
	}
	return maxConnected + 1 // +1 for the current node
}

func part2(contents string) interface{} {
	reCodes := regexp.MustCompile(`([a-z][a-z])\-([a-z][a-z])`)
	found := reCodes.FindAllStringSubmatch(contents, -1)
	nodes := map[string]bool{}            // set of all nodes
	links := map[string]map[string]bool{} // all links, both directions are specified
	for _, link := range found {
		a, b := link[1], link[2]
		nodes[a] = true
		nodes[b] = true
		if _, exists := links[a]; !exists {
			links[a] = map[string]bool{}
		}
		links[a][b] = true
		if _, exists := links[b]; !exists {
			links[b] = map[string]bool{}
		}
		links[b][a] = true
	}

	graph := Graph{Nodes: nodes, Links: links}
	graph.LargestConnectedSubgroup(0, map[string]bool{})
	// fmt.Printf("largest: %d, maxList has %d\n", largest, len(maxList))

	computers := make([]string, 0, len(maxList))
	for k, _ := range maxList {
		computers = append(computers, k)
	}
	slices.Sort(computers)
	return strings.Join(computers, ",")
}
