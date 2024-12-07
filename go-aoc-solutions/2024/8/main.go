package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/8
	aoc.Local(part1, "part1", "sample.aoc", 14)
	aoc.Local(part1, "part1", "input.aoc", 396)
	aoc.Local(part2, "part2", "sample.aoc", 34)
	aoc.Local(part2, "part2", "input.aoc", 1200)
}

//
// Calculate the impact of the signal. How many unique locations within the
// bounds of the map contain an antinode?
//

func part1(contents string) interface{} {
	var lines = aoc.ParseLines(contents)
	colCount := len(lines)
	rowCount := len(lines[0])
	sets := map[byte][]int{}
	output := map[int]bool{}

	pos := func(r, c int) int {
		return r*colCount + c
	}

	loc := func(p int) (int, int) {
		return p / colCount, p % colCount
	}

	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			at := lines[r][c]
			if at != '.' {
				p := pos(r, c)
				if sets[at] == nil {
					sets[at] = []int{}
				}
				sets[at] = append(sets[at], p)
			}
		}
	}

	place := func(ch byte, r, c int) {
		if r < 0 || r >= rowCount || c < 0 || c >= colCount {
			return
		}
		p := pos(r, c)
		output[p] = true
	}

	for k, v := range sets {
		for i := 0; i < len(v)-1; i++ {
			r1, c1 := loc(v[i])
			for j := i + 1; j < len(v); j++ {
				r2, c2 := loc(v[j])
				// 5,5 to 7,7
				// dr1,dc1 are 2,2
				// sub from 1, add to 2
				dr := r2 - r1
				dc := c2 - c1
				lr, lc := r1-dr, c1-dc
				place(k, lr, lc)
				lr, lc = r2+dr, c2+dc
				place(k, lr, lc)
			}
		}
	}

	total := 0
	for _, _ = range output {
		total++
	}

	return total
}

// part 2 almost the same, but handles multiples of
// distance and includes the nodes themselves
func part2(contents string) interface{} {
	var lines = aoc.ParseLines(contents)
	colCount := len(lines)
	rowCount := len(lines[0])
	sets := map[byte][]int{}
	output := map[int]bool{}

	pos := func(r, c int) int {
		return r*colCount + c
	}

	loc := func(p int) (int, int) {
		return p / colCount, p % colCount
	}

	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			at := lines[r][c]
			if at != '.' {
				p := pos(r, c)
				if sets[at] == nil {
					sets[at] = []int{}
				}
				sets[at] = append(sets[at], p)
			}
		}
	}

	place := func(ch byte, r, c int) {
		if r < 0 || r >= rowCount || c < 0 || c >= colCount {
			return
		}
		p := pos(r, c)
		output[p] = true
	}

	for k, v := range sets {
		for i := 0; i < len(v)-1; i++ {
			r1, c1 := loc(v[i])
			for j := i + 1; j < len(v); j++ {
				r2, c2 := loc(v[j])

				// we found a pair, so the two are definitely included
				place(k, r1, c1)
				place(k, r2, c2)

				// 5,5 to 7,7
				// dr1,dc1 are 2,2
				// sub from 1, add to 2
				dr := r2 - r1
				dc := c2 - c1
				lr, lc := r1-dr, c1-dc
				for lr >= 0 && lr < rowCount && lc >= 0 && lc <= colCount {
					place(k, lr, lc)
					lr -= dr
					lc -= dc
				}
				lr, lc = r2+dr, c2+dc
				for lr >= 0 && lr < rowCount && lc >= 0 && lc <= colCount {
					place(k, lr, lc)
					lr += dr
					lc += dc
				}
			}
		}
	}

	total := 0
	for _, _ = range output {
		total++
	}

	return total
}
