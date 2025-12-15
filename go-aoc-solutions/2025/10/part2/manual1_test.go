package part2

import (
	"fmt"
	"testing"
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
