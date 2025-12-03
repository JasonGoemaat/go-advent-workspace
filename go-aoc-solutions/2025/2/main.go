package main

import (
	"regexp"
	"strconv"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 1227775554)
	aoc.Local(part1, "part1", "input.aoc", 40214376723)
	aoc.Local(part2, "part2", "sample.aoc", 4174379265)
	aoc.Local(part2, "part2", "input.aoc", 50793864718)
}

type IntRange struct {
	Start int
	End   int
}

func ParseIntRanges(content string) []IntRange {
	reWholeNumberRange := regexp.MustCompile(`(\d+)-(\d+)`)
	matches := reWholeNumberRange.FindAllStringSubmatch(content, -1)
	results := make([]IntRange, len(matches))
	for i, match := range matches {
		if len(match) == 3 {
			start, _ := strconv.Atoi(match[1])
			end, _ := strconv.Atoi(match[2])
			results[i] = IntRange{Start: start, End: end}
		}
	}
	return results
}

func getDigitCount(i int) int {
	// get number of digits in number
	digit_count := 0
	for n := i; n > 0; n /= 10 {
		digit_count++
	}
	return digit_count
}

func powerOfTen(exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= 10
	}
	return result
}

func part1Validate(i int) bool {
	digit_count := getDigitCount(i)
	if (digit_count % 2) != 0 {
		return false
	}
	power_of_ten := powerOfTen(digit_count / 2)
	a, b := i/power_of_ten, i%power_of_ten
	return a == b
}

func part1(contents string) interface{} {
	ranges := ParseIntRanges(contents)
	sum := 0
	for _, r := range ranges {
		// Process pairs of integers
		for i := r.Start; i <= r.End; i++ {
			if part1Validate(i) {
				sum += i
			}
		}
	}
	return sum
}

func part2Validate(i int) bool {
	digit_count := getDigitCount(i)
	for digits := 1; digits <= digit_count/2; digits++ {
		if digits == 1 || ((digit_count % digits) == 0) {
			power_of_ten := powerOfTen(digits)
			n := i
			valid := true
			first := n % power_of_ten
			n = n / power_of_ten
			for ; n > 0; n /= power_of_ten {
				next := n % power_of_ten
				if next != first {
					valid = false
					break
				}
			}
			if valid {
				return true
			}
		}
	}
	return false
}

func part2(contents string) interface{} {
	ranges := ParseIntRanges(contents)
	sum := 0
	for _, r := range ranges {
		// Process pairs of integers
		for i := r.Start; i <= r.End; i++ {
			if part2Validate(i) {
				sum += i
			}
		}
	}
	return sum
}
