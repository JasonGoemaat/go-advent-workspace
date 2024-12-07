package main

import (
	"strconv"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/11
	aoc.Local(part1, "part1", "sample.aoc", 55312)
	aoc.Local(part1, "part1", "input.aoc", 216996)
	aoc.Local(part2, "part2", "sample.aoc", 65601038650482)
	aoc.Local(part2, "part2", "input.aoc", 257335372288947)
}

// RULES:
// 0 -> 1
// Even Digit Count -> split to left and right
// Odd Digit Count -> multiply by 2024

type TreeNode struct {
	Left, Right int // -1 for none
}

type Tree map[int]*TreeNode

func (tree Tree) GetTreeNode(value int) *TreeNode {
	if value < 0 {
		return nil
	}
	if tree[value] != nil {
		return tree[value]
	}
	left, right := -1, -1

	// HERE are the rules from the puzzle
	if value == 0 {
		left = 1
	} else {
		// maybe more efficient to count digits another way
		s := strconv.Itoa(value)
		if (len(s) & 1) == 0 {
			// even digits, split
			left, _ = strconv.Atoi(s[:len(s)/2])
			right, _ = strconv.Atoi(s[len(s)/2:])
		} else {
			// odd digits, *2024
			left = value * 2024
		}
	}

	newTreeNode := TreeNode{left, right}
	tree[value] = &newTreeNode
	return &newTreeNode
}

func part1(contents string) interface{} {
	return part1Depth(contents, 25)
}

var usingDepth = 0

func spaces(count int) string {
	if count < 1 {
		return ""
	}
	return strings.Repeat(" ", count)
}

func part1Depth(contents string, depth int) interface{} {
	usingDepth = depth // for pretty printing
	nums := aoc.ParseLinesToInts([]string{contents})[0]
	tree := Tree{}

	var totalAtDepth func(value, depth int) int
	totalAtDepth = func(value, depth int) int {
		if value < 0 {
			return 0
		}
		if depth == 0 {
			// fmt.Printf("%s%d@depth(%d): returning 1\n", spaces(usingDepth-depth), value, depth)
			return 1
		}

		node := tree.GetTreeNode(value)
		leftTotal := totalAtDepth(node.Left, depth-1)
		rightTotal := totalAtDepth(node.Right, depth-1)

		// fmt.Printf("%s%d@depth(%d): (%d, %d) = (%d, %d) total %d\n", spaces(usingDepth-depth), value, depth, node.Left, node.Right, leftTotal, rightTotal, leftTotal+rightTotal)

		result := leftTotal + rightTotal
		return result
	}

	total := 0
	for _, num := range nums {
		_ = tree.GetTreeNode(num)
		total += totalAtDepth(num, depth)
	}

	// report total count of keys
	// fmt.Printf("Depth %d created %d keys\n", depth, len(tree))
	return total
}

func part2(contents string) interface{} {
	return part2Depth(contents, 75) // should be 75
}

func part2Depth(contents string, depth int) interface{} {
	nums := aoc.ParseLinesToInts([]string{contents})[0]
	current := map[int]int{}
	for _, num := range nums {
		current[num] = 1
	}

	tree := Tree{} // not used as a tree here, but to calculate and cache left/right for numbers

	var addToNew = func(m map[int]int, value, count int) {
		if value >= 0 {
			if c, ok := m[value]; ok {
				m[value] = c + count
			} else {
				m[value] = count
			}
		}
	}

	for blinks := 0; blinks < depth; blinks++ {
		// fmt.Printf("After %d blinks, %d unique values, %d keys in tree\n", blinks, len(current), len(tree))
		new := map[int]int{}
		for num, count := range current {
			node := tree.GetTreeNode(num)
			addToNew(new, node.Left, count)
			addToNew(new, node.Right, count)
		}
		current = new
	}

	total := 0
	for _, count := range current {
		total += count
	}
	return total
}
