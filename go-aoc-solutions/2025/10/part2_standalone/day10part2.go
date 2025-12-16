package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// --------------------------------------------------------------------------------
// Set 'explain' to true for lots of output showing matrix operations and to
// display all results.
//
// Set 'outputEach' to true and it will display line number and puzzle along with
// the resulting button press counts for each minimal solution.
//
// Set 'input' to a single puzzle, will use standard input if zero length string,
// so pipe your input into it or paste into terminal.
// --------------------------------------------------------------------------------
var explain = true
var outputEach = true
var input = ""

// var input = "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
// var input = "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
// var input = "[###..#...#] (3,6,7) (1,2,6) (0,2,3,4,5,6,9) (1) (0,1,2,5,6,7) (0,1,2,3,6,7,8,9) (0,1,2,3,5,6,8,9) (1,2,3,4,5,6,8,9) (0,5,6,8,9) (1,2,4,7,9) (0,3,8,9) (0,2,4,5,6,7,8) (2,3,5,6,8,9) {56,74,68,51,33,39,58,48,52,69}"

func main() {
	if len(input) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input += scanner.Text() + "\r\n"
		}
	}
	re := regexp.MustCompile("[^\r\n]+")
	lines := re.FindAllString(input, -1)
	sum := 0
	for i, line := range lines {
		if outputEach {
			fmt.Printf("Line %d: %s\n", i, line)
		}
		minPresses, solutions := solveLine(line)
		sum += minPresses
		if outputEach {
			for _, solution := range solutions {
				if solution.TotalPresses != minPresses {
					break
				}
				fmt.Printf("  %d presses: %v\n", solution.TotalPresses, solution.Presses)
			}
		}
	}
	fmt.Println()
	fmt.Printf("Result: %d\n", sum)
}

func solveLine(line string) (int, []MatrixSolution) {
	m := ParsePuzzle(line)
	if explain {
		fmt.Println("Initial matrix")
		fmt.Println(m)
	}
	m.RREFRecurse(0, 0, explain)
	solutions := m.Solve()
	return solutions[0].TotalPresses, solutions
}

//--------------------------------------------------------------------------------
// Base Matrix Operations
//--------------------------------------------------------------------------------

type Matrix struct {
	Rows, Cols   int
	Data         [][]int
	OriginalRows []int
	OriginalCols []int
	MaxPresses   []int
}

func (m *Matrix) SwapCols(c1, c2 int) {
	m.OriginalCols[c1], m.OriginalCols[c2] = m.OriginalCols[c2], m.OriginalCols[c1]
	m.MaxPresses[c1], m.MaxPresses[c2] = m.MaxPresses[c2], m.MaxPresses[c1]
	for r := range m.Rows {
		a, b := m.Get(r, c1), m.Get(r, c2)
		m.Set(r, c1, b)
		m.Set(r, c2, a)
	}
}

func NewMatrix(rows, cols int) *Matrix {
	data := make([][]int, rows)
	for i := range rows {
		data[i] = make([]int, cols)
	}
	originalRows := make([]int, rows)
	for i := range rows {
		originalRows[i] = i
	}
	originalCols := make([]int, cols)
	for i := range cols {
		originalCols[i] = i
	}
	maxPresses := make([]int, cols)
	return &Matrix{Rows: rows, Cols: cols, Data: data, OriginalRows: originalRows, OriginalCols: originalCols, MaxPresses: maxPresses}
}

func (m *Matrix) Get(row, col int) int {
	return m.Data[row][col]
}

func (m *Matrix) Set(row, col, value int) {
	m.Data[row][col] = value
}

func (m *Matrix) AddRow(source, dest, factor int) {
	for col := 0; col < m.Cols; col++ {
		m.Data[dest][col] += m.Data[source][col] * factor
	}
}

func (m *Matrix) MultiplyRow(row, factor int) {
	for col := 0; col < m.Cols; col++ {
		m.Data[row][col] *= factor
	}
}

func (m *Matrix) SwapRows(r1, r2 int) {
	m.Data[r1], m.Data[r2] = m.Data[r2], m.Data[r1]
	m.OriginalRows[r1], m.OriginalRows[r2] = m.OriginalRows[r2], m.OriginalRows[r1]
}

//--------------------------------------------------------------------------------
// Matrix solving - reduce matrix, constrain max presses, solve
//--------------------------------------------------------------------------------

type Solution struct {
	Presses []int // how many times each button is pressed
	Total   int   // total presses
	MS      int   // milliseconds
}

// Only look at matrix from row,col to the right and down.
// Any values in the range of rows to the left of col must
// be zero.
//
//  1. Find pivot value '1' or '-1' in 'col' between 'row' and bottom.
//     If -1, multiply row by -1 to make it 1
//     If not in 'row', swap the found row with 'row' so 'row','col' value is 1
//     Add multiple of 'row' to rows below to make sure they have 0 in 'col'
//     Call RREFRecurse(row+1, col+1)
//  2. If all values are 0, move on to next column calling RREFRecurse(row1, col + 1)
//  3. If there are no '1' or '-1', but there are higher values (i.e. '-3' or '2'), panic
func (m *Matrix) RREFRecurse(row, col int, explain bool) {
	m.ConstrainMaxPresses()
	// exit if we're past the last button column or the bottom row
	if row >= m.Rows || col >= m.Cols-1 {
		if explain {
			fmt.Printf("Explain %d,%d: done\n", row, col)
			fmt.Println(m)
		}
		return
	}
	allZero := true
	goodRow := -1
	for r := row; r < m.Rows; r++ {
		value := m.Get(r, col)
		if value == 1 {
			goodRow = r
			allZero = false
			break
		}
		if value == -1 {
			m.MultiplyRow(r, -1) // fix so it is +1
			if explain {
				fmt.Printf("Multiplying row %d by -1\n", r)
				fmt.Println(m)
			}
			goodRow = r
			allZero = false
			break
		}
		if value != 0 {
			allZero = false
		}
	}

	// happy case, we have goodRow
	if goodRow >= row {
		// not first row, so swap
		if goodRow > row {
			m.SwapRows(row, goodRow)
			if explain {
				fmt.Printf("Explain %d,%d: sawpping good row %d with expected %d\n", row, col, goodRow, row)
				fmt.Println(m)
			}
		}

		// fix-up non-zero rows below
		for r := row + 1; r < m.Rows; r++ {
			value := m.Get(r, col)
			if value != 0 {
				m.AddRow(row, r, -value)
				if explain {
					fmt.Printf("Explain %d,%d: Adding row %d with factor %d to row %d\n", row, col, row, -value, r)
					fmt.Println(m)
				}
			}
		}
		// fix-up rows above to change their values in col to 0
		for r := row - 1; r >= 0; r-- {
			value := m.Get(r, col)
			if value != 0 {
				m.AddRow(row, r, -value)
				if explain {
					fmt.Printf("Explain %d,%d: Adding row %d to row %d with factor %d\n", row, col, row, r, -value)
					fmt.Println(m)
				}
			}
		}

		// move to next row and col
		m.RREFRecurse(row+1, col+1, explain)
		return
	}

	// if all zeros, move on to next column, stay on current row
	if allZero {
		m.RREFRecurse(row, col+1, explain)
	}

	// not all zeros and no 1s found, we can't continue
	// TODO:  As an edge case, we may be able to continue if there exists
	// a '1' or '-1' *somewhere* in the submatrix, by just swapping columns
	// to put that '1' or '-1' in column 'col'

	// // TESTING: Move on to next row AND column - WORKS!
	// m.RREFRecurse(row+1, col+1, explain)

	// OPTIMIZING: if there is a 1 in a future column in any row, swap the columns and try again
	for r := row; r < m.Rows; r++ {
		for c := col + 1; c < m.Cols-1; c++ {
			if m.Get(r, c) == 1 || m.Get(r, c) == -1 {
				if explain {
					if explain {
						fmt.Printf("Explain %d,%d: swapping columns %d and %d to put a '1' first\n", row, col, col, c)
						fmt.Println(m)
					}
				}
				m.SwapCols(c, col)
				m.RREFRecurse(row, col, explain)
				return
			}
		}
	}
}

// Lower MaxPresses[].  Look for rows that are all one sign.  If all negative,
// multiply by -1 to make all positive. Constrain the buttons specified in
// that row because pressing them more would overflow the joltage for the row.
// TODO: Not currently calling, but it may be useful each time we RREFRecurse
// or just after doing all recursion.
// TODO: It may be helpful to try different things to, for example in the
// first sample I end up with button 0 being constrained to 7, but if I add
// row 3 to row 0, that will let me constrain it to 5.   Doesn't matter in
// this case because button 0 is not one of the buttons I need to brute-force.
func (m *Matrix) ConstrainMaxPresses() {
	for r := range m.Rows {
		foundNegative := false
		foundPositive := false
		for c := range m.Cols {
			value := m.Get(r, c)
			if value < 0 {
				foundNegative = true
			}
			if value > 0 {
				foundPositive = true
			}
		}
		// if found both, move on to next row
		if foundPositive && foundNegative {
			// TODO - if there is exactly 1 of a type and it matches sign of result column, it could be a minimum constraint
			break
		}
		if foundNegative && !foundPositive {
			// all negative, multiply by -1 to make all positive
			m.MultiplyRow(r, -1)
			foundPositive = true
			foundNegative = false
		}

		// all row values are the same sign, any buttons are constrained by the result
		answer := m.Get(r, m.Cols-1)
		for c := 0; c < m.Cols-1; c++ {
			value := m.Get(r, c)
			if value == 0 {
				continue
			}
			maxPresses := answer / value
			if m.MaxPresses[c] > maxPresses {
				m.MaxPresses[c] = maxPresses
			}
		}
	}
}

type MatrixSolution struct {
	Presses      []int // presses for each button
	TotalPresses int   // total presses - lower is better
}

func MatrixSolutionSortFunc(a, b MatrixSolution) int {
	if a.TotalPresses < b.TotalPresses {
		return -1
	}
	if a.TotalPresses > b.TotalPresses {
		return 1
	}
	return 0
}

func MatrixSolutionLess() {
	solutions := make([]MatrixSolution, 5)
	slices.SortFunc(solutions, MatrixSolutionSortFunc)
}

func (m *Matrix) Solve() []MatrixSolution {
	// first I need to figure out buttons that can be dependent, these have
	// a single '1' in the column and the rest are '0'.   These are the
	// pivots.  Also buttons that are variable, they affect multiple joltages.
	// I can brute-force the variable buttons first and then each dependent
	// button will have to be pressed the same number of times as the remaining
	// joltage for the answer.
	dependentButtons := make([]int, 0, m.Cols-1)
	dependentButtonJoltage := make([]int, 0, m.Cols-1)
	variableButtons := make([]int, 0, m.Cols-1)
	for c := range m.Cols - 1 {
		oneRow := -1
		for r := range m.Rows {
			if m.Get(r, c) == 1 {
				if oneRow == -1 {
					// first '1' we found, record location
					oneRow = r
				} else {
					// uh oh, there is more than one '1', break with failure
					oneRow = -1
					break
				}
			} else {
				if m.Get(r, c) != 0 && m.Get(r, c) != 1 {
					oneRow = -1
					break
				}
			}
		}
		if oneRow == -1 {
			variableButtons = append(variableButtons, c)
		} else {
			// button index
			dependentButtons = append(dependentButtons, c)
			// joltage affected by that button alone
			dependentButtonJoltage = append(dependentButtonJoltage, oneRow)
		}
	}

	// for each variable button, get a list of joltages affected by the button
	variableButtonJoltages := make([][]int, len(variableButtons))
	for i, c := range variableButtons {
		joltages := make([]int, 0, m.Rows)
		for r := range m.Rows {
			if m.Get(r, c) != 0 {
				joltages = append(joltages, r)
			}
		}
		variableButtonJoltages[i] = joltages
	}

	// now we permute the variable buttons through their ranges and try them all,
	// keeping track of presses which will be copied to solutions
	presses := make([]int, m.Cols-1)
	donePermuting := false
	permuteVariableButtons := func() bool {
		// increase from first to last, when last overflows we are done
		for _, c := range variableButtons {
			presses[c]++
			if presses[c] <= m.MaxPresses[c] {
				return true // no overflow, we are fine
			}
			presses[c] = 0
		}
		donePermuting = true
		return false
	}
	joltages := make([]int, m.Rows)
	permutationCount := 1
	for _, c := range variableButtons {
		permutationCount *= (m.MaxPresses[c] + 1)
	}

	solutions := make([]MatrixSolution, 0, permutationCount)
	for ; !donePermuting; permuteVariableButtons() {
		// initialize joltages
		for r := range m.Rows {
			joltages[r] = 0
		}
		for _, c := range dependentButtons {
			presses[c] = 0
		}
		for i, c := range variableButtons {
			if presses[c] != 0 {
				for _, r := range variableButtonJoltages[i] {
					joltages[r] += m.Get(r, c) * presses[c]
				}
			}
		}

		// check that joltages are less than required since remaining
		// dependant buttons can only add 1 per press
		valid := true
		for r, joltage := range joltages {
			if joltage > m.Get(r, m.Cols-1) {
				valid = false
				break
			}
		}

		if !valid {
			continue
		}

		// now dependent buttons will be 1:1
		for i, c := range dependentButtons {
			r := dependentButtonJoltage[i]
			p := m.Get(r, m.Cols-1) - joltages[r]
			joltages[r] += p
			presses[c] = p
		}

		// make sure all joltages now match - this may not be required...
		// if any joltages affected by ONLY variable buttons remain, that
		// would cause this.
		for r, joltage := range joltages {
			if joltage != m.Get(r, m.Cols-1) {
				valid = false
			}
		}

		if !valid {
			continue
		}

		// we have a solution
		totalPresses := 0
		for _, p := range presses {
			totalPresses += p
		}

		// put presses in original column (button) order to save with solution
		savePresses := make([]int, len(presses))
		for i, p := range presses {
			originalIndex := m.OriginalCols[i]
			savePresses[originalIndex] = p
		}
		solutions = append(solutions, MatrixSolution{savePresses, totalPresses})
	}

	slices.SortFunc(solutions, MatrixSolutionSortFunc)
	return solutions
}

// Simplified function to handle the parsing and boilerplate calls
// and return the minimum presses from all solutions we found.  For
// final result and individual testing when we just care about
// comparing one result to another method.
func SimpleSolve(input string) int {
	matrix := ParsePuzzle(input)
	matrix.RREFRecurse(0, 0, false)
	solutions := matrix.Solve()
	return solutions[0].TotalPresses
}

//--------------------------------------------------------------------------------
// Matrix helpers - parsing, tostring
//--------------------------------------------------------------------------------

func ParsePuzzle(input string) *Matrix {
	// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
	var rxButtons = regexp.MustCompile(`\(([^\)]*)\)`)
	var rxJoltages = regexp.MustCompile(`{[^}]*}`)

	found := rxButtons.FindAllStringSubmatch(input, -1)
	buttons := make([][]int, len(found))
	for i, submatch := range found {
		numberStrings := strings.Split(submatch[1], ",")
		value := make([]int, len(numberStrings))
		for j, numberString := range numberStrings {
			intValue, _ := strconv.Atoi(numberString)
			value[j] = intValue
		}
		buttons[i] = value
	}

	joltagesString := rxJoltages.FindString(input)
	joltagesString = joltagesString[1 : len(joltagesString)-1]
	joltagesStrings := strings.Split(joltagesString, ",")
	joltages := make([]int, len(joltagesStrings))
	for i, s := range joltagesStrings {
		joltages[i], _ = strconv.Atoi(s)
	}

	matrix := NewMatrix(len(joltages), len(buttons)+1)
	for x, button := range buttons {
		for _, joltage := range button {
			matrix.Set(joltage, x, 1)
		}
	}
	for y := range joltages {
		matrix.Set(y, len(buttons), joltages[y])
	}
	for c := range len(buttons) {
		max := math.MaxInt
		for r := range matrix.Rows {
			if matrix.Get(r, c) == 1 && joltages[r] < max {
				max = joltages[r]
			}
		}
		matrix.MaxPresses[c] = max
	}

	return matrix
}

func (m *Matrix) String() string {
	rows := make([]string, m.Rows+2) // extra rows for first and last
	cells := make([]string, m.Rows*m.Cols)
	maxColumnLengths := make([]int, m.Cols)

	// get string representation for each matrix cell and calculate maximum
	// lengths for each column to align them
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			index := i*m.Cols + j
			cells[index] = strconv.Itoa(m.Get(i, j))
			length := len(cells[index])
			if length > maxColumnLengths[j] {
				maxColumnLengths[j] = length
			}
		}
	}

	// set max column lengths for all to be the same, if we don't do this it
	// will be more compressed but maybe harder to read
	maxColumnLength := 2 // make at least 2 for negatives and double-digit column indexes
	for i := range maxColumnLengths {
		if maxColumnLengths[i] > maxColumnLength {
			maxColumnLength = maxColumnLengths[i]
		}
	}
	for i := range maxColumnLengths {
		maxColumnLengths[i] = maxColumnLength
	}

	// left-pad each cell for alignment
	for i, value := range cells {
		length := len(value)
		diff := maxColumnLengths[i%m.Cols] - length
		if diff > 0 {
			cells[i] = strings.Repeat(" ", diff) + cells[i]
		}
	}

	// calculate total row length
	totalRowLength := len(maxColumnLengths) - 1 // separating spaces
	for _, l := range maxColumnLengths {
		totalRowLength += l
	}

	// add column identifier cells for top row
	identifierCells := make([]string, m.Cols)
	for i := range m.Cols {
		identifier := strconv.Itoa(i)
		identifierCells[i] = strings.Repeat(" ", maxColumnLengths[i]-len(identifier)) + identifier
		if len(identifierCells[i]) > maxColumnLengths[i] {
			maxColumnLengths[i] = len(identifierCells[i])
		}
	}

	// add max press cells for bottom row
	maxPressCells := make([]string, m.Cols)
	for i := range m.Cols {
		s := strconv.Itoa(m.MaxPresses[i])
		maxPressCells[i] = strings.Repeat(" ", maxColumnLengths[i]-len(s)) + s
	}

	rows[0] = "      ┌" + strings.Join(identifierCells, " ") + "┐"
	rows[len(rows)-1] = fmt.Sprintf(" MaxP └%s┘", strings.Join(maxPressCells, " "))
	for i := 0; i < m.Rows; i++ {
		rows[i+1] = fmt.Sprintf("%2d(%2d)│%s│", i, m.OriginalRows[i], strings.Join(cells[i*m.Cols:(i+1)*m.Cols], " "))
	}
	return strings.Join(rows, "\n")
}

func (m *Matrix) GetWolframString() string {
	// see: https://www.wolframalpha.com/input/?i=Reduced+row+echelon+form&f1=%7B%7B3%2C+i%2C+2%2C3%7D%2C+%7B2%2Bi%2C+1%2C+3%2C1%7D%7D&f=ReducedRowEchelon.theMatrix_%7B%7B3%2C+i%2C+2%2C3%7D%2C+%7B2%2Bi%2C+1%2C+3%2C1%7D%7D
	// format: {{r0c0, r0c1, r0c2},{r1c0, r1c1, r1c2}}
	rows := make([]string, m.Rows)
	for r, values := range m.Data {
		valuesAsStrings := make([]string, len(values))
		for i, value := range values {
			valuesAsStrings[i] = strconv.Itoa(value)
		}
		rows[r] = fmt.Sprintf("{%s}", strings.Join(valuesAsStrings, ", "))
	}
	return fmt.Sprintf("{%s}", strings.Join(rows, ", "))
}

func (m *Matrix) Report(message string) {
	fmt.Println(message)
	fmt.Println(m)
}
