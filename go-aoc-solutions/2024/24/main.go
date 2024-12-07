package main

import (
	"bufio"
	"fmt"
	"io"
	_ "io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/24
	aoc.Local(part1, "part1", "sample.aoc", 4)
	aoc.Local(part1, "part1", "sample2.aoc", 2024)
	aoc.Local(part1, "part1", "input.aoc", 51837135476040)
	// see subdirectory 'naming2' for how I solved Part 2 by looking manually
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	// aoc.Local(part2, "part2", "input.aoc", 0)
}

func parseMe(contents string) (map[string]int, map[string]Operation) {
	groups := aoc.ParseGroups(contents)
	reInput := regexp.MustCompile(`(...): ([01])`)
	inputMatches := reInput.FindAllStringSubmatch(groups[0], -1)

	// reOperation := regexp.MustCompile(`([a-z0-9]+) ([XORAND]+) -> ([a-z0-9]+)`)
	reOperation := regexp.MustCompile(`([a-z0-9]+) ([A-Z]+) ([a-z0-9]+) -> ([a-z0-9]+)`)
	operationMatches := reOperation.FindAllStringSubmatch(groups[1], -1)

	values := map[string]int{}
	for _, arr := range inputMatches {
		if arr[2] == "0" {
			values[arr[1]] = 0
		} else {
			values[arr[1]] = 1
		}
	}

	operations := map[string]Operation{}
	for _, arr := range operationMatches {
		operations[arr[4]] = Operation{a: arr[1], b: arr[3], op: arr[2]}
	}
	return values, operations
}

func part1(contents string) interface{} {
	values, operations := parseMe(contents)
	for len(operations) > 0 {
		newOperations := map[string]Operation{}
		for k, op := range operations {
			a, haveA := values[op.a]
			b, haveB := values[op.b]
			if haveA && haveB {
				switch op.op {
				case "AND":
					values[k] = a & b
				case "OR":
					values[k] = a | b
				case "XOR":
					values[k] = a ^ b
				}
			} else {
				newOperations[k] = op
			}
		}
		operations = newOperations
	}

	result := 0
	for i := 0; true; i++ {
		k := fmt.Sprintf("z%02d", i)
		v, exists := values[k]
		if !exists {
			break
		}
		result = result | (v << i)
	}
	return result
}

func outputPart2(w io.Writer, values map[string]int, operations map[string]Operation, key string, depth int) {
	op, hasOperation := operations[key]
	indent := strings.Repeat("    ", depth)
	if hasOperation {
		fmt.Fprintf(w, "%s%s = %s %s %s\n", indent, key, op.a, op.op, op.b)
		outputPart2(w, values, operations, op.a, depth+1)
		outputPart2(w, values, operations, op.b, depth+1)
	} else {
		// fmt.Fprintf(w, "%s%s = %v\n", indent, key, values[key])
	}
}

func outputPart2Known(w io.Writer, values map[string]int, operations map[string]Operation, key string, depth int, knownGood map[string]bool) {
	if !knownGood[key] {
		op, hasOperation := operations[key]
		indent := strings.Repeat("    ", depth)
		if hasOperation {
			fmt.Fprintf(w, "%s%s = %s %s %s\n", indent, key, op.a, op.op, op.b)
			outputPart2Known(w, values, operations, op.a, depth+1, knownGood)
			outputPart2Known(w, values, operations, op.b, depth+1, knownGood)
		}
	}
}

func part2Output(contents string) interface{} {
	values, operations := parseMe(contents)
	for i := range 46 { // up to Z45
		// xk, yk, zk := fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i), fmt.Sprintf("z%02d", i)
		// fmt.Printf("%s=%v, %s=%v, %s: %v\n", xk, values[xk], yk, values[yk], zk, operations[zk])
		fullName := aoc.GetSubPath(filepath.Join("zs", fmt.Sprintf("z%02d.md", i)))
		f, _ := os.Create(fullName)
		defer f.Close()
		w := bufio.NewWriter(f)
		defer w.Flush()
		outputPart2(w, values, operations, fmt.Sprintf("z%02d", i), 0)
		w.Flush()
	}
	return 0
}

func solve(values map[string]int, operations map[string]Operation) int {
	for len(operations) > 0 {
		newOperations := map[string]Operation{}
		for k, op := range operations {
			a, haveA := values[op.a]
			b, haveB := values[op.b]
			if haveA && haveB {
				switch op.op {
				case "AND":
					values[k] = a & b
				case "OR":
					values[k] = a | b
				case "XOR":
					values[k] = a ^ b
				}
			} else {
				newOperations[k] = op
			}
		}
		operations = newOperations
	}

	result := 0
	for i := 0; true; i++ {
		k := fmt.Sprintf("z%02d", i)
		v, exists := values[k]
		if !exists {
			break
		}
		result = result | (v << i)
	}
	return result
}

func tryCombos(values map[string]int, operations map[string]Operation, bit int, knownGood map[string]bool) {
	// clear
	for i := range 45 {
		xk, yk := fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i)
		values[xk] = 0
		values[yk] = 0
	}
	zk1, zk2 := fmt.Sprintf("z%02d", bit), fmt.Sprintf("z%02d", bit+1)
	xk, yk := fmt.Sprintf("x%02d", bit), fmt.Sprintf("y%02d", bit)

	isGood := true
	for i := 0; i < 4; i++ {
		newValues := map[string]int{}
		for k, v := range values {
			newValues[k] = v
		}
		newOperations := map[string]Operation{}
		for k, v := range operations {
			newOperations[k] = v
		}
		v1, v2 := 0, 0
		if (i & 1) > 0 {
			newValues[xk] = 1
			v1 = 1 << bit
		}
		if (i & 2) > 0 {
			newValues[yk] = 1
			v2 = 1 << bit
		}
		result := solve(newValues, newOperations)
		sum := v1 + v2
		if result != sum {
			fmt.Printf("Bit %02d: %d + %d = %d%d got %d expected %d\n", bit, newValues[xk], newValues[yk], newValues[zk2], newValues[zk1], result, sum)
			isGood = false
		}
	}

	var markGood func(name string)
	markGood = func(name string) {
		if op, exists := operations[name]; exists {
			if knownGood[name] {
				return
			}
			knownGood[name] = true
			markGood(op.a)
			markGood(op.b)
		}
	}
	if isGood {
		// mark all touched nodes as good
		markGood(zk1)
	}
}

// I'm going to rename nodes, if they're in my list,
// when outputting I will output their name

func swap(operations map[string]Operation, k1, k2 string) {
	a, b := operations[k1], operations[k2]
	operations[k1], operations[k2] = b, a

}

func fix(operations map[string]Operation) {
	// candidates: z14, vss
	swap(operations, "z14", "vss")
	// swap(operations, "kdh", "sbg") // Maybe kdh and sbg should be swapped?
}

func part2(contents string) interface{} {
	values, operations := parseMe(contents)
	fix(operations)
	knownGood := map[string]bool{}
	for i := range 45 { // up to x/y 44
		// try 4 combinations for each bit, 00, 01, 10, 11 and see what we get
		tryCombos(values, operations, i, knownGood)
	}

	// now output all that we don't have knownGood for
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := range 46 { // up to Z45
		// xk, yk, zk := fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i), fmt.Sprintf("z%02d", i)
		// fmt.Printf("%s=%v, %s=%v, %s: %v\n", xk, values[xk], yk, values[yk], zk, operations[zk])
		key := fmt.Sprintf("z%02d", i)
		if !knownGood[key] {
			outputPart2Known(w, values, operations, key, 0, knownGood)
		}
	}

	return 0
}

type Operation struct {
	a  string
	b  string
	op string
}
