package main

import (
	"encoding/json"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2015/day/12
	aoc.Local(part1, "part1", "sample.aoc", 3)
	aoc.Local(part1, "part1", "input.aoc", 156366)
	aoc.Local(part2, "part2", "sample.aoc", 3)
	aoc.Local(part2, "part2", "input.aoc", 96852)
}

func part1(contents string) interface{} {
	nums := aoc.ParseInts(contents)
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func removeRedObjects(obj interface{}) interface{} {
	switch value := obj.(type) {
	case []interface{}:
		for i := range value {
			value[i] = removeRedObjects(value[i])
		}
	case map[string]interface{}:
		// it's an object, it it has a property with the VALUE 'red', return nil
		for _, v := range value {
			switch s := v.(type) {
			case string:
				if s == "red" {
					return nil
				}
			}
		}
		// ok, it doesn't have a 'red' property itself, but check each property
		// for them
		for k, v := range value {
			value[k] = removeRedObjects(v)
		}
	}
	return obj
}

func part2(contents string) interface{} {
	// need escaping I think: .^$*+?()[{\|
	// 	rxRedObjects := regexp.MustCompile(`\{[^}]*}`)
	// probably not that simple, what if an object contains an array?
	// this would be easier in javascript or something with dynamic processing.
	// I don't even know how I would do this in go.  Well, maybe I should
	// look into GO json marshalling more, from what I've seen so far it is
	// mostly for well-defined structs.

	// let's see what it does:
	// var data map[string]interface{} // this works on sample.aoc because it's an object
	// var data map[string][]interface{} // this works on input.aoc because it's an array

	var data interface{} // try fully generic - Works for both, either get []interface{} for an array or map[string]interface{} for object
	contentBytes := []byte(contents)
	err := json.Unmarshal(contentBytes, &data)
	if err != nil {
		panic(err)
	}
	// switch t := data.(type) {
	// case map[string]interface{}:
	// 	fmt.Printf("It's an object!\n")
	// case []interface{}:
	// 	fmt.Printf("It's an array!\n")
	// 	fmt.Printf("Slice length: %d\n", len(t))
	// default:
	// 	fmt.Printf("It's an unknown type!  (%T)\n", t)
	// }
	// well, that just throws an error on input, let's try sample
	removed := removeRedObjects(data)
	newJsonBytes, err := json.Marshal(removed)
	if err != nil {
		panic(err)
	}

	nums := aoc.ParseInts(string(newJsonBytes))
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
