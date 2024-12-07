## TUI - Text User Interface

```go
import tui "github.com/JasonGoemaat/go-aoc/aoc/tui"
```

Example: [../samples/TuiAStar/main.go](../samples/TuiAStar/main.go)

This is where I found out how using interfaces in go worked for the most part.
This package is made so you can create a struct implementing the 
`tui.TuiActions` interface:

```go
TuiActions       interface {
    Render() string
    IsDone() bool
    Step()
    GetSolution() interface{}
}
```

So you create a struct that contains the data you need to render and solve
the puzzle in steps.  The sample has this:

```go
type MyState struct {
	Title    string
	Contents string
	astar    *astar.AStar
}
```

Then you implement the functions on the interface and run the program like
this in the sample:

```go
func main() {
	tm := NewMyState(sample)
	tui.RunProgram(tm)
	fmt.Printf("The solution is: %s\n", tm.GetSolution())
}
```

Alternatively if you don't want the visualization, you should be able to use
your structure like this:

```go
func main() {
	tm := NewMyState(sample)
	// tui.RunProgram(tm)
	for !tm.IsDone() {
		tm.Step()
	}
	fmt.Printf("The solution is: %s\n", tm.GetSolution())
}
```

The methods for my solution are defined to fit the `tui.Actions` interface
which allows us to pass a value of `MyState` to `tui.RunProgram()` because
the structure implements those interfaces.   I use a pointer as the base
for the functions so they can modify the underlying structure.

```go
func (state *MyState) Render() string
func (state *MyState) IsDone() bool
func (state *MyState) Step()
func (state *MyState) GetSolution() interface{}
```

### UI

When you call `tui.RunProgram()`, what it does is call your `Render()` function
which returns a string to display the current state of the solution, initially
before any step has been taken.  This can use the
[BubbleTea](https://github.com/charmbracelet/bubbletea)
and [LipGloss](https://github.com/charmbracelet/lipgloss) packages which let
you do all kinds of cool things like colorizing and drawing borders.

The UI starts in a paused state and accepts these keypresses:

* `ESC`, `q` - Quit
* `s` - Single step, disable auto-running if enabled
* `SPACE` - Toggle auto-running
* `r` - Toggle rendering
* `UP` - Increase auto-run delay * 10 (starts at 100ms, up to 1000ms)
* `DOWN` - Decrease auto-run delay / 10 (starts at 100ms, down to 10ms, 1ms, 0ms)

If doing a single step 's' or auto-running steps, it will call your `Step()`
function and stop when your `IsDone()` function returns true and return from the
call to `RunProgram()`.
