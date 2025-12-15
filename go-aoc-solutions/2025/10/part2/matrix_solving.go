package part2

type Solution struct {
	Buttons []int // how many times each button is pressed
	Presses int   // total presses
}

// prep matrix, attempting to put in reduced row echelon form
func (m *Matrix) Prep() {

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
func (m *Matrix) RREFRecurse(row, col int) {
	// exit if we're past the last button column or the bottom row
	if row >= m.Rows || col >= m.Cols-1 {
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
		}

		// fix-up rows above to make change their values in col to 0
		for r := row - 1; r >= 0; r-- {
			value := m.Get(r, col)
			if value != 0 {
				m.AddRow(row, r, -value)
			}
		}

		// move to next row and col
		m.RREFRecurse(row+1, col+1)
		return
	}

	// if all zeros, move on to next column, stay on current row
	if allZero {
		m.RREFRecurse(row, col+1)
	}

	// not all zeros and no 1s found, we can't continue
	// TODO:  As an edge case, we may be able to continue if there exists
	// a '1' or '-1' *somewhere* in the submatrix, by just swapping columns
	// to put that '1' or '-1' in column 'col'
	return
}

// Attempt to convert to Reduced Row Echelon Form
func (m *Matrix) RREF() {
	m.RREFRecurse(0, 0)
}

// Lower MaxPresses[].  Look for rows that are all one sign
// If all negative, multiply by -1 to make all positive.
// Constrain the buttons specified in that row so that they
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
