package main

import (
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/14
	aoc.LocalArgs(part1, "part1", "sample.aoc", 1120, 1000)
	aoc.LocalArgs(part1, "part1", "input.aoc", 2640, 2503)
	aoc.LocalArgs(part2, "part2", "sample.aoc", 689, 1000)
	aoc.LocalArgs(part2, "part2", "input.aoc", 1102, 2503) // 1064 is too low
}

type Reindeer struct {
	Name           string
	Speed          int
	ActiveSeconds  int
	RestingSeconds int
	CycleSeconds   int
	CycleDistance  int
}

func NewReindeer(line string) *Reindeer {
	spaceIndex := strings.IndexRune(line, ' ')
	nums := aoc.ParseInts(line)
	name := line[0:spaceIndex]
	speed := nums[0]
	activeSeconds := nums[1]
	restingSeconds := nums[2]
	cycleSeconds := activeSeconds + restingSeconds
	cycleDistance := activeSeconds * speed
	return &Reindeer{
		Name:           name,
		Speed:          speed,
		ActiveSeconds:  activeSeconds,
		RestingSeconds: restingSeconds,
		CycleSeconds:   cycleSeconds,
		CycleDistance:  cycleDistance,
	}
}

func ParseReindeer(contents string) []*Reindeer {
	lines := aoc.ParseLines(contents)
	results := make([]*Reindeer, len(lines))
	for i, line := range lines {
		results[i] = NewReindeer(line)
	}
	return results
}

func (r *Reindeer) CalculateDistance(seconds int) int {
	cycles := seconds / r.CycleSeconds
	remainingSeconds := seconds % r.CycleSeconds
	return (cycles * r.CycleDistance) + min(r.CycleDistance, r.Speed*remainingSeconds)
}

func part1(contents string, args ...interface{}) interface{} {
	seconds := args[0].(int)
	reindeer := ParseReindeer(contents)
	largestDistance := 0
	for _, r := range reindeer {
		distance := r.CalculateDistance(seconds)
		largestDistance = max(distance, largestDistance)
	}
	return largestDistance
}

// well, well...   I thought it would be something like who wins after 20
// trillion seconds or something like that.   Turns out we give one point to
// the reindeer in the lead at each second over the 2503 seconds, and report
// the point total for the winning reindeer.
func part2(contents string, args ...interface{}) interface{} {
	seconds := args[0].(int)
	reindeer := ParseReindeer(contents)
	points := make([]int, len(reindeer))
	for i := 1; i <= seconds; i++ {
		largestDistance := 0
		for _, r := range reindeer {
			largestDistance = max(largestDistance, r.CalculateDistance(i))
		}
		for j, r := range reindeer {
			distance := r.CalculateDistance(i)
			if distance == largestDistance {
				points[j]++
			}
		}
	}
	maxPoints := 0
	for _, p := range points {
		maxPoints = max(maxPoints, p)
	}
	return maxPoints
}
