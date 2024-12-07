// What is the lowest positive initial value for register A that causes the program to output a copy of itself?
package main

import "fmt"

// oh, that's devious, so you need to actually figure out what the program does,
// or I guess you could try a bunch of things.   You need the lowest value that will
// output the program itself.

// program: 2,4,1,1,7,5,1,4,0,3,4,5,5,5,3,0
// 2,4: bst 4 	B = A & 7 // 3 bits from A
// 1,1: bxl 1 	(B = B ^ 1)
// 7,5: cdv 5
// 1,4: bxl 4 	(B = B ^ 4)
// 0,3: adv 3 	A = A >> 3 // strip 3 bits from A
// 4,5: bxc 	B = B ^ C
// 5,5: out 	B
// 3,0: JMP START

// re-organizing, we take 3 bits from A and store in B each iteration
// 2,4: bst 4 	B = A & 7 // 3 bits from A
// 0,3: adv 3 	A = A >> 3 // strip 3 bits from A

// here's the complicated part, b has 3 bits from A, we xor with 1
// 1,1: bxl 1 	(B = B ^ 1) // b has 3 bits from A with last (lowest) bit flipped
// 7,5: cdv 5   (C = A>>B) // C takes entire value in A (before stripping bits) and shifts right by B bits
// 1,4: bxl 4 	(B = B ^ 4) // flip 2nd bit in B
// 4,5: bxc 	B = B ^ C // xor with C,
// 5,5: out 	B

// 3,0: JMP START

// want to turn	5,1,4,0,5,1,0,2,6
// into 		2,4,1,1,7,5,1,4,0,3,4,5,5,5,3,0
// ugly, 16 digits = up to 48 bits required

// one thing I note is that C is reset each round
// also the only thing it can shift A by is up to 7 bits, but that's after taking 3 so 10 bits of A are important
// ok, maybe for each B output I need a list of 1024 possible values for the next 10 bits in A
// so I have some things for 10 bits in A and record which ones provide each of the 8 outputs,
// or 8192 total candidate bits
// maybe it's just simpler to run it and use memoizing.   Try 0-1024 until I get a '2' for instance,
// then add each 3 bits to A

// B and C are trashed each loop, so the only thing that matters is A's starting, or just the
// low 10 bits of A actually.   A cannot be 0 until we have a 0

var required = [8][]int{}

// returns the new value of a and the value output
// if this return 0 for A we should exit because of the jnz at the end
func fastLoopSample(A int) (int, int) {
	A = A >> 1      // 0,1    A = A >> 1
	return A, A & 7 // 5,4 output A mod 8
	// 3,0 NOOP (jnz, caller will end when A = 0, otherwise will call again
}

func part2FastLoopSample(contents string) interface{} {
	var state = NewDay17(contents)
	var A = state.A
	var output int
	for A > 0 {
		A, output = fastLoopSample(A)
		state.outputs = append(state.outputs, output)
	}
	return state.Output()
}

// returns the new value of a and the value output
// if this return 0 for A we should exit because of the jnz at the end
func fastLoopInput(A int) (int, int) {
	var B, C int
	B = A & 7  // 2,4: bst 4 	B = A & 7 // 3 bits from A
	B = B ^ 1  // 1,1: bxl 1 	(B = B ^ 1)
	C = A >> B // 7,5: cdv 5 	C = A >> B
	C = C ^ 4  // 1,4: bxl 4 	(B = B ^ 4)
	A = A >> 3 // 0,3: adv 3 	A = A >> 3 // strip 3 bits from A
	B = B ^ C  // 4,5: bxc 		B = B ^ C
	// output does mod 8
	B = B & 7
	return A, B
}

func part2FastLoopInput(contents string) interface{} {
	var state = NewDay17(contents)
	var A = state.A
	var output int
	for A > 0 {
		A, output = fastLoopInput(A)
		state.outputs = append(state.outputs, output)
	}
	return state.Output()
}

func part2Test(contents string) interface{} {
	// var state = NewDay17(contents)
	// ok, let's first find a result where A=0 and B=0
	count := 0
	for i := range 1 << 19 {
		A, output := fastLoopInput(i)

		// 2,4,1,1,7,5,1,4,0,3,4,5,5,5,3,0 <-- start from end

		// // FINAL OUTPUT
		// if A == 0 && output == 0 { // ONLY 1 SOLUTION: A = 5
		// 	fmt.Printf("%3x\n", i)
		// }

		// // 2nd to last OUTPUT
		// if A == 5 && output == 3 { // sweet, 1 solution, A = 0x2e
		// 	fmt.Printf("%3x\n", i)
		// 	count++
		// }

		// // 3rd to last OUTPUT
		// if A == 0x2e && output == 5 { // sweet, 2 solutions, A = 170,171
		// 	fmt.Printf("%3x\n", i)
		// 	count++
		// }

		// // 4rd to last OUTPUT, 5 possibilities?   why are these > 10 digits?, oh, derp, I'm looking for 8 with 0x170,0x171
		// // so let's see what these give
		// // b80
		// // b81
		// // b84
		// // b89
		// // b8c
		// if (A == 0x170 || A == 0x171) && output == 5 { // 5 solutions now
		// 	fmt.Printf("%3x\n", i)
		// 	count++
		// }

		// how to leave 0xb84
		if (A == 0xb80) && output == 5 { //
			fmt.Printf("%3x\n", i)
			count++
		}
	}

	// THESE are the possibilities to leave 171 or 170
	// 0000 0010 1110 // 0x2e
	// 0001 0111 0000 // 0x170 - note how the top 5 bits are all the same, shifted
	// 0001 0111 0001 // 0x171
	// 1011 1000 0000 // 0xb80 - note how the top 8 bits are all the same, shifted 3 from above
	// 1011 1000 0001
	// 1011 1000 0101
	// 1011 1000 1001
	// 1011 1000 1100
	for _, v := range []int{0xb80, 0xb81, 0xb84, 0xb89, 0xb8c} {
		fmt.Printf("Trying 0x%x\n", v)
		var output int
		var A = v
		for A > 0 {
			A, output = fastLoopInput(A)
			fmt.Printf("  Output '%d', A now 0x%x\n", output, A)
		}
	}

	fmt.Printf("FOUND %d SOLUTIONS\n", count)

	// ok, ONLY thing that matters for the programs we are given is the
	// value of A, that's nice to know.   ALSO we want the SMALLEST value
	// that will produce the complete output.   That means if we go from
	// lowest A to highest, we can quit once we find the output.
	// HOWEVER, the 'cdv' instruction can  use more bits.   At least it happens
	// before shifting A 3 bits, so it is limited to 256 options, or 3 future iterations
	// I DO THINK we should start from the right
	return 0
}

// func (state *model) findOriginal(A, position, bits int) int {
// 	if position >= len(state.code) {
// 		return A // all previous positions matched, we good
// 	}
// 	check := state.code[len(code)]
// 	shiftLeft := position * 3
// 	for i := range 8 {
// 		// OHHHHHH, I only need to add 3 bits every time, the A that is passed
// 		// includes 10 bits for position 0, 13 for 1, etc...
// 		newA, output := fastLoopInput(i)
// 	}
// 	return 0
// }

var recursions, calls = 0, 0

func (state *day17) findResult(requiredA, position int) int {
	recursions++
	if position >= len(state.code) {
		return requiredA // we already found it
	}
	requiredOutput := state.code[len(state.code)-position-1]

	// start with required A shifted 3 bits, then mangle the bottom 10 bits
	// to try and get the necessary output
	shifted := requiredA << 3
	for i := range 1 << 10 {
		// try out our shifted and mangled A to see if it produces
		// the required A and output
		testA := shifted ^ i
		calls++
		newA, output := fastLoopInput(testA)
		if newA == requiredA && output == requiredOutput {
			result := state.findResult(testA, position+1)
			if result > 0 {
				return result
			}
		}
	}
	return -1 // no chance?  thought it should be possible...   Wait, not necessarily since there's multiple possibilities?
	// OHHHHHH,  bet this will trap me, I bet there will be end results that might be less somehow...   We'll see
	// OH, this happens with requiredA 0 position 5, but how did I get to requiredA 0?   that would end it, no?
}

func part2(contents string) interface{} {
	var state = NewDay17(contents)

	// look for A=0, output of position 0 from the end
	result := state.findResult(0, 0)

	fmt.Printf("%d recursions, %d loops\n", recursions, calls)
	return result
}

// func (state *day17) findResult(requiredA, position int) int {
// 	if position >= len(state.code) {
// 		return requiredA // we already found it
// 	}
// 	requiredOutput := state.code[len(state.code)-position-1]

// 	shifted := requiredA << 3
// 	for i := range 1 << 10 {
// 		testA := shifted ^ i
// 		newA, output := fastLoopInput(testA)
// 		if newA == requiredA && output == requiredOutput {
// 			result := state.findResult(testA, position+1)
// 			if result > 0 {
// 				return result
// 			}
// 		}
// 	}
// 	return -1
// }
