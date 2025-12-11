# Part 2 Thinking

Part 2 would take way too long to try every combination...
The maximum combinations I could possibly have to check
with my input is 8113918308230581796, that's including
pressing every button the maximum number of times it can
be pressed without overflowing any joltages.  That's
8 million trillion.

The actual number needed to test would be much smaller.
The maximum number of times needed to press a button
should cause other button presses to overflow.   Maybe it's
worth a shot to see how it performs.

Depth of the recursion would still be for each button.
I would try each button starting with the maximum number of
times and going down.  I feel that would be best because
some of the higher puzzles have certain high values:

Worst sample (parts split to lines by me):

    [###...##..]
    (2,6)
    (0,2,3,4,5,6,7,8,9)
    (0,1,2,5,6)
    (0,1,2,3,4,7,8,9)
    (0,1,4,5,6,9)
    (3,5)
    (0,2,9)
    (0,1,2,4,6,7)
    (6,9)
    (0,1,8,9)
    (0,1,2,3,4,5,8)
    {287,74,298,75,76,77,83,57,55,243}

Buttons 0, 2, and 9 are all over 200, the rest are 55-83.
Starting with button 1 (`(0,2,3,4,5,6,7,8,9)`) which affects
the most joltages is probably smart.   Using a greedy
algorithm

## Linear Equation

Sample:
    
    [..#.] (1,3) (0,3) (0,2,3) {33,7,14,40}

Buttons:

* a - 1,3
* b - 0,3
* c - 0,2,3

Answer: a + b + c (total presses of all three buttons)

* Joltage 0: b + c = 33
* Joltage 1: a = 7
* Joltage 2: c = 14
* Joltage 3: a + b + c = 40 (this has all buttons, so this IS the answer)


Matrix for buttons (rows) and joltages affected (columns)

    [0 1 0 1
     1 0 0 1
     1 0 1 1]

Matrix for joltages required (exactly):

    [33
     7
     14
     40]


# Second example

    [..#.##.##]
    a (0,1,2,3,4,5,8)
    b (2,3,5,7)
    c (2,5,6,7)
    d (0,3,4,5,6,7,8)
    e (0,2,4,6,7,8)
    f (0,1,2,4,5,8)
    g (0,2,4,5,6)
    h (0,1,3,4,7,8)
    i (0,1)
    j (2,3,6)
    {55,25,58,54,55,53,44,43,42}

* Joltage 0: a +         d + e + f + g + h + i     = 55
* Joltage 1: a +                 f +     h + i     = 25
* Joltage 2: a + b + c +     e + f + g +         j = 58
* Joltage 3: a + b +     d +             h +     j = 54
* Joltage 4: a +         d + e + f + g + h         = 55
* Joltage 5: a + b + c + d +     f + g             = 53
* Joltage 6:         c + d + e +     g +         j = 44
* Joltage 7:     b + c + d + e +         h         = 43
* Joltage 8: a +         d + e + f +     h         = 42

Not sure if this is helpful, but I can look at the total
voltage and it should equal the number of time each button
is pressed times the number of joltages affected by
that button.   For example, when I press button 'a', it
adds 7 to the total voltage output.

    7a + 4b + 4c + 7d + 6e + 6f + 5g + 6h + 2i + 3j
        = 55 + 25 + 58 + 54 + 55 + 53 + 44 + 43 + 42
        = 429


