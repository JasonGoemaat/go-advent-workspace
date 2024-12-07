package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/13
	aoc.Local(part1, "part1", "sample.aoc", 480)
	aoc.Local(part1, "part1", "input.aoc", 31897)
	aoc.Local(part2, "part2", "sample.aoc", 875318608908)
	aoc.Local(part2, "part2", "input.aoc", 85868156972635)
}

// NOTE: Don't appear to be any minus values, so we can just use
// parse numbers
func solve(content string) int {
	numsList := aoc.ParseLinesToInts(aoc.ParseLines(content))
	// no more than 100 button presses, let's brute-force it
	ax, ay := numsList[0][0], numsList[0][1]
	bx, by := numsList[1][0], numsList[1][1]
	goalX, goalY := numsList[2][0], numsList[2][1]
	lowestCost := 0
	for aCount := 0; aCount < 100; aCount++ {
		x1, y1 := ax*aCount, ay*aCount
		if (x1 > goalX) || (y1 > goalY) {
			break
		}
		for bCount := 0; bCount < 100; bCount++ {
			x2, y2 := bx*bCount, by*bCount
			if ((x1 + x2) >= goalX) && ((y1 + y2) >= goalY) {
				if ((x1 + x2) == goalX) && ((y1 + y2) == goalY) {
					cost := (aCount * 3) + (bCount * 1)
					if (lowestCost == 0) || (cost < lowestCost) {
						lowestCost = cost
					}
				}
				break
			}
		}
	}
	return lowestCost
}

func part1(contents string) interface{} {
	total := 0
	for _, group := range aoc.ParseGroups(contents) {
		total += solve(group)
	}
	return total
}

// I wondered how they'd make it harder, now we have to add 10 trillion (1e13)
// to the prize x and y coordinates.  That would take hours probably to
// brute-force each one.

// thought 1: is there a formula?   Take A(94,34) and B(22,67) with
// GOAL(8400,5400).   That becomes GOAL(10000000008400,10000000005400)
// So   a*94 + b*22 = 10000000008400
// Also a*34 + b*67 = 10000000005400
// Ok, reverse brute-force, but we also need a stopping point to realize
// if there is no solution.  We can find the least common multiple of two
// values.   If we just multiply 94*22 for the two possible X values, that
// is = 2068.   We should not have to go more than that far below
// 10000000008400 to find a solution (stop when we get below 10000000006332).
// at that point we will be adding some number of 2068s, either pressing the
// A button 22 times or the B button 94 times.  Least common multiple means
// we can merge common factors.   22 is 2*11 and 94 is 2*47, so we can use
// 2*11*47 or 1034.
// Let's go through an easier example.   Say the puzzle is finding the solution for this:
//		Button A: X+5, Y+2
//		Button B: X+3, Y+7
//		Prize: X=300, Y=180
// There will be some count N representing how many times we push both buttons,
// with one button being pushed more times than the other.  So A will be pushed
// N+Na and B will be pushed N+Nb.  We want to try and minimize Na because that
// costs 3x as much as Nb.

// ok, thought about for a minute.  There may be cases where we want to get
// to a point using As becaue B might take > 3 times as many pushes to get there.
// An easy example - A: X+50, Y+50 and B: X+15, Y+15
// Is that a 'special' case?   So maybe there's not a way to calculate exactly,
// but can we get to a 'close' point and then take-over with the initial solver?

// solve by brute force, but just small window 2 lower to 2 higher for a and b given
func solveBrute(content string, a, b int) int {
	numsList := aoc.ParseLinesToInts(aoc.ParseLines(content))
	// no more than 100 button presses, let's brute-force it
	ax, ay := numsList[0][0], numsList[0][1]
	bx, by := numsList[1][0], numsList[1][1]
	goalX, goalY := numsList[2][0], numsList[2][1]
	lowestCost := 0
	for aCount := a - 100; aCount < a+100; aCount++ {
		x1, y1 := ax*aCount, ay*aCount
		if (x1 > goalX) || (y1 > goalY) {
			break
		}
		for bCount := b - 100; bCount < b+100; bCount++ {
			x2, y2 := bx*bCount, by*bCount
			if ((x1 + x2) >= goalX) && ((y1 + y2) >= goalY) {
				if ((x1 + x2) == goalX) && ((y1 + y2) == goalY) {
					cost := (aCount * 3) + (bCount * 1)
					if (lowestCost == 0) || (cost < lowestCost) {
						lowestCost = cost
					}
				}
				break
			}
		}
	}
	return lowestCost
}

func solve2(content string, addIt bool) int {
	numsList := aoc.ParseLinesToInts(aoc.ParseLines(content))
	// no more than 100 button presses, let's brute-force it
	ax, ay := numsList[0][0], numsList[0][1]
	bx, by := numsList[1][0], numsList[1][1]
	goalX, goalY := numsList[2][0], numsList[2][1]
	if addIt {
		goalX += 10000000000000
		goalY += 10000000000000
	}
	aMax := min(goalX/ax, goalY/ay)

	aLow := 0
	aHigh := aMax

	for aLow <= aHigh {
		aMid := (aLow + aHigh) / 2
		axValue := aMid * ax
		ayValue := aMid * ay
		bxPresses := (goalX - axValue) / bx
		byPresses := (goalY - ayValue) / by
		bxValue := bxPresses * bx
		byValue := bxPresses * by
		// SECTION 1
		if (bxPresses == byPresses) && (bxValue+axValue == goalX) && (byValue+ayValue == goalY) {
			return (aMid * 3) + bxPresses
		}
		// move right if bxPresses > byPresses, testing if that will shorten the gap
		if bxPresses > byPresses {
			aLow = aMid + 1
		} else if bxPresses < byPresses {
			aHigh = aMid - 1
		} else {
			// // bxPresses and byPresses are the same, see if we can adjust a presses
			// // one of my failures has aMid=81 here and it needs to be 82
			// axPresses, ayPresses := (goalX-bxValue)/ax, (goalY-byValue)/ay
			// if axPresses == ayPresses {
			// 	axValue := axPresses * ax
			// 	ayValue := ayPresses * ay
			// 	if ((axValue + bxValue) == goalX) && ((ayValue + byValue) == goalY) {
			// 		return (axPresses * 3) + bxPresses
			// 	}
			// }
			return solveBrute(content, aMid, bxPresses)
		}
	}

	aLow = 0
	aHigh = aMax

	for aLow <= aHigh {
		aMid := (aLow + aHigh) / 2
		axValue := aMid * ax
		ayValue := aMid * ay
		bxPresses := (goalX - axValue) / bx
		byPresses := (goalY - ayValue) / by
		bxValue := bxPresses * bx
		byValue := bxPresses * by
		// SECTION 2
		if (bxPresses == byPresses) && (bxValue+axValue == goalX) && (byValue+ayValue == goalY) {
			return (aMid * 3) + bxPresses
		}
		// swapping comparison from above
		if bxPresses < byPresses {
			aLow = aMid + 1
		} else if bxPresses > byPresses {
			aHigh = aMid - 1
		} else {
			// // bxPresses and byPresses are the same, see if we can adjust a presses
			// // one of my failures has aMid=81 here and it needs to be 82
			// axPresses, ayPresses := (goalX-bxValue)/ax, (goalY-byValue)/ay
			// if axPresses == ayPresses {
			// 	axValue := axPresses * ax
			// 	ayValue := ayPresses * ay
			// 	if ((axValue + bxValue) == goalX) && ((ayValue + byValue) == goalY) {
			// 		return (axPresses * 3) + bxPresses
			// 	}
			// }
			// return 0 // impossible
			return solveBrute(content, aMid, bxPresses)
		}
	}

	return 0 // no solution
}

func part2(contents string) interface{} {
	total := 0
	for _, group := range aoc.ParseGroups(contents) {
		// s1 := solve(group)
		// s2 := solve2(group, false)
		// if s1 != s2 {
		// 	fmt.Printf("Difference %d => %d for group %d\n", s1, s2, groupId)
		// 	fmt.Printf("%s\n", group)
		// }
		total += solve2(group, true) // dang, not working: 85868156972635,
	}
	return total
}
