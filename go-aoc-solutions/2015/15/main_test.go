package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestCalculateScore(t *testing.T) {
	// func calculateScore(ingredients [][]int, counts []int) int {
	// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	// Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
	ingredients := [][]int{
		{-1, -2, 6, 3},
		{2, 3, -2, -1},
	}
	counts := []int{44, 56}
	score := calculateScore(ingredients, counts)
	aoc.ExpectJson(t, 62842880, score)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}
