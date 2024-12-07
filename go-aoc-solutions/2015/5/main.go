package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/X
	// aoc.Local(part1, "part1", "sample.aoc", 0)
	aoc.Local(part1, "part1", "input.aoc", 236)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 51)
}

// A nice string is one with all of the following properties:

var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

// map second character to first for a bad match
var badCombinations = map[byte]byte{
	'b': 'a',
	'd': 'c',
	'q': 'p',
	'y': 'x',
}

// It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
// It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
func isNice(text string) bool {
	vowelCount := 0
	hasDupe := false
	last := byte(0)
	bytes := []byte(text)

	for _, ch := range bytes {
		if ch == last {
			hasDupe = true
		}
		if vowels[ch] {
			vowelCount++
		}
		firstBad := badCombinations[ch]
		if (firstBad > 0) && (firstBad == last) {
			return false // immediately know it's bad
		}
		last = ch
	}
	return hasDupe && (vowelCount >= 3)
}

func part1(contents string) interface{} {
	count := 0
	for _, line := range aoc.ParseLines(contents) {
		if isNice(line) {
			count++
		}
	}
	return count
}

// NEW CONDITIONS
// It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
// It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.

func isNice2(text string) bool {
	found := false
	for i := 0; i < len(text)-3; i++ {
		for j := i + 2; j < len(text)-1; j++ {
			if (text[i] == text[j]) && (text[i+1] == text[j+1]) {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return false
	}
	found = false
	for i := 2; i < len(text); i++ {
		if text[i] == text[i-2] {
			found = true
			break
		}
	}
	return found
}

func part2(contents string) interface{} {
	count := 0
	for _, line := range aoc.ParseLines(contents) {
		if isNice2(line) {
			count++
		}
	}
	return count
}
