package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// An IntMatrix struct consists of 'm' rows and 'n' columns
// with integer values initialized to 0.
type IntMatrix struct {
	M, N                int
	Data                []int
	Identifiers         []string
	RowIdentifiers      []string
	MaxPivotColumn      int
	OriginalButtonIndex []int
	MinPresses          []int
	MaxPresses          []int
	OriginalRowIndex    []int
}

func (im *IntMatrix) Value(r, c int) int {
	index := r*im.N + c
	if index >= len(im.Data) {
		panic("OUT OF RANGE")
	}
	return im.Data[index]
}

func (im *IntMatrix) Set(r, c, value int) {
	im.Data[r*im.N+c] = value
}

func NewIntMatrix(m, n int) *IntMatrix {
	data := make([]int, m*n)
	identifiers := make([]string, n)
	identifiers[n-1] = "TOTAL"
	for i := 0; i < n-1; i++ {
		identifiers[i] = string(rune('A' + i))
	}
	rowIdentifiers := make([]string, m)
	for i := 0; i < m; i++ {
		rowIdentifiers[i] = fmt.Sprintf("%2s", fmt.Sprintf("%d", i))
	}
	originalButtonIndex := make([]int, n-1)
	maxPresses := make([]int, n-1)
	minPresses := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		originalButtonIndex[i] = i
		maxPresses[i] = math.MaxInt64
		minPresses[i] = 0
	}
	originalRowIndex := make([]int, m)
	for i := 0; i < m; i++ {
		originalRowIndex[i] = i
	}
	return &IntMatrix{m, n, data, identifiers, rowIdentifiers, n - 2, originalButtonIndex, minPresses, maxPresses, originalRowIndex}
}

// FindMaxPresses will set the values in the MaxPresses array if called
// after initializing the matrix.
func (im *IntMatrix) FindMaxPresses() {
	for i := 0; i < im.N-1; i++ {
		for j := 0; j < im.M; j++ {
			if im.Value(j, i) > 0 && im.Value(j, im.N-1) < im.MaxPresses[i] {
				im.MaxPresses[i] = im.Value(j, im.N-1)
			}
		}
	}
}

func (im *IntMatrix) ReportConstraints() {
	for i := 0; i < len(im.MinPresses); i++ {
		fmt.Printf("Button %s: MinPresses = %d, MaxPresses = %d\n", im.Identifiers[i], im.MinPresses[i], im.MaxPresses[i])
	}
}

// Possibly lower MaxPresses based on current matrix state.
// 1. all numbers are same sign (+ or -) then max is TOTAL/value in that column
// 2. number and total are same sign, ALL others are different sign or 0 then MIN is (total+value-1)/value
func (im *IntMatrix) ConstrainPresses() {
	signOf := func(value int) (int, bool) {
		if value == 0 {
			return 0, false
		}
		if value < 0 {
			return -1, true
		}
		return 1, true
	}
	for i := 0; i < im.N-1; i++ { // i is 0 to last column before totals
		for j := 0; j < im.M; j++ { // j is 0 to last row
			sign, check := signOf(im.Value(j, i))
			if !check || sign == 0 {
				continue // this value does not constrain the result, should be both or neither
			}

			signOfAnswer, checkAnswer := signOf(im.Value(j, im.N-1))
			if !checkAnswer || sign != signOfAnswer {
				continue // sign must be the same as answer to constrain
			}

			allSame := true         // initialize to true, set to false if a difference is found
			othersDifferent := true // initialize to true, set to false if non-zero opposite sign detected
			for k := 0; k < im.N-1; k++ {
				if i == k {
					continue // ignore our own cell
				}
				otherSign, otherCheck := signOf(im.Value(j, k))
				if !otherCheck {
					continue // other is 0, meaningless
				}
				if otherSign == sign {
					othersDifferent = false // same sign, not all others are differeent
				} else {
					allSame = false // different sign, not all are the same
				}
			}

			if allSame {
				// if all are the same, max is constrained by how many times this
				// button can be pressed without overflowing
				max := im.Value(j, im.N-1) / im.Value(j, i) // joltage divided by joltage set from this button
				if max < im.MaxPresses[i] {
					im.MaxPresses[i] = max
				}
			} else if othersDifferent {
				// sign is same as answer and others are different, NEEDs to be pressed at least enough times
				// to get the joltage, others can be pressed more to reverse it
				min := (im.Value(j, im.N-1) + im.Value(j, i) - 1) / im.Value(j, i)
				if min > im.MinPresses[i] {
					im.MinPresses[i] = min
				}
			}
		}
	}
}

// for pivotCount pivot, sort columns after to place zeroes first
func (im *IntMatrix) SortColumnsByPivotZeroes(pivotCount int) {
	columnLess := func(c1, c2 int) bool {
		v1 := im.Value(pivotCount, c1)
		if v1 == 0 {
			return true
		}
		return false
	}
	for i := pivotCount + 1; i <= im.MaxPivotColumn-1; i++ {
		least := i
		for j := i + 1; j < im.MaxPivotColumn; j++ {
			if columnLess(j, least) {
				least = j
			}
		}
		if least != i {
			im.SwapColumns(i, least)
		}
	}
}

// Sort potential pivot columns, finding one with at least one 1
// and then what?   Most zeroes?
func (im *IntMatrix) SortColumnsForPivot(pivotCount int) {
	columnLess := func(c1, c2 int) bool {
		t1, t2 := 0, 0
		potential1, potential2 := false, false
		for j := pivotCount; j < im.M; j++ {
			v1 := im.Value(j, c1)
			v2 := im.Value(j, c2)
			if j == pivotCount && (v1 == 1 || v1 == -1) {
				potential1 = true
			}
			if j == pivotCount && (v2 == 1 || v2 == -1) {
				potential2 = true
			}
			if v1 == 0 {
				t1++
			}
			if v2 == 0 {
				t2++
			}
		}
		if potential1 && !potential2 {
			return true
		}
		return t1 > t2 // most zeroes
	}
	for i := pivotCount; i < im.N-2; i++ {
		least := i
		for j := i + 1; j < im.N-1; j++ {
			if columnLess(j, least) {
				least = j
			}
		}
		if least != i {
			im.SwapColumns(i, least)
		}
	}
}

// for column pivotCount, sort rows so that 1s and -1s appear before other
// values, and they are ordered so that the row with the most zeros is at
// the top
func (im *IntMatrix) SortPotentialPivotRows(pivotCount int) {
	rowLess := func(r1, r2 int) bool {
		value1 := im.Value(r1, pivotCount)
		if value1 == -1 {
			im.MultiplyRow(r1, -1)
			value1 = 1
		}
		value2 := im.Value(r2, pivotCount)
		if value2 == -1 {
			im.MultiplyRow(r2, -1)
			value2 = 1
		}
		if value1 == 1 && value2 != 1 {
			return true
		}
		if value1 != 1 && value2 == 1 {
			return false
		}
		if value1 == 1 && value2 == 1 {
			// order by most zeroes after pivot
			zeroCount1 := 0
			zeroCount2 := 0
			for i := pivotCount + 1; i < im.N-1; i++ {
				if im.Value(r1, i) == 0 {
					zeroCount1++
				}
				if im.Value(r2, i) == 0 {
					zeroCount2++
				}
			}
			if zeroCount1 > zeroCount2 {
				return true
			}
			if zeroCount1 < zeroCount2 {
				return false
			}
		}
		if value1 == 0 && value2 == 0 {
			return false // no less so don't swap
		}
		return value1 == 0 // zeroes go between 1,-1 values and non-zero values at bottom
	}
	for i := pivotCount; i < im.M-1; i++ {
		least := i
		for j := i + 1; j < im.M; j++ {
			if rowLess(j, least) {
				least = j
			}
		}
		if least != i {
			im.SwapRows(i, least)
		}
	}

	// get range of potential pivots (rows with value 1 in pivotCount column)
	minRow := pivotCount
	maxRow := minRow - 1
	for i := minRow; i < im.M; i++ {
		if im.Value(i, pivotCount) == 1 {
			maxRow = i
		}
	}

	// sort remaining button columns by zero counts
	countZeroes := func(column int) int {
		count := 0
		for i := minRow; i <= maxRow; i++ {
			if im.Value(i, column) == 0 {
				count++
			}
		}
		return count
	}

	columnLess := func(c1, c2 int) bool {
		count1 := countZeroes(c1)
		count2 := countZeroes(c2)
		return count1 < count2
	}

	for column := pivotCount + 1; column < im.N-2; column++ {
		least := column
		for j := column + 1; j < im.N-1; j++ {
			if columnLess(column, column+1) {
				least = j
			}
		}
		if least != column {
			im.SwapColumns(column, least)
		}
	}
}

func (im *IntMatrix) ReduceBelowPivot(pivotCount int) {
	for i := pivotCount + 1; i < im.M; i++ {
		value := im.Value(i, pivotCount)
		im.AddRows(pivotCount, i, -value)
	}
}

func (im *IntMatrix) RemoveDuplicateRows() {
	areRowsEqual := func(r1, r2 int) bool {
		for c := 0; c < im.N; c++ {
			if im.Value(r1, c) != im.Value(r2, c) {
				return false
			}
		}
		return true
	}
	for i := 0; i < im.M-1; i++ {
		for j := i + 1; j < im.M; {
			if areRowsEqual(i, j) {
				// j is a dupe, swap j with last row and decrease size
				lastRow := im.M - 1
				// CHANGED to use SwapRows to preserve original index and identifiers
				// for c := 0; c < im.N; c++ {
				// 	im.Set(j, c, im.Value(lastRow, c))
				// }
				im.SwapRows(j, lastRow)
				im.M--
				im.Data = im.Data[:im.M*im.N]
			} else {
				j++
			}
		}
	}
}

// Sort will order the rows to put zeros in the bottom-left of the matrix.
// This proceeds looking at each column from left to right and ordering the
// rows by value descending.   We proceed to the next column only looking at
// the range of rows from the bottom up that have zero values.
func (im *IntMatrix) Sort() {
	var rowGreater = func(r1, r2 int) bool {
		for c := 0; c < im.N; c++ {
			v1 := im.Value(r1, c)
			v2 := im.Value(r2, c)
			if v1 > v2 {
				return true
			}
			if v2 > v1 {
				return false
			}
		}
		return false
	}

	for i := 0; i < im.M-1; i++ {
		largest := i
		for j := i + 1; j < im.M; j++ {
			if rowGreater(j, largest) {
				largest = j
			}
		}
		if largest != i {
			// fmt.Printf("Swapping row %d with largest %d\n", i, largest)
			im.SwapRows(i, largest)
			// fmt.Println(im)
		}
	}
}

// SwapRows will do just that, swap any two rows
func (im *IntMatrix) SwapRows(r1, r2 int) {
	im.RowIdentifiers[r1], im.RowIdentifiers[r2] = im.RowIdentifiers[r2], im.RowIdentifiers[r1]
	for i := 0; i < im.N; i++ {
		im.Data[r1*im.N+i], im.Data[r2*im.N+i] = im.Data[r2*im.N+i], im.Data[r1*im.N+i]
	}
}

// SwapColumns will do just that, swap any two columns
func (im *IntMatrix) SwapColumns(c1, c2 int) {
	for i := 0; i < im.M; i++ {
		im.Data[i*im.N+c1], im.Data[i*im.N+c2] = im.Data[i*im.N+c2], im.Data[i*im.N+c1]
	}
	im.Identifiers[c1], im.Identifiers[c2] = im.Identifiers[c2], im.Identifiers[c1]
	im.OriginalButtonIndex[c1], im.OriginalButtonIndex[c2] = im.OriginalButtonIndex[c2], im.OriginalButtonIndex[c1]
	im.MaxPresses[c1], im.MaxPresses[c2] = im.MaxPresses[c2], im.MaxPresses[c1]
	im.MinPresses[c1], im.MinPresses[c2] = im.MinPresses[c2], im.MinPresses[c1]
}

// MultiplyRow will multiply all values in the row by the given factor
func (im *IntMatrix) MultiplyRow(r1 int, factor int) {
	for i := 0; i < im.N; i++ {
		im.Data[r1*im.N+i] *= factor
	}
}

// AddRows adds the source row's values to the destination row
func (im *IntMatrix) AddRows(source, destination int, factor int) {
	for i := 0; i < im.N; i++ {
		im.Data[destination*im.N+i] += im.Data[source*im.N+i] * factor
	}
}

func (im *IntMatrix) Reduce() {
	for i := 1; i <= im.MaxPivotColumn; i++ {
		// pivot Value(i,i) is always 1, add/substract from higher rows
		// to give 0s in that column
		for j := 0; j < i; j++ {
			value := im.Value(j, i)
			im.AddRows(i, j, -value)
		}
	}
}

// FindPivot will find a pivot value '1' for column 'index', moving
// columns around if necessary.   It then changes the rows below to
// make sure they have zeroes in that column.  It returns true if it
// was successful, or false if no pivot coult be found.
func (im *IntMatrix) FindPivot(index int) bool {
	// find the column with the most zeros on or below pivot cell
	// and swap to the index if different
	zeroCounts := make([]int, im.N)
	least := index
	for i := index; i <= im.MaxPivotColumn; i++ {
		for j := index; j < im.M; j++ {
			if im.Value(j, i) == 0 {
				zeroCounts[i]++
			}
		}
	}
	for i := index + 1; i < im.N-1; i++ {
		if zeroCounts[i] < zeroCounts[least] {
			least = i
		}
	}
	if least != index {
		im.SwapColumns(index, least)
	}

	// get three groups of rows on or below pivot:
	// 1. zeros
	// 2. ones or negative ones
	// 3. others
	zeroes := make([]int, 0, im.M-index)
	ones := make([]int, 0, im.M-index)
	others := make([]int, 0, im.M-index)
	for i := index; i < im.M; i++ {
		if im.Value(i, index) == 0 {
			zeroes = append(zeroes, i)
		} else if im.Value(i, index) == 1 || im.Value(i, index) == -1 {
			ones = append(ones, i)
		} else {
			others = append(others, i)
		}
	}
	if len(ones) == 0 {
		// note: this might happen, if so, we could move this column to the end
		// and try again, in fact that is what my first manual test does...
		// panic("NO ONES!")
		if index <= im.MaxPivotColumn {
			// no pivotable value in this column, move to end and freeze,
			// then call again
			im.SwapColumns(index, im.MaxPivotColumn)
			im.MaxPivotColumn--
			return im.FindPivot(index)
		}
		return false
	}

	// TODO: PICK pivot with the most zeroes instead of the first?
	// Or use some algorithm to prefer pivots that will keep the most
	// values below in the range -1 to 1?
	pivotRow := ones[0]

	// convert -1 to +1
	if im.Value(pivotRow, index) == -1 {
		im.MultiplyRow(pivotRow, -1)
	}

	// move pivot row into position
	if pivotRow != index {
		im.SwapRows(pivotRow, index)
	}

	// fix remaining rows so that they have zeroes
	for i := index + 1; i < im.M; i++ {
		value := im.Value(i, index)
		if value == 0 {
			continue
		}
		im.AddRows(index, i, -value) // will change pivot column in that row to 0
	}
	return true
}

func (im *IntMatrix) String() string {
	rows := make([]string, im.M+2) // extra rows for first and last
	cells := make([]string, len(im.Data))
	maxColumnLengths := make([]int, im.N)
	for i, value := range im.Data {
		cells[i] = strconv.Itoa(value)
		length := len(cells[i])
		col := i % im.N
		if length > maxColumnLengths[col] {
			maxColumnLengths[col] = length
		}
	}

	// include identifier length in calculation
	for i, maxLength := range maxColumnLengths {
		if len(im.Identifiers[i]) > maxLength {
			maxColumnLengths[i] = len(im.Identifiers[i])
		}
	}

	for i, value := range cells {
		length := len(value)
		diff := maxColumnLengths[i%im.N] - length
		if diff > 0 {
			cells[i] = strings.Repeat(" ", diff) + cells[i]
		}
	}
	totalRowLength := len(maxColumnLengths) - 1 // separating spaces
	for _, l := range maxColumnLengths {
		totalRowLength += l
	}
	identifierCells := make([]string, im.N)
	for i, identifier := range im.Identifiers {
		identifierCells[i] = strings.Repeat(" ", maxColumnLengths[i]-len(identifier)) + identifier
	}
	rows[0] = "  ┌" + strings.Join(identifierCells, " ") + "┐"
	rows[len(rows)-1] = fmt.Sprintf("  └%s┘", strings.Repeat(" ", totalRowLength))
	for i := 0; i < im.M; i++ {
		rows[i+1] = fmt.Sprintf("%s│%s│", im.RowIdentifiers[i], strings.Join(cells[i*im.N:(i+1)*im.N], " "))
	}
	return strings.Join(rows, "\n")
}

/*

    top := "┌───┐"
    mid := "│   │"
    bot := "└───┘"

Row Echelon Form and Reduced Row Echelon Form

REF:
	Can have non-1 pivots
	Can have non-zero entries above pivots
	Not unique
	Used for quick solutions via back-substitution.

RREF:
	Pivots are 1
	zeros everywhere else in pivot columns
	Unique
	Directly reveals solutions (like identity matrices).



https://en.wikipedia.org/wiki/Matrix_(mathematics)
	A matrix with m rows and n columns is called an m × n matrix
	So m is rows and n is columns

For sorting, I might want to start with leading 1s, but then prefer 0s
for other columns.

## IDEA -

Looking at floating point sample and going all the way, I ended up
when I had 10 buttons and 9 joltages with 7 joltages I could form
in REF, so I could simplify 7 buttons.   I could iterate the other
three.   Picking the optimum joltage to iterate over first MAY
help.   Here it was the one I could eliminate, the only possible
value for button 8 is 0 because it is the only button affecting joltage
7 AFTER REDUCTION.

I wonder how well this works with all machines...

Iteration:

Step 1: Find all pivots we can have as a '1'.  Take this example:
 ┌                                                       ┐
0│1.00 1.00 1.00 1.00 0.00 1.00 1.00 0.00 0.00 0.00 53.00│
1│1.00 1.00 1.00 0.00 1.00 1.00 1.00 0.00 0.00 1.00 58.00│
2│1.00 1.00 0.00 1.00 0.00 0.00 0.00 1.00 0.00 1.00 54.00│
3│1.00 0.00 0.00 1.00 1.00 1.00 1.00 1.00 1.00 0.00 55.00│
4│1.00 0.00 0.00 1.00 1.00 1.00 1.00 1.00 0.00 0.00 55.00│
5│1.00 0.00 0.00 1.00 1.00 1.00 0.00 1.00 0.00 0.00 42.00│
6│1.00 0.00 0.00 0.00 0.00 1.00 0.00 1.00 1.00 0.00 25.00│
7│0.00 1.00 1.00 1.00 1.00 0.00 0.00 1.00 0.00 0.00 43.00│
8│0.00 0.00 1.00 1.00 1.00 0.00 1.00 0.00 0.00 1.00 44.00│
 └                                                       ┘

Pick out pivots - Rows 0-6 work for 0.   Row 7 works for 1.
Row 8 works for 2.

Here I ***MIGHT***, that is ***MIGHT*** be able to optimize
picking wich row to ***USE*** for pivot 0.  I prefer rows where
other columns are 0.   Then when I subtract, I'm less likely
to get negative numbers.

***ACTUALLY***, re-arranging columns might help as well.
In this example I have pivots for 0, 1, and 2.  This would
equate to having buttons with the fewest joltages first.
In this example I've re-ordered the columns by number of
1s in them, and sorted the resulting possible pivots to the
top.   For instance column I has 1 in rows 3 and 6 so I moved
them to the top.

 ┌ I    J    B    C    G    E    F    H    D    A   JOLTS┐
3│1.00 0.00 0.00 0.00 1.00 1.00 1.00 1.00 1.00 1.00 55.00│
6│1.00 0.00 0.00 0.00 0.00 0.00 1.00 1.00 0.00 1.00 25.00│
1│0.00 1.00 1.00 1.00 1.00 1.00 1.00 0.00 0.00 1.00 58.00│
2│0.00 1.00 1.00 0.00 0.00 0.00 0.00 1.00 1.00 1.00 54.00│
8│0.00 1.00 0.00 1.00 1.00 1.00 0.00 0.00 1.00 0.00 44.00│
0│0.00 0.00 1.00 1.00 1.00 0.00 1.00 0.00 1.00 1.00 53.00│
7│0.00 0.00 1.00 1.00 0.00 1.00 0.00 1.00 1.00 0.00 43.00│
4│0.00 0.00 0.00 0.00 1.00 1.00 1.00 1.00 1.00 1.00 55.00│
5│0.00 0.00 0.00 0.00 0.00 1.00 1.00 1.00 1.00 1.00 42.00│
 └                                                       ┘

Looking at rows 3 and 6, I would prefer to use row "6"
(now actually row 1) because they have the same values in
columns J, B, and C, but row "3" has a 1 in column G and
row "6" has a 0.   If I subtracted 3 from 6 I would get
a negative number left-most here.

*/
