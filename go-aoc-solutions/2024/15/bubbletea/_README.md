# Trying to add bubbletea

* Reddit post: https://www.reddit.com/r/adventofcode/comments/1hem5hg/2024_day_15_part_1_go_move_little_robot/
* sirgwain's advent repo: https://github.com/sirgwain/advent-of-code-2024
* BubbleTea: https://github.com/charmbracelet/bubbletea
* Blog explaining commands: https://charm.sh/blog/commands-in-bubbletea/
* lipgloss: https://github.com/charmbracelet/lipgloss

Installed BubbleTea with:

    go get github.com/charmbracelet/bubbletea

He starts with this:

    p := tui.NewViewportProgram(tui.NewModel("Day 15"))

And has a custom Model:

    type Model struct {
        ready         bool
        viewport      viewport.Model
        title         string
        viewportLines []string
        minWidth      int
        windowWidth   int
        windowHeight  int
    }

    func NewModel(title string) Model {
        return Model{title: title}
    }

'viewport' wasn't defined, it suggested and added this import, but it wasn't found:

    "github.com/charmbracelet/bubbles/viewport"

Running this gave an error:

    go get package github.com/charmbracelet/bubbles/viewport
    go: malformed module path "package": missing dot in first path element

So I don't know what that syntax is, but doing that from the quick fix tooltip on the import worked.

## Hello, world!

BAREBONES tea program.   `tea1()` is called by `main()` in another file.

```go
package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type tea1model struct {
}

func (m tea1model) Init() tea.Cmd {
	return nil
}

func (m tea1model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m tea1model) View() string {
	return "Hello, world!"
}

func tea1() {
	p := tea.NewProgram(tea1model{})
	m, err := p.Run()
	if err != nil {
		panic("tea did not run!")
	}
	fmt.Printf("tea did run!  model is: %v\n", m)
}
```

So to be a `tea.Model` you need to  the methods `Init()`, `Update()`, and `View()`
and that's pretty much it.  This program will never end and just sit there after
displaying 'Hello, world!'.

## Update Loop

It seems the `Update()` method is the heart of the 'program' you run.   Looks like
tea can handle keyboard input and send you commands.   You can do different things
based on command type.

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
```

`tea.KeyMsg` is what it sends for keyboard input.   You can call `msg.String()`
and access it in a friendly way, such as comparing with "ctrl+c" or "up" for the
up arrow.  I'm going to try to add x and y to my model and let you move around
the message.  Also adding ctrl+c and esc to quit.

That worked pretty well.   I store the x and y in the model:

```go
type tea2model struct {
	x, y int
}
```

My update method will quit or change x and y based on keys:

```go
func (m tea2model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up":
			m.y = max(0, m.y-1)
		case "down":
			m.y = m.y + 1
		case "left":
			m.x = max(0, m.x-1)
		case "right":
			m.x = m.x + 1
		}
	}
	return m, nil
}
```

And my `View()` now alters the string output accordingly and draws
'lines' leading to my output.

```go
func (m tea2model) View() string {
	sb := strings.Builder{}
	spaces := strings.Repeat(" ", m.x+6)
	for _ = range m.y {
		sb.WriteString(spaces)
		sb.WriteRune('|')
		sb.WriteRune('\n')
	}
	for _ = range m.x {
		sb.WriteRune('-')
	}
	sb.WriteString("Hello, world!")
	return sb.String()
}
```

## Process

How sirgwain does it:

```go
// run the solver in a gouroutine and Send a message to the bubbletea program to update the viewport on each step
go func() {
    // update the UI
    board.onStep = func() {
        content := fmt.Sprintf("%s\n%s", board.view(), board.viewSolution())
        p.Send(tui.UpdateViewport(content, board.width))

        if d.Options.Delay > 0 {
            time.Sleep(time.Millisecond * time.Duration(d.Options.Delay))
        }

    }

    board.solve()
}()
```

This has a 'board' that represents the puzzle and sets an 'onStep' function.  This function
sends an ' UpdateViewport' message to the bubbletea program.   I don't think `Run()` is ever
called on the program, so the updates are done in this loop.

The `board.solve()`  method loops through each step and calls that `onStep()` function
after each iteration.   So it just runs normally and uses the code in that function to draw
the display by sending an `updateViewport` message created by that function.

```go
func (b *day15Board) solve() {
	for _, dir := range b.moves {
		b.moveRobot(dir)
		b.solution = b.gps()
		if b.onStep != nil {
			b.onStep()
		}	
    }
}

type updateViewport struct {
    content string
    width   int
    height  int
}

func UpdateViewport(content string, width int) tea.Msg {
	return updateViewport{content: content, width: width}
}
```

Ah, `tea.Msg` is just `interface{}` so it can be ANYTHING. (maybe not?)

```go
type Msg interface{}
```

The reason we have a type declared is so it can be used in our 'Update()' method
on the model in the switch statement:

```go
switch msg := msg.(type) {
case updateViewport:
```

## Stepping?

I think I want to do it so that I can pause and step through one at a time.
Undo would be nice also, but would be more of a pain.   Having the previous
result left or right of the current state would be cool.

To do this I think I need to change my loops so that I have a function that
will operate on them.  I can have a command that will auto-run.   Model
will have flags for rendering and automatic.

For automatic, you can send a `step` command which will call the step function
from the solver, then send a `render` command.

The `render` command will render things and if automatic is selected it
will send another `step` command.

Yeah, that works pretty well.   I added delay property to the model.
GO idiosyncrasy - I had to use `func (m *model)` and not `func (m model)`
when I did the update or things I changed in the function would not update
the existing struct as it's passed by value.   No wonder there were warning
squigglies :)

So the puzzle starts stopped.   Here are the commands:

* `S` - single step (pauses if running auto)
* `SPACE` - toggle auto/pause
* `UP/DOWN` - change delay by a factor of 10 (starts at 10ms, down to 1, up to 1000)
* `ESC/CTRL+C` - quit


