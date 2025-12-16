[LANGUAGE: GO]

I hadn't used matrices with linear equations really before, so I decided
to use this puzzle to understand it.   It was kind of hard debugging since
I would have issues with some puzzles due to a small error that was
hard to find and didn't affect the majority of puzzles.

It was hard finding solutions here I could use to verify my number to even
find what lines of my input were the problem.   There are a lot of
solutions here, but just re-arranging them to run an individual puzzle
isn't easy as a lot of them have custom libraries for handling AOC.
Also I haven't seen one that can be made to easily give you the details
like how many times you have to press each button.

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

Here is the running code:

```go
var explain = false
var outputEach = true
var input = ""
// var input = "[..#.##.##] (0,1,2,3,4,5,8) (2,3,5,7) (2,5,6,7) (0,3,4,5,6,7,8) (0,2,4,6,7,8) (0,1,2,4,5,8) (0,2,4,5,6) (0,1,3,4,7,8) (0,1) (2,3,6) {55,25,58,54,55,53,44,43,42}"
// var input = "[###..#...#] (3,6,7) (1,2,6) (0,2,3,4,5,6,9) (1) (0,1,2,5,6,7) (0,1,2,3,6,7,8,9) (0,1,2,3,5,6,8,9) (1,2,3,4,5,6,8,9) (0,5,6,8,9) (1,2,4,7,9) (0,3,8,9) (0,2,4,5,6,7,8) (2,3,5,6,8,9) {56,74,68,51,33,39,58,48,52,69}"

func main() {
	if len(input) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input += scanner.Text() + "\r\n"
		}
	}
	re := regexp.MustCompile("[^\r\n]+")
	lines := re.FindAllString(input, -1)
	sum := 0
	for i, line := range lines {
		if outputEach {
			fmt.Printf("Line %d: %s\n", i, line)
		}
		minPresses, solutions := solveLine(line)
		sum += minPresses
		if outputEach {
			for _, solution := range solutions {
				if solution.TotalPresses != minPresses {
					break
				}
				fmt.Printf("  %d presses: %v\n", solution.TotalPresses, solution.Presses)
			}
		}
	}
	fmt.Println()
	fmt.Printf("Result: %d\n", sum)
}

func solveLine(line string) (int, []MatrixSolution) {
	m := ParsePuzzle(line)
	m.RREFRecurse(0, 0, explain)
	solutions := m.Solve()
	return solutions[0].TotalPresses, solutions
}
```

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

```
      ┌ 0  1  2  3  4  5  6┐
 0( 0)│ 0  0  0  0  1  1  3│
 1( 1)│ 0  1  0  0  0  1  5│
 2( 2)│ 0  0  1  1  1  0  4│
 3( 3)│ 1  1  0  1  0  0  7│
 MaxP └ 7  5  4  4  3  3  0┘
 ```

So you can see button 4 (column 4) has '1's in rows 0 (joltage value 3) and 2 (joltage value 4).
That means pressing button 4 will increase each of those joltages by 1.   That also lets me
constrain the maximum presses (MaxP, which is a separate array, not part of the matrix)
of that button, it can't be pressed more than 3 times or it would overflow the joltage in row 2.

`RREFRecurse` attempts to reduce the matrix to try and leave most columns with a single
value.   The buttons for those columns are then dependent, if only those columns remain then
each button must be pressed an exact number of times to achieve the required remaining joltage
after other buttons are pressed.   Here is what the matrix looks like after I reduce:

```
      ┌ 0  1  2  3  4  5  6┐
 0( 3)│ 1  0  0  0  1 -1  2│
 1( 1)│ 0  1  0  0  0  1  5│
 2( 2)│ 0  0  1  0  1 -1  1│
 3( 0)│ 0  0  0  1  0  1  3│
 MaxP └ 7  5  4  3  4  3  0┘
```

So there are two button columns (4 and 5, column 6 is the joltage remember) that have 
multiple numbers in them.  I permute the possible presses for those buttons, button
4 can be pressed 0-4 times and button 5 can be pressed 0-3 times.   I first copy the
jojltages to a new array and for each of those combinations I adjust the remaining
joltage.   Then I go through each of the remaining columns which will have their presses
defined by the remaining joltage.   No presses should be negative and all joltages
should be 0 for a valid solution.  Theses are the valid possible presses to get the
minimum number of presses (10), the first one is the result shown on the puzzle page:

    10 presses: [1 3 0 3 1 2]
    10 presses: [1 5 0 1 3 0]
    10 presses: [1 2 0 4 0 3]
    10 presses: [1 4 0 2 2 1]

I was happy with my code for permuting combinations:

```go
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
    // adjust joltages[] for variable button values set by permuting
    // assign presses[] for dependant buttons based on joltages and subtract from joltages
    // verify no presses are negative and all joltages are 0 - add solution to list
}
```

Since I wanted to track the presses, I create a list of variable and dependent columns.
The permuting happens over the indexes representing variable buttons in the presses[]
array.

The result of the solve function is an array of `MatrixSolution` sorted by total presses
which I can then examine to check and help find out where my code might have gone wrong.
So if you're having trouble and getting a valid solution but not the minimum solution,
this might help figure out why you don't have the minimum solution, or if your solution
is even a valid one.

```go
type MatrixSolution struct {
	Presses      []int // presses for each button, in original button order
	TotalPresses int   // total presses - lower is better
}
```

Links:

* https://www.math.odu.edu/~bogacki/cgi-bin/lat.cgi?c=rref - nice page that will take a matrix and reduce it, though it uses floats
* https://www.wolframalpha.com/input?i=Reduced+row+echelon+form - WolframAlpha page - doesn't seem to work for large matrices
* https://www.youtube.com/watch?v=2GKESu5atVQ - Good video on Gaussian Elimination in series on Algebra, next in series on Gauss-Jordan, next visualizes as planes in 3d space, etc.
