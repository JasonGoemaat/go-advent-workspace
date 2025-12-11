package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/JasonGoemaat/go-aoc/aoc"
)

func TestLightsRegex(t *testing.T) {
	input := "[..#.] (1,3) (0,3) (0,2,3) {33,7,14,40}"
	expected := "..#."
	values := rxLights.FindStringSubmatch(input)
	if values[1] != expected {
		t.Errorf("rxLights found '%s' but we expected '%s'", values[1], expected)
	}
}

func TestGetLights(t *testing.T) {
	input := ".###"
	expected := 14
	value := GetLights(input)
	if value != expected {
		t.Errorf("GetLights(\"%s\") returned '%d' but we expected '%d'", input, value, expected)
	}
}

func TestRxButtons(t *testing.T) {
	input := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	expected := []string{"3", "1,3", "2", "2,3", "0,2", "0,1"}
	value := rxButtons.FindAllStringSubmatch(input, -1)
	for i, stringArray := range value {
		if expected[i] != stringArray[1] {
			t.Errorf("RxButtons index %d is \"%s\" but we expected \"%s\"", i, stringArray[1], expected[i])
		}
	}
}

func TestGetButtons(t *testing.T) {
	input := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	expected := []int{8, 10, 4, 12, 5, 3}
	buttons := GetButtons(input)
	for i, button := range buttons {
		if expected[i] != button {
			t.Errorf("GetButtons() index %d is %d but we expected %d", i, button, expected[i])
		}
	}
}

func TestCalculateFewestButtonPresses(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"[.##] (1,2)", 1},
		{"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}", 2},
		{"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}", 3},
		{"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}", 2},
	}
	for i, test := range tests {
		machine := ParseMachine(test.input)
		value := CalculateFewestButtonPresses(machine)
		if value != test.expected {
			t.Errorf("CalculateFewestButtonPresses with machine %d '%s' returned %d but expected %d", i, test.input, value, test.expected)
		}
	}
}

func TestJoltagesRegex(t *testing.T) {
	tests := []struct {
		line     string
		expected string
	}{
		{"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}", "3,5,4,7"},
	}
	for i, test := range tests {
		value := rxJoltages.FindString(test.line)
		value = value[1 : len(value)-1]
		if value != test.expected {
			t.Errorf("Test %d got `%s` but expected `%s`", i, value, test.expected)
		}
	}

}

func TestParseMachine2(t *testing.T) {
	tests := []struct {
		line     string
		expected Machine2
	}{
		{
			"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
			Machine2{
				[][]int{{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2}},
				[]int{10, 11, 11, 5, 10, 5},
				[]int{5, 5, 5, 0, 5, 0, 5, 11, 5, 0, 5, 11, 5, 5, 0, 0, 5, 5, 5, 0, 0, 0, 5, 0},
			},
		},
	}
	for i, test := range tests {
		machine := ParseMachine2(test.line)
		if !reflect.DeepEqual(machine, &test.expected) {
			t.Errorf("Test %d: '%s'\n  Got: %v\n  Expected: %v\n", i, test.line, machine, test.expected)
		}
	}
}

func TestPart2Counts(t *testing.T) {
	contents := aoc.GetLocalFile("input.aoc")
	lines := aoc.ParseLines(contents)
	max := 0
	var maxMachine *Machine2
	for i, line := range lines {
		machine := ParseMachine2(line)
		presses := calculatePresses(machine)
		if presses > max {
			max = presses
			maxMachine = machine
		}
		fmt.Printf("Line %d could take %d presses\n", i, presses)
	}
	fmt.Printf("Absolute max presses is %d\n", max)
	fmt.Printf("%v", maxMachine)
}

func calculatePresses(machine *Machine2) int {
	maxPresses := make([]int, len(machine.Buttons))
	for i, button := range machine.Buttons {
		minJoltage := math.MaxInt64
		for _, index := range button {
			if machine.Joltages[index] < minJoltage {
				minJoltage = machine.Joltages[index]
			}
		}
		maxPresses[i] = minJoltage
	}
	possible := 1
	for _, p := range maxPresses {
		possible *= p
	}
	return possible
}
