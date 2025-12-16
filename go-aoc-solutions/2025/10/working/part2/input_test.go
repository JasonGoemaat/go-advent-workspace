package part2

import (
	"fmt"
	"testing"
	"time"

	o "github.com/JasonGoemaat/go-aoc-solutions/2025/10/working/lib"
	"github.com/JasonGoemaat/go-aoc/aoc"
)

func TestSingle(t *testing.T) {
	// input.aoc line 2 - takes too long
	input := "[##..#..] (0,3) (0,1,2,5,6) (0,1,2,4,5) (1,3,4,6) (1,2,3,4,5) (1,6) (1,2,4,5) (0,5) (3,4,5,6) {151,96,61,49,74,197,61}" // test 2, should be 12 presses, but I find 10
	matrix := ParsePuzzle(input)
	matrix.Report(fmt.Sprintf("Parsed: %s", input))
	fmt.Println("Input for Wolfram:", matrix.GetWolframString())
	// matrix.RREF()
	matrix.RREFRecurse(0, 0, false)
	// I see that I'm stopped by 4,4 which has a 2, however if I see
	// column 5 has a leading -1 that could be used in row 6 and
	// column 7 has a leading 1 that can be used in row 5...
	// I'm gonna try solving it MY way by swapping columns to ensure
	// there is a 1 in position 4,4...
	matrix.Report("RREF()")
	// matrix.Solve()
}

func TestLine2(t *testing.T) {
	input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
	fmt.Println(SimpleSolve(input))
}

func Test28(t *testing.T) {
	input := "[###..#...#] (3,6,7) (1,2,6) (0,2,3,4,5,6,9) (1) (0,1,2,5,6,7) (0,1,2,3,6,7,8,9) (0,1,2,3,5,6,8,9) (1,2,3,4,5,6,8,9) (0,5,6,8,9) (1,2,4,7,9) (0,3,8,9) (0,2,4,5,6,7,8) (2,3,5,6,8,9) {56,74,68,51,33,39,58,48,52,69}"
	m := ParsePuzzle(input)
	fmt.Println(m)
	m.RREFRecurse(0, 0, true)

	// I see right away I can get rid of one more column.   Row 9
	// has a 2 as the pivot, but the only other thing in the row is the
	// joltage which is 8, so it can be reduced to 1 and 4 and take 1/39th the time
	// That also has a value in EVERY row.   Also I could swap columns 10 and 8
	// before that to get rid of one more button.

	start := time.Now()
	min1 := m.Solve()[0]
	d1 := time.Since(start)

	// now perform operations
	fmt.Println("Before optimizing")
	fmt.Println(m)

	// fmt.Println("Divided row 9 by 2")
	// m.Set(9, 8, 1)  // only button value in row, originally 2
	// m.Set(9, 13, 4) // joltage, originally 4
	// // this could be handled by an optimization checking if there is only 1 button in a row also,
	// // or by checking for LCD and doing that at least if multiple values are divisible
	// fmt.Println(m)

	// m.SwapCols(8, 10)
	// fmt.Println("Swapped columns 8 and 10")
	// fmt.Println(m)

	// m.RREFRecurse(8, 8, true)
	// fmt.Println("Re-ran RREF from 8,8")

	// problems when I allow column swapping:
	// input.aoc line 78 I have 94 but other says 88
	// input.aoc line 84 I have 89 but other says 85
	// input.aoc line 163 I have 76 but other says 74
	// input.aoc line 187 I have 79 but other says 62

	start = time.Now()
	min2 := m.Solve()[0]
	d2 := time.Since(start)

	fmt.Printf("Original: %d in %s\n", min1, d1)
	fmt.Printf("Optimized: %d in %s\n", min2, d2)
}

func TestSolve(t *testing.T) {
	aoc.Local(SolvePart2, "SolvePart2", "input.aoc", 20172)
}

func SolvePart2(content string) any {
	// Initial: 18613ms SolvePart2("input.aoc") = 20172 (BAD - Expected <nil>)
	// ConstrainMaxPresses() 18119ms SolvePart2("input.aoc") = 20172 (BAD - Expected <nil>)
	// Line 28 takes by FAR the longest
	// Solved 28: 95 presses - [###..#...#] (3,6,7) (1,2,6) (0,2,3,4,5,6,9) (1) (0,1,2,5,6,7) (0,1,2,3,6,7,8,9) (0,1,2,3,5,6,8,9) (1,2,3,4,5,6,8,9) (0,5,6,8,9) (1,2,4,7,9) (0,3,8,9) (0,2,4,5,6,7,8) (2,3,5,6,8,9) {56,74,68,51,33,39,58,48,52,69}
	lines := aoc.ParseLines(content)
	sum := 0
	for _, line := range lines {
		matrix := ParsePuzzle(line)
		matrix.RREF()
		solutions := matrix.Solve()
		sum += solutions[0].TotalPresses
		// fmt.Printf("Solved %d: %d presses - %s\n", i, solutions[0].TotalPresses, line)
	}
	return sum
}

func TestCompare(t *testing.T) {
	lines := aoc.ParseLines(aoc.GetLocalFile("input.aoc"))
	for i, line := range lines {
		mine := SimpleSolve(line)
		theirs := o.Solution10B(line)
		if mine != theirs {
			fmt.Printf("input.aoc line %d I have %d but other says %d\n", i+1, mine, theirs)
		}
	}
}
