package part2

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
