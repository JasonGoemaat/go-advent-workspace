package main

import (
	"fmt"
	"sort"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 50)
	aoc.Local(part1, "part1", "input.aoc", 4776487744)
	aoc.Local(part2, "part2", "sample.aoc", 24)
	aoc.Local(part2, "part2", "input.aoc", 111936242)
}

func CalculateArea(a, b []int) int {
	dx := a[0] - b[0] + 1
	dy := a[1] - b[1] + 1
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx * dy
}

func part1(content string) interface{} {
	ints := aoc.ParseIntsPerLine(content)
	maxArea := 0
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			area := CalculateArea(ints[i], ints[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

type Side struct {
	X1, Y1       int
	X2, Y2       int
	IsHorizontal bool
}

func CreateSide(ints [][]int, i, j int) Side {
	X1, Y1 := ints[i][0], ints[i][1]
	X2, Y2 := ints[j][0], ints[j][1]
	if X2 < X1 {
		X1, X2 = X2, X1
	}
	if Y2 < Y1 {
		Y1, Y2 = Y2, Y1
	}
	IsHorizontal := Y1 == Y2
	return Side{X1, Y1, X2, Y2, IsHorizontal}
}

type IntRange struct {
	Start, End int
}

type LineRanges struct {
	Ranges []IntRange
}

type Rectangle struct {
	Corner1, Corner2 int
	Area             int
}

type Part2Puzzle struct {
	Ints            [][]int
	Sides           []Side
	HorizontalSides []Side
	VerticalSides   []Side
	Rectangles      []Rectangle
}

func IsInside(puzzle *Part2Puzzle, x, y int) bool {
	count := 0
	for _, side := range puzzle.VerticalSides {
		if side.X1 < x && side.Y1 <= y && side.Y2 >= y {
			count++
		}
	}
	return (count % 2) == 1
}

func part2(content string) interface{} {
	ints := aoc.ParseIntsPerLine(content)
	sides := make([]Side, len(ints))
	horizontalSides := make([]Side, 0)
	verticalSides := make([]Side, 0)
	for i := range len(ints) {
		sides[i] = CreateSide(ints, i, (i+1)%len(ints))
		if sides[i].IsHorizontal {
			horizontalSides = append(horizontalSides, sides[i])
		} else {
			verticalSides = append(verticalSides, sides[i])
		}
	}
	rectangles := make([]Rectangle, 0, len(ints)*len(ints))
	for i := range len(ints) - 1 {
		for j := i + 1; j < len(ints); j++ {
			rectangle := Rectangle{i, j, CalculateArea(ints[i], ints[j])}
			rectangles = append(rectangles, rectangle)
		}
	}
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Area > rectangles[j].Area
	})
	puzzle := Part2Puzzle{ints, sides, horizontalSides, verticalSides, rectangles}
	// for _, rectangle := range puzzle.Rectangles {
	// 	if IsValidRectangle(puzzle, rectangle) {
	// 		return rectangle.Area
	// 	}
	// }

	maxArea := 0
	invalidCount := 0
	for i := 0; i < len(puzzle.HorizontalSides)-1; i++ {
		for j := i + 1; j < len(puzzle.HorizontalSides); j++ {
			s1 := puzzle.HorizontalSides[i]
			s2 := puzzle.HorizontalSides[j]
			if s1.X1 > s2.X2 || s2.X1 > s1.X2 {
				// non-overlapping
				break
			}
			minY, maxY := s1.Y1, s2.Y1
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			minX, maxX := s1.X1, s1.X2
			if s2.X1 < minX {
				minX = s2.X1
			}
			if s2.X2 > maxX {
				maxX = s2.X2
			}
			dx := maxX - minX + 1
			dy := maxY - minY + 1
			area := dx * dy

			valid := true

			// check for hoizontal overlapping lines between
			for _, side := range puzzle.HorizontalSides {
				// check horizontal lines that intersect the left or right
				// sides of our rectangle

				if side.Y1 <= minY {
					// ignore horizontal lines 'above' our rectangle (or on it)
					continue
				}

				if side.Y1 >= maxY {
					// ignore horizontal lines 'below' our rectangle (or on it)
					continue
				}

				// intersect right side if right is on or greater and left is less
				if side.X1 < maxX && side.X2 >= maxX {
					valid = false
					break
				}

				// intersect left side if left is on or less and right is greater
				if side.X1 <= minX && side.X2 > minX {
					valid = false
					break
				}
			}

			// check for vertical overlapping lines between
			for _, side := range puzzle.VerticalSides {
				if side.X1 <= minX {
					continue // completely leftof our box
				}

				if side.X1 >= maxX {
					continue // completely right of our box
				}

				if side.Y1 < maxY && side.Y2 >= maxY {
					valid = false // intersects box bottom
					break
				}

				if side.Y1 <= minY && side.Y2 > minY {
					valid = false // intersects box top
					break
				}
			}

			// if central point is outside polygon, is invalid
			if valid {
				centerX := (minX + maxX) / 2
				centerY := (minY + maxY) / 2
				isInside := IsInside(&puzzle, centerX, centerY)
				if !isInside {
					valid = false
				}
			}

			if valid {
				if area > maxArea {
					maxArea = area
				}
			} else {
				invalidCount++
			}
		}
	}

	fmt.Printf("Skipped %d invalid rectangles\n", invalidCount)
	return maxArea
}

func IsValidRectangle(puzzle Part2Puzzle, rectangle Rectangle) bool {
	// 1. Any corner is inside the two values
	x1, x2 := puzzle.Ints[rectangle.Corner1][0], puzzle.Ints[rectangle.Corner2][0]
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	y1, y2 := puzzle.Ints[rectangle.Corner1][1], puzzle.Ints[rectangle.Corner2][1]
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	for _, corner := range puzzle.Ints {
		x, y := corner[0], corner[1]
		if x > x1 && x < x2 && y > y1 && y < y2 {
			return false
		}
	}

	// 2. Any horizontal line within the y bounds
	// that involves the range
	// for _, side := range puzzle.HorizontalSides {
	// 	if side.x >=
	// }
	return true
}
