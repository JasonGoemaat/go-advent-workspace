package main

import (
	"fmt"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/14
	aoc.Local(part1, "part1", "sample.aoc", 12)
	aoc.Local(part1, "part1", "input.aoc", 210587128)
	// part 2 is odd, wants you to say how many for a christmas tree
	// I *could* maybe see if it is symmetrical left/right, let's try some input though
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	// aoc.Local(part2, "part2", "input.aoc", 0) // manual solving
	// aoc.Local(part2b, "part2b", "input.aoc", 0)
}

type (
	Position struct{ x, y int }
)

func part1(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 11, 7 // sample size
	if len(numLists) > 20 {
		w, h = 101, 103 // actual input size
	}
	endingPositions := make([]Position, len(numLists))
	for i, nums := range numLists {
		endingPositions[i] = Position{(nums[0] + nums[2]*100) % w, (nums[1] + nums[3]*100) % h}
		if endingPositions[i].x < 0 {
			endingPositions[i].x += w
		}
		if endingPositions[i].y < 0 {
			endingPositions[i].y += h
		}
	}
	quadrants := []int{0, 0, 0, 0}
	mx := w / 2
	my := h / 2
	for _, pos := range endingPositions {
		if pos.x < mx {
			if pos.y < my {
				quadrants[0]++
			} else if pos.y > my {
				quadrants[1]++
			}
		} else if pos.x > mx {
			if pos.y < my {
				quadrants[2]++
			} else if pos.y > my {
				quadrants[3]++
			}
		}
	}
	total := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	return total
}

func part2(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 11, 7 // sample size
	if len(numLists) > 20 {
		w, h = 101, 103 // actual input size
	}
	quadrants := []int{0, 0, 0, 0}

	seconds := 488
	for ; seconds < 1000000; seconds += 103 {
		finder := map[Position]int{}
		endingPositions := make([]Position, len(numLists))
		for i, nums := range numLists {
			endingPositions[i] = Position{(nums[0] + nums[2]*seconds) % w, (nums[1] + nums[3]*seconds) % h}
			if endingPositions[i].x < 0 {
				endingPositions[i].x += w
			}
			if endingPositions[i].y < 0 {
				endingPositions[i].y += h
			}
			finder[endingPositions[i]] += 1
		}
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if finder[Position{x, y}] > 0 {
					fmt.Printf("%d", finder[Position{x, y}])
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println("==================== Seconds:", seconds, "====================")
		fmt.Println()
		fmt.Println()
		fmt.Scanln()
	}
	endingPositions := make([]Position, len(numLists))
	for i, nums := range numLists {
		endingPositions[i] = Position{(nums[0] + nums[2]*100) % w, (nums[1] + nums[3]*100) % h}
		if endingPositions[i].x < 0 {
			endingPositions[i].x += w
		}
		if endingPositions[i].y < 0 {
			endingPositions[i].y += h
		}
	}
	mx := w / 2
	my := h / 2
	for _, pos := range endingPositions {
		if pos.x < mx {
			if pos.y < my {
				quadrants[0]++
			} else if pos.y > my {
				quadrants[1]++
			}
		} else if pos.x > mx {
			if pos.y < my {
				quadrants[2]++
			} else if pos.y > my {
				quadrants[3]++
			}
		}
	}
	total := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	return total
}

func mostConnected(positions map[Position]int, includeDiagonals bool) int {
	// duplicate positions
	newPositions := map[Position]bool{}
	for pos, _ := range positions {
		newPositions[pos] = true
	}

	var recurse func(p Position) int
	recurse = func(p Position) int {
		if !newPositions[p] {
			return 0
		}
		newPositions[p] = false
		count := 1 +
			recurse(Position{p.x, p.y - 1}) +
			recurse(Position{p.x, p.y + 1}) +
			recurse(Position{p.x - 1, p.y}) +
			recurse(Position{p.x + 1, p.y})
		if includeDiagonals {
			count += recurse(Position{p.x + 1, p.y + 1}) +
				recurse(Position{p.x + 1, p.y - 1}) +
				recurse(Position{p.x - 1, p.y + 1}) +
				recurse(Position{p.x - 1, p.y - 1})
		}
		return count
	}

	maxCount := 0
	for position, _ := range newPositions {
		count := recurse(position)
		maxCount = max(count, maxCount)
	}

	return maxCount
}

func part2b(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 11, 7 // sample size
	if len(numLists) > 20 {
		w, h = 101, 103 // actual input size
	}
	maxCount, secondMax := 0, 0
	maxSeconds := 0

	seconds := 0
	for ; seconds < (w*h+w+h)*2; seconds++ {
		finder := map[Position]int{}
		endingPositions := make([]Position, len(numLists))
		for i, nums := range numLists {
			endingPositions[i] = Position{(nums[0] + nums[2]*seconds) % w, (nums[1] + nums[3]*seconds) % h}
			if endingPositions[i].x < 0 {
				endingPositions[i].x += w
			}
			if endingPositions[i].y < 0 {
				endingPositions[i].y += h
			}
			finder[endingPositions[i]] += 1
		}
		count := mostConnected(finder, true)
		if count >= maxCount {
			secondMax, maxCount, maxSeconds = maxCount, count, seconds
			fmt.Printf("%d seconds: %d - (max %d, second %d)\n", seconds, count, maxCount, secondMax)
		} else if count >= secondMax {
			secondMax = count
			// fmt.Printf("%d seconds: %d - (max %d, second %d)\n", seconds, count, maxCount, secondMax)
		}
		// for y := 0; y < h; y++ {
		// 	for x := 0; x < w; x++ {
		// 		if finder[Position{x, y}] > 0 {
		// 			fmt.Printf("%d", finder[Position{x, y}])
		// 		} else {
		// 			fmt.Printf(" ")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println("==================== Seconds:", seconds, "====================")
		// fmt.Println()
		// fmt.Println()
		// fmt.Scanln()
	}
	// endingPositions := make([]Position, len(numLists))
	// for i, nums := range numLists {
	// 	endingPositions[i] = Position{(nums[0] + nums[2]*100) % w, (nums[1] + nums[3]*100) % h}
	// 	if endingPositions[i].x < 0 {
	// 		endingPositions[i].x += w
	// 	}
	// 	if endingPositions[i].y < 0 {
	// 		endingPositions[i].y += h
	// 	}
	// }
	// mx := w / 2
	// my := h / 2
	// for _, pos := range endingPositions {
	// 	if pos.x < mx {
	// 		if pos.y < my {
	// 			quadrants[0]++
	// 		} else if pos.y > my {
	// 			quadrants[1]++
	// 		}
	// 	} else if pos.x > mx {
	// 		if pos.y < my {
	// 			quadrants[2]++
	// 		} else if pos.y > my {
	// 			quadrants[3]++
	// 		}
	// 	}
	// }
	// total := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	fmt.Printf("Max at seconds = %d (%d) second highest %d\n", maxSeconds, maxCount, secondMax)
	return maxCount
}
