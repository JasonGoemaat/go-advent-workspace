package part2

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	m := NewMatrix(3, 3)
	m.Set(1, 2, -1)
	fmt.Println(m)
}

func TestSample1(t *testing.T) {
	input := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	matrix := ParsePuzzle(input)
	matrix.Report("Initial Matrix")
	matrix.AddRow(3, 0, 1)
	matrix.Report("After Adding Row 3 to Row 0 to make 0,0 a pivot")
	matrix.AddRow(0, 3, -1)
	matrix.Report("After Adding -1 * Row 0 to Row 3 so that values below 0,0 are zero")
	// 1,1 is also a good pivot
	// 2,2 is also a good pivot
	matrix.AddRow(1, 0, -1)
	matrix.Report("After Adding -1 * Row 1 to Row 0 to make 0-2 triangular")
	fmt.Println("Now need to solve using combinations of buttons 3-5, then 0-2 have one possible value")
	fmt.Println("If swap columns 3 and 4 then multiply row 3 by -1, then I could make 3,3 triangular too,")
	fmt.Println("but that doesn't help much as I still need to try combinations of buttons 3-5")
}
