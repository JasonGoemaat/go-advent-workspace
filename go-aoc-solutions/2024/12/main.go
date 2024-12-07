package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/12
	aoc.Local(part1, "part1", "sample.aoc", 772)
	aoc.Local(part1, "part1", "sample2.aoc", 1930)
	aoc.Local(part1, "part1", "input.aoc", 1396562)
	aoc.Local(part2, "part2", "sample.aoc", 436)
	aoc.Local(part2, "part2", "sample2.aoc", 1206)
	aoc.Local(part2, "part2", "sample3.aoc", 80)
	aoc.Local(part2, "part2", "sample4.aoc", 236)
	aoc.Local(part2, "part2", "sample5.aoc", 368)
	aoc.Local(part2, "part2", "input.aoc", 844132)
}

// Puzzle doesn't seem too difficult.   I think we can process in one sweep.
// For terminology, a 'plot' is an individual square.   A 'region' is a
// connected set of plots that are touching.

// So I start with an int array the same size as Area.Data representing the
// region number.   Array is initialized to 0, so regions start numbering
// at 1.

// I have two options.   A single pass solution could simply test left and up
// and if it's the same character, use that for the region.   And calculate
// the number of fences (look up,down,left,right for different characters
// or the edge of the map) and add that to a map/array for that region.
// Also if the character above or left is the same but a different region,
// I can merge the regions or note that they should be merged using
// a map.

// The other way would be when I found a new region to map it out completely,
// that may be easier.

var (
	left       = []int{0, -1}
	top        = []int{-1, 0}
	right      = []int{0, 1}
	bottom     = []int{1, 0}
	directions = [][]int{left, top, right, bottom}
)

func scanRegion(area *aoc.Area, ids []int, index int, count int) (int, int) {
	ids[index] = count
	r, c := area.IndexToRowCol(index)
	totalArea := 1      // for me
	totalPerimeter := 0 // not known yet
	for _, dir := range directions {
		r2, c2 := r+dir[0], c+dir[1]
		if area.Inside(r2, c2) && (area.Get(r2, c2) == area.Data[index]) {
			if ids[area.RowColToIndex(r2, c2)] == 0 {
				newArea, newPerimeter := scanRegion(area, ids, area.RowColToIndex(r2, c2), count)
				totalArea += newArea
				totalPerimeter += newPerimeter
			}
		} else {
			totalPerimeter += 1
		}
	}
	return totalArea, totalPerimeter
}

func part1(contents string) interface{} {
	area := aoc.ParseArea(contents) // bug here?
	regionIds := make([]int, len(area.Data))
	count := 0
	totalArea := 0
	totalPerimeter := 0
	totalResult := 0
	for index, _ := range area.Data {
		if regionIds[index] == 0 {
			count++ // start at 1
			na, np := scanRegion(area, regionIds, index, count)
			totalArea += na
			totalPerimeter += np
			totalResult += (na * np) // for puzzle
		}
	}
	return totalResult
}

// a TWIST - basically we count 'sides' for perimeter instead of
// counting each square.   What that means is if we have AA that
// is 4 instead of 6 because the As share the top and bottom.
// Using my current solution I don't think that's very easy
// without storing more information.  How could I store that and
// keep my same recursive structure?    Maybe it would be
// better to store a pointer to the area and calculate that
// later.   That could be done off the int array created for
// the first part.   I could keep my current recursive function
// because it returns perimeter and area separately and throw
// away the perimeter value.   Sounds good, let's create
// recalculateParameter(index)

func recalculateParameter(area *aoc.Area, ids []int, index int) int {
	// I think I can do this in row/col fashion.  A top perimeter exists when
	// we are in our region and two things are true:
	//		1. The plot above is not in the region
	//		2. The plot left is not in the region OR the plot up-left is not in the region
	//	AAA
	//  BAB
	//  BBB
	// here when scanning the B region we count the two top Bs because:
	//		The plot above is not in the region
	//		The plot left is not in the region
	// If the middle 'A' were a B we wouldn't count it or the right B
	// However if the middle 'A' AND the 'A' above it were 'B', we would
	//		need to still count the second 'B' because the top-left is also
	//		in the region
	// Maybe pre-processing would be useful, or two-step.  First flag all the
	// positions where there is a perimeter to the top.   Then only count a 'new'
	// perimeter if the one to the left doesn't also have it flagged.   Yeah...
	// insideUp := false
	// insideDown := false
	// counter := 0
	// for r := range area.Height {
	// 	inside = false
	// 	for c := range area.Width {
	// 		pos := area.RowColToIndex(r, c)
	// 		if ids[pos] != index {
	// 			insideUp := false
	// 			insideDown := false
	// 		} else {
	// 			if (!insideUp) {

	// 			}
	// 		}
	// 		if inside {
	// 			if !area.Is(r-1, c, b) {
	// 				inside := false
	// 			}
	// 		} else {}
	// 	}
	// }
	return 0
}

type puzzle struct {
	area *aoc.Area
	ids  []int
}

func (p puzzle) leftOf(pos int) int {
	_, c := p.area.IndexToRowCol(pos)
	if c <= 0 {
		return -1
	}
	return pos - 1
}

func (p puzzle) rightOf(pos int) int {
	_, c := p.area.IndexToRowCol(pos)
	if (c + 1) >= p.area.Width {
		return -1
	}
	return pos + 1
}

func (p puzzle) topOf(pos int) int {
	r, _ := p.area.IndexToRowCol(pos)
	if r <= 0 {
		return -1
	}
	return pos - p.area.Width
}

func (p puzzle) bottomOf(pos int) int {
	r, _ := p.area.IndexToRowCol(pos)
	if (r + 1) >= p.area.Height {
		return -1
	}
	return pos + p.area.Width
}

func scanRegion2(p puzzle, index int, count int) (int, int) {
	p.ids[index] = count
	r, c := p.area.IndexToRowCol(index)
	totalArea := 1      // for me
	totalPerimeter := 0 // not known yet
	for dirIndex, dir := range directions {
		r2, c2 := r+dir[0], c+dir[1]
		if p.area.Inside(r2, c2) && (p.area.Get(r2, c2) == p.area.Data[index]) {
			if p.ids[p.area.RowColToIndex(r2, c2)] == 0 {
				newArea, newPerimeter := scanRegion2(p, p.area.RowColToIndex(r2, c2), count)
				totalArea += newArea
				totalPerimeter += newPerimeter
			}
		} else {
			// the one to this direction (dirIndex, dir) at r2,c2 is different than the one
			// we're looking at (r, c)
			//
			// for part1 we count a perimeter every time this happens
			// for part2 we make sure for horizontal perimeters that we have not
			//		already counted the top perimeter on the block to the left,
			//		which we can verify in two ways:
			//			1. left is different (no way it can be a shared perimeter)
			//			2. left is same, but left and up/bottom is different
			if (dirIndex == 1) || (dirIndex == 3) { // top or bottom
				leftIndex := p.leftOf(index)
				if (leftIndex < 0) || (p.area.Data[leftIndex] != p.area.Data[index]) {
					// left isn't part of our region, so this is the start of a new perimeter
					totalPerimeter += 1
				} else {
					// left is part of our region, this is only the start of a new perimeter if
					// the plot above or below is also part of the region
					otherIndex := -1
					if dirIndex == 1 {
						otherIndex = p.topOf(leftIndex)
					} else {
						otherIndex = p.bottomOf(leftIndex)
					}
					if (otherIndex != -1) && (p.area.Data[otherIndex] == p.area.Data[index]) {
						totalPerimeter += 1
					}
				}
			} else { // left or right
				topIndex := p.topOf(index)
				if (topIndex < 0) || (p.area.Data[topIndex] != p.area.Data[index]) {
					// top isn't part of our region, so this is the start of a new perimeter
					totalPerimeter += 1
				} else {
					// top is part of our region, this is only the start of a new perimeter if
					// the plot left/right of that top plot is also part of the region
					otherIndex := -1
					if dirIndex == 0 { // 0 or left
						otherIndex = p.leftOf(topIndex)
					} else { // must be 2 or right
						otherIndex = p.rightOf(topIndex)
					}
					if (otherIndex != -1) && (p.area.Data[otherIndex] == p.area.Data[index]) {
						totalPerimeter += 1
					}
				}
			}
			// totalPerimeter += 1
		}
	}
	return totalArea, totalPerimeter
}

func part2(contents string) interface{} {
	area := aoc.ParseArea(contents) // bug here?
	ids := make([]int, len(area.Data))
	p := puzzle{area, ids}
	count := 0
	totalArea := 0
	totalPerimeter := 0
	totalResult := 0
	perimeters := map[int]int{}
	areas := map[int]int{}
	for index, _ := range area.Data {
		if ids[index] == 0 {
			count++ // start at 1
			na, np := scanRegion2(p, index, count)
			totalArea += na
			totalPerimeter += np
			perimeters[count] = np
			areas[count] = na
			totalResult += (na * np) // for puzzle
		}
	}
	return totalResult
}
