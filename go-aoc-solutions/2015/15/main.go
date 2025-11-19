package main

import (
	year2015day15part2 "github.com/JasonGoemaat/go-aoc-solutions/2015/15/part2"
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/15
	aoc.Local(part1, "part1", "sample.aoc", 62842880)
	aoc.Local(part1, "part1", "input.aoc", 13882464)
	aoc.Local(year2015day15part2.Part2, "part2", "sample.aoc", 57600000)
	aoc.Local(year2015day15part2.Part2, "part2", "input.aoc", 11171160)
}

func calculateScore(ingredients [][]int, counts []int) int {
	scores := make([]int, len(ingredients[0]))
	for i, ingredient := range ingredients {
		for j := range len(ingredient) {
			score := ingredient[j] * counts[i]
			scores[j] += score
		}
	}
	total := scores[0]
	for i := 1; i < len(scores); i++ {
		total *= scores[i]
	}
	return total
}

func part1(contents string) interface{} {
	ingredients := aoc.ParseIntsPerLine(contents) // name doesn't matter
	counts := make([]int, len(ingredients))

	// for now, ignore calories
	for i := range len(ingredients) {
		ingredients[i] = ingredients[i][0 : len(ingredients[i])-1]
	}

	for i := range len(ingredients) {
		counts[i] = 100
	}
	totalCount := 100 * len(ingredients)

	// remove one teaspoon at a time, removing the ingredient that leaves us
	// with the best score
	for totalCount > 100 {
		bestScore := 0
		bestIndex := 0
		for i := range len(ingredients) {
			if counts[i] == 0 {
				// can't have negative ingredients
				continue
			}
			counts[i]--
			score := calculateScore(ingredients, counts)
			if score > bestScore {
				bestScore = score
				bestIndex = i
			}
			counts[i]++ // replace for now
		}
		counts[bestIndex]--
		totalCount--
	}
	return calculateScore(ingredients, counts)
	// 182206 was the total calls to recurse()
	// 176952 was the total calls to the end of the recurse() line

	// 4 ways to use 100

	// 3*4 ways to use 99,1,0,0

	// 3*4 ways to use 98,2,0,0
	// 3*3 ways to use 98,1,1,0

	// 3*4 ways to use 97,3,0,0
	//     ways to use 97,2,1,0

}
