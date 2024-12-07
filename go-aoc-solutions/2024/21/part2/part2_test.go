package part2

import (
	"testing"

	"github.com/JasonGoemaat/go-aoc/aoc"
)

func TestPlayer(t *testing.T) {
	state := CreateState(0) // no robot controllers, just digits
	cost := state.CalculateCost(0, "379A")
	aoc.ExpectJson(t, 14, cost)
}

func TestOneController(t *testing.T) {
	state := CreateState(1) // ONE robot controller
	cost := state.CalculateCost(0, "379A")
	aoc.ExpectJson(t, 28, cost)
}

func TestTwoRobotControllers(t *testing.T) {
	state := CreateState(2) // ONE robot controller
	cost := state.CalculateCost(0, "379A")
	aoc.ExpectJson(t, 64, cost)
}

// for testing a bad one with display -- FIXED, A9 was ^^, should have been ^^^
func TestTwoRobotControllersBad(t *testing.T) {
	state := CreateState(2)
	cost := state.CalculateCostDisplay(0, "980A") // should be 60, is 59
	aoc.ExpectJson(t, 60, cost)
}

// for testing a bad one with display
func TestTwoRobotControllersBad2(t *testing.T) {
	state := CreateState(2)
	cost := state.CalculateCostDisplay(0, "179A") // should be 68, is 67
	aoc.ExpectJson(t, 68, cost)
}

func TestTwoRobotControllersSamples(t *testing.T) {
	state := CreateState(2) // TWO robot controllers (SAMPLE/PART 1)
	aoc.ExpectJson(t, 68, state.CalculateCost(0, "029A"))
	aoc.ExpectJson(t, 60, state.CalculateCost(0, "980A")) // bad, got 59
	aoc.ExpectJson(t, 68, state.CalculateCost(0, "179A")) // bad, got 67
	aoc.ExpectJson(t, 64, state.CalculateCost(0, "456A"))
	aoc.ExpectJson(t, 64, state.CalculateCost(0, "379A"))
}
