package main

import (
	"math"
	"strconv"
	"strings"
)

func ParseMachine2Bifurcate(line string) *Machine2 {
	found := rxButtons.FindAllStringSubmatch(line, -1)
	buttons := make([][]int, len(found))
	for i, submatch := range found {
		numberStrings := strings.Split(submatch[1], ",")
		value := make([]int, len(numberStrings))
		for j, numberString := range numberStrings {
			intValue, _ := strconv.Atoi(numberString)
			value[j] = intValue
		}
		buttons[i] = value
	}

	joltagesString := rxJoltages.FindString(line)
	joltagesString = joltagesString[1 : len(joltagesString)-1]
	joltagesStrings := strings.Split(joltagesString, ",")
	joltages := make([]int, len(joltagesStrings))
	for i, s := range joltagesStrings {
		joltages[i], _ = strconv.Atoi(s)
	}

	maxPresses := make([]int, len(joltages)*len(buttons))
	for col := range buttons {
		maxValue := math.MaxInt64
		for _, i := range buttons[col] {
			if joltages[i] < maxValue {
				maxValue = joltages[i]
			}
		}
		for _, joltageIndex := range buttons[col] {
			maxPresses[joltageIndex*len(buttons)+col] = maxValue
		}
	}

	return &Machine2{buttons, joltages, maxPresses}
}

type BifurcateCombo struct {
	Presses        int
	JoltageChanges []int
	BitMask        int
	Buttons        []int // list of indexes of buttons contained in combo, for reference not use in solving
}

type BifurcatePuzzle struct {
	Combos         map[int][]BifurcateCombo
	Joltages       []int
	ButtonJoltages [][]int // ButtonJoltages[0] is array of joltage changes when pressing button 0
}

// Return minimum number of button presses
func recursePart2Bifurcate(puzzle *BifurcatePuzzle, joltages []int) int {
	totalRemainingJoltage := 0
	for _, joltage := range joltages {
		if joltage < 0 {
			return math.MaxInt64
		}
		totalRemainingJoltage += joltage
	}
	if totalRemainingJoltage == 0 {
		return 0 // nothing left to do!
	}

	// oddMask will be mask of joltages that are odd, i.e. {79,144,32,17,8}
	// would be 10010 reading left-right, binary 01001 or decimal 9
	oddMask := 0
	for i, joltage := range joltages {
		if (joltage & 1) == 1 {
			oddMask |= 1 << i
		}
	}

	// see if we can use combos to make them even so we can divide by 2
	minPresses := math.MaxInt64
	combos, ok := puzzle.Combos[oddMask]
	if ok {
		for _, combo := range combos {
			nextJoltages := make([]int, len(joltages))
			isValid := true
			for i := range joltages {
				nextJoltages[i] = joltages[i] - combo.JoltageChanges[i]
				if nextJoltages[i] < 0 || (nextJoltages[i]&1) > 0 {
					isValid = false
					break
				}
				nextJoltages[i] = nextJoltages[i] >> 1
			}
			if isValid {
				presses := combo.Presses + recursePart2Bifurcate(puzzle, nextJoltages)*2
				if presses < minPresses {
					minPresses = presses
				}
			}
		}
	}

	// found a solution using recursion
	if minPresses < math.MaxInt64 {
		return minPresses
	}

	// NO SOLUTION POSSIBLE USING PARITY, DO BRUTE FORCE HERE!
	return math.MaxInt64
}

func part2Bifurcate(content string) any {
	return nil
}
