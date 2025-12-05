package aoc

type Area struct {
	Width  int
	Height int
	Data   []byte
}

func (area Area) Get(row, col int) byte {
	return area.GetIndex(row*area.Width + col)
}

func (area Area) Set(row, col int, value byte) {
	if !area.Inside(row, col) {
		return
	}
	index := (row*area.Width + col)
	area.Data[index] = value
}

func (area Area) GetIndex(index int) byte {
	return area.Data[index]
}

func (area Area) IndexToRowCol(index int) (int, int) {
	row := index / area.Width
	col := index % area.Width
	return row, col
}

func (area Area) RowColToIndex(row, col int) int {
	return row*area.Width + col
}

func (area Area) Inside(row, col int) bool {
	return (row >= 0 && row < area.Height) && (col >= 0 && col < area.Width)
}

func (area Area) InsideIndex(index int) bool {
	return area.Inside(area.IndexToRowCol(index))
}

func (area Area) Is(row, col int, b byte) bool {
	return area.Inside(row, col) && (area.Get(row, col) == b)
}

func (area Area) Clone() *Area {
	newArea := Area{area.Width, area.Height, make([]byte, len(area.Data))}
	copy(newArea.Data, area.Data)
	return &newArea
}

func ParseArea(content string) *Area {
	lines := ParseLines(content) // slice of strings for each line in content
	area := Area{len(lines[0]), len(lines), make([]byte, len(lines[0])*len(lines))}
	for row, line := range lines {
		copy(area.Data[row*area.Width:], []byte(line))
	}
	return &area
}
