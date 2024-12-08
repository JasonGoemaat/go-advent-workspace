package aoc

import (
	"testing"
)

func TestParseLines(t *testing.T) {
	content := "1\n4\r\n\r\n153\r\n"
	expected := []string{"1", "4", "153"}
	actual := ParseLines(content)
	ExpectJson(t, expected, actual)
}

func TestParseLinesToInts(t *testing.T) {
	lines := []string{"13|7", "8    5 3   12", "47,999,1234,9999"}
	expected := [][]int{{13, 7}, {8, 5, 3, 12}, {47, 999, 1234, 9999}}
	actual := ParseLinesToInts(lines)
	ExpectJson(t, expected, actual)
}
