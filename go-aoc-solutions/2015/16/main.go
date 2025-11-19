package main

import (
	"regexp"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/XX
	// aoc.Local(part1, "part1", "sample.aoc", 0)
	aoc.Local(part1, "part1", "input.aoc", 373)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 0)
}

var detected = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func isMatch(line string) {
	// the things detected
}

func part1(contents string) interface{} {
	// lines like: Sue 1: cars: 9, akitas: 3, goldfish: 0
	reCounts := regexp.MustCompile(`([a-z]+): (\d+)`)
	matchedLines := []int{}
	for i, line := range aoc.ParseLines(contents) {
		allMatches := reCounts.FindAllStringSubmatch(line, -1)
		isMatch := true
		for _, matches := range allMatches {
			for i := 1; i < len(matches); i++ {
				rememberedName := matches[1]
				rememberedCount := aoc.ParseInt(matches[2])
				if detectedCount, exists := detected[rememberedName]; exists {
					if detectedCount != rememberedCount {
						isMatch = false
					}
				} else {
					panic("uh oh")
				}

			}
		}
		if isMatch {
			matchedLines = append(matchedLines, i)
		}
	}
	return matchedLines[0] + 1 // start numbering lines at 0, Sues at 1
}

func part2(contents string) interface{} {
	// In particular, the cats and trees readings indicates that there are
	// greater than that many (due to the unpredictable nuclear decay of cat
	// dander and tree pollen), while the pomeranians and goldfish readings
	// indicate that there are fewer than that many (due to the modial
	// interaction of magnetoreluctance).
	// lines like: Sue 1: cars: 9, akitas: 3, goldfish: 0
	reCounts := regexp.MustCompile(`([a-z]+): (\d+)`)
	matchedLines := []int{}
	for i, line := range aoc.ParseLines(contents) {
		allMatches := reCounts.FindAllStringSubmatch(line, -1)
		isMatch := true
		for _, matches := range allMatches {
			for i := 1; i < len(matches); i++ {
				rememberedName := matches[1]
				rememberedCount := aoc.ParseInt(matches[2])
				if detectedCount, exists := detected[rememberedName]; exists {
					if rememberedName == "cats" || rememberedName == "trees" {
						// should be MORE cats and trees than detected
						if rememberedCount <= detectedCount {
							isMatch = false
						}
					} else if rememberedName == "pomeranians" || rememberedName == "goldfish" {
						// should be FEWER cats and trees than detected
						if rememberedCount >= detectedCount {
							isMatch = false
						}
					} else {

						if detectedCount != rememberedCount {
							isMatch = false
						}
					}
				} else {
					panic("uh oh")
				}

			}
		}
		if isMatch {
			matchedLines = append(matchedLines, i)
		}
	}
	return matchedLines[0] + 1 // start numbering lines at 0, Sues at 1
}
