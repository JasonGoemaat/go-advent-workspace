package main

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

var (
	reColors = regexp.MustCompile("[wubrg]+")
)

func main() {
	// https://adventofcode.com/2015/day/X
	// aoc.Local(part1, "part1", "sample.aoc", 6)
	aoc.Local(part1, "part1", "input.aoc", 0)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	// aoc.Local(part2, "part2", "input.aoc", 0)
}

var allowDupes = true
var totalLength = 0

func can(pattern string, towels []string) bool {
	lp := len(pattern)
	if lp == 0 {
		return true // we've covered the entire pattern
	}

	// refilter our list of towels
	matchingTowels := make([]string, 0, len(towels))
	matchingIndex := make([]int, 0, len(towels))
	for _, towel := range towels {
		index := strings.Index(pattern, towel)
		if index >= 0 {
			matchingTowels = append(matchingTowels, towel)
			matchingIndex = append(matchingIndex, index)
		}
	}
	for i, towel := range matchingTowels {
		// remove towel from possible array
		var newTowels = matchingTowels
		if !allowDupes {
			lmt := len(matchingTowels)
			lnt := len(matchingTowels) - 1
			newTowels = make([]string, lnt)
			copy(newTowels, matchingTowels[0:i])
			copy(newTowels[i:lnt], matchingTowels[i+1:lmt])
		}

		// remove string
		newPattern := pattern[0:matchingIndex[i]] + pattern[matchingIndex[i]+len(towel):len(pattern)]
		if can(newPattern, newTowels) {
			return true
		}
	}
	return false
}

func part1(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	towels := reColors.FindAllString(groups[0], -1)
	// sort towels with largest first
	slices.SortFunc(towels, func(a, b string) int {
		return cmp.Compare(len(b), len(a))
	})
	reTowels := make([]regexp.Regexp, len(towels))
	for i, towel := range towels {
		reTowels[i] = *regexp.MustCompile(towel)
	}
	patterns := aoc.ParseLines(groups[1])
	goodPatterns := 0

	tm := map[string]int{}
	for _, t := range towels {
		tm[t]++
	}
	// fmt.Printf("%d towels, %d unique\n", len(towels), len(tm)) // 447 unique

	for _, pattern := range patterns {
		totalLength = len(pattern)
		if can(pattern, towels) {
			goodPatterns++
			fmt.Printf("'%s': YES!\n", pattern)
		} else {
			fmt.Printf("'%s': NO!\n", pattern)
		}
	}
	return goodPatterns
}

func part1Info(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	towels := reColors.FindAllString(groups[0], -1)
	reTowels := make([]regexp.Regexp, len(towels))
	for i, towel := range towels {
		reTowels[i] = *regexp.MustCompile(towel)
	}
	patterns := aoc.ParseLines(groups[1])
	largestTowel := 0
	for _, t := range towels {
		largestTowel = max(largestTowel, len(t))
	}

	for i, pattern := range patterns {
		towelCount := 0
		matchCount := 0
		for _, re := range reTowels {
			matches := re.FindAllStringIndex(pattern, -1)
			if len(matches) > 0 {
				towelCount++
				matchCount += len(matches)
			}
		}
		fmt.Printf("%03d: %d towels, %d matches\n", i, towelCount, matchCount)
	}
	return len(towels) + len(patterns)
}

func part2(contents string) interface{} {
	return 0
}
