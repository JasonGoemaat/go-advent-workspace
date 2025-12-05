package aoc

import (
	"regexp"
	"strconv"
)

var reDoubleCRLF = regexp.MustCompile("[\r]?[\n][\r]?[\n]")

// ParseGroups takes a puzzle's content and splits it into separate strings
// for each section separated by an empty line.  Returns a slice of strings
// representing each group.
func ParseGroups(content string) []string {
	// remove final crlf if there is one, files usually
	// end in crlf and would have a blank line.   Splitting
	// on two here means other groups won't have them
	if content[len(content)-1] == '\n' {
		content = content[:len(content)-1]
	}
	if content[len(content)-1] == '\r' {
		content = content[:len(content)-1]
	}
	groups := reDoubleCRLF.Split(content, -1)

	// omit last group if blank
	if len(groups[len(groups)-1]) == 0 {
		groups = groups[:len(groups)-1]
	}

	return groups
}

var reNumbers = regexp.MustCompile(`-?\d+`)

// ParseLinesToInts takes a string representing lines from a puzzle input
// or a groups returned by ParseGroups and returns a slice where each element
// is a slice of int representing numbers found in the line.
//
// # Example
//
// ParseLinesToInts("345|742\n143: 972 81 43")
//
// {{345,742},{972,81,43}}
func ParseLinesToInts(lines []string) [][]int {
	results := make([][]int, len(lines))
	for i, line := range lines {
		// parts := reNumbers.FindAllString(line, -1)
		// result := make([]int, len(parts))
		// results[i] = result
		// for j, numberString := range parts {
		// 	result[j], _ = strconv.Atoi(numberString)
		// }
		results[i] = ParseInts(line)
	}
	return results
}

// ParseLines will take the content of a puzzle or group and parse it into a
// slice of NON-EMPTY lines.
func ParseLines(content string) []string {
	re := regexp.MustCompile("[^\r\n]+")
	lines := re.FindAllString(content, -1)
	return lines
}

func ParseInt(numberString string) int {
	i, _ := strconv.Atoi(numberString)
	return i
}

type IntRange struct {
	Start int
	End   int
}

var reNumberRange = regexp.MustCompile(`(\d+)-(\d+)`)

func ParseIntRange(rangeString string) IntRange {
	matches := reNumberRange.FindSubmatch([]byte(rangeString))
	start, _ := strconv.Atoi(string(matches[1]))
	end, _ := strconv.Atoi(string(matches[2]))
	return IntRange{Start: start, End: end}
}

func ParseIntRanges(contents string) []IntRange {
	lines := ParseLines(contents)
	results := make([]IntRange, len(lines))
	for i, line := range lines {
		results[i] = ParseIntRange(line)
	}
	return results
}

// Parse all numbers in possibly multi-line content into a slice.  Used by
// parseIntsPerLine and ParseLinesToInts.
func ParseInts(content string) []int {
	parts := reNumbers.FindAllString(content, -1)
	result := make([]int, len(parts))
	for j, numberString := range parts {
		result[j], _ = strconv.Atoi(numberString)
	}
	return result
}

// Parse multi-line content returning array of numbers for each line
func ParseIntsPerLine(content string) [][]int {
	return ParseLinesToInts(ParseLines(content))
}
