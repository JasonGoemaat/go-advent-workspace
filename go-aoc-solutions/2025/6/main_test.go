package main

import (
	"fmt"
	"regexp"
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestFindAllString(t *testing.T) {
	rxColumns := regexp.MustCompile(`[^\w ]+`)
	sampleString := "+   * +   *"
	found := rxColumns.FindAllString(sampleString, -1)
	fmt.Println(found)
}

func TestPreprocessPart2(t *testing.T) {
	contents := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	result := PreprocessPart2(contents)
	groups := aoc.ParseGroups(result)
	lines := aoc.ParseLines(groups[0])
	expected := []string{
		"1*",
		"24",
		"356",
		// "",
		// "369+",
		// "248",
		// "8",
		// "",
		// "32*",
		// "581",
		// "175",
		// "",
		// "623+",
		// "431",
		// "4",
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Expecting line %d to be '%s' but got '%s'", i, expected[i], line)
		}
	}
}
