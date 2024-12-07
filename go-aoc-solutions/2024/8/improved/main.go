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
	// use new area parser
	area := aoc.ParseArea(contents)

	// sets is a map of non-'.' characters with slice of index (row*Width+col) where they are
	sets := map[byte][]int{}

	// output is a map of index (row*Width+col) to a flag if it is an antinode
	output := map[int]bool{}

	// Find any non-empty locations.  Create a set for each unique non-empty
	// character that contains a list of the indexes where that character
	// occurs.  For example a 10x10 area with 'A' at 0,5 and 3,7 would be:
	// sets['A'] = []int{5, 37}
	for r := range area.Height {
		for c := range area.Width {
			at := area.Get(r, c)
			if at != '.' {
				if sets[at] == nil {
					sets[at] = []int{}
				}
				sets[at] = append(sets[at], area.RowColToIndex(r, c))
			}
		}
	}

	// helper function, place antinode if inside area
	place := func(r, c int) {
		if area.Inside(r, c) {
			output[area.RowColToIndex(r, c)] = true
		}
	}

	// for each unique character, iterate with list of indexes
	for _, v := range sets {
		// loop through all but the last
		for i := 0; i < len(v)-1; i++ {
			// get row,col from index into area of first antenna
			r1, c1 := area.IndexToRowCol(v[i])

			// loop through the remaining indexes from current position so we
			// end up trying every pair
			for j := i + 1; j < len(v); j++ {
				// get row,col of second antenna
				r2, c2 := area.IndexToRowCol(v[j])

				// delta row and delta column are what you add to r1,c1 to get to r2,c2
				dr := r2 - r1
				dc := c2 - c1

				// place antinode backwards from r1,c1 the same distance to r2,c2 (if in area)
				lr, lc := r1-dr, c1-dc
				place(lr, lc)

				// place antinode further from r2,c2 than distance to r1,c1  (if in area)
				// (or backwards from r2,c2 if you think about it that way)
				lr, lc = r2+dr, c2+dc
				place(lr, lc)
			}
		}
	}

	// actually this works
	return len(output) // count keys
}

// part 2 almost the same, but handles multiples of
// distance and includes the nodes themselves
func part2(contents string) interface{} {
	// use new area parser
	area := aoc.ParseArea(contents)

	// sets is a map of non-'.' characters with slice of index (row*Width+col) where they are
	sets := map[byte][]int{}

	// output is a map of index (row*Width+col) to a flag if it is an antinode
	output := map[int]bool{}

	// Find any non-empty locations.  Create a set for each unique non-empty
	// character that contains a list of the indexes where that character
	// occurs.  For example a 10x10 area with 'A' at 0,5 and 3,7 would be:
	// sets['A'] = []int{5, 37}
	for r := range area.Height {
		for c := range area.Width {
			at := area.Get(r, c)
			if at != '.' {
				if sets[at] == nil {
					sets[at] = []int{}
				}
				sets[at] = append(sets[at], area.RowColToIndex(r, c))
			}
		}
	}

	// helper function, place antinode if inside area
	// modified to return true if inside so it can be used in
	// a loop
	place := func(r, c int) bool {
		if area.Inside(r, c) {
			output[area.RowColToIndex(r, c)] = true
			return true
		}
		return false
	}

	// for each unique character, iterate with list of indexes
	for _, v := range sets {
		// loop through all but the last
		for i := 0; i < len(v)-1; i++ {
			// get row,col from index into area of first antenna
			r1, c1 := area.IndexToRowCol(v[i])

			// loop through the remaining indexes from current position so we
			// end up trying every pair
			for j := i + 1; j < len(v); j++ {
				// get row,col of second antenna
				r2, c2 := area.IndexToRowCol(v[j])

				// delta row and delta column are what you add to r1,c1 to get to r2,c2
				dr := r2 - r1
				dc := c2 - c1

				// start at r1,c1 and move forwards in steps of distance
				lr, lc := r1, c1
				for place(lr, lc) {
					lr += dr
					lc += dc
				}

				// start at r2,c2 and move opposite direction
				lr, lc = r2, c2
				for place(lr, lc) {
					lr -= dr
					lc -= dc
				}
			}
		}
	}

	// actually this works
	return len(output) // count keys
}
