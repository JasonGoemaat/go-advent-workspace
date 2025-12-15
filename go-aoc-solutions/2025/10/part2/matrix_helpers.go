package part2

import (
	"fmt"
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

	return matrix
}

func (m *Matrix) String() string {
	rows := make([]string, m.Rows+2) // extra rows for first and last
	cells := make([]string, m.Rows*m.Cols)
	maxColumnLengths := make([]int, m.Cols)
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

	for i, value := range cells {
		length := len(value)
		diff := maxColumnLengths[i%m.Cols] - length
		if diff > 0 {
			cells[i] = strings.Repeat(" ", diff) + cells[i]
		}
	}
	totalRowLength := len(maxColumnLengths) - 1 // separating spaces
	for _, l := range maxColumnLengths {
		totalRowLength += l
	}
	identifierCells := make([]string, m.Cols)
	for i := range m.Cols {
		identifier := strconv.Itoa(i)
		identifierCells[i] = strings.Repeat(" ", maxColumnLengths[i]-len(identifier)) + identifier
		if len(identifierCells[i]) > maxColumnLengths[i] {
			maxColumnLengths[i] = len(identifierCells[i])
		}
	}
	// set max column lengths for all to be the same, if we don't do this it
	// will be more compressed but maybe harder to read
	maxColumnLength := maxColumnLengths[0]
	for i := range maxColumnLengths {
		if maxColumnLengths[i] > maxColumnLength {
			maxColumnLength = maxColumnLengths[i]
		}
	}
	for i := range maxColumnLengths {
		maxColumnLengths[i] = maxColumnLength
	}

	rows[0] = "  ┌" + strings.Join(identifierCells, " ") + "┐"
	rows[len(rows)-1] = fmt.Sprintf("  └%s┘", strings.Repeat(" ", totalRowLength))
	for i := 0; i < m.Rows; i++ {
		rows[i+1] = fmt.Sprintf("%2d│%s│", i, strings.Join(cells[i*m.Cols:(i+1)*m.Cols], " "))
	}
	return strings.Join(rows, "\n")
}
