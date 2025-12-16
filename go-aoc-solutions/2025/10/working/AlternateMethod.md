# Alternate Method

Reddit - [Bifurcate your way to victory](https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/)

A      B   C     D
(0,1) (1) (0,2) (1,3)  {30,15,27,14}

 0 = No presses         ....
 1 = Press A            ##..
 2 = Press B            .#..
 3 = Press A+B          #...
 4 = Press C            #.#.
 5 = Press A+C          .##.
 6 = Press B+C          ###.
 7 = Press A+B+C        ..#.
 8 = Press D            .#.#
 9 = Press A     +D     #..#
10 = Press B     +D     ...#
11 = Press A+B   +D     ##.#
12 = Press C     +D     ####
13 = Press A+C   +D     ..##
14 = Press B+C   +D     #.##
15 = Press A+B+C +D     .###

- presses=2 and multiplier=2, divide joltages by 2
    {14,7,13,7} - D - 1+3, A+C

Need to flip last 3 - F is the only one
so 4 presses (A+B+C+D), giving {12,4,12,6}

Multiplier is now 4x as we divide to {6,2,6,3}
We need 2 presses(B+D), giving {6,0,6,2}

Multiplier is now 8x as we divide to {3,0,3,1}
Pattern is #.##.   But that is not solvable!
The only pattern to match that is 14 which requires pressing B+C+D, but that
would flip joltage 1 twice an it is 0.

So I have two ideas for that.  One is that we just return a high number of
presses to indicate it is unsolvable.   If we are trying every combination
of presses each iteration, we could just not try ones that go negative.

Another option would be to switch to brute forcing when any number gets down
below a threshold.   We could count for each combination not just the pattern,
but also the number of decreases.   If the remaining joltages are not at least
that number, switch to brute force.  In this example, A+B+D all trigger
the joltage 1, so A+B+D and A+B+C+D would trigger 3 joltage changes.  So we would
not even attempt 6,2,6,3 because joltage 1 is lower than that.


## Pseudo-code

```go
// is 2**buttonCount length, can be represented by bitmask
// we can create bitmask of odd values for comparison
type ButtonCombination struct {
    JoltageFlipMaskMask     int // mask of bits for flipped joltages
    Joltages                []int // count to subtract joltages by
    Presses                 int // how many buttons are being pressed
}

