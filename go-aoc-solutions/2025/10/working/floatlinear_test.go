package main

import (
	"fmt"
	"testing"
)

// [..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}
func TestFloatMatrixSolve(t *testing.T) {
	input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
	machine := ParseMachine2(input)
	matrix := NewFloatMatrix(len(machine.Joltages), len(machine.Buttons)+1)
	for i, button := range machine.Buttons {
		for _, joltageIndex := range button {
			matrix.Set(joltageIndex, i, 1)
		}
	}
	for i, joltage := range machine.Joltages {
		matrix.Set(i, len(machine.Buttons), float64(joltage))
	}
	fmt.Println(matrix)
	matrix.RemoveDuplicateRows()
	matrix.RESort()
	fmt.Println(matrix)

	// ok, subtracting row 0 from every ther row with a 1 in the first column
	matrix.AddRows(0, 1, -1)
	matrix.AddRows(0, 2, -1)
	matrix.AddRows(0, 3, -1)
	matrix.AddRows(0, 4, -1)
	matrix.AddRows(0, 5, -1)
	matrix.AddRows(0, 6, -1)
	matrix.RESort()
	// now row 1 has pivot in column 1
	// add row 1 to rows 5-8 to make them 0 in column 1
	fmt.Println(matrix)
	fmt.Println("Adding row 1 to rows 5-8")
	matrix.AddRows(1, 5, 1)
	matrix.AddRows(1, 6, 1)
	matrix.AddRows(1, 7, 1)
	matrix.AddRows(1, 8, 1)
	matrix.RESort()

	fmt.Println("Should be one row with column 0 a 1.0")
	fmt.Println(matrix)

	fmt.Println("Row 2 column 2 has pivot, add to row 8 to make it a zero")
	matrix.AddRows(2, 8, 1)
	matrix.RESort()
	fmt.Println(matrix)

	fmt.Println("Row 6 is a good one for next pivot, it has 1 in the next column, swap")
	matrix.SwapRows(3, 6)
	fmt.Println(matrix)
	fmt.Println("Now subtracting from 4,5,6 and add to 8 to make them 0")
	matrix.AddRows(3, 4, -1)
	matrix.AddRows(3, 5, -1)
	matrix.AddRows(3, 6, -1)
	matrix.AddRows(3, 8, 1)
	matrix.RESort()
	fmt.Println(matrix)

	// for column 4, row 7 is a good pivot, it has a leading 1 and negative value,
	// other rows have positive, so subtracting will increase their positive
	// joltages
	matrix.SwapRows(4, 7)
	/*
		│1.00 1.00 1.00 1.00 0.00  1.00  1.00 0.00 0.00  0.00  53.00│
		│0.00 1.00 1.00 1.00 1.00  0.00  0.00 1.00 0.00  0.00  43.00│
		│0.00 0.00 1.00 1.00 1.00  0.00  1.00 0.00 0.00  1.00  44.00│
		│0.00 0.00 0.00 1.00 1.00 -1.00  0.00 1.00 0.00  2.00  45.00│
		│0.00 0.00 0.00 0.00 1.00  1.00 -1.00 1.00 0.00 -2.00 -13.00│
		│0.00 0.00 0.00 0.00 2.00 -1.00  0.00 1.00 0.00  3.00  50.00│
		│0.00 0.00 0.00 0.00 1.00  1.00  0.00 1.00 1.00 -2.00   0.00│
		│0.00 0.00 0.00 0.00 1.00  1.00  0.00 1.00 0.00 -2.00   0.00│
		│0.00 0.00 0.00 0.00 1.00  0.00 -1.00 2.00 1.00  0.00  15.00│
	*/
	// make other rows into 0 in that column
	fmt.Println("Swapping 4 and 7, adding 4 to 5-8 with (-1, -1, -2, 1)")
	fmt.Println(matrix)

	matrix.AddRows(4, 5, -1)
	matrix.AddRows(4, 6, -1)
	matrix.AddRows(4, 7, -2)
	matrix.AddRows(4, 8, -1)
	matrix.RESort()
	fmt.Println(matrix)

	// Row 7 is a good candidate, multiply by -1
	// then move to 5, then add to 8 with factor-3
	matrix.MultiplyRow(7, -1)
	matrix.SwapRows(5, 7)
	matrix.AddRows(5, 8, 3)
	matrix.RESort()
	fmt.Println(matrix)

	// For row 6, modify rows 6 and 7, then use 8
	matrix.AddRows(8, 6, -2)
	matrix.AddRows(8, 7, -1)
	matrix.SwapRows(6, 8)
	matrix.RESort()
	fmt.Println(matrix)

	// for row 7, use r8 lmultiplied by -1/4
	fmt.Println("<----- ABOUT TO DO STUPID ----->")
	fmt.Println("Multiplying row 8 by -.25 and swapping with 7")
	matrix.MultiplyRow(8, -0.25)
	matrix.SwapRows(7, 8)
	matrix.RESort()
	fmt.Println(matrix)

	// at the end, the last row has a 1 in column 8, but 0 in answer column,
	// I ***THINK*** that means that button doesn't matter, or shouldn't be pushed?
	// I can't add/subtract it from other rows because that would alter the values
	// for other equations using that button, but not alter the joltage produced.
	// so here I'm going to start with row 7 with pivot in column 7 and ensure other
	// values in column 7 above are 0
	matrix.AddRows(7, 5, 1)
	matrix.AddRows(7, 4, -1)
	matrix.AddRows(7, 3, -1)
	matrix.AddRows(7, 1, -1)
	fmt.Println("Column 7 done?")
	fmt.Println(matrix)

	// not column 6 row 6 so that rows 0-5 are 0
	matrix.AddRows(6, 4, 1)
	matrix.AddRows(6, 2, -1)
	matrix.AddRows(6, 0, -1)
	fmt.Println("Column 6 done?")
	fmt.Println(matrix)

	// now pivot 5
	/*
	    ┌                                                                 ┐
	   0│ 1.00  1.00  1.00  1.00  0.00  1.00  0.00 0.00  0.00  0.00  40.00│
	   1│ 0.00  1.00  1.00  1.00  1.00  0.00  0.00 0.00 -0.75  0.25  34.50│
	   2│ 0.00  0.00  1.00  1.00  1.00  0.00  0.00 0.00  0.00  1.00  31.00│
	   3│ 0.00  0.00  0.00  1.00  1.00 -1.00  0.00 0.00 -0.75  2.25  36.50│
	   4│ 0.00  0.00  0.00  0.00  1.00  1.00  0.00 0.00 -0.75 -1.75  -8.50│
	   5│-0.00 -0.00 -0.00 -0.00 -0.00  1.00 -0.00 0.00 -0.25 -2.25 -19.50│
	   6│ 0.00  0.00  0.00  0.00  0.00  0.00  1.00 0.00  0.00  0.00  13.00│
	   7│-0.00 -0.00 -0.00 -0.00 -0.00 -0.00 -0.00 1.00  0.75 -0.25   8.50│
	   8│ 0.00  0.00  0.00  0.00  0.00  0.00  0.00 0.00  1.00  0.00   0.00│
	    └                                                                 ┘
	*/
	matrix.AddRows(5, 3, 1)
	matrix.AddRows(5, 4, -1)
	matrix.AddRows(5, 0, -1)
	fmt.Println("Column 5 done?")
	fmt.Println(matrix)

	matrix.AddRows(4, 3, -1)
	matrix.AddRows(4, 2, -1)
	matrix.AddRows(4, 1, -1)
	fmt.Println("Column 4 done?")
	fmt.Println(matrix)
	matrix.AddRows(3, 2, -1)
	matrix.AddRows(3, 1, -1)
	matrix.AddRows(3, 0, -1)
	matrix.AddRows(2, 1, -1)
	matrix.AddRows(2, 0, -1)
	matrix.AddRows(1, 0, -1)
	fmt.Println("Done?")
	fmt.Println(matrix)
}

/*
Subtract row 0 from rows 1-6 so they have leading zeros, then resort
┌                                                       ┐
│1.00 1.00 1.00 1.00 0.00 1.00 1.00 0.00 0.00 0.00 53.00│
│1.00 1.00 1.00 0.00 1.00 1.00 1.00 0.00 0.00 1.00 58.00│
│1.00 1.00 0.00 1.00 0.00 0.00 0.00 1.00 0.00 1.00 54.00│
│1.00 0.00 0.00 1.00 1.00 1.00 1.00 1.00 1.00 0.00 55.00│
│1.00 0.00 0.00 1.00 1.00 1.00 1.00 1.00 0.00 0.00 55.00│
│1.00 0.00 0.00 1.00 1.00 1.00 0.00 1.00 0.00 0.00 42.00│
│1.00 0.00 0.00 0.00 0.00 1.00 0.00 1.00 1.00 0.00 25.00│
│0.00 1.00 1.00 1.00 1.00 0.00 0.00 1.00 0.00 0.00 43.00│
│0.00 0.00 1.00 1.00 1.00 0.00 1.00 0.00 0.00 1.00 44.00│
└                                                       ┘
*/
