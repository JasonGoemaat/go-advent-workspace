#   Trying something

For Part 2 it should require some memoization.   Maybe I should just do that
to start with.  Each depth needs to figure out the fastest way for the given
direction strings to be run on the child.  Just try them all, maybe avoiding
zig-zags, and definitely avoiding hitting the ' '.  I was having a little
trouble thiking about how to precisely count multiple presses at once.
The first would be to force parent to go from A->direction, then to press 'A'
however many times that direction is repeated.   Then force parent to go
back to 'A' and press it.

To do '<<^^' and press it for example, what do I push to the child for counting?
Child has to push one of two options:

    v<<AA>^AA>A  // move parent to '<' and press twice, move to '^' and press twice, move to 'A' and press once.
    <AAv<AA>>^A  // move parent to '^' and press twice, move to '<' and press twice, move to 'A' and press once.

Should the parent make sure those are valid?  Who counts the extras for the
second 'A' press?   Who puts the 'A's into the string?   Should the parent pass
'<<', '>^' alone?   There's really only two options.   The parent is the one
that knows where they currently are, so maybe.  What does that child pass to
it's child?   `v<<`, `>^`, `>` and `<`, `v<`, `>>^`?  The child needs to report
how long it will take in total to process those moves, which includes the `A` to
press the key at the end?   And the parent then adds 1 for duplicate `A`s in a row?
Or should parent pass `v<<AA`, `>^AA`, `>A`?   I think that is the best, it tells
the child what keys need pressed and the child can figure it out from there.  The
only problem then is the keypad where I thought that would be a problem because
it shouldn't move to 'A', but that's not a problem, I'm telling the child to
press 'A'.  Maybe I don't need to worry about dupes that way either.

## Why is this so hard?

    v<<A>>^AvA^Av<<A>>^AAv<A<A>>^AAvAA<^A>Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A // 3   68
    <   A > A   <   AA  v <   AA >>  ^ A  v  AA ^ A  v <   AAA >  ^ A // 2   28
        ^   A       ^^        <<       A     >>   A        vvv      A // 1   14
            3                          7          9                 A // 0   4
    <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A     // 3   64
    <   A > A  v <<   AA >  ^ AA > A  v  AA ^ A   < v  AAA >  ^ A     // 2   28
        ^   A         <<      ^^   A     >>   A        vvv      A     // 1   14
            3                      7          9                 A     // 0   4

So keypad depth 0 gets `379A` as KEYPRESSES, sequence where we start at `A`.
It steps through each key and calculates the moves required to press it and
finds the cheapest way by passing it to the child at depth 1.  These have
to be pressed in order.   It loops through and does this:

* 00 - Calculate `A`->`3` and press, child needs only option `^A`
* 01,02 - Calculate `3`->`7` and press, two options: `^^<<A` and `<<^^A`.  Find cheapest and use that
* 03 - Calculate `7`->`9` and press, one option: `>>A`
* 04 - Calculate `9`->`A` and press, one option: `vvvA`, find cheapest and use that

Depth 1

* Starting at `A`, need to press `^A`, so send down option `<A`

* `A`->`^` = `<A` - child needs to press `<A`
* `^`->`A` = `>A` - child needs to press `>A`
