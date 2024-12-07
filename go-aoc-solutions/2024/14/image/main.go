package main

//--------------------------------------------------------------------------------
// THIS is to produce images for part 2
//--------------------------------------------------------------------------------

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/X
	// aoc.Local(part1, "part1", "sample.aoc", 0)
	// aoc.Local(part2, "part2b", "input.aoc", 0)
	// aoc.Local(part2, "part2b", "sample.aoc", 0) // trying with sample
	// aoc.Local(part2c, "part2c", "sample.aoc", 0) // one big image with sample and bigger map
	aoc.Local(part2e, "part2e", "input.aoc", 0)
}

type (
	Position struct{ x, y int }
	Robots   map[Position]int
)

func getRobots(numLists [][]int, w, h, seconds int) Robots {
	robots := Robots{}
	for _, nums := range numLists {
		p := Position{(nums[0] + nums[2]*seconds) % w, (nums[1] + nums[3]*seconds) % h}
		if p.x < 0 {
			p.x += w
		}
		if p.y < 0 {
			p.y += h
		}
		robots[p] += 1
	}
	return robots
}

func part1(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 11, 7 // actual input size

	for method := 0; method < 6; method++ {
		seconds := 0
		for seconds < w*h {
			myw, myh := w, h
			if (method & 0x1) > 0 {
				myw, myh = myh, myw
			}
			m := image.NewRGBA(image.Rect(0, 0, w*myw, h*myh))
			// start := seconds
			// 200 per image
			for index := 0; index < 200; index++ {
				row := index / myw
				col := index % myw
				top := row * h
				left := col * w
				robots := getRobots(numLists, w, h, seconds)
				if (method / 2) == 0 {
					// NOOP
				} else if (method / 2) == 1 {
					// delete middles
					for y := 0; y < h; y++ {
						delete(robots, Position{w / 2, y})
					}
					for x := 0; x < w; x++ {
						delete(robots, Position{x, h / 2})
					}
				} else if (method / 2) == 2 {
					// populate middles
					for y := 0; y < h; y++ {
						robots[Position{w / 2, y}] = 5
					}
					for x := 0; x < w; x++ {
						robots[Position{x, h / 2}] = 5
					}
				} else {
					panic(fmt.Sprintf("Bad method: %d", method))
				}
				for pos, _ := range robots {
					m.Set(pos.x+left, pos.y+top, color.RGBA{128, 255, 128, 255})
				}
				seconds++
			}
			pngPath := aoc.GetSubPath(fmt.Sprintf("output/part1_%d.png", method))
			file, err := os.Create(pngPath)
			if err != nil {
				panic(fmt.Sprintf("ERROR CREATING FILE: %s", pngPath))
			}
			defer file.Close()
			png.Encode(file, m)
		}
	}
	return 0
}

func part2(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 101, 103 // actual input size

	seconds := 0
	for seconds < w*h {
		m := image.NewRGBA(image.Rect(0, 0, w*20, h*10))
		start := seconds
		// 200 per image
		for index := 0; index < 200; index++ {
			row := index / 20
			col := index % 20
			top := row * h
			left := col * w
			robots := getRobots(numLists, w, h, seconds)
			for pos, _ := range robots {
				m.Set(pos.x+left, pos.y+top, color.RGBA{128, 255, 128, 255})
			}
			seconds++
		}
		pngPath := aoc.GetSubPath(fmt.Sprintf("output2/%05d.png", start))
		file, err := os.Create(pngPath)
		if err != nil {
			panic(fmt.Sprintf("ERROR CREATING FILE: %s", pngPath))
		}
		defer file.Close()
		png.Encode(file, m)
	}
	return 0
}

// same as part2 drawing pngs, but use lines and one big one that is 10430x10430
func part2c(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 101, 103 // actual input size
	iw, ih := w*w+w-1, h*h+h-1

	seconds := 0
	for seconds < w*h {
		m := image.NewRGBA(image.Rect(0, 0, iw, ih))
		start := seconds
		for index := 0; index < w*h; index++ {
			row := index / w
			col := index % w
			top := row * (h + 1)
			left := col * (w + 1)
			robots := getRobots(numLists, w, h, seconds)
			for pos, _ := range robots {
				m.Set(pos.x+left, pos.y+top, color.RGBA{192, 255, 192, 255})
			}
			if top > 0 {
				for x := 0; x < w; x++ {
					m.Set(x+left, top-1, color.RGBA{255, 255, 255, 255})
				}
			}
			if left > 0 {
				for y := 0; y < h; y++ {
					m.Set(left-1, y+top, color.RGBA{255, 255, 255, 255})
				}
			}
			seconds++
		}
		pngPath := aoc.GetSubPath(fmt.Sprintf("output2/part2c_%d.png", start))
		file, err := os.Create(pngPath)
		if err != nil {
			panic(fmt.Sprintf("ERROR CREATING FILE: %s", pngPath))
		}
		defer file.Close()
		png.Encode(file, m)
	}
	return 0
}

// same as part2 drawing pngs, but use lines and one big one that is 10430x10430
func part2d(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 103, 101 // actual input size
	iw, ih := w*w+w-1, h*h+h-1

	seconds := 0
	for seconds < w*h {
		m := image.NewRGBA(image.Rect(0, 0, iw, ih))
		start := seconds
		for index := 0; index < w*h; index++ {
			row := index / w
			col := index % w
			top := row * (h + 1)
			left := col * (w + 1)
			robots := getRobots(numLists, w, h, seconds)
			for pos, _ := range robots {
				m.Set(pos.x+left, pos.y+top, color.RGBA{192, 255, 192, 255})
			}
			if top > 0 {
				for x := 0; x < w; x++ {
					m.Set(x+left, top-1, color.RGBA{255, 255, 255, 255})
				}
			}
			if left > 0 {
				for y := 0; y < h; y++ {
					m.Set(left-1, y+top, color.RGBA{255, 255, 255, 255})
				}
			}
			seconds++
		}
		pngPath := aoc.GetSubPath(fmt.Sprintf("output2/part2c_%d.png", start))
		file, err := os.Create(pngPath)
		if err != nil {
			panic(fmt.Sprintf("ERROR CREATING FILE: %s", pngPath))
		}
		defer file.Close()
		png.Encode(file, m)
	}
	return 0
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

// like original part2, output to 'output3' directory, have white lines between
func part2e(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 101, 103         // actual input size
	numLists = numLists[0:1] // just one pixel, checking values

	seconds := 0
	for seconds < min(103, w*h) {
		m := image.NewRGBA(image.Rect(0, 0, w*20+20, h*10+10))
		// start := seconds
		// 200 per image
		for index := 0; index < 200; index++ {
			row := index / 20
			col := index % 20
			top := row*h + row
			left := col*w + col
			robots := getRobots(numLists, w, h, seconds)
			for pos, _ := range robots {
				fmt.Printf("%d: at %d, %d\n", seconds, pos.x, pos.y)
				m.Set(pos.x+left, pos.y+top, color.RGBA{128, 255, 128, 255})
			}
			if left > 0 {
				for y := range h + 1 {
					m.Set(left, top+y-1, color.RGBA{255, 255, 255, 255})
				}
			}
			if top > 0 {
				for x := range w + 1 {
					m.Set(left+x-1, top, color.RGBA{255, 255, 255, 255})
				}
			}
			seconds++
		}
		// pngPath := aoc.GetSubPath(fmt.Sprintf("output3/%05d.png", start))
		// file, err := os.Create(pngPath)
		// if err != nil {
		// 	panic(fmt.Sprintf("ERROR CREATING FILE: %s", pngPath))
		// }
		// defer file.Close()
		// png.Encode(file, m)
	}
	return 0
}
