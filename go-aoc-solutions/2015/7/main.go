package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/X
	aoc.Local(part1, "part1", "sample.aoc", 507) // note: added line to copy e to 'a' for solution
	aoc.Local(part1, "part1", "input.aoc", 46065)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 14134)
}

// probably should build graph of dependencies...   We have a bunch of
// commands I'll assume can appear in any order like this:
// ab LSHIFT c -> ky
// this means we have two other values 'ab' and 'c' we need to find before finding 'ky'
// A value will only appear on the right side of one of the lines.
// either could be a const number instead of a variable:
// 123 -> ab
// or a special 'NOT instruction that takes only one argument:
// NOT ab -> df
// Operations: LSHIFT, RSHIFT, AND, OR

type node struct {
	name      string // right-side
	operation string // "" for node we haven't discovered on the right yet, "CONST" for numbers, "NOT", "AND", "OR", "LSHIFT", "RSHIFT" for operations
	args      []string
	haveValue bool // after we've calculated
	value     int
}

var nodes = map[string]*node{}

// Take string from left side (i.e. "a", "123", "xy") and return the node if
// it exists, or create a new one if it does not.   For numbers it will be
// a "CONST" type node and start with haveValue = true and value as the 16 bit
// value.  If the node is already added to the nodes map, return it
func nodeFrom(s string) *node {
	if nodes[s] != nil {
		return nodes[s]
	}
	i, err := strconv.Atoi(s)
	var n node
	if err == nil {
		n = node{s, "CONST", []string{}, true, i}
	} else {
		n = node{s, "", []string{}, false, 0}
	}
	nodes[s] = &n
	return &n
}

func nodeValue(name string) int {
	if n, ok := nodes[name]; ok {
		if n.operation == "" {
			return n.value
		}
		return n.calculate()
	} else {
		// not a node, must be an int
		i, err := strconv.Atoi(name)
		if err != nil {
			panic(fmt.Sprintf("Bad int for name '%s'", name))
		}
		return i
	}
}

var depth = 0

func (n *node) calculate() int {
	// fmt.Printf("%d: calculating '%s': '%s' %v\n", depth, n.name, n.operation, n.args)
	depth++
	value := 0
	if n.haveValue {
		value = n.value
	} else if n.operation == "CONST" {
		value = nodes[n.args[0]].calculate()
	} else if n.operation == "NOT" {
		value = nodes[n.args[0]].calculate() ^ 0xffff
	} else if n.operation == "OR" {
		value = nodes[n.args[0]].calculate() | nodes[n.args[1]].calculate()
	} else if n.operation == "AND" {
		value = nodes[n.args[0]].calculate() & nodes[n.args[1]].calculate()
	} else if n.operation == "LSHIFT" {
		value = nodes[n.args[0]].calculate() << nodes[n.args[1]].calculate()
	} else if n.operation == "RSHIFT" {
		value = nodes[n.args[0]].calculate() >> nodes[n.args[1]].calculate()
	} else {
		panic(fmt.Sprintf("INVALID OPERATION: %v", n.operation))
	}
	depth = depth - 1
	// fmt.Printf("%d: '%s' = %d\n", depth, n.name, value)
	n.haveValue = true
	n.value = value
	return value
}

// // Try to convert to number and return "", <number> otherwise
// // return "VAR", -1
// func getValueOrVariable(s string) (string, int) {
// 	value, err := strconv.Atoi(s)
// 	if err == nil {
// 		return "", value
// 	} else {
// 		return s, -1
// 	}
// }

// func constantNode(value int) node {
// 	valueS := string(value)
// 	if n, ok := nodes[valueS]; ok {
// 		return n
// 	} else {
// 		n = node{valueS, "", "", "", value}
// 		nodes[valueS] = n
// 		return n
// 	}
// }

func parseNode(line string) *node {
	parts := strings.Split(line, " -> ")
	if len(parts) < 2 {
		panic("BAD LEN FOR: '" + line + "'")
	}
	name := parts[1]

	parts = strings.Split(parts[0], " ")

	// what will THIS node have for:
	// 123 -> ? - {?, "CONST", []string{"123"}, false, 0 } // '123' node added as it's own CONST node with value known
	// xyz -> ? - {?, "CONST", []string{"xyz"}, false, 0 } // 'xyz' created as "" type, or set to existing
	if len(parts) == 1 {
		// just pointer to another node
		n := node{name, "CONST", []string{nodeFrom(parts[0]).name}, false, 0}
		nodes[name] = &n
		return &n
	}

	// must be NOT
	// NOT 123 -> ? - {?, "NOT", []string{"123"}, false, 0 } // '123' node added as it's own CONST node with value known if it doesn't exist
	// NOT xyz -> ? - {?, "NOT", []string{"xyz"}, false, 0 } // 'xyz' created as "" type if it doesn't exist
	if len(parts) == 2 {
		n := node{name, parts[0], []string{nodeFrom(parts[1]).name}, false, 0}
		nodes[name] = &n
		return &n
	}

	// some operation with two parameters
	// 123 OR y -> ? - {?, "OR", []string{"123","y"}, false, 0 } // '123' and 'x' will be created if they don't exist
	// xyz LSHIFT 2 -> ? - {?, "LSHIFT", []string{"xyz","2"}, false, 0 } // 'xyz' and '2' created if they dont' exist
	n := node{name, parts[1], []string{nodeFrom(parts[0]).name, nodeFrom(parts[2]).name}, false, 0}
	nodes[name] = &n
	return &n
}

func part1(contents string) interface{} {
	lines := aoc.ParseLines(contents)
	for _, line := range lines {
		n := parseNode(line)
		nodes[n.name] = n
	}
	return nodes["a"].calculate()
}

// is this cheesy?  I'm updating the value manually...
func part2(contents string) interface{} {
	lines := aoc.ParseLines(contents)
	for _, line := range lines {
		n := parseNode(line)
		nodes[n.name] = n
	}
	nodes["b"] = &node{"b", "CONST", []string{"46065"}, true, 46065}
	return nodes["a"].calculate()
}
