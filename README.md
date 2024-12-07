# Workspace - Advent of Code

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
