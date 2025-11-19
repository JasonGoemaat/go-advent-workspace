package main

import (
	"fmt"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/13
	aoc.Local(part1, "part1", "sample.aoc", 330)
	aoc.Local(part1, "part1", "input.aoc", 733)
	aoc.Local(part2, "part2", "sample.aoc", 286) // unverified
	aoc.Local(part2, "part2", "input.aoc", 725)
}

func part1(contents string) interface{} {
	// need to parse this (gain might be lose):
	//		Alice would gain 2 happiness units by sitting next to Bob.
	// probably fastest to replace
	text := strings.ReplaceAll(contents, "would gain ", "")
	text = strings.ReplaceAll(text, "would lose ", "-")
	text = strings.ReplaceAll(text, "happiness units by sitting next to ", "")
	text = strings.ReplaceAll(text, ".", "")
	changes := map[string]map[string]int{}
	for _, line := range aoc.ParseLines(text) {
		parts := strings.Split(line, " ")
		name := parts[0]
		happiness := aoc.ParseInt(parts[1])
		other := parts[2]
		if m, exists := changes[name]; exists {
			m[other] = happiness
		} else {
			m = map[string]int{}
			m[other] = happiness
			changes[name] = m
		}
	}
	names := make([]string, 0, len(changes))
	for k := range changes {
		names = append(names, k)
	}

	calculateHappiness := func() int {
		happiness := 0
		for i, name := range names {
			left := i - 1
			if left < 0 {
				left = len(names) - 1
			}
			right := (i + 1) % len(names)
			happiness += changes[name][names[left]]
			happiness += changes[name][names[right]]
		}
		return happiness
	}

	var recurse func(index int) int
	recurse = func(index int) int {
		// we try all combinations at here to end of names
		if index == len(names)-1 {
			// we're on last name, nothing to swap
			return calculateHappiness()
		}
		maxRecurse := -1000000
		for i := index; i < len(names); i++ {
			if i == index {
				maxRecurse = max(maxRecurse, recurse(index+1))
				continue
			}
			names[index], names[i] = names[i], names[index]
			maxRecurse = max(maxRecurse, recurse(index+1))
		}
		return maxRecurse
	}

	result := recurse(1) // start at 1, move other people around 0
	return result
}

func display(names []string, changes map[string]map[string]int, better bool) {
	total := 0
	for i, name := range names {
		leftIndex := i - 1
		if leftIndex < 0 {
			leftIndex = len(names) - 1
		}
		leftName := names[leftIndex]
		rightIndex := (i + 1) % len(names)
		rightName := names[rightIndex]
		leftValue := changes[name][leftName]
		rightValue := changes[name][rightName]
		fmt.Printf("%3d %s %-3d ", leftValue, name, rightValue)
		total += leftValue
		total += rightValue
	}
	fmt.Printf("= %d", total)
	if better {
		fmt.Printf(" *\n")
	} else {
		fmt.Printf("\n")
	}
}

// part 2 is pretty easy, just add "me" into list and all relationships
// have value 0
func part2(contents string) interface{} {
	// need to parse this (gain might be lose):
	//		Alice would gain 2 happiness units by sitting next to Bob.
	// probably fastest to replace
	text := strings.ReplaceAll(contents, "would gain ", "")
	text = strings.ReplaceAll(text, "would lose ", "-")
	text = strings.ReplaceAll(text, "happiness units by sitting next to ", "")
	text = strings.ReplaceAll(text, ".", "")
	changes := map[string]map[string]int{}
	for _, line := range aoc.ParseLines(text) {
		parts := strings.Split(line, " ")
		name := parts[0]
		happiness := aoc.ParseInt(parts[1])
		other := parts[2]
		if m, exists := changes[name]; exists {
			m[other] = happiness
		} else {
			m = map[string]int{}
			m[other] = happiness
			changes[name] = m
		}
	}
	changes["ME"] = map[string]int{}
	names := make([]string, 0, len(changes))
	for k := range changes {
		names = append(names, k)
	}

	calculateHappiness := func() int {
		happiness := 0
		for i, name := range names {
			left := i - 1
			if left < 0 {
				left = len(names) - 1
			}
			right := (i + 1) % len(names)
			leftName := names[left]
			rightName := names[right]
			leftHappiness := changes[name][leftName]
			rightHappiness := changes[name][rightName]
			totalChange := leftHappiness + rightHappiness
			if totalChange != 0 {
				happiness += (leftHappiness + rightHappiness)
			}
		}
		return happiness
	}

	maxHappiness := -1000000
	var recurse func(index int)
	recurse = func(index int) {
		// we try all combinations at here to end of names
		if index == len(names)-1 {
			// we're on last name, nothing to swap
			happiness := calculateHappiness()
			if happiness > maxHappiness {
				// fmt.Printf("Setting maxHappiness to %d from %d\n", happiness, maxHappiness)
				maxHappiness = happiness
				// display(names, changes, true)
				// fmt.Printf("calculateHappiness: %d\n", calculateHappiness())
				// } else {
				// 	display(names, changes, false)
			}
		}
		for i := index; i < len(names); i++ {
			if i == index {
				recurse(index + 1)
				continue
			}
			t := names[i]
			names[i] = names[index]
			names[index] = t
			// names[index], names[i] = names[i], names[index]
			recurse(index + 1)
		}
	}

	recurse(1) // start at 1, move other people around 0
	return maxHappiness
}
