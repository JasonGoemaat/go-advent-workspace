package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/X
	aoc.Local(part1, "part1", "sample.aoc", 609043)
	aoc.Local(part1, "part1", "sample.aoc", 1048970)
	aoc.Local(part1, "part1", "input.aoc", 282749)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 9962624)
}

func check(key string, addon int) bool {
	text := key + strconv.Itoa(addon)
	hash := md5.Sum([]byte(text))
	first := hex.EncodeToString(hash[:3])[:5]
	return first == "00000"
}

func check2(key string, addon int) bool {
	text := key + strconv.Itoa(addon)
	hash := md5.Sum([]byte(text))
	first := hex.EncodeToString(hash[:3])
	return first == "000000"
}

func part1(contents string) interface{} {
	for i := 1; true; i++ {
		if check(contents, i) {
			return i
		}
	}
	return -1
}

func part2(contents string) interface{} {
	for i := 1; true; i++ {
		if check2(contents, i) {
			return i
		}
	}
	return -1
}
