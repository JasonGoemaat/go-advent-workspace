package main

import (
	"fmt"
	"math"
	"testing"
)

// [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}

func TestMatrixCreation(t *testing.T) {
	m := NewIntMatrix(6, 5)
	m.Set(0, 0, 1) // joltage 0, button 0
	m.Set(1, 0, 1) // joltage 1, button 0
	m.Set(2, 0, 1)
	m.Set(3, 0, 1)
	m.Set(4, 0, 1)
	m.Set(0, 1, 1) // button 1, joltage 0
	m.Set(3, 1, 1)
	m.Set(4, 1, 1)
	m.Set(0, 2, 1) // button 2, joltage 0
	m.Set(1, 2, 1)
	m.Set(2, 2, 1)
	m.Set(4, 2, 1)
	m.Set(5, 2, 1)
	m.Set(1, 3, 1)
	m.Set(2, 3, 1)
	m.Set(0, 4, 10)
	m.Set(1, 4, 11)
	m.Set(2, 4, 11)
	m.Set(3, 4, 5)
	m.Set(4, 4, 10)
	m.Set(5, 4, 5)
	fmt.Println(m)

	fmt.Printf("\nRemoving duplicate rows...\n")
	m.RemoveDuplicateRows()
	fmt.Println(m)

	fmt.Printf("\nSorting...\n")
	m.Sort()
	fmt.Println(m)
}

// [..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}
func TestMatrixSolve(t *testing.T) {
	// input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
	input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
	machine := ParseMachine2(input)
	matrix := NewIntMatrix(len(machine.Joltages), len(machine.Buttons)+1)
	for i, button := range machine.Buttons {
		for _, joltageIndex := range button {
			matrix.Set(joltageIndex, i, 1)
		}
	}
	for i, joltage := range machine.Joltages {
		matrix.Set(i, len(machine.Buttons), joltage)
	}
	matrix.Report("Initial matrix")
	matrix.RemoveDuplicateRows()
	matrix.Report("After removing duplicate rows")
	matrix.SortColumnsForPivot(0)
	matrix.Report("After sorting columns by fewest values")
	matrix.Sort()
	matrix.Report("After sorting rows")
	matrix.SwapRows(0, 1)
	matrix.AddRows(0, 1, -1)
	matrix.Sort()
	matrix.Report("Swapped rows 0 and 1, subtracted row 0 from row 1, sorted rows")
	matrix.SwapRows(1, 3)
	matrix.AddRows(1, 2, -1)
	matrix.AddRows(1, 3, -1)
	matrix.Sort()
	matrix.Report("Swapped rows 1 and 3, subtracted row 1 from rows 2,3, sorted rows")
	matrix.SwapRows(4, 2)
	matrix.AddRows(2, 3, -1)
	matrix.AddRows(2, 4, -1)
	matrix.AddRows(2, 5, -1)
	matrix.Sort()
	matrix.Report("Swapped rows 4 and 2, subtracted row 2 from rows 3,4,5, sorted rows")
	matrix.MultiplyRow(8, -1)
	matrix.Sort()
	matrix.Report("Negated row 8 because it had a leading -1, resorted")
	matrix.SwapRows(4, 3)
	matrix.AddRows(3, 4, -1)
	matrix.AddRows(3, 5, -1)
	matrix.Sort()
	matrix.Report("Swapped rows 4 and 3, subtracted row 3 from rows 4 and row 5, sorted rows")
	matrix.MultiplyRow(8, -1)
	matrix.Sort()
	matrix.Report("Negated row 8 because it had a leading -1, resorted")
	matrix.SwapRows(4, 6)
	matrix.AddRows(4, 5, -1)
	matrix.AddRows(4, 6, -1)
	matrix.Sort()
	matrix.Report("Swapped rows 4 and 6, subtracted row 4 from rows 5 and row 6, sorted rows")
	matrix.SwapRows(5, 7)
	matrix.AddRows(5, 6, -1)
	matrix.AddRows(5, 7, -1)
	matrix.AddRows(5, 8, -1)
	matrix.Sort()
	matrix.Report("Swapped rows 5 and 7, subtracted row 5 from rows 5,6,7, sorted rows")
	matrix.MultiplyRow(8, -1)
	matrix.Sort()
	matrix.Report("Negated row 8 because it had a leading -2, resorted")
	matrix.AddRows(7, 6, -2)
	matrix.AddRows(7, 8, -1)
	matrix.Sort()
	matrix.Report("Pick row 7 (leading 1, followed by 0,0), subtract 2x from 6 and 1x from 7, resort")
	// thinking about that last one, no need to swap, just adjust and sorting will handle it
	fmt.Println("================================================================================")
	fmt.Println("Now have pivots (1 values) in rows 0-6 in columns 0-6 of the matching row")
	matrix.AddRows(6, 4, -1)
	matrix.AddRows(6, 3, 1)
	matrix.AddRows(6, 2, -1)
	matrix.Report("Subtracted row 6 from rows 2 and 4, added to row 3, all 0s above G6 pivot and by definition below")
	matrix.AddRows(5, 2, -1)
	matrix.AddRows(5, 0, -1)
	matrix.Report("Subtracted row 5 from rows 0 and 2")
	matrix.AddRows(4, 2, -1)
	matrix.Report("Subtracted row 4 from row 2")
	matrix.AddRows(3, 1, -1)
	matrix.Report("Subtracted row 3 from row 1")
	fmt.Println("================================================================================")
	fmt.Println(`LOOKING at the matrix, I see some interesting constraints:
	1. Joltage 0 has a total of 0, meaning button I cannot be pressed
		This row could then be removed and treated separately
	2. I can brute-force or recurse values for H, A, and D
	3. D has to be at LEAST 6 to make the -12 in row 1 viable
	4. D has to be at MOST 22 to make the 66 in row 3 viable
	5. D has to be at MOST 17 to make row 6 viable
	6. D has to be at MOST 19 to make row 8 viable
	7. Therefore D is constrained between 6 and 19
	8. Likewise A is constrained by row 3 to <=66, row 5 to <=25 and row 8 to <=77
	9. Therefore A is between 0 and 25
	10. Likewise H is constrained by rows 3,5,8 and must be <= 25
	11. So the three variable buttons are 6-19, 0-25, and 0-25, or 9464 possible combinations
	11. Row 4 shows 13 and has a single entry for button G, so G (button 6)=13
		This row could then be removed and treated separately`)
	fmt.Println("================================================================================")
	fmt.Println(`But checking, column H cannot have a pivot because the values are 2 and -2
	and the total is not even, HOWEVER, D7 and A8 have -1s, and row 7 had a 0 in A,
	so let me try swapping D and H and negating row 7`)
	matrix.SwapColumns(7, 9)
	matrix.Report("Swapped columns 7 and 9")
	matrix.MultiplyRow(7, -1)
	matrix.Report("Negated row 7")
	matrix.AddRows(7, 8, 4)
	matrix.Report("Added 4 times row 7 to row 8")
	matrix.AddRows(7, 6, -1)
	matrix.AddRows(7, 3, -3)
	matrix.AddRows(7, 2, 2)
	matrix.AddRows(7, 1, 2)
	matrix.Report("Added some factor of row 7 to rows 1,2,3,6 to make it a good pivot")
	fmt.Println("Now I see H is constrained to 6<=H<=12 by 7 and 8")
	matrix.MultiplyRow(8, -1)
	matrix.AddRows(8, 5, -1)
	matrix.AddRows(8, 3, -1)
	matrix.AddRows(8, 2, 1)
	matrix.Report("Negated row 8 to make it a +1 pivot and fixed rows above to have 0 in column A")
	fmt.Println("Now I have only one button (Button H) that I have to check!")
	fmt.Println("And I know from rows 7 and 8 that it is between 6 and 12!")
	fmt.Println("I'm not sure if I can ALWAYS count on that being the case...")
	fmt.Println("It may not NECESSARILY be a 1 though, it would work for any value")
	fmt.Println("    IF the values above are divisible by it")
	fmt.Println("I was thinking of looking at factors for a row, but the fact that each row")
	fmt.Println("    starts with a '1' means the row isn't divisible by anything else")
	fmt.Println(`IDEA! - since leading 0s after pivot help prevent negatives when subtracting
	from the rows below to make them 0, maybe I sort the columns at that point to put the most
	zeroes after the one each time to start with? CHECK NEXT TEST
	IT may be useful to just start looking at how many 0s are in each row and
	then swapping columns so that there is a leading '1', followed by
	all the zeros then the remaining 1s`)
}

// [..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}
func TestMatrixSolve2(t *testing.T) {
	// input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"

	// I think this should be 80, I get 78?
	input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"

	// result should be 331 according to other solution
	// dang, I calculate 333, but it should be 331?
	// input := "[###...##..] (2,6) (0,2,3,4,5,6,7,8,9) (0,1,2,5,6) (0,1,2,3,4,7,8,9) (0,1,4,5,6,9) (3,5) (0,2,9) (0,1,2,4,6,7) (6,9) (0,1,8,9) (0,1,2,3,4,5,8) {287,74,298,75,76,77,83,57,55,243}"
	machine := ParseMachine2(input)
	matrix := NewIntMatrix(len(machine.Joltages), len(machine.Buttons)+1)
	for i, button := range machine.Buttons {
		for _, joltageIndex := range button {
			matrix.Set(joltageIndex, i, 1)
		}
	}
	for i, joltage := range machine.Joltages {
		matrix.Set(i, len(machine.Buttons), joltage)
	}
	matrix.Report("Initial matrix")
	matrix.RemoveDuplicateRows()
	matrix.Report("After removing duplicate rows")
	matrix.ConstrainPresses()
	matrix.ReportConstraints()

	for i := 0; i < matrix.N-1; i++ {
		success := matrix.FindPivot(i)
		matrix.Report(fmt.Sprintf("Found pivot %d?  %v", i, success))
		if !success {
			break
		}
		matrix.ConstrainPresses()
		matrix.ReportConstraints()
	}
	matrix.Reduce()
	matrix.Report("REDUCED!")
	matrix.ConstrainPresses()
	matrix.ReportConstraints()

	// SWEET - now I need to brute-force columns past IntMatrix.MaxPivotColumn
	// IF there are no more column, we can calculate an answer directly
	// Calculating will take modified joltages after those presses (if any).
	// Any columns without a pivot should be 0 after the presses.
	// Any pivot columns should have a '1' in the pivot and presses are equal to the
	// modified joltages.   At this point there should be no negatives.
	minPresses := 0
	// calculate := func(remaining []int) (int, bool) {
	// 	presses := 0
	// 	for i := 0; i <= matrix.MaxPivotColumn; i++ {
	// 		if remaining[i] < 0 {
	// 			return 0, false
	// 		}
	// 		presses += remaining[i]
	// 	}
	// 	return presses, true
	// }

	// TODO: Thought is to use for all columns, going right to left and using
	// special logic for pivot columns that does not create a new joltages
	// array and instead adds specified presses

	presses := make([]int, matrix.N-1)

	// return the minimum number of presses if successful, and a flag indicating success
	var recurse func([]int, int) (int, bool)
	recurse = func(remaining []int, column int) (int, bool) {
		// past end of buttons to brute force, a success means joltages are positive
		// and each one is one press of one of the pivot buttons
		if column > matrix.N-2 {
			// we're after the brute-force columns, remaining joltages cannot
			// be negative for a valid solution, and each value is a single press
			// of a pivot button
			finalPresses := 0
			for i, joltage := range remaining {
				if joltage < 0 {
					return math.MaxInt64, false
				}
				presses[i] = joltage
				finalPresses += joltage
			}
			return finalPresses, true
		}

		// copy remaining joltages, add one value at a time as we loop through possible presses
		// for this column
		minPresses := math.MaxInt64
		joltages := make([]int, len(remaining))
		copy(joltages, remaining)
		for current := matrix.MinPresses[column]; current <= matrix.MaxPresses[column]; current++ {
			for i := 0; i < matrix.M; i++ {
				joltages[i] -= matrix.Value(i, column)
			}
			presses[column] = current
			futurePresses, success := recurse(joltages, column+1)
			if success {
				presses := futurePresses + current
				if presses < minPresses {
					minPresses = presses
				}
			}
		}
		return minPresses, minPresses < math.MaxInt64
	}

	remaining := make([]int, matrix.M)
	for i := range matrix.M {
		remaining[i] = matrix.Value(i, matrix.N-1)
	}
	minPresses, success := recurse(remaining, matrix.MaxPivotColumn+1)
	if !success {
		panic("solution not found")
	}
	fmt.Printf("Min presses: %d\n", minPresses)
}

func (im *IntMatrix) Report(action string) {
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(action)
	fmt.Println(im)
}

func TestMatrixSolve3(t *testing.T) {
	// input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
	input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
	machine := ParseMachine2(input)
	matrix := NewIntMatrix(len(machine.Joltages), len(machine.Buttons)+1)
	for i, button := range machine.Buttons {
		for _, joltageIndex := range button {
			matrix.Set(joltageIndex, i, 1)
		}
	}
	for i, joltage := range machine.Joltages {
		matrix.Set(i, len(machine.Buttons), joltage)
	}
	matrix.Report("Initial matrix")
	matrix.RemoveDuplicateRows()
	matrix.Report("RemoveDuplicateRows() called")
	for pivotColumn := 0; pivotColumn <= matrix.MaxPivotColumn; pivotColumn++ {
		fmt.Printf("========== processing pivotColumn %d ==========\n", pivotColumn)
		// sort columns starting at pivotColumn to the end so that pivotColumn has the most zeroes
		matrix.SortColumnsForPivot(pivotColumn)
		matrix.Report(fmt.Sprintf("SortColumnsForPivot(%d) called", pivotColumn))

		// sort rows starting at pivotColumn so that -1s are changed to 1s and 1s are first,
		// followed by 0s, then other values
		matrix.SortPotentialPivotRows(pivotColumn)
		matrix.Report(fmt.Sprintf("SortPotentialPivotRows(%d) called", pivotColumn))

		// the pivot should now be 1, if not we are done
		if matrix.Value(pivotColumn, pivotColumn) != 1 {
			matrix.MaxPivotColumn = pivotColumn - 1
			continue // should break
		}

		// fort columns AFTER pivotColumn so that the pivot rows has a '1' then zeroes, then other
		matrix.SortColumnsByPivotZeroes(pivotColumn)
		matrix.Report(fmt.Sprintf("SortColumnsByPivotZeroes(%d) called", pivotColumn))

		// ensure all zeroes in pivot column below pivot row
		matrix.ReduceBelowPivot(pivotColumn)
		matrix.Report(fmt.Sprintf("ReduceBelowPivot(%d) called", pivotColumn))
	}

	matrix.Reduce()
	matrix.Report("Reduce() called")
}
