package main

import (
	"fmt"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/X
	aoc.Local(part1, "part1", "sample.aoc", 480)
	aoc.Local(part1, "part1", "input.aoc", 31897)
	aoc.Local(part2, "part2", "sample.aoc", 875318608908)
	aoc.Local(part2, "part2", "input.aoc", 87596249540359)
}

func solveEasy(content string, sa, sb int) (int, int, int) {
	numsList := aoc.ParseLinesToInts(aoc.ParseLines(content))
	// no more than 100 button presses, let's brute-force it
	ax, ay := numsList[0][0], numsList[0][1]
	bx, by := numsList[1][0], numsList[1][1]
	goalX, goalY := numsList[2][0], numsList[2][1]
	lowestCost := 0
	aPresses, bPresses := 0, 0
	for aCount := sa; aCount < (sa + 100); aCount++ {
		x1, y1 := ax*aCount, ay*aCount
		if (x1 > goalX) || (y1 > goalY) {
			break
		}
		for bCount := sb; bCount < (sb + 100); bCount++ {
			x2, y2 := bx*bCount, by*bCount
			if ((x1 + x2) >= goalX) && ((y1 + y2) >= goalY) {
				if ((x1 + x2) == goalX) && ((y1 + y2) == goalY) {
					cost := (aCount * 3) + (bCount * 1)
					if (lowestCost == 0) || (cost < lowestCost) {
						aPresses = aCount
						bPresses = bCount
						lowestCost = cost
					}
				}
				break
			}
		}
	}
	return lowestCost, aPresses, bPresses
}

// WAIT 2, let's solve with  math
func solveMath(content string, extraFlag bool) (int, int, int) {
	numsList := aoc.ParseLinesToInts(aoc.ParseLines(content))
	ax, ay := numsList[0][0], numsList[0][1]
	bx, by := numsList[1][0], numsList[1][1]
	extra := 0
	if extraFlag {
		extra = 10000000000000
	}
	goalX, goalY := numsList[2][0]+extra, numsList[2][1]+extra

	// SPECIAL CASE: test for exact a and b presses
	// {xpresses, ypresses, xremainder, yremainder}
	aPresses := []int{goalX / ax, goalY / ay, goalX % ax, goalY % ay}
	bPresses := []int{goalX / bx, goalY / by, goalX % bx, goalY % by}
	if (bPresses[2] == 0) && (bPresses[3] == 0) && (bPresses[0] == bPresses[1]) {
		// exact match on b
		if (aPresses[2] == 0) && (aPresses[3] == 0) && (aPresses[0] == aPresses[1]) {
			// also on a, use a if cheaper
			if (aPresses[0] * 3) < bPresses[0] {
				return aPresses[0] * 3, aPresses[0], 0
			}
		}
		return bPresses[0], 0, bPresses[0]
	}
	// exact match on a
	if (aPresses[2] == 0) && (aPresses[3] == 0) && (aPresses[0] == aPresses[1]) {
		return aPresses[0] * 3, aPresses[0], 0
	}

	// find slopes and intercepts
	// y = aM*x + aB for button a
	// aM = y - aB * x for button a
	aM := float64(ay) / float64(ax)
	aB := float64(goalY) - (aM * float64(goalX))
	bM := float64(by) / float64(bx)
	bB := float64(goalY) - (bM * float64(goalX))

	// if both intercepts are posititive or negative there's no way to find
	// a solution
	if ((aB > 0) && (bB > 0)) || ((aB < 0) && (bB < 0)) {
		// unsolveable
		return 0, 0, 0
	}

	// now let's take the line from 0,0 using the slope of b, and find the
	// location where line a will intersect it.  Solve for x
	// y = bM*x
	// y = aM*x + aB
	// bM*x = aM*x + aB
	// bM*x - aM*x = aB
	// x * (bM - aM) = aB
	// x = aB / (bM - aM)

	// SPECIAL CASE: if (bM-aM) is 0, slopes are the same, no intercept
	if (bM - aM) == 0 {
		return 0, 0, 0
	}
	x := aB / (bM - aM)

	// ok, this should be an exact number if we have a solution I think
	// x is where we break apart from line B and start pressing A
	bPRESSES := (int(x) + 1) / bx // +1 for rounding errors
	aPRESSES := (goalY - (bPRESSES * by)) / ay
	cost := bPRESSES + aPRESSES*3
	// if not exact, screw it
	TARGETX := bPRESSES*bx + aPRESSES*ax
	TARGETY := bPRESSES*by + aPRESSES*ay
	if (TARGETX == goalX) && (TARGETY == goalY) {
		return cost, aPRESSES, bPRESSES
	}
	return 0, 0, 0 // no cigar
}

func part1(contents string) interface{} {
	total := 0
	for _, group := range aoc.ParseGroups(contents) {
		c1, a1, b1 := solveEasy(group, 0, 0)
		c2, a2, b2 := solveMath(group, false)
		if c1 != c2 {
			fmt.Printf("\n%s\n", group)
			fmt.Printf("  Expected: %d, %d, %d\n", c1, a1, b1)
			fmt.Printf("    Actual: %d, %d, %d\n", c2, a2, b2)
		}
		total += c2
		// fmt.Printf("%d: %d, %d\n", cost, aPresses, bPresses)
	}
	return total
}

func part2(contents string) interface{} {
	total := 0
	for _, group := range aoc.ParseGroups(contents) {
		// c1, a1, b1 := solveEasy(group, 0, 0)
		c2, _, _ := solveMath(group, true)
		// if c1 != c2 {
		// 	fmt.Printf("\n%s\n", group)
		// 	fmt.Printf("  Expected: %d, %d, %d\n", c1, a1, b1)
		// 	fmt.Printf("    Actual: %d, %d, %d\n", c2, a2, b2)
		// }
		total += c2
		// fmt.Printf("%d: %d, %d\n", cost, aPresses, bPresses)
	}
	return total
}
