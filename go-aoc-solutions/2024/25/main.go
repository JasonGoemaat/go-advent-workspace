package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/25
	aoc.Local(part1, "part1", "sample.aoc", 3)
	aoc.Local(part1, "part1", "input.aoc", 2824)
	aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 0)
}

type Entry struct {
	EntryType string
	Values    [5]byte
}

func ParseEntry(entryContent string) *Entry {
	lines := aoc.ParseLines(entryContent)
	heights := [5]byte{0, 0, 0, 0, 0}
	for _, line := range lines {
		for j := range len(line) {
			if line[j] == '#' {
				heights[j]++
			}
		}
	}

	if lines[0] == "#####" {
		return &Entry{"lock", heights}
	}
	return &Entry{"key", heights}
}

func (entry *Entry) Fits(other *Entry) bool {
	if entry.EntryType == other.EntryType {
		return false
	}
	for i, b := range entry.Values {
		if (other.Values[i] + b) > 7 {
			return false
		}
	}
	return true
}

func part1(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	entries := make([]*Entry, len(groups))
	for i, content := range groups {
		entries[i] = ParseEntry(content)
	}

	fitCount := 0
	for i := range len(entries) - 1 {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].Fits(entries[j]) {
				fitCount++
			}
		}
	}
	return fitCount
}

func part2(contents string) interface{} {
	return 0
}
