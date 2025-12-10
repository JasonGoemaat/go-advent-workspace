package main

import (
	"fmt"
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

// ..............
// .......#XXX#.. // 3. create third area - 7,1 to 11,1 (height 1, width 5)
// .......X...X.. // 2. create second area - 7,3 to 7,1 (height 3, width 1)
// ..#XXXX#...X.. // 1. Create first area - 2,3 to 7,3 (height 1, width 5)
// ..X........X..
// ..#XXXXXX#.X..
// .........X.X..
// .........#X#..
// ..............

// PROBLEM AREA
// ..............
// .......AXXXB.. // top-left of largest ara
// .......X...X..
// ..CXXXXD...+..
// ..X........X..
// ..EXXXF....+..
// ......X....X..
// ......GXXXXH.. // bottom-right of largest area
// ..............
// Areas: A-H, C-F, F-H, D-H
// Area is not valid if:
// 1. There is a horizontal line that CROSSES inside of it (start, end inside (on is ok), or if any space inside)
var sample = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestSpatial(t *testing.T) {
	ints := aoc.ParseIntsPerLine(sample)
	sides := make([]Side, len(ints))
	for i := range len(ints) {
		sides[i] = CreateSide(ints, i, (i+1)%len(ints))
	}

	fmt.Println(sides)
}
