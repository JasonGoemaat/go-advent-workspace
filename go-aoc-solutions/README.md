# Advent of Code Solutions (In GO)

This is my repo for solving [Advent of Code](https://adventofcode.com/)
using GO.

Each year has their own directory, with a subdirectory for each day.

Generally I have a single 'main.go' file you can run after downloading
your input to 'input.aoc' in the same directory and creating a file called
'sample.aoc' with the copied sample from the puzzle page.

I have a separate repository with a library I wrote with helper functions
for parsing, and for a few puzzles displaying a terminal UI:
[github.com/JasonGoemaat/go-aoc](https://github.com/JasonGoemaat/go-aoc).

To get started, run `go mod tidy` in the root of this repository with
the `go.mod` file.  Then with vscode (opened from PowerShell or Command
Prompt, not git bash) and the GO extension.   You can F5 to debug
one of the `main.go` files, I like this `launch.json` file with
`"console": "externalTerminal"` added to open up a separate terminal
which is handy for the few with TUI (text user interface) options.

    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Launch Package",
                "type": "go",
                "request": "launch",
                "mode": "auto",
                "program": "${fileDirname}",
                "console": "externalTerminal"
            }
        ]
    }

## Special Days

### Day 14

Here for part 2 you need to find the Christmas Tree image.  There's
several different code files for outputting images.   The key is to find
horizontal and vertical patterns in the images.   Since everything is
mod 101 and 103, there are horizontal and vertical clumping at those
intervals.  I also have a lot of code looking for an actual easter egg
in the puzzle trying different values from 101 and 103 and checking
standard deviation.   The 101 and 103 images pop out immediately when
looking at the low standard deviation values compared to the second
lowest.  I did find some other things that looked like clumping, but
never found values for width and height that had a stark contrast.

I *THINK* it could be done programatically looking for the 101 and 103
patterns for horizontal/vertical standard deviations and calculating,
but I saw the image and didn't bother automating it.

## Differences

Sometimes when I have some trouble I'll create a subdirectory
for making some changes to work on the problem.   For example
on day 24 I created code in the `2024/24/naming2` directory
that outputs files for the various outputs and as I worked on
it I manually found the right ones to swap.

2024 Day 15 I started messing around with 
[BubbleTea](https://github.com/charmbracelet/bubbletea)
which is a library for writing TUIs (Terminal User Interfaces).
There's a 'bubbletea' subdirectory with some things I was trying
out.  The 'main.go' in that directory shows the boxes moving
around, kinda cool.

2024 Day 21 I never finished.   There's a 'cheat' subdirectory
with someone else's solution I used.   I don't see a way to 
take my stars back so I will have to finish it some day.   The
problem was I spent way too much time writing the code for the
visualization, which you can see by running it, though it
gives the wrong answer :).   I think I know the track to take, 
but some of these puzzles have wracked my brain and I just
don't have the energy to go back and fix that, especially after
writing the visualization since it's not really useful
in the actual solution.  It's pretty cool for TUI though, showing
how to have columns and aligning things with styles based on
a model.  

I think I know the track to take documented somewhere for day 21
in the several READMEs.   You actually want to send '<'s before 'v's,
but only if it won't go over a space.   Should be no zig-zag in
the controller inputs (except the last one which doesn't count
anyway).  That will be good for Part 1 I believe.
