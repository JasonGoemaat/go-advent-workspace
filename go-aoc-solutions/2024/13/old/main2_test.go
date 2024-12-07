package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestMain(t *testing.T) {
	aoc.ExpectJson(t, 1, 1)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}

func TestBad1(t *testing.T) {
	contents := `Button A: X+36, Y+17
Button B: X+43, Y+77
Prize: X=3253, Y=1933`
	c1, a1, b1 := solveEasy(contents, 0, 0)
	c2, a2, b2 := solveMath(contents, false) // should Cost 253, A 82, B 7 times
	aoc.ExpectJson(t, []int{c1, a1, b1}, []int{c2, a2, b2})
}

func TestBad2(t *testing.T) {
	contents := `Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450`
	c1, a1, b1 := solveEasy(contents, 0, 0)
	c2, a2, b2 := solveMath(contents, false) // should be Cost 200, A 38, B 86
	aoc.ExpectJson(t, []int{c1, a1, b1}, []int{c2, a2, b2})
}

func TestBad3(t *testing.T) {
	contents := `Button A: X+11, Y+14
Button B: X+80, Y+28
Prize: X=4187, Y=2450`
	c1, a1, b1 := solveEasy(contents, 0, 0)
	c2, a2, b2 := solveMath(contents, false) // should be Cost 330, A 97, B 39
	aoc.ExpectJson(t, []int{c1, a1, b1}, []int{c2, a2, b2})
}

func TestBad4(t *testing.T) {
	// should be no solution
	contents := `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`
	c1, a1, b1 := solveEasy(contents, 0, 0)
	c2, a2, b2 := solveMath(contents, false) // should be Cost ?, A ?, B ?
	aoc.ExpectJson(t, []int{c1, a1, b1}, []int{c2, a2, b2})
}

func TestBadMath1(t *testing.T) {
	// should be 112, 23, 43
	contents := `Button A: X+48, Y+58
Button B: X+86, Y+22
Prize: X=4802, Y=2280`
	c1, a1, b1 := solveEasy(contents, 0, 0)
	c2, a2, b2 := solveMath(contents, false) // should be Cost ?, A ?, B ?
	aoc.ExpectJson(t, []int{c1, a1, b1}, []int{c2, a2, b2})
}

func TestBadMath2(t *testing.T) {
	// should be 0, 0, 0 (no solution)
	contents := `Button A: X+23, Y+74
Button B: X+74, Y+24
Prize: X=15709, Y=15466`
	c1, a1, b1 := solveEasy(contents, 0, 0)
	c2, a2, b2 := solveMath(contents, false) // should be Cost ?, A ?, B ?
	aoc.ExpectJson(t, []int{c1, a1, b1}, []int{c2, a2, b2})
}

func TestBadMath3(t *testing.T) {
	// should be 173, 50, 23 (got 323, 0, 323)
	contents := `Button A: X+90, Y+74
Button B: X+15, Y+50
Prize: X=4845, Y=4850`
	c1, a1, b1 := solveEasy(contents, 0, 0)
	c2, a2, b2 := solveMath(contents, false) // should be Cost ?, A ?, B ?
	aoc.ExpectJson(t, []int{c1, a1, b1}, []int{c2, a2, b2})
}
