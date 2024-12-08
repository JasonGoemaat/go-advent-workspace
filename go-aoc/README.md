# go-aoc

Library with helper functions for solving [Advent of Code](https://adventofcode.com/)
puzzles.

## Easy quickstart

Add create an empty directory for your project and initialize a go module
inside, then add this module:

    mkdir mysolvers
    cd mysolvers
    go mod init example.com/mysolvers
    go get github.com/JasonGoemaat/go-aoc

If you want to just start solving, create a directory for each day and add
a 'main.go' file, i.e. `day1/main.go`.   Create `part1` and `part2` functions
to solve the puzzles.   Download the input to `input.txt` and create a
`sample.txt` with the sample from the instructions.   It should have these:

    input.txt
    main.go
    sample.txt

Use this for the `main.go` file:

```go
package main

import "github.com/JasonGoemaat/go-aoc"

func main() {
    aoc.SolveLocal(part1, part2)
}

func part1(contents string) string {
    return fmt.Sprintf("%d", 0)
}

func part2(contents string) string {
    return "tbd"
}
```

Run from the `2024/1` directory with:

    go run .

Run from the root with:

    go run 2024/1/main.go

The `part1` and `part2` functions should take the entire contents of the
puzzle data you would get from `sample.txt` or `input.txt` and
return the puzzle answer as a string.   `aoc.SolveLocal()` will look for
`sample.txt` and `input.txt` in the same directory as the `.go` file
calling it and call the `part1` and `part2` functions you pass for
each of them and report their returned strings.  It'll also report the
time it took for each.

