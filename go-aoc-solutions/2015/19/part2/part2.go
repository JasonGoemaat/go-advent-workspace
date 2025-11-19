package year2015day19part2

import (
	"regexp"
	"strings"

	"github.com/JasonGoemaat/go-aoc/aoc"
)

type Replacement struct {
	Original, Replaced string
	Rx                 *regexp.Regexp
}

var replacements []Replacement
var fewestStepsTo map[string]int // fewest steps to get to this string
var parent map[string]string     // when setting fewestStepsTo, set this to the parent string
var distanceFromE map[string]int // how many steps to get to 'E', use in combination with possible
var possible map[string]bool
var checked map[string]bool

var (
	queue     map[string]bool
	handled   map[string]bool
	cost      map[string]int
	maxLength int
)

func recurse(current string, steps int) {
	// loop from part 1
	for _, replacement := range replacements {
		replaced := replacement.Replaced
		all := replacement.Rx.FindAllStringIndex(current, -1)
		if all != nil {
			for _, values := range all {
				newString := current[0:values[0]] + replaced + current[values[1]:]
				if !handled[newString] || cost[newString] > (steps+1) {
					handled[newString] = true
					cost[newString] = steps + 1
					queue[newString] = true
				}
			}
		}
	}
}

// hmmm....   part2 is finding the fewest steps to get from 'e' to the final
// molecule.  I think I can do this by going in reverse.  This time I can just
// reverse the directions.  I can do this in steps maybe?
func Part2(contents string) interface{} {
	groups := aoc.ParseGroups(contents)
	replacementLines := aoc.ParseLines(groups[0])
	replacements = make([]Replacement, len(replacementLines))
	for i, line := range replacementLines {
		parts := strings.Split(line, " => ")
		rx := regexp.MustCompile(parts[1])
		replacements[i] = Replacement{parts[1], parts[0], rx} // reversed

	}
	queue = map[string]bool{}
	handled = map[string]bool{}
	cost = map[string]int{}
	queue[groups[1]] = true
	handled[groups[1]] = true
	cost[groups[1]] = 0
	for len(queue) > 0 {
		q2 := queue
		queue = map[string]bool{}
		for k := range q2 {
			recurse(k, cost[k])
		}
	}
	return cost["e"]
}
