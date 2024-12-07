# go-aoc

Library with helper functions for solving [Advent of Code](https://adventofcode.com/)
puzzles.

See:

* [General](doc/General.md) - library functions for running tests with expected output, manipulating directories/files
* [Parsing](doc/Parsing.md) - parsing functions to help with input
* [Area](doc/Area.md) - structs and funcs to help with maps/grids
* [AStar](doc/AStar.md) - structs and funcs to do A* pathfinding on Areas
* [Tui](doc/Tui.md) - package for running a step-based algorithm and displaying a nice text UI

## Easy quickstart

Add create an empty directory for your project and initialize a go module
inside, then add this module:

    mkdir mysolvers
    cd mysolvers
    go mod init example.com/mysolvers
    go get github.com/JasonGoemaat/go-aoc

You can put your files anywhere, the important thing is to have a '.go'
file with `func main()` in it and `package main` at the top.  The sample
inputs go in the same directory.   Here's how i do it:

    mysolvers
    ├── go.mod
    └── 2024
        ├── 1
        │   ├── input.aoc
        │   ├── main.go
        │   └── sample.aoc
        └── 2
            ├── input.aoc
            ├── main.go
            └── sample.aoc

Then I run from the 'mysolvers' directory with:

    go run 2024/2/main.go

Which produces output like this:

    PS C:\git\go\advent-workspace\go-aoc-solutions> go run 2024/2/main.go
    0ms part1("sample.aoc") = 2 (GOOD)
    4ms part1("input.aoc") = 660 (GOOD)
    0ms part2("sample.aoc") = 4 (GOOD)
    4ms part2("input.aoc") = 689 (BAD - Expected 889)

`2024/2/main.go` looks like this, calling `aoc.Local` for each run passing
the solving function, a description, the input file name (that is in the
same directory as the go file), and the expected output.

```go
package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	// https://adventofcode.com/2024/day/2
	aoc.Local(part1, "part1", "sample.aoc", 2)
	aoc.Local(part1, "part1", "input.aoc", 660)
	aoc.Local(part2, "part2", "sample.aoc", 4)
	aoc.Local(part2, "part2", "input.aoc", 889) // intentionally wrong
}

func part1(contents string) interface{} {
    return 0 // should actually solve puzzle part 1
}

func part2(contents string) interface{} {
    return 0 // should actually solve puzzle part 2
}
```

## Parsing

The `aoc.Local` call handles loading the named file from the same directory
as the '.go' file and passes the text contents of the file as a string to
the solver function (`part1` and `part2` is what I named them above).

Using regular expressions to split strings in go is pretty easy.  As an
example, this is used in `aoc.ParseGroups()` to split into a string that
contains empty lines dividing groups of input:

    var reDoubleCRLF = regexp.MustCompile("[\r]?[\n][\r]?[\n]")
    var sections := reDoubleCRLF.Split(text)

The `aoc.ParseGroups()` function omits the last group if empty, I was thinking
there might be a blank line at the end of the file, but the last line doesn't
end in a newline so I don't think that's necessary.

See [parsing_test.go](https://github.com/JasonGoemaat/go-aoc/blob/main/aoc/parsing_test.go)
for some examples.

### aoc.ParseGroups(content string) []string

Splits sample input with empty lines.   An example is [2024 Day 5](https://adventofcode.com/2024/day/5)
which has first a list of rules of which page must be before another, then
a section with CSV of page numbers on each line:

    47|53
    97|13

    13,53,47,97
    47,53,97,13

`aoc.ParseGroups` will return a slice with each section in it's own string.

### aoc.ParseLinesToInts(content string) [][]int

This takes a string and parses each line into a slice of numbers, identifying
numbers by consecutive digits using the regex `\d+`.

As an example, the input:

    13,53,47,97
    93: 8 4 172

Would return a slice of two slices containing integers:

```go
result := [][]int{
    {13, 53, 47, 97},
    {93, 8, 4, 172}
}
```

An example would be the same input shown in `aoc.ParseGroups` above
from [2024 Day 5](https://adventofcode.com/2024/day/5) and using
`aoc.parseLinesToints()` to get at the numbers in each:

```go
groups := aoc.ParseGroups(content)
rulesInput := aoc.ParseLinesToInts(aoc.ParseLines(groups[0]))
updates := aoc.ParseLinesToInts(aoc.ParseLines(groups[1]))
```

## Testing

There is a `aoc.ExpectJson(t *testing.T, expected, actual interface{})`
function that makes it easy to compare complex go structures and slices
to make sure they contain the same data.   It fails the test if the
JSON serialization of whatever two values that are passed is different.
It falls back to go's `%v` formatting if serialization fails for some
reason (maybe if there is a cycle?).  If it fails, it logs the expected
and actual values and since it calls `t.Helper()`, the line number
reported is from the calling test.

