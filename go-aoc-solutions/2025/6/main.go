package main

import (
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 4277556)
	aoc.Local(part1, "part1", "input.aoc", 6757749566978)
	aoc.Local(part2, "part2", "sample.aoc", 3263827)
	aoc.Local(part2, "part2", "input.aoc", nil)
}

type Problem struct {
	Operation string
	Numbers   []int
	Result    int
}

func ParseProblems(contents string) []Problem {
	rxColumns := regexp.MustCompile(`[^\s ]+`)
	lines := aoc.ParseLines(contents)
	numberCount := len(lines) - 1
	operations := rxColumns.FindAllString(lines[len(lines)-1], -1)
	problemCount := len(operations)
	problems := make([]Problem, problemCount)
	for i, operation := range operations {
		problems[i] = Problem{Operation: operation, Numbers: make([]int, numberCount)}
	}
	for i := 0; i < numberCount; i++ {
		line := lines[i]
		numberStrings := rxColumns.FindAllString(line, -1)
		for j := 0; j < problemCount; j++ {
			number, _ := strconv.Atoi(numberStrings[j])
			problems[j].Numbers[i] = number
		}
	}
	return problems
}

func CalculateResults(problems []Problem) int {
	sum := 0
	for _, problem := range problems {
		problem.Result = problem.Numbers[0]
		for i := 1; i < len(problem.Numbers); i++ {
			switch problem.Operation {
			case "+":
				problem.Result = problem.Result + problem.Numbers[i]
			case "*":
				problem.Result = problem.Result * problem.Numbers[i]
			default:
				panic("UNEXPECTED OPERATION")
			}
		}
		sum += problem.Result
	}
	return sum
}

func part1(contents string) interface{} {
	problems := ParseProblems(contents)
	return CalculateResults(problems)
}

func PreprocessPart2(contents string) string {
	originalLines := aoc.ParseLines(contents)
	originalColumnCount := len(originalLines[0]) // if not padded at the end, find max
	outputLines := make([]string, originalColumnCount)
	for i := 0; i < originalColumnCount; i++ {
		outputLines[i] = ""
		for j := 0; j < len(originalLines); j++ {
			ch := originalLines[j][i]
			if ch != ' ' {
				outputLines[i] += string(ch)
			}
		}
	}
	result := strings.Join(outputLines, "\r\n")
	return result
}

func ParseProblemsPart2(contents string) []Problem {
	contents = PreprocessPart2(contents)
	groups := aoc.ParseGroups(contents)
	problems := make([]Problem, len(groups))
	for i, group := range groups {
		lines := aoc.ParseLines(group)
		line0Length := len(lines[0])
		operation := lines[0][line0Length-1 : line0Length]
		lines[0] = lines[0][0 : line0Length-1]
		problem := Problem{Operation: operation, Numbers: make([]int, len(lines)), Result: 0}
		for j, line := range lines {
			problem.Numbers[j], _ = strconv.Atoi(line)
		}
		problems[i] = problem
	}
	return problems
}

func part2(contents string) interface{} {
	problems := ParseProblemsPart2(contents)
	return CalculateResults(problems)
}
