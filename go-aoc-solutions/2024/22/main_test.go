package main

import (
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

// Test sample next 10 secrets starting at 123
func TestNext(t *testing.T) {
	secret := 123
	secret = next(secret)
	aoc.ExpectJson(t, 15887950, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 16495136, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 527345, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 704524, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 1553684, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 12683156, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 11100544, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 12249484, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 7753432, secret)
	secret = next(secret)
	aoc.ExpectJson(t, 5908254, secret)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}

func TestSample(t *testing.T) {
	data := [4][2]int{{1, 8685429}, {10, 4700978}, {100, 15273692}, {2024, 8667524}}
	for _, sample := range data {
		secret := sample[0]
		for range 2000 {
			secret = next(secret)
		}
		aoc.ExpectJson(t, sample[1], secret)
	}
}
