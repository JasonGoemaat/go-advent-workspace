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
	fmt.Println(matrix)
}
