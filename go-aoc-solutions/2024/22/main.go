package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/XX
	aoc.Local(part1, "part1", "sample.aoc", 37327623)
	aoc.Local(part1, "part1", "input.aoc", 13584398738)
	aoc.Local(part2, "part2", "sample2.aoc", 23)
	aoc.Local(part2, "part2", "input.aoc", 1612)
}

func next(secret int) int {
	// modulo 1000000
	secret = ((secret << 6) ^ secret) & 0xffffff
	secret = ((secret >> 5) ^ secret) & 0xffffff
	secret = ((secret << 11) ^ secret) & 0xffffff
	return secret
}

func part1(contents string) interface{} {
	nums := aoc.ParseInts(contents)
	total := 0
	for _, num := range nums {
		for range 2000 {
			num = next(num)
		}
		total += num
	}
	return total
}

var (
	modVal = 19 * 19 * 19
	size   = 19 * 19 * 19 * 19
)

func addChange(currentChanges, newChange int) int {
	// // simple bit-wise, 0x100000 or about 1 million possibilities
	// return ((currentChanges << 5) | (newChange + 9)) & 0xfffff

	// modulo, 19**4 possibilities
	a := currentChanges % (modVal)
	b := a * 19
	c := b + 9 + newChange // convert -9 to 9 into 0 to 18
	return c
}

func helpChangeString(changes int) string {
	list := [4]int{}
	list[0] = (changes % 19) - 9
	changes = changes / 19
	list[1] = (changes % 19) - 9
	changes = changes / 19
	list[2] = (changes % 19) - 9
	changes = changes / 19
	list[3] = (changes % 19) - 9
	changes = changes / 19
	return fmt.Sprintf("%d,%d,%d,%d", list[3], list[2], list[1], list[0])
}

func helpChangesFromString(s string) int {
	ss := strings.Split(s, ",")
	is := [4]int{}
	tot := 0
	for i := range 4 {
		is[i], _ = strconv.Atoi(ss[i])
		tot = tot*19 + (is[i]) + 9
	}
	return tot
}

func part2(contents string) interface{} {
	nums := aoc.ParseInts(contents)

	s := "-2,1,-1,3"
	a := helpChangesFromString(s)
	b := helpChangeString(a)
	if b != s {
		panic("BAD")
	}

	bananas := make([][]int32, len(nums))
	for mi, secret := range nums {
		bananas[mi] = make([]int32, size)
		changes := 0
		// fmt.Printf("%8d: %3d\n", secret, secret%10)
		for di := range 2000 {
			newSecret := next(secret)
			oldPrice := secret % 10
			newPrice := newSecret % 10
			secret = newSecret
			// no change on first value
			if di > -1 {
				change := newPrice - oldPrice
				changes = addChange(changes, change)
			}
			if di > 2 {
				// have at least 4 changes, set price for these changes
				// if it hasn't already been set (add 1)
				if bananas[mi][changes] == 0 {
					bananas[mi][changes] = int32((newSecret % 10) + 1) // NOTE ADD 1 TO PRICE SO 0 MEANS UNSET
				}
				// cs := helpChangeString(changes)
				// fmt.Printf("%8d: %3d (%d) - %s\n", newSecret, newPrice, newPrice-oldPrice, cs)
			} else {
				// fmt.Printf("%8d: %3d (%d)\n", newSecret, newPrice, newPrice-oldPrice)
			}
		}
	}

	// ok, let's examine the supposed best changes in 'a'
	// OH SWEET JESUS!   I just realized they changed the sample for part 2.

	best := int32(0)
	for si := range size {
		current := int32(0)
		for mi := range bananas {
			// ignore 0 which means not found
			if bananas[mi][si] > 0 {
				// subtract 1 because we added 1
				current += bananas[mi][si] - 1
			}
		}
		best = max(current, best)
	}
	return best
}
