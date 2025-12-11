package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 7)
	aoc.Local(part1, "part1", "input.aoc", 530)
	aoc.Local(part2, "part2", "sample.aoc", 33)
	aoc.Local(part2, "part2", "input.aoc", nil)
}

// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
type Machine struct {
	Lights  int
	Buttons []int
}

var rxLights = regexp.MustCompile(`\[([^\]]*)\]`)

func GetLights(line string) int {
	lightsString := rxLights.FindStringSubmatch(line)[1]
	bitValue := 1
	result := 0
	for _, rune := range lightsString {
		if rune == '#' {
			result = result | bitValue
		}
		bitValue = bitValue << 1
	}
	return result
}

// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
var rxButtons = regexp.MustCompile(`\(([^\)]*)\)`)

func GetButtons(line string) []int {
	found := rxButtons.FindAllStringSubmatch(line, -1)
	result := make([]int, len(found))
	for i, submatch := range found {
		numberStrings := strings.Split(submatch[1], ",")
		value := 0
		for _, numberString := range numberStrings {
			intValue, _ := strconv.Atoi(numberString)
			value = value | (1 << intValue)
		}
		result[i] = value
	}
	return result
}

func ParseMachine(line string) *Machine {
	lights := GetLights(line)
	buttons := GetButtons(line)
	return &Machine{lights, buttons}
}

func calculateRecursive(machine *Machine, currentValue, depth, presses int) int {
	if currentValue == machine.Lights {
		return presses
	}

	if depth >= len(machine.Buttons) {
		return -1
	}

	offPresses := calculateRecursive(machine, currentValue, depth+1, presses)

	buttonValue := machine.Buttons[depth]
	newValue := currentValue ^ buttonValue
	onPresses := calculateRecursive(machine, newValue, depth+1, presses+1)

	if offPresses == -1 {
		return onPresses
	}

	if onPresses == -1 {
		return offPresses
	}

	if offPresses < onPresses {
		return offPresses
	} else {
		return onPresses
	}
}

func CalculateFewestButtonPresses(machine *Machine) int {
	return calculateRecursive(machine, 0, 0, 0)
}

func part1(content string) interface{} {
	lines := aoc.ParseLines(content)
	sum := 0
	for _, line := range lines {
		machine := ParseMachine(line)
		sum += CalculateFewestButtonPresses(machine)
	}
	return sum
}

// --------------------------------------------------------------------------------
// PART 2

type Machine2 struct {
	Buttons    [][]int
	Joltages   []int
	MaxPresses []int // row is joltage, column is button
}

var rxJoltages = regexp.MustCompile(`{[^}]*}`)

func ParseMachine2(line string) *Machine2 {
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

type Machine2State struct {
	Machine  *Machine2
	Joltages []int
	Presses  int
}

func areZeroes(values []int) bool {
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

func calculate2Recursive(state *Machine2State, depth int) int {
	// no more buttons, return high value
	if depth >= len(state.Machine.Buttons) {
		return math.MaxInt64
	}

	buttonJoltages := state.Machine.Buttons[depth]
	lowestJoltage := math.MaxInt64
	for _, joltageIndex := range buttonJoltages {
		joltage := state.Joltages[joltageIndex]
		if joltage < lowestJoltage {
			lowestJoltage = joltage
		}
	}

	nextJoltages := make([]int, len(state.Machine.Joltages))
	copy(nextJoltages, state.Joltages)

	// start by pressing the button the maximum number of times,
	// which is the lowest joltage remaining for any joltages
	// affected by the button
	presses := lowestJoltage
	for _, joltageIndex := range buttonJoltages {
		nextJoltages[joltageIndex] -= presses
	}

	// check for finish, will always occur when pressing a button
	// the maximum number of times allowed by current voltages
	if areZeroes(nextJoltages) {
		return state.Presses + presses // previous presses + current
	}

	// if we're the last button, no use trying fewer presses
	if depth >= len(state.Machine.Buttons)-1 {
		return math.MaxInt64
	}

	// create next state
	next := Machine2State{state.Machine, nextJoltages, state.Presses + presses}

	min := math.MaxInt64
	for ; ; presses-- {
		lowest := calculate2Recursive(&next, depth+1)
		if lowest < min {
			min = lowest
		}
		if presses <= 0 {
			break
		}
		next.Presses--
		for _, joltageIndex := range buttonJoltages {
			next.Joltages[joltageIndex]++
		}
	}
	return min
}

func CalculateMachine2(machine *Machine2) int {
	state := Machine2State{
		machine,
		machine.Joltages,
		0,
	}
	return calculate2Recursive(&state, 0)
}

func part2(content string) interface{} {
	lines := aoc.ParseLines(content)
	sum := 0
	timings := make([]int64, len(lines))
	for i, line := range lines {
		startTime := time.Now()

		machine := ParseMachine2(line)
		sum += CalculateMachine2(machine)

		endTime := time.Now()
		ms := endTime.Sub(startTime).Milliseconds()
		timings[i] = ms

	}
	return sum
}
