package part2

type Matrix struct {
	Rows, Cols int
	Data       [][]int
}

func NewMatrix(rows, cols int) *Matrix {
	data := make([][]int, rows)
	for i := range rows {
		data[i] = make([]int, cols)
	}
	return &Matrix{Rows: rows, Cols: cols, Data: data}
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
