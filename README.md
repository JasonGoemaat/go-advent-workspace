# Workspace - Advent of Code

## Quickstart

Work in `go-aoc-solutions`.  One directory per year, one subdirectory
per day, i.e. `go-aoc-solutions/2015/1`.

Then I will create `sample.aoc` and paste in the sample puzzle data and
download the input to `input.aoc`.

And finally a `main.go`, this will use my `go-aoc` package with helper
functions for parsing files and running test functions with the inputs.
See `go-aoc-solutions/25/2` for an example.  

The basic format will be something like this, imagine a file `sample.aoc`
in the same directory as `main.go` with one line '25,-7' and the solution
to part 1 is to add the two numbers together and expecting '18' as the
answer:

```go
package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 18)
}

func part1(string contents) {
    // aoc.ParseInts returns an array of integers
    ints := aoc.ParseInts(contents)
    return ints[0] + ints[1]
}
```

## go-aoc module

To install the module:

    go get github.com/JasonGoemaat/go-aoc

To use it, import like so:

```go
import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)
```

The `Local` function takes four arguments:

1. a function to run, passing it the string contents of the file
2. a string to display as the name in the output
3. a file name, relative to the directory your go file is in
4. the expected resulting value

This example calls the function `part`, passing it a string
with the contents of the file `sample.aoc` in the same directory,
and expects the return value to be the integer 18:

```go
func main() {
	aoc.Local(part1, "part1", "sample.aoc", 18)
}
```

This is a sample `part1` function that expects a string
containing two integers and adds them together.   Using the
call above with the contents of `sample.aoc` being `27 -9`
should result in the two values being added together and
returning the integer `18`:

```go
func part1(string contents) {
    // aoc.ParseInts returns an array of integers
    ints := aoc.ParseInts(contents)
    return ints[0] + ints[1]
}
```

