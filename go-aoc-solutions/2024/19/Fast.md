# How to fast

Start with player pad, determine fastest path and cost for every combination of keys.

Say KEYPAD talks directly to CONTROLLER

And using sample:

    v<<A>>^AvA^Av<<A>>^AAv<A<A>>^AAvAA<^A>Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A // 3   68
    <   A > A   <   AA  v <   AA >>  ^ A  v  AA ^ A  v <   AAA >  ^ A // 2   28
        ^   A       ^^        <<       A     >>   A        vvv      A // 1   14
            3                          7          9                 A // 0   4
    <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A     // 3   64
    <   A > A  v <<   AA >  ^ AA > A  v  AA ^ A   < v  AAA >  ^ A     // 2   28
        ^   A         <<      ^^   A     >>   A        vvv      A     // 1   14
            3                      7          9                 A     // 0   4

KEYPAD has to enter 379A

    789
    456
    123
     0A

KEYPAD finds best way to move from A to 3 and press it.  Possibilities:

The only possibility for A3 is '^'.  This will be the case for any single
direction.  It passes '^' to PLAYER.   PLAYER returns 2 (the length of string
plus the 1 to  hit 'A' and press it) for total cost.   

Keypad needs to move 37.   Possibilities:

    <<^^A
    ^^<<A

If one would hit the ' ', ignore it.   There's no need to try anything but
two directions in a row.   Doing a bend should always take more time!  (I HOPE)
Player returns 5 for both, the moves and hitting 'A'.

## multiple

    v<<A>>^AvA^Av<<A>>^AAv<A<A>>^AAvAA<^A>Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A // 3   68
    <   A > A   <   AA  v <   AA >>  ^ A  v  AA ^ A  v <   AAA >  ^ A // 2   28
        ^   A       ^^        <<       A     >>   A        vvv      A // 1   14
            3                          7          9                 A // 0   4
    <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A     // 3   64
    <   A > A  v <<   AA >  ^ AA > A  v  AA ^ A   < v  AAA >  ^ A     // 2   28
        ^   A         <<      ^^   A     >>   A        vvv      A     // 1   14
            3                      7          9                 A     // 0   4

* KEYPAD passes two options to CONTROLLER0:
    * `^^<<A`
        * CONTROLLER0 calculates it needs these moves:
            * `A^` (or <) +2 for 2 presses
            * `^<` (or v<) + 2 for two presses
            * `<A` (or >>^) + 1 for one press
    * `<<^^A`
        * CONTROLLER0 calculates it needs these moves:
            * `A<` (or v<<) +2 for 2 presses
            * `<^` (or >^) + 2 for two presses
            * `^A` (or >) + 1 for one press
