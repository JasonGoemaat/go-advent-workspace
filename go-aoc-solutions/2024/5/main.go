package main

import (
	"slices"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/5
	aoc.Local(part1, "part1", "sample.aoc", 143)
	aoc.Local(part1, "part1", "input.aoc", 5108)
	aoc.Local(part2, "part2", "sample.aoc", 123)
	aoc.Local(part2, "part2", "input.aoc", 7380)
}

func part1(content string) interface{} {
	groups := aoc.ParseGroups(content)
	rulesInput := aoc.ParseLinesToInts(aoc.ParseLines(groups[0]))
	updates := aoc.ParseLinesToInts(aoc.ParseLines(groups[1]))
	total := 0

	// rules is map of page number to list of pages that cannot appear
	// before it in an update
	rules := map[int][]int{}
	for _, rule := range rulesInput {
		if rules[rule[0]] == nil {
			rules[rule[0]] = []int{}
		}
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	isOk := func(pages []int) bool {
		seen := map[int]bool{}
		for _, page := range pages {
			seen[page] = true
			for _, after := range rules[page] {
				if seen[after] {
					return false
				}
			}
		}
		return true
	}

	for _, update := range updates {
		if isOk(update) {
			total += update[len(update)/2] // add middle page
		}
	}

	return total
}

func part2(content string) interface{} {
	groups := aoc.ParseGroups(content)
	rulesInput := aoc.ParseLinesToInts(aoc.ParseLines(groups[0]))
	updates := aoc.ParseLinesToInts(aoc.ParseLines(groups[1]))
	total := 0

	// ok, rules is now a map of maps.   The sample '97|13' means that 97
	// must appear before 13, so isBeforeRules[97][13] will be true.  My
	// less(97, 13) function will just return that rule, less(13, 97)
	// will return the default of false since that will either be false
	// and we create an empty map for 13 here in case it doesn't exist
	// in the data
	isBeforeRules := map[int]map[int]bool{}
	for _, rule := range rulesInput {
		before := rule[0]
		after := rule[1]
		if isBeforeRules[before] == nil {
			isBeforeRules[before] = map[int]bool{}
		}
		if isBeforeRules[after] == nil {
			// setting this for that one page that doesn't
			// exist as a before rule so we don't get errors
			isBeforeRules[after] = map[int]bool{}
		}
		isBeforeRules[before][after] = true
	}

	isOk := func(pages []int) bool {
		seen := map[int]bool{}
		for _, page := range pages {
			seen[page] = true
			for k, _ := range isBeforeRules[page] {
				if seen[k] {
					return false
				}
			}
		}
		return true
	}

	sortFunc := func(a, b int) int {
		if isBeforeRules[a][b] {
			return -1
		}
		return 1 // shouldn't ever be equal in our case
	}

	sortUpdate := func(pages []int) []int {
		sorted := slices.Clone(pages)
		slices.SortFunc(sorted, sortFunc)
		return sorted
	}

	for _, update := range updates {
		if !isOk(update) {
			sorted := sortUpdate(update)
			total += sorted[len(sorted)/2] // add middle page number
		}
	}

	return total
}
