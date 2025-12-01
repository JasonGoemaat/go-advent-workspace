# Workspace - Advent of Code

## Quickstart

Work in `go-aoc-solutions`.  One directory per year, one subdirectory
per day, i.e. `go-aoc-solutions/2015/1`.

Then I will create `sample.aoc` and paste in the sample puzzle data and
download the input to `input.aoc`.

And finally a `main.go`

## Notes

Thinking of creating a `go-aoc` module that I will use in other modules for
creating solvers.

A lot of it is creating a way to call them.  A thought is to keep using
`cobra-cli` and creating commands in a more concise way in my main app.

## https://go.dev/doc/tutorial/workspaces

Creating library (from root here where README.md exists):

    mkdir go-aoc
    cd go-aoc
    go mod init github.com/JasonGoemaat/go-aoc

Creating app for 2024 (from root here where README.md exists):

    mkdir go-aoc-2024
    cd go-aoc-2024
    go mod init github.com/JasonGoemaat/go-aoc

Creating workspace and adding the two modules (from root):

    go work init .\go-aoc .\go-aoc-2024

In the 'go-aoc' folder for my library I just created a 'main.go'
file with the package name 'aoc':

```go
package aoc

import "fmt"

func SayHello() {
	fmt.Println("Hello, World! (from go-aoc/main.go SayHello())")
}

func init() {
	fmt.Println("go-aoc/main.go init() running")
}
```

In the 'go-aoc-2024' folder I created a 'main.go' for the app
with a main function.   I had to add the import manually, but it
found it in the workspace without me ever having to actually
put the code on github.   The workspace file knows where they
are.

```go
package main

import (
	"fmt"

	"github.com/JasonGoemaat/go-aoc"
)

func main() {
	fmt.Println("go-aoc-2024 main() running")
	fmt.Println("calling aoc.SayHello() from other module in workspace")
	aoc.SayHello()
}
```

Running the app in `go-aoc-2024` shows the order of execution:

```go
PS C:\git\go\advent-workspace\go-aoc-2024> go run .
go-aoc/main.go init() running
go-aoc-2024 init() running
go-aoc-2024 main() running
calling aoc.SayHello() from other module in workspace
Hello, World! (from go-aoc/main.go SayHello())
```

## .go file location

I added a function in the `aoc` library.   This uses `runtime.Caller(n)` which
will give the path to the code file.  If called with `0` it would return
the current `.go` file.   When called with `1`, it returns the path of the
`.go` file with the function that made the call to `GetDir()`.   I'm guessing
this can walk back up the call stack.

```go
func GetDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}
```

## Re-organization of puzzle repos

This will be handy for me how I want to re-arrange the puzzles.  I want to
be able to create a directory for the day and just create a single file.
I'm thinking one of two things though:

1. By convention have 'sample.aoc' and 'index.txt' in the same directory for
the tests.  Run the tests with `go run 2024\1\main.go`
2. Have a main app that uses `cobra-cli` and can be called easily with arguments
like `go run . 2024day1`.

For the simple one, I'm thinking of calling it like this:

```go
func main() {
    aoc.SolveLocal(part1, part2)
}

func part1(contents string) string {
    return aoc.Output(123) // use %v for any type
}

func part2(contents string) string {
    return "tbd"
}
```

I to use all the command-line options, I would need a little more.
I like the simplicity above though.  It keeps my solution files
small.

```go
func init() {
    aoc.RegisterDay(2024, 1, part1, part2)
}

// ... Part1 and Part2 from above
```

This is pretty easy too.   In this, I think I would have a file replacing
`root.go` in my old repo to create the root command:

```go
func init() {
	cobra.OnInitialize(initConfig)
    //... create root command, handle config and parameters
```

And RegisterDay would create the command and add it to root:

```go
// use map for year so day is sub-command of appropriate year?
// map[int]cobra.Command

func RegisterDay(year, day int, part1, part2 func(string) string) {
    var cmd = &cobra.Command{
        Use:   fmt.Sprintf("%d-%d", year, day),
        Short: fmt.Sprintf(`Advent of Code %d day %d
https://adventofcode.com/%d/day/%d`, year, day, year, day),
        Long:  ``,
        Run: func(cmd *cobra.Command, args []string) {
        }
    }
    rootCmd.AddCommand
}
```

Thinking now I may want to have each year be a subcommand, and each day a
subcommand of that year's command.   Then this would run 2024 day 1

    go run . 2024 1

