package main

import (
	"fmt"
	"testing"
)

func TestInput1(t *testing.T) {
	input := "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
	result := solution10B(input)
	fmt.Println(result)
	// result is 80
}
