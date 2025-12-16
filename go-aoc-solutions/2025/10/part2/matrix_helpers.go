package part2

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

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
