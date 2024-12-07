package main

import (
	"bufio"
	"fmt"
	"io"
	_ "io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/24
	// aoc.Local(part1, "part1", "sample.aoc", 4)
	// aoc.Local(part1, "part1", "sample2.aoc", 2024)
	// aoc.Local(part1, "part1", "input.aoc", 51837135476040)
	// aoc.Local(part2, "part2", "sample.aoc", 0)
	aoc.Local(part2, "part2", "input.aoc", 0)
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

func outputPart2(w io.Writer, values map[string]int, operations map[string]Operation, key string, depth int) {
	op, hasOperation := operations[key]
	indent := strings.Repeat("    ", depth)
	if isNamed[key] {
		fmt.Fprintf(w, "%s%s = %s\n", indent, key, named[key])
		return
	}
	if hasOperation {
		if isNamed[op.a] && isNamed[op.b] {
			fmt.Fprintf(w, "%s%s = %s %s %s\n", indent, key, named[op.a], op.op, named[op.b])
		} else {
			fmt.Fprintf(w, "%s%s = %s %s %s\n", indent, key, op.a, op.op, op.b)
			outputPart2(w, values, operations, op.a, depth+1)
			outputPart2(w, values, operations, op.b, depth+1)
		}
	} else {
		// fmt.Fprintf(w, "%s%s = %v\n", indent, key, values[key])
	}
}

var (
	values     map[string]int
	operations map[string]Operation
)

func nameOperations() int {
	changed := 0
	for k, o := range operations {
		// can only name if not already named and if both children are named
		if isNamed[k] || !isNamed[o.a] || !isNamed[o.b] {
			continue
		}
		// oa, o.b := operations[o.a], operations[o.b]
		// keep in name-sorted order
	}
	return changed
}

func part2(contents string) interface{} {
	values, operations := parseMe(contents)
	fix(operations)
	nameInitial(values, operations)
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

type Operation struct {
	a            string
	b            string
	op           string
	name         string
	subOperation string
	largestBit   int
}

var named = map[string]string{}
var isNamed = map[string]bool{}
var isRoot = regexp.MustCompile(`[xy]\d\d|RAW[xy]\d\d`)

func nameNode(key, name string) {
	isNamed[key] = true
	named[key] = name
}

var reNumbers = regexp.MustCompile(`[-]?\d+`)

func parseInt(s string) int {
	numbers := reNumbers.FindString(s)
	v, err := strconv.Atoi(numbers)
	if err != nil {
		panic("Bad integer string")
	}
	return v
}

func nameInitial(values map[string]int, operations map[string]Operation) {
	for k, _ := range values {
		nameNode(k, fmt.Sprintf("RAW%s", k))
	}

	for k, o := range operations {
		if isRoot.MatchString(o.a) && isRoot.MatchString(o.b) {
			if o.a[1:3] == o.b[1:3] {
				isNamed[k] = true
				if o.op == "XOR" {
					named[k] = fmt.Sprintf("ADD(%s)", o.a[1:3])
					o.subOperation = "ADD"
					o.largestBit = parseInt(o.a[1:3])
				} else if o.op == "OR" {
					o.subOperation = "EITHER"
					o.largestBit = parseInt(o.a[1:3])
					named[k] = fmt.Sprintf("EITHER(%s)", o.a[1:3])
				} else if o.op == "AND" {
					o.subOperation = "BOTH"
					o.largestBit = parseInt(o.a[1:3])
					named[k] = fmt.Sprintf("BOTH(%s)", o.a[1:3])
				} else {
					panic("BAD OPERATION!")
				}
			}
		}
	}
	nameNode("cpb", "CARRY(01)")
	nameNode("tdw", "CARRY(02)")
	nameNode("vbb", "CARRY(03)")
	nameNode("wht", "CARRY(04)")
	nameNode("tng", "CARRY(05)")
	nameNode("gmw", "CARRY(06)")
	nameNode("sws", "CARRY(07)")
	nameNode("vpd", "CARRY(08)")
	nameNode("msb", "CARRY(09)")
	nameNode("ntv", "CARRY(10)")
	nameNode("wpk", "CARRY(11)")
	nameNode("dkh", "CARRY(12)")
	nameNode("nhg", "CARRY(13)")
	nameNode("vss", "CARRY(14)")
	nameNode("rgp", "CARRY(15)")
	nameNode("qtn", "CARRY(16)")
	nameNode("jhw", "CARRY(17)")
	nameNode("ZZZ", "CARRY(00)")
	nameNode("ZZZ", "CARRY(00)")
	nameNode("ZZZ", "CARRY(00)")
}

// func tryNaming(values map[string]int, operations map[string]Operation) {
// 	hasNamed := false
// 	for k, o := range operations {
// 		if isNamed[k] || !isNamed[o.a] || !isNamed[o.b] {
// 			continue
// 		}
// 		oa, ob := operations[o.a], operations[o.b]
// 		// look for CARRY(b-1) AND ADD(b), replace with PRECARRY(b)
// 		if o.op == "XOR" {
// 			if oa.subOperation != "ADD" {
// 				oa, ob = ob, oa
// 			}
// 			if oa.subOperation == "ADD" && ob.op == "OR" {

// 			}
// 		}
// 		if ob.subOperation !=
// 	}
// }

// CHECK FOR A SITUATION LIKE THIS:
// z19 = twb XOR tsm
// 		twb = rqv OR pns <--- look here
// 			rqv = CARRY(17) AND ADD(18) <-- rename 'PRECARRY(18)'
// 			pns = BOTH(18)
// 		tsm = ADD(19)
