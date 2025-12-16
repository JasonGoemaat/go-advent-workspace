[LANGUAGE: GO]

QuickStart:

* Download https://github.com/JasonGoemaat/go-advent-workspace/blob/main/go-aoc-solutions/2025/10/part2_standalone/day10part2.go
* Edit the variables 'explain', 'output', and 'input' how you wish, it is setup to take puzzles from the standard input and do extensive reporting
* Install go
* From the command line go to the directory with you file and run `go run day10part2.go`
* Paste your puzzle(s) and hit the key combination to terminate input (`CTRL+Z` for windows cmd, `CTRL+D` for bash)

I hadn't used matrices with linear equations really before, so I decided
to use this puzzle to understand it.   It was kind of hard debugging since
I would have issues with some puzzles due to a small error that was
hard to find and didn't affect the majority of puzzles.

It was hard finding solutions here I could use to verify my number to even
find what lines of my input were the problem.   There are a lot of
solutions here, but just re-arranging them to run an individual puzzle
isn't easy as a lot of them have custom libraries for handling AOC.
Also I haven't seen one that can be made to easily give you the details
like how many times you have to press each button.  Thanks to `___ciaran`
for [his post](https://www.reddit.com/r/adventofcode/comments/1pity70/comment/nu6v5fq/)
which I was at least able to easily manipulate to run and give answers for
each puzzle that I could compare to my own.

So what I came up with was a single '.go' file that will solve either
from a string variable 'input' in the file or from the standard
input if the variable is an empty string.   I have two global flags
that can be set:

1. `outputEach` - if set to true, it will output each input line and
the line number it is on as well as the min presses for that puzzle
and an array showing how many times each button was pressed for one
of the possible minimum solutions (some puzzles have several)

2. `explain` - if set to true it will go through each step of reducing
the matrix and print what it is doing along with the resulting matrix
after performing the operation.

So if you are having trouble with one of your inputs you can just set
the `input` variable in the file and set `explain` and `outputEach` to
true.  You'll then get output showing each operation to reduce the matrix
and details results with all button presses for all the ways that result
in the minimum number of presses.   You can change it to output all
solutions if you want to see if a solution that seems valid is actually
valid but does not result in the minimum number of presses.

I use an integer matrix, I didn't want to deal with floats and wanted to
be able to reason through exactly what was happening since this is the
first time I'm doing an operation like this.  This makes it easy to reason
through each step and do it by hand.

Using only integers also made it a little more difficult to program.  I am
new to this and everything I find online deals with floats.  They also avoid
swapping columns because when using floats or fractions there is no need.
From what I read swapping columns is frowned upon and isn't proven to result
in an equivalent matrix.  For this puzzle however I don't see how that is an
issue.  I just need to keep track of the original order so when I report
the number of presses for each button they are in the correct order.  I don't
think I even need to re-arrange, but I do to keep the same stair-step 
pattern for the pivots as I see everywhere that talks about this.

The matrix has one column for each button and one final column for the 
resulting joltage.   Each joltage has it's own row.  When a button 
affects a joltage I put a '1' in that column and row, otherwise it has
a zero.   The last column has the whole joltage numbers.

Here is the first sample after parsing:

        ┌ 0  1  2  3  4  5  6┐
    0( 0)│ 0  0  0  0  1  1  3│
    1( 1)│ 0  1  0  0  0  1  5│
    2( 2)│ 0  0  1  1  1  0  4│
    3( 3)│ 1  1  0  1  0  0  7│
    MaxP └ 7  5  4  4  3  3  0┘

So you can see button 4 (column 4) has '1's in rows 0 (target joltage 3) and 2 (target joltage 4).
That means pressing button 4 will increase each of those joltages by 1.   That also lets me
constrain the maximum presses (MaxP, which is a separate array, not part of the matrix)
of that button, it can't be pressed more than 3 times or it would overflow the joltage in row 2.

`RREFRecurse` attempts to reduce the matrix to try and leave most columns with a single
value.   The buttons for those columns are then dependent, if only those columns remain then
each button must be pressed an exact number of times to achieve the required remaining joltage
after other buttons are pressed.   Here is what the matrix looks like after I reduce:

        ┌ 0  1  2  3  4  5  6┐
    0( 3)│ 1  0  0  0  1 -1  2│
    1( 1)│ 0  1  0  0  0  1  5│
    2( 2)│ 0  0  1  0  1 -1  1│
    3( 0)│ 0  0  0  1  0  1  3│
    MaxP └ 7  5  4  3  4  3  0┘

So there are two button columns (4 and 5, column 6 is the target joltage remember) that have 
multiple numbers in them.  I permute the possible presses for those buttons, button
4 can be pressed 0-4 times and button 5 can be pressed 0-3 times.   I first create an empty
joltages[] array and for each of those combinations I multiply the number of presses by the
value in that button's column for each joltage and add it to the joltages[] array value.

For example I start with joltages `[0 0 0 0]`.   When I press button '4' 3 times it changes
to `[3 0 3 0]` (because the values in column 4 are 1,0,1,0 times 3 presses).   If I then
press button '5' 2 times it changes to `[1 2 1 2]` (because the values in column 5 are
-1,1,-1,1 times 2 presses).  `presses[]` is then `[0 0 0 0 3 2]`

Then I go through each of the remaining dependent columns (buttons) which will have their presses
defined by the remaining joltage: `[targetJoltages[i] - joltages[i]) / matrix[i, button]`


* Column 0 (button 0) is `[1 0 0 0]` so it has to be pressed 1 time to change joltages to `[2 2 1 2]` and presses[] is `[1 0 0 0 3 2]`
* Column 1 (button 1) is `[0 1 0 0]` so it has to be pressed 3 times to change joltages to `[2 5 1 2]` and presses[] is `[1 3 0 0 3 2]`
* Column 2 (button 2) is `[0 0 1 0]`, but `joltage[2]` is 1 and that's the target os it's pressed 0 times, not changing the arrays
* Column 3 (button 3) is `[0 0 0 1]`, so it has to be pressed 1 time to change joltages to `[2 5 1 3]` and presses[] is `[1 3 0 1 3 2]`
* No buttons had to be pressed a negative number of times and the joltages array is the same as the initial joltages (column 6) so this is valid solution

That series represents one of the minimum combinations and the first one in the list
which is the one specified on the advent of code puzzle sample.  Here are all the possible
ways to press the buttons 10 times to achieve the required joltages:

    10 presses: [1 3 0 3 1 2]
    10 presses: [1 5 0 1 3 0]
    10 presses: [1 2 0 4 0 3]
    10 presses: [1 4 0 2 2 1]

I was happy with my code for permuting combinations.   I initialize the `presses[]`
array with zeroes and permute the values for the variable buttons (`variableButtons[]`
is an array of button indexes/column indexes for the variable buttons).   I just
go in order and when the first overflows I reset and add to the next, etc.   When
the last overflows I am done.

    donePermuting := false
    permuteVariableButtons := func() bool {
        // increase from first to last, when last overflows we are done
        for _, c := range variableButtons {
            presses[c]++
            if presses[c] <= m.MaxPresses[c] {
                return true // no overflow, we are fine
            }
            presses[c] = 0
        }
        donePermuting = true
        return false
    }

    for ; !donePermuting; permuteVariableButtons() {
        // ... clear 'presses' for dependent columns, variable columns remain
        // initialize joltages[] array from the last matrix column
        // subtract from joltages[] for based on variable button presses[]
        // assign presses[] for dependant buttons based on remaining joltages and subtract
        // verify no presses are negative and all joltages are 0 - add solution to list
    }

The result of the solve function is an array of `MatrixSolution` sorted by total presses
which I can then examine to check and help find out where my code might have gone wrong.
So if you're having trouble and getting a valid solution but not the minimum solution,
this might help figure out why you don't have the minimum solution, or if your solution
is even a valid one.

    type MatrixSolution struct {
        Presses      []int // presses for each button, in original button order
        TotalPresses int   // total presses - lower is better
    }

Links:

* https://www.math.odu.edu/~bogacki/cgi-bin/lat.cgi?c=rref - nice page that will take a matrix and reduce it, though it uses floats
* https://www.wolframalpha.com/input?i=Reduced+row+echelon+form - WolframAlpha page - doesn't seem to work for large matrices
* https://www.youtube.com/watch?v=2GKESu5atVQ - Good video on Gaussian Elimination in series on Algebra, next in series on Gauss-Jordan, next visualizes as planes in 3d space, etc.
