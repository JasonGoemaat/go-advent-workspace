package main

import (
	"regexp"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// aoc.Local(part1, "part1", "sample.aoc", 5)
	// aoc.Local(part1, "part1", "input.aoc", 772)
	// aoc.Local(part2, "part2", "sample2.aoc", 2)
	aoc.Local(part2, "part2", "input.aoc", nil)
}

func part1(content string) any {
	lines := aoc.ParseLines(content)
	rxIds := regexp.MustCompile("[a-z]+")
	paths := make(map[string][]string)
	for _, line := range lines {
		ids := rxIds.FindAllString(line, -1)
		paths[ids[0]] = ids[1:]
	}

	total := 0
	var goFrom func(string)
	goFrom = func(id string) {
		if id == "out" {
			total++
			return
		}
		for _, nextId := range paths[id] {
			goFrom(nextId)
		}
	}
	goFrom("you")
	return total
}

func countPathsFrom(paths map[string][]string, from, to string, excludes []string) int {
	return 0
}

func backtrack(backpaths map[string][]string, end, start string, count int) int {
	if end == start {
		return count
	}
	total := 0
	for _, parent := range backpaths[end] {
		total += backtrack(backpaths, parent, start, count)
	}
	return total
}

func part2(content string) any {
	lines := aoc.ParseLines(content)
	rxIds := regexp.MustCompile("[a-z]+")
	paths := make(map[string][]string)
	backpaths := make(map[string][]string)
	for _, line := range lines {
		ids := rxIds.FindAllString(line, -1)
		from := ids[0]
		to := ids[1:]
		paths[from] = to
		for _, child := range to {
			parents, exists := backpaths[child]
			if exists {
				parents = append(parents, from)
			} else {
				parents = []string{from}
			}
			backpaths[child] = parents
		}
	}

	// dacParents := backpaths["dac"]
	// fftParents := backpaths["fft"]
	// fmt.Printf("dac parents: %v\n", dacParents)
	// fmt.Printf("fft parents: %v\n", fftParents)
	dacFirst := backtrack(backpaths, "fft", "dac", 1)
	fftFirst := backtrack(backpaths, "dac", "fft", 1)
	total := 0
	if fftFirst > 0 {
		backToSvr := backtrack(backpaths, "fft", "svr", 1)
		backToDac := backtrack(backpaths, "out", "dac", 1)
		total = backToDac * backToSvr * fftFirst
	} else {
		backToSvr := backtrack(backpaths, "dac", "svr", 1)
		backToDac := backtrack(backpaths, "out", "fft", 1)
		total = backToDac * backToSvr * dacFirst
	}

	return total
}
