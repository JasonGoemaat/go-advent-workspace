package main

import (
	"sort"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 3)
	aoc.Local(part1, "part1", "input.aoc", 567)
	aoc.Local(part2, "part2", "sample.aoc", 14)
	aoc.Local(part2, "part2", "input.aoc", 354149806372909)
}

func part1(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	freshRanges := aoc.ParseIntRanges(groups[0])
	ids := aoc.ParseInts(groups[1])
	count := 0
	for _, id := range ids {
		for _, r := range freshRanges {
			if id >= r.Start && id <= r.End {
				count++
				break
			}
		}
	}
	return count
}

func mergeRanges(ranges []aoc.IntRange) []aoc.IntRange {
	count := len(ranges)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	for i := 1; i < count; {
		if ranges[i].Start <= ranges[i-1].End {
			// overlapping ranges
			if ranges[i-1].End < ranges[i].End {
				// current range would expand previous ranges
				ranges[i-1].End = ranges[i].End
			}
			count--
			for j := i; j < count; j++ {
				ranges[j] = ranges[j+1]
			}
		} else {
			i++
		}
	}
	return ranges[0:count]
}

func part2(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	freshRanges := aoc.ParseIntRanges(groups[0])
	results := mergeRanges(freshRanges)
	for len(freshRanges) > len(results) {
		freshRanges = results
		results = mergeRanges(freshRanges)
	}
	totalIds := 0
	for _, r := range results {
		totalIds += (r.End - r.Start + 1)
	}
	return totalIds
}
