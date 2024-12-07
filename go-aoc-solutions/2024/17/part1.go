// what is part 2?   change in instructions?   new instructions?
// oh, maybe it lets you modify instruction code?
package main

import (
	"fmt"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func part1(contents string) interface{} {
	var state = NewDay17(contents)
	for !state.Done() {
		state.Step()
	}
	return state.Output()
}

type day17 struct {
	A, B, C int
	code    []int
	ip      int
	outputs []int
}

func (state *day17) Output() string {
	ss := make([]string, len(state.outputs))
	for i, v := range state.outputs {
		ss[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(ss, ",")
}

func NewDay17(contents string) *day17 {
	state := day17{0, 0, 0, nil, 0, []int{}}
	groups := aoc.ParseGroups(contents)
	alines := aoc.ParseLinesToInts(aoc.ParseLines(groups[0]))
	state.A = alines[0][0]
	state.B = alines[1][0]
	state.C = alines[2][0]
	state.code = aoc.ParseLinesToInts(aoc.ParseLines(groups[1]))[0]
	return &state
}

func (state *day17) Done() bool {
	return state.ip >= len(state.code)
}

func (state *day17) Step() {
	switch state.code[state.ip] {
	case 0:
		state.adv()
	case 1:
		state.bxl()
	case 2:
		state.bst()
	case 3:
		state.jnz()
	case 4:
		state.bxc()
	case 5:
		state.out()
	case 6:
		state.bdv()
	case 7:
		state.cdv()
	default:
		panic("UNKNOWN INSTRUCTION")
	}
}

// for part 2, run through to jnz instruction,
// output the value that was output and a flag that is true
// if the jnz fell through (no more loops, required for last iteration
// when 0 is returned)
func (state *day17) Run1() (int, bool) {
	output := 0
	for true {
		switch state.code[state.ip] {
		case 0:
			state.adv()
		case 1:
			state.bxl()
		case 2:
			state.bst()
		case 3:
			// state.jnz()
			return output, state.jnzFake()
		case 4:
			state.bxc()
		case 5:
			// state.out()
			output = state.outFake()
		case 6:
			state.bdv()
		case 7:
			state.cdv()
		default:
			panic("UNKNOWN INSTRUCTION")
		}
	}
	return 0, true
}

func (state *day17) combo() int {
	arg := state.code[state.ip+1]
	if arg == 7 {
		panic("ARGUMENT RESERVED")
	}
	if arg < 0 {
		panic("ARGUMENT LESS THAN 0")
	}
	if arg < 4 {
		return arg
	}
	if arg == 4 {
		return state.A
	}
	if arg == 5 {
		return state.B
	}
	if arg == 6 {
		return state.C
	}
	panic("ARGUMENT > 7")
}

func (state *day17) literal() int {
	return state.code[state.ip+1]
}

func (state *day17) adv() {
	// The adv instruction (opcode 0) performs division. The numerator is the
	// value in the A register. The denominator is found by raising 2 to the
	// power of the instruction's combo operand. (So, an operand of 2 would
	// divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result
	// of the division operation is truncated to an integer and then written
	// to the A register.

	// this was silly, dividing by 2**x is the same as >> x
	// arg := state.combo()
	// operand := 1 << arg // this could be massive in some cases
	// state.A = state.A / operand
	x := state.combo()
	y := state.A >> x
	state.A = y
	state.ip = state.ip + 2
}

func (state *day17) bxl() {
	// The bxl instruction (opcode 1) calculates the bitwise XOR of register B
	// and the instruction's literal operand, then stores the result in register B.
	x := state.literal()
	y := state.B ^ x
	state.B = y
	state.ip = state.ip + 2
}

func (state *day17) bst() {
	// The bst instruction (opcode 2) calculates the value of its combo operand
	// modulo 8 (thereby keeping only its lowest 3 bits), then writes that value
	// to the B register.
	x := state.combo()
	y := x & 7
	state.B = y
	state.ip = state.ip + 2
}

// Returns true if program is done, required for last output and should stop recursion
func (state *day17) jnzFake() bool {
	if state.A == 0 {
		state.ip = state.ip + 2
		return true
	}
	y := state.literal()
	state.ip = y
	return false
}

func (state *day17) jnz() {
	// The jnz instruction (opcode 3) does nothing if the A register is 0.
	// However, if the A register is not zero, it jumps by setting the
	// instruction pointer to the value of its literal operand; if this
	// instruction jumps, the instruction pointer is not increased by 2 after
	// this instruction.
	if state.A == 0 {
		state.ip = state.ip + 2
		return
	}
	y := state.literal()
	state.ip = y
}

func (state *day17) bxc() {
	// The bxc instruction (opcode 4) calculates the bitwise XOR of register
	// B and register C, then stores the result in register B. (For legacy
	// reasons, this instruction reads an operand but ignores it.)
	x := state.C
	y := state.B ^ x
	state.B = y
	state.ip = state.ip + 2
}

func (state *day17) outFake() int {
	// like output, but return value instead of adding to array
	x := state.combo()
	y := x & 7
	return y
}

func (state *day17) out() {
	// The out instruction (opcode 5) calculates the value of its combo
	// operand modulo 8, then outputs that value. (If a program outputs
	// multiple values, they are separated by commas.)
	x := state.combo()
	y := x & 7
	state.outputs = append(state.outputs, y)
	state.ip = state.ip + 2
}

func (state *day17) bdv() {
	// The bdv instruction (opcode 6) works exactly like the adv instruction
	// except that the result is stored in the B register. (The numerator is
	// still read from the A register.)
	x := state.combo()
	y := state.A >> x
	state.B = y
	state.ip = state.ip + 2
}

func (state *day17) cdv() {
	// The cdv instruction (opcode 7) works exactly like the adv instruction
	// except that the result is stored in the C register. (The numerator is
	// still read from the A register.)
	x := state.combo()
	y := state.A >> x
	state.C = y
	state.ip = state.ip + 2
}
