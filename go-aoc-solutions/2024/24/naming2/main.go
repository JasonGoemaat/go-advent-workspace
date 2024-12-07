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

func parseMe(contents string) (map[string]int, map[string]*Operation) {
	groups := aoc.ParseGroups(contents)
	reInput := regexp.MustCompile(`(...): ([01])`)
	inputMatches := reInput.FindAllStringSubmatch(groups[0], -1)

	// reOperation := regexp.MustCompile(`([a-z0-9]+) ([XORAND]+) -> ([a-z0-9]+)`)
	reOperation := regexp.MustCompile(`([a-z0-9]+) ([A-Z]+) ([a-z0-9]+) -> ([a-z0-9]+)`)
	operationMatches := reOperation.FindAllStringSubmatch(groups[1], -1)

	values := map[string]int{}
	operations := map[string]*Operation{}

	for _, arr := range inputMatches {
		operations[arr[1]] = &Operation{name: arr[1], a: arr[1], b: arr[1], op: "VALUE", subOperation: "VALUE", largestBit: parseInt(arr[1])}
		isNamed[arr[1]] = true
		named[arr[1]] = arr[1]
		if arr[2] == "0" {
			values[arr[1]] = 0
		} else {
			values[arr[1]] = 1
		}
	}

	for _, arr := range operationMatches {
		operations[arr[4]] = &Operation{name: arr[4], a: arr[1], b: arr[3], op: arr[2], largestBit: -1}
	}
	return values, operations
}

func tryNaming() int {
	changed := 0
	for k, o := range operations {
		if isNamed[k] || !isNamed[o.a] || !isNamed[o.b] {
			continue
		}
		oa, ob := operations[o.a], operations[o.b]
		if named[o.b] < named[o.a] { // in name sorted order
			oa, ob = ob, oa
		}
		// replace XOR of two values of same bit number and chainge
		// to ADD of that bit number
		if o.op == "XOR" && oa.largestBit == ob.largestBit && oa.op == "VALUE" && ob.op == "VALUE" {
			o.subOperation = "ADD"
			isNamed[k] = true
			o.largestBit = oa.largestBit
			named[k] = fmt.Sprintf("ADD(%d)", o.largestBit)
			changed++
			continue
		}

		if o.op == "AND" && oa.largestBit == ob.largestBit && oa.op == "VALUE" && ob.op == "VALUE" {
			o.subOperation = "BOTH"
			isNamed[k] = true
			o.largestBit = oa.largestBit
			named[k] = fmt.Sprintf("BOTH(%d)", o.largestBit)
			changed++
			continue
		}

		// z02 = cpb XOR wmr
		// 	cpb = djn OR tnr
		// 		djn = ADD(1) AND BOTH(0)
		// 		tnr = BOTH(1)
		// 	wmr = ADD(2)

		// ADD(1) AND BOTH(0) = PRECARRY(1)
		if o.op == "AND" && oa.subOperation == "ADD" && ob.subOperation == "BOTH" && oa.largestBit == (ob.largestBit+1) {
			o.subOperation = "PRECARRY"
			isNamed[k] = true
			o.largestBit = oa.largestBit
			named[k] = fmt.Sprintf("PRECARRY(%d)", o.largestBit)
			changed++
			continue
		}

		// now change this to CARRY(1)
		// cpb = BOTH(1) OR PRECARRY(1)
		if o.op == "OR" && oa.subOperation == "BOTH" && ob.subOperation == "PRECARRY" && oa.largestBit == (ob.largestBit) {
			o.subOperation = "CARRY"
			isNamed[k] = true
			o.largestBit = oa.largestBit
			named[k] = fmt.Sprintf("CARRY(%d)", o.largestBit)
			changed++
			continue
		}

		// now we get the actual result for a Z value:
		// z02 = ADD(2) XOR CARRY(1)
		if o.op == "XOR" && oa.subOperation == "ADD" && ob.subOperation == "CARRY" && oa.largestBit == (ob.largestBit+1) {
			o.subOperation = "ANSWER"
			isNamed[k] = true
			o.largestBit = oa.largestBit
			named[k] = fmt.Sprintf("ANSWER(%d)", o.largestBit)
			changed++
			continue
		}

		// ADD and CARRY(-1) is also CARRY
		// ADD(2) AND CARRY(1)
		if o.op == "AND" && oa.subOperation == "ADD" && ob.subOperation == "CARRY" && oa.largestBit == (ob.largestBit+1) {
			o.subOperation = "CARRY"
			isNamed[k] = true
			o.largestBit = oa.largestBit
			named[k] = fmt.Sprintf("CARRY(%d)", o.largestBit)
			changed++
			continue
		}

		// BOTH OR CARRY is CARRY
		// BOTH(2) OR CARRY(2)
		if o.op == "OR" && oa.subOperation == "BOTH" && ob.subOperation == "CARRY" && oa.largestBit == (ob.largestBit) {
			o.subOperation = "CARRY"
			isNamed[k] = true
			o.largestBit = oa.largestBit
			named[k] = fmt.Sprintf("CARRY(%d)", o.largestBit)
			changed++
			continue
		}
	}
	return changed
}

func outputPart2(w io.Writer, values map[string]int, operations map[string]*Operation, key string, depth int) {
	op, hasOperation := operations[key]
	indent := strings.Repeat("    ", depth)
	if isNamed[key] && depth > 3 {
		fmt.Fprintf(w, "%s%s = %s\n", indent, key, named[key])
		if op.subOperation != "ANSWER" {
			return
		}
	}
	if hasOperation {
		if isNamed[op.a] && isNamed[op.b] {
			oa, ob := operations[op.a], operations[op.b]
			if ob.subOperation < oa.subOperation {
				oa, ob = ob, oa
			}
			fmt.Fprintf(w, "%s%s = %s %s %s", indent, key, named[oa.name], op.op, named[ob.name])
			fmt.Fprintf(w, " (%s %s %s)\n", oa.name, op.op, ob.name)
			if depth < 3 && (op.subOperation != "BOTH" && op.subOperation != "ADD") {
				outputPart2(w, values, operations, op.a, depth+1)
				outputPart2(w, values, operations, op.b, depth+1)
			}
		} else {
			fmt.Fprintf(w, "%s%s = %s %s %s\n", indent, key, op.a, op.op, op.b)
			outputPart2(w, values, operations, op.a, depth+1)
			outputPart2(w, values, operations, op.b, depth+1)
		}
	} else {
		// fmt.Fprintf(w, "%s%s = %v\n", indent, key, values[key])
	}
}

func outputPart2b(w io.Writer, values map[string]int, operations map[string]*Operation, key string, depth int) {
	op, hasOperation := operations[key]
	indent := strings.Repeat("    ", depth)
	if isNamed[key] {
		fmt.Fprintf(w, "%s%s = %s\n", indent, key, named[key])
		if op.subOperation != "ANSWER" {
			return
		}
	}
	if hasOperation {
		if isNamed[op.a] && isNamed[op.b] {
			oa, ob := operations[op.a], operations[op.b]
			if ob.subOperation < oa.subOperation {
				oa, ob = ob, oa
			}
			fmt.Fprintf(w, "%s%s = %s %s %s", indent, key, named[oa.name], op.op, named[ob.name])
			fmt.Fprintf(w, " (%s %s %s)\n", oa.name, op.op, ob.name)
		} else {
			fmt.Fprintf(w, "%s%s = %s %s %s\n", indent, key, op.a, op.op, op.b)
			outputPart2(w, values, operations, op.a, depth+1)
			outputPart2(w, values, operations, op.b, depth+1)
		}
	} else {
		// fmt.Fprintf(w, "%s%s = %v\n", indent, key, values[key])
	}
}

func part2(contents string) interface{} {
	values, operations = parseMe(contents)
	fix(operations)
	for i := range 100 {
		fmt.Printf("%02d: %d changed\n", i, tryNaming())
	}
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

	// now try combos
	knownGood := map[string]bool{}
	tryCombos(0, knownGood)
	return 0
}

var (
	values     map[string]int
	operations map[string]*Operation
)

func solve(values map[string]int, operations map[string]*Operation) int {
	for len(operations) > 0 {
		newOperations := map[string]*Operation{}
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

func tryCombos(bit int, knownGood map[string]bool) {
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
		newOperations := map[string]*Operation{}
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

func swap(operations map[string]*Operation, k1, k2 string) {
	a, b := operations[k1], operations[k2]
	operations[k1], operations[k2] = b, a

}

func fix(operations map[string]*Operation) {
	// candidates: z14, vss
	swap(operations, "z14", "vss")

	// candidates, see README - hjf, kdh
	swap(operations, "hjf", "kdh")

	// bit 31/32:
	swap(operations, "z31", "kpp")

	// swap z35 and sgj
	swap(operations, "z35", "sgj")
}

type Operation struct {
	name         string
	a            string
	b            string
	op           string
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
