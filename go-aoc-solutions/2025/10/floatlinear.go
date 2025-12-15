package main

import (
	"fmt"
	"strings"
)

// An FloatMatrix struct consists of 'm' rows and 'n' columns
// with integer values initialized to 0.
type FloatMatrix struct {
	M, N int
	Data []float64
}

func (fm *FloatMatrix) Value(r, c int) float64 {
	index := r*fm.N + c
	if index >= len(fm.Data) {
		panic("OUT OF RANGE")
	}
	return fm.Data[index]
}

func (fm *FloatMatrix) Set(r, c int, value float64) {
	fm.Data[r*fm.N+c] = value
}

func NewFloatMatrix(m, n int) *FloatMatrix {
	data := make([]float64, m*n)
	return &FloatMatrix{m, n, data}
}

func (fm *FloatMatrix) RemoveDuplicateRows() {
	areRowsEqual := func(r1, r2 int) bool {
		for c := 0; c < fm.N; c++ {
			if fm.Value(r1, c) != fm.Value(r2, c) {
				return false
			}
		}
		return true
	}
	for i := 0; i < fm.M-1; i++ {
		for j := i + 1; j < fm.M; {
			if areRowsEqual(i, j) {
				// j is a dupe, copy last row to j and decrease size
				lastRow := fm.M - 1
				for c := 0; c < fm.N; c++ {
					fm.Set(j, c, fm.Value(lastRow, c))
				}
				fm.M--
				fm.Data = fm.Data[:fm.M*fm.N]
			} else {
				j++
			}
		}
	}
}

// RESort will order the rows to put zeros in the bottom-left of the matrix.
// This proceeds looking at each column from left to right and ordering the
// rows by value descending.   We proceed to the next column only looking at
// the range of rows from the bottom up that have zero values.
func (fm *FloatMatrix) RESort() {
	var rowGreater = func(r1, r2 int) bool {
		for c := 0; c < fm.N; c++ {
			v1 := fm.Value(r1, c)
			v2 := fm.Value(r2, c)
			if v1 > v2 {
				return true
			}
			if v2 > v1 {
				return false
			}
		}
		return false
	}

	for i := 0; i < fm.M-1; i++ {
		largest := i
		for j := i + 1; j < fm.M; j++ {
			if rowGreater(j, largest) {
				largest = j
			}
		}
		if largest != i {
			fm.SwapRows(i, largest)
		}
	}
}

// SwapRows will do just that, swap any two rows
func (fm *FloatMatrix) SwapRows(r1, r2 int) {
	for i := 0; i < fm.N; i++ {
		fm.Data[r1*fm.N+i], fm.Data[r2*fm.N+i] = fm.Data[r2*fm.N+i], fm.Data[r1*fm.N+i]
	}
}

// MultiplyRow will multiply all values in the row by the given factor
func (fm *FloatMatrix) MultiplyRow(r1 int, factor float64) {
	for i := 0; i < fm.N; i++ {
		fm.Data[r1*fm.N+i] *= factor
	}
}

// AddRows adds the source row's values to the destination row
func (fm *FloatMatrix) AddRows(source, destination int, factor float64) {
	for i := 0; i < fm.N; i++ {
		fm.Data[destination*fm.N+i] += fm.Data[source*fm.N+i] * factor
	}
}

func (fm *FloatMatrix) String() string {
	rows := make([]string, fm.M+2) // extra rows for first and last
	cells := make([]string, len(fm.Data))
	maxColumnLengths := make([]int, fm.N)
	for i, value := range fm.Data {
		cells[i] = fmt.Sprintf("%0.2f", value)
		length := len(cells[i])
		col := i % fm.N
		if length > maxColumnLengths[col] {
			maxColumnLengths[col] = length
		}
	}
	for i, value := range cells {
		length := len(value)
		diff := maxColumnLengths[i%fm.N] - length
		if diff > 0 {
			cells[i] = strings.Repeat(" ", diff) + cells[i]
		}
	}
	totalRowLength := len(maxColumnLengths) - 1 // separating spaces
	for _, l := range maxColumnLengths {
		totalRowLength += l
	}
	rows[0] = "┌" + strings.Repeat(" ", totalRowLength) + "┐"
	rows[len(rows)-1] = "└" + strings.Repeat(" ", totalRowLength) + "┘"
	for i := 0; i < fm.M; i++ {
		rows[i+1] = "│" + strings.Join(cells[i*fm.N:(i+1)*fm.N], " ") + "│"
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
for other columns
*/
