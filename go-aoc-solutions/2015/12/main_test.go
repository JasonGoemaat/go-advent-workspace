package main

import (
	"encoding/json"
	"testing"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func TestRemoveRedObjects(t *testing.T) {
	s := `[1, {"a": 2, "b": "red", "c": 3}, 4]`
	var data interface{}
	_ = json.Unmarshal([]byte(s), &data)
	result := removeRedObjects(data)
	aoc.ExpectJson(t, `[1,null,4]`, result)
	// aoc.ExpectJson(t, 609043, part1("abcdef"))
	// aoc.ExpectJson(t, 1048970, part1("pqrstuv"))
}
