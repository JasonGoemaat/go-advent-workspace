package main

import (
	"fmt"
	"time"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/X
	// aoc.Local(part1, "part1", "sample.aoc", 0)
	// aoc.Local(part1, "part1", "input.aoc", 492982) // whew - 26 seconds
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 0)
}

// works, but is so slow because new strings are allocated and freed
func handle(s string) string {
	next := ""
	for i := 0; i < len(s); {
		ch := s[i]
		cnt := 0
		for ; i < len(s) && s[i] == ch; i++ {
			cnt++
		}
		next = next + fmt.Sprintf("%d%c", cnt, ch)
	}
	return next
}

// might use extra memory (14mb allocatio at end), but in big
// chunks and is very fast
func handle2(s string) string {
	behind := make([]byte, len(s)*2)
	size := 0
	for i := 0; i < len(s); {
		ch := s[i]
		cnt := 0
		for ; i < len(s) && s[i] == ch; i++ {
			cnt++
		}
		behind[size] = byte(0x30 + cnt)
		behind[size+1] = ch
		size += 2
		// next = next + fmt.Sprintf("%d%c", cnt, ch)
	}
	return string(behind[0:size])
}

func part1(contents string) interface{} {
	current := contents
	list := []string{current}
	fmt.Println(current)
	for i := 0; i < 40; i++ {
		current = handle(current)
		if len(current) > 60 {
			fmt.Printf("%d: %s  (+%d...)\n", i, current[0:50], len(current)-50)
		} else {
			fmt.Printf("%d: %s\n", i, current)
		}
		list = append(list, current)
	}
	return len(current)
}

// func handleRegexp(contents string) string {
// // 	const runLengthEncode = (input: string) =>
// //   input.replaceAll(/((\d)\2*)/g, (match) => `${match.length}${match[0]}`);

// // export const partOne = (input: AOCInput, repetition = 40): number => {
// //   let speak = input.lines().filter(Boolean).nth();
// //   while (repetition--) speak = runLengthEncode(speak);
// //   return speak.length;
// // };
// 	// rxMain := regexp.MustCompile(`((\d)\2*)`)
// }

func part2(contents string) interface{} {
	current := contents
	list := []string{current}
	fmt.Println(current)
	for i := 0; i < 50; i++ {
		start := time.Now()
		current = handle2(current)
		end := time.Now()
		ms := end.Sub(start).Seconds()
		if len(current) > 60 {
			fmt.Printf("%d: %s  (+%d...) in %1.3fs\n", i, current[0:50], len(current)-50, ms)
		} else {
			fmt.Printf("%d: %s in %1.3fs\n", i, current, ms)
		}
		list = append(list, current)
	}
	return len(current)
}
