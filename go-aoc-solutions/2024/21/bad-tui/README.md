# Day 21 part 1

Created some structs to help me define the keypads:

```go
type (
	Keypad struct {
		Layout     []string
		CurrentPos KeypadPos
		InitialPos KeypadPos
		Preference []string
	}
	KeypadPos struct {
		X, Y int
	}
)
```

And functions to initialize the two types:

```go
func NewNumericKeypad() *Keypad {
	keypad := Keypad{
		Layout:     []string{"789", "456", "123", " 0A"},
		CurrentPos: KeypadPos{X: 2, Y: 3},
		InitialPos: KeypadPos{X: 2, Y: 3},
		Preference: []string{"^", "<", ">", "v"},
	}
	return &keypad
}

func NewControlKeypad() *Keypad {
	keypad := Keypad{
		Layout:     []string{" ^A", "<v>"},
		CurrentPos: KeypadPos{X: 2, Y: 0},
		InitialPos: KeypadPos{X: 2, Y: 0},
		Preference: []string{"v", "<", ">", "^"},
	}
	return &keypad
}
```

## Logic 

I think for the moves we can follow the preferences and always find the
shortest path.  This is an assumption, but I think it's valid when I think
through how it works.

The most efficient path to get to a number means I want to press the same key
as many times as I can because the controlling robot can keep pressing 'A'
to press that key and doesn't have to move to another key.

For example if I want to press `7` starting at `A`, that requires 3 up and 2
left.   The controlling controller needs to have these pressed: `^^^<<A`
The controller controlling that needs to then only move to the `^` and press
it three times, move to the `<` and press it twice, then move to the `A`
and press it once: `<AAA`, `v<AA`, `>>^A`.

     ^A
    <v>

If I were to instead use `^<^<^A` that would require a lot more movement on
the controller.   `<A v<A >^A v<A >^A >A` with spaces between the moves or
`<Av<A>^Av<A>^A>A` vs. `<AAAv<AA>>^A`

By using the preverence list I can keep from moving over the empty space in
a keypad.  For instance if I'm going down and left on the controller, going
down first then left will never hit the space.  Likewise if I'm going
up and/or right, going right first will prevent hitting the space.

Could there be a situation when it would make sense to move say left first
before moving down?  hmm...   Maybe if the parent controller were already
on the left?   But that shouldn't happen.   Each step before moving should
require tha the parent is on the `A` because it would have hit that to
actually press the key on the child.   I don't think it can be a problem.

## Visualizing

Each keypad could be visualized separately.  I'm thinking in columns.
Maybe with the list of moves below.  It would be nice to be able to group
them somehow I think, but how?   Is there a point in playing it out?
I guess I could start at the last level with the numbers.   Then specify
the series of moves needed to get to that number.   Then pass that down
to the next which would generate a series of moves for each.   Each move
is then passed down the line to the next.

So let's look at the keypads, going from top (Parent) down to children:

1. Numeric Keypad "depressurized" - Queue up '029A'
2. Controller 1 "high levels of radiation" - empty queue
3. Controller 2 "-40 degrees" - empty queue
4. Controller 3 "full of Historians" - empty queue - this is the answer

So to process, I start with the keypad being #1 and queue up '029A'.
I can submit the moves required to press '0' to the child, then move
to the child #2 which will have `^^^<<A` in it's queue.  It will submit
the moves required to press `^` to it's child.   It's currently on `A`
so it submits `<A` to the child #3 and move the current pointer to #3.
#3 then calculates the moves required for each.   These don't get
submitted anywhere else, but instead form the answer.

I think it would be cool visually to show which one was active somehow
(color? border?), and show the active elements and history for each.  Yeah!

Ok, so if I store the history of things queued for visualization, I just
have to look at that history for #3 and I have my answer.

## Speed

This isn't necessary now, but I imagine it will be in part 2.  How would
memoizing work?   Well I know each of the children start at 'A' all the
way down when submitting new presses, so each keypad only has to
remember it's own.   I'm thinking the real way to do this might be
similar to Day 19 Part 2.

## States:

Actions to process in Step(), first we always clear keypresses so that
the only ones actually PRESSED will be set by this action

* `StartLine()`
    * Reset keypads
    * Submit `Enqueue(0, "029A")` action
* `Enqueue(pad, str)`
    * set `Queue[pad] = str`, `QueueIndex[pad] = 0`, `CurrentPad = pad`
    * Submit `ProcessQueue(pad)` action
* `FinishLine()`
    * Look at final keypad and calculate score based on line and total length of history
    * Submit StartLine() action
* `ProcessQueue(pad)`
    * If `QueueIndex[pad] > 0`, `set KeyPressed[pad] = Queue[Pad][QueueIndex[pad] - 1]` to mark the previous entry as pressed
    * If queue empty:
        * add queue to history and clear queue
        * Submit `FinishLine()` if pad 0
        * Submit `ProcessQueue(pad-1)` if not pad 0
    * ElseIf `CurrentPad == 3` (last pad):
        * Increment `QueueIndex[pad]`
        * Submit `ProcessQueue[pad]`
    * Else
        * Calculate moves to get from current position to next char in queue
        * Increment `QueueIndex[pad]`
        * Submit `Enqueue(pad+1, str)` to let next handle


## BUG!

Damnit!  Everything looks good except for the last one.   The sample says 64,
but I get 68.  Let's inspect that one and see where the problem is.
I'll move it first and break on finish line to show the history.
The order doesn't affect the solution, so I'll change my input.

    029A: 68 * 29 = 1972
    980A: 60 * 980 = 58800
    179A: 68 * 179 = 12172
    456A: 64 * 456 = 29184
    379A: 68 * 379 = 25772

Me/them:

```
v<<A>>^AvA^Av<<A>>^AAv< A<A>>^A  AvAA<^A>Av<A>^AA<A>Av< A<A>>^AAAvA<^A>A // 3
   <   A > A   <   AA   v <   A  A >>  ^ A  v  AA ^ A   v <   AAA >  ^ A // 2
       ^   A       ^^         <  <       A     >>   A         vvv      A // 1
           3                             7          9                  A
<v<A>>^AvA^A<vA <  AA>>^A Av  A<^A>AAv A^A<vA>^AA<A>A<v<A>A>^ AAAvA<^A>A // 3
   <   A > A  v    <<   A A   >  ^ AA  > A  v  AA ^ A   < v   AAA >  ^ A // 2
       ^   A            < <        ^^    A     >>   A         vvv      A // 1
           3                             7          9                  A // 0

v<<A>>^AvA^Av<<A>>^AAv<A<A>>^AAvAA<^A>Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A // 3   68
   <   A > A   <   AA  v <   AA >>  ^ A  v  AA ^ A  v <   AAA >  ^ A // 2   28
       ^   A       ^^        <<       A     >>   A        vvv      A // 1   14
           3                          7          9                 A // 0   4
<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A     // 3   64
   <   A > A  v <<   AA >  ^ AA > A  v  AA ^ A   < v  AAA >  ^ A     // 2   28
       ^   A         <<      ^^   A     >>   A        vvv      A     // 1   14
           3                      7          9                 A     // 0   4

```

Ok why is it faster to hit the two lefts before the two tops?   Do I want
to go as far as possible?  It's the same keys to press both, but since the ^
is closer to the 'A' button, there is less to do after hitting the second?

But if you have to move multiple times, maybe it's still faster to do them
all at once, i.e. `v<<` is faster than `<v<`

## No way around it

see [Recursion.md](Recursion.md)

