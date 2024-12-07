package main

import (
	"fmt"
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestNOT(t *testing.T) {
	aoc.ExpectJson(t, 0xff00, part1("NOT 255 -> a"))
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}

func TestOR(t *testing.T) {
	aoc.ExpectJson(t, 68|3, part1("68 OR 3 -> a"))
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}

func TestAND(t *testing.T) {
	aoc.ExpectJson(t, 0x27, part1("255 AND 39 -> a"))
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}

func TestAND_OR(t *testing.T) {
	aoc.ExpectJson(t, 0x27, part1("255 AND x -> a\n6 OR 35 -> x")) // make x 39
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}

var samples = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`

func TestSamples(t *testing.T) {
	check := func(w string, v int) {
		s := fmt.Sprintf("%s\n%s OR 0 -> a", samples, w)
		aoc.ExpectJson(t, v, part1(s))
	}
	// d: 72
	// e: 507
	// f: 492
	// g: 114
	// h: 65412
	// i: 65079
	// x: 123
	// y: 456
	check("d", 72)
	check("e", 507)
	check("f", 492)
	check("g", 114)
	check("h", 65412)
	check("i", 65079)
	check("x", 123)
	check("y", 456)
}
