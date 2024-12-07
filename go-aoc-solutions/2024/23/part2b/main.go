package main

import (
	"fmt"
	"regexp"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/23
	aoc.Local(part2, "part2", "sample.aoc", 0) // 16 nodes
	aoc.Local(part2, "part2", "input.aoc", 0)  // 520 nodes
}

// depth-first-search:
// take in network - map of node to set of connected nodes
// loop through each

type Node struct {
	Name  string
	Links map[string]bool
}

type Graph struct {
	Nodes map[string]Node
}

// remove return new graph with only nodes that are connected
// directly to the one that is passed.
func (g *Graph) WithoutNode(name string) *Graph {
	result := Graph{map[string]Node{}}
	node := g.Nodes[name]

	// create new nodes for any connected to this one
	for k, _ := range node.Links {
		result.Nodes[k] = Node{k, map[string]bool{}}
	}

	// keep all links pointing to nodes in new list (that our removed node point to)
	for name, node := range result.Nodes {
		// now populate Links for remaining nodes, only for nodes
		// connected to the one passed
		newNode := result.Nodes[name]
		for k, _ := range node.Links {
			if node.Links[k] {
				newNode.Links[k] = true
			}
		}
	}
	return &result
}

// Called with how many nodes we've already removed that are definitely
// connected to the remaining nodes
func (g *Graph) MostConnected(included int) {

}

// Given a graph (starting with the original and depth 0), find the maximum
// subset that is all connected to eath other and return the size.
func (g *Graph) GetMaxConnected(depth int) int {
	// for each node in graph,
	return 0
}

// Starting with the base graph, I start one node at a time:
// Take first node:
//		Create copy of graph with only all nodes connected to this node
//

// Return new graph that consists of only nodes connected to the named
// one (NOT including it) and their connections to each other.
func (graph *Graph) GetConnected(name string) *Graph {
	return nil
}

// depth is now many nodes we've removed, that are connected to all remaining
// nodes in the graph
func getLargestConnected(graph *Graph, depth int) int {
	return depth
}

func part2(contents string) interface{} {
	reCodes := regexp.MustCompile(`([a-z][a-z])\-([a-z][a-z])`)
	found := reCodes.FindAllStringSubmatch(contents, -1)
	nodes := map[string]bool{} // set of all nodes
	links := map[string]map[string]bool{}
	for _, link := range found {
		a, b := link[1], link[2]
		nodes[a] = true
		nodes[b] = true
		if links[a] == nil {
			links[a] = make(map[string]bool)
		}
		if links[b] == nil {
			links[b] = make(map[string]bool)
		}
		links[a][b] = true
		links[b][a] = true
	}

	// max links is 13
	maxLength := 0
	for name, _ := range nodes {
		maxLength = max(maxLength, len(links[name]))
	}
	fmt.Printf("%d links, %d nodes, %d maximum links\n", len(found), len(nodes), maxLength)

	// there are only 16 nodes in the sample with

	// use this as a Set where key is sorted list of nodes in triplet
	// resultMap := map[string]bool{}
	// for _, tn := range tNodes {
	// 	local := m.NodeLinks[tn]
	// 	for i := 0; i < len(local)-1; i++ {
	// 		n1 := local[i].To
	// 		for j := i + 1; j < len(local); j++ {
	// 			n2 := local[j].To
	// 			if m.Links[Link{n1, n2}] {
	// 				a, b, c := tn, n1, n2
	// 				if a > b {
	// 					a, b = b, a // [1] is largest of [0] or [1]
	// 				}
	// 				if b > c {
	// 					b, c = c, b // [2] is largest of [1] or [2]
	// 				}
	// 				if a > b {
	// 					a, b = b, a // [1] is largest of [0] or [1]
	// 				}
	// 				resultMap[a+"-"+b+"-"+c] = true
	// 			}
	// 		}
	// 	}
	// }
	// for k, _ := range resultMap {
	// 	fmt.Printf("%s\n", k)
	// }
	return len(nodes)
}
