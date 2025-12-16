package part2

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	m := NewMatrix(3, 3)
	m.Set(1, 2, -1)
	fmt.Println(m)
}

func TestSample1(t *testing.T) {
	input := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	matrix := ParsePuzzle(input)
	matrix.Report(fmt.Sprintf("Parsed: %s", input))
	fmt.Println("Input for Wolfram:", matrix.GetWolframString())
	matrix.RREF()
	matrix.Report("RREF()")
	solutions := matrix.Solve()
	for i, solution := range solutions {
		fmt.Printf("Solution %d has %d total presses: %v\n", i, solution.TotalPresses, solution.Presses)
	}
}

func TestSample2(t *testing.T) {
	input := "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}" // test 2, should be 12 presses, but I find 10
	matrix := ParsePuzzle(input)
	matrix.Report(fmt.Sprintf("Parsed: %s", input))
	fmt.Println("Input for Wolfram:", matrix.GetWolframString())
	// matrix.RREF()
	matrix.RREFRecurse(0, 0, false)
	matrix.Report("RREF()")

	start := time.Now()
	var solutions []MatrixSolution
	iterations := 1000
	for _ = range iterations {
		solutions = matrix.Solve()
	}
	duration := time.Since(start)

	for i, solution := range solutions {
		fmt.Printf("Solution %d has %d total presses: %v\n", i, solution.TotalPresses, solution.Presses)
	}

	minPresses := solutions[0].TotalPresses
	minCount := 0
	for _, solution := range solutions {
		if solution.TotalPresses == minPresses {
			minCount++
		} else {
			break
		}
	}
	fmt.Printf("Found %d solutions %d times with %d presses in %s\n", minCount, iterations, minPresses, duration)
}

func TestSample3(t *testing.T) {
	// sample 3 input - minimum presses 11: 5,0,5,1
	input := "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}"
	matrix := ParsePuzzle(input)
	matrix.Report(fmt.Sprintf("Parsed: %s", input))
	fmt.Println("Input for Wolfram:", matrix.GetWolframString())
	// matrix.RREF()
	matrix.RREFRecurse(0, 0, true)
	matrix.Report("RREF()")

	start := time.Now()
	var solutions []MatrixSolution
	iterations := 1000
	for _ = range iterations {
		solutions = matrix.Solve()
	}
	duration := time.Since(start)

	fmt.Printf("Checked %d solutions\n", len(solutions))
	for i, solution := range solutions {
		fmt.Printf("Solution %d has %d total presses: %v\n", i, solution.TotalPresses, solution.Presses)
	}

	if len(solutions) == 0 {
		panic("no solutions")
	}
	minPresses := solutions[0].TotalPresses
	minCount := 0
	for _, solution := range solutions {
		if solution.TotalPresses == minPresses {
			minCount++
		} else {
			break
		}
	}
	fmt.Printf("Found %d solutions %d times with %d presses in %s\n", minCount, iterations, minPresses, duration)
}

/*
Wolfram answer reducing to reduced row echelon form:
(1 | 0 | 0 | 1 | 0 | -1 | 2
 0 | 1 | 0 | 0 | 0 |  1 | 5
 0 | 0 | 1 | 1 | 0 | -1 | 1
 0 | 0 | 0 | 0 | 1 |  1 | 3)

This solver gives the same result and shows 3 operations done:
https://www.math.odu.edu/~bogacki/cgi-bin/lat.cgi?c=rref

Seeing this, it's cool that the last row starts with a '1' and
the 4th column doesn't have a pivot.  So if a column is all
zeros when looking for a pivot, we can skip it.

Sample answer on AOC: https://adventofcode.com/2025/day/10

Configuring the first machine's counters requires a minimum of 10 button presses.
One way to do this is by pressing (3) once, (1,3) three times, (2,3) three times,
(0,2) once, and (0,1) twice.

Pressing button 0 (3) once gives joltages (3,5,4,6)
Pressing button 1 (1,3) three times gives (3,2,4,3)
button 2 is not included
Pressing button 3 (2,3) three times gives (3,2,1,0)
Pressing button 4 (0,2) one time gives    (2,2,0,0)
Pressing button 5 (0,1) twice gives       (0,0,0,0) - solved
*/
