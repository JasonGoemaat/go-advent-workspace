package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/14
	// aoc.Local(part1, "part1", "sample.aoc", 3)
	// aoc.Local(part1, "part1", "input.aoc", 1158)
	// aoc.Local(part2, "part2", "sample.aoc", 6)
	// aoc.Local(part2, "part2", "input.aoc", 6860)
	//	aoc.Local(part2Faster, "part2Faster", "sample.aoc", 6)
	aoc.Local(part2Faster, "part2Faster", "input.aoc", 6860)
	aoc.Local(part2Simpler, "part2Simpler", "input.aoc", 6860)
}

type Command struct {
	sign   int
	number int
}

func parseCommands(lines []string) []Command {
	commands := make([]Command, len(lines))
	for k, line := range lines {
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}
		var number int
		fmt.Sscanf(line[1:], "%d", &number)
		commands[k] = Command{sign, number}
	}
	return commands
}

// part1 solves day 1, part 1 of advent of code 2025
// number of times a turn ends in position 0
func part1(contents string) interface{} {
	commands := parseCommands(aoc.ParseLines(contents))
	position := 50
	count := 0
	for _, command := range commands {
		position = (position + command.sign*command.number) % 100
		if position < 0 {
			position = (position + 100) % 100
		}
		if position == 0 {
			count++
		}
	}
	return count
}

// part2 solves day 1, part 2 of advent of code 2025
// number of times a CLICK ends in position 0, including
// when a single turn passes 0
func part2(contents string) interface{} {
	_, caller, _, _ := runtime.Caller(0)
	callerDir := filepath.Dir(caller)
	outputPath := filepath.Join(callerDir, "part2.txt")
	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	commands := parseCommands(aoc.ParseLines(contents))
	position := 50
	count := 0
	for _, command := range commands {
		for i := command.number; i > 0; i-- {
			position += command.sign
			if position < 0 {
				position += 100
			} else if position >= 100 {
				position -= 100
			}
			if position == 0 {
				count++
			}
		}
		fmt.Fprintf(f, "After command %+v, position: %d, count: %d\n", command, position, count)
	}
	return count
}

func part2Faster(contents string) interface{} {
	_, caller, _, _ := runtime.Caller(0)
	callerDir := filepath.Dir(caller)
	outputPath := filepath.Join(callerDir, "part2Faster.txt")
	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	commands := parseCommands(aoc.ParseLines(contents))
	position := 50
	count := 0
	for _, command := range commands {
		// count 100s as full turns and count, then ignore
		count += command.number / 100
		number := command.number % 100
		sign := command.sign

		// treat left as a right spin for this if it will cross or equal 0
		if sign == -1 && number >= position {
			number = 100 - number
			sign = 1
			if (position+number) < 100 && position != 0 {
				count++
			}
		}

		position += number * sign

		// count if we crossed or landed on 0
		if position >= 100 {
			count++
			position -= 100
		}
		fmt.Fprintf(f, "After command %+v, position: %d, count: %d\n", command, position, count)
	}
	return count
}

func part2Simpler(contents string) interface{} {
	commands := parseCommands(aoc.ParseLines(contents))
	position := 50
	count := 0
	for _, command := range commands {
		count += command.number / 100  // full spins
		number := command.number % 100 // remaining clicks
		if command.sign == -1 {
			// turn left

			// special case, already on 0 so don't count, act as if we're 100
			if position == 0 {
				position = 100
			}
			position -= number
			if position <= 0 {
				count++
				if position < 0 {
					position += 100
				}
			}
		} else {
			// turn right
			position += number
			if position >= 100 {
				position -= 100
				count++
			}
		}
	}
	return count
}
