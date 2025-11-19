package main

import (
	"regexp"
	"strings"

	year2015day19part2 "github.com/JasonGoemaat/go-aoc-solutions/2015/19/part2"
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/19
	// aoc.Local(part1, "part1", "sample.aoc", 7)
	// aoc.Local(part1, "part1", "input.aoc", 518)
	aoc.Local(year2015day19part2.Part2, "part2", "sample.aoc", 6)
	aoc.Local(year2015day19part2.Part2, "part2", "input.aoc", 0)
}

type Replacement struct {
	Original, Replaced string
}

func part1(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	replacementLines := aoc.ParseLines(groups[0])
	replacements := make([]Replacement, len(replacementLines))
	for i, line := range replacementLines {
		parts := strings.Split(line, " => ")
		replacements[i] = Replacement{parts[0], parts[1]}
	}
	results := map[string]int{} // keys will be unique, value will be ways to form it doing 1 replacement
	stringToUse := groups[1]
	for _, replacement := range replacements {
		original := replacement.Original
		replaced := replacement.Replaced
		reReplacement := regexp.MustCompile(original)
		all := reReplacement.FindAllStringIndex(stringToUse, -1)
		if all != nil {
			for _, values := range all {
				newString := stringToUse[0:values[0]] + replaced + stringToUse[values[1]:]
				results[newString]++
			}
		}
	}

	return len(results)
}

// hmmm....   part2 is finding the fewest steps to get from 'e' to the final
// molecule.  I think I can do this by going in reverse.  This time I can just
// reverse the directions.  I can do this in steps maybe?
func part2(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	replacementLines := aoc.ParseLines(groups[0])
	replacements := make([]Replacement, len(replacementLines))
	for i, line := range replacementLines {
		parts := strings.Split(line, " => ")
		replacements[i] = Replacement{parts[1], parts[0]} // reversed
	}
	results := map[string]int{} // keys will be unique, value will be minimum number of steps to get there
	stringToUse := groups[1]
	for _, replacement := range replacements {
		original := replacement.Original
		replaced := replacement.Replaced
		reReplacement := regexp.MustCompile(original)
		all := reReplacement.FindAllStringIndex(stringToUse, -1)
		if all != nil {
			for _, values := range all {
				newString := stringToUse[0:values[0]] + replaced + stringToUse[values[1]:]
				results[newString]++
			}
		}
	}

	return len(results)
}
