# Recursion

I think there's no way around it, going to have to do some recursion and
memoization.  Each layer can keep track of the shortest way to get from
any starting position to any other and press it maybe?

Final robot layer defines moves to  send to the player layer and that
represets the final length.  Let's name player P0 and the first 
controller robot C0, the second robot controller C1, and the keypad
robot K0.

So C0 needs to calculate all of these.   Since C0's child is P0, we
can just use my algorithm which prefers v<>^ to avoid the space.
Here there are 20 options of moving the child from 1 key to another
and pressing the 'A' key on the child to press the key on the
parent:

    A< - 4 - v<<A will be one of the fastest
    A> - 2 - vA will be the fastest
    A^ - 2 - <A
    Av - 3 - v<A
    <A - 4 - v<<A
    <> - 3 - >>A
    <^ - 3 - >^A
    <v - 2 - >A
    >A - 2 - ^A
    >^ - 3 - ^<A
    >< - 3 - <<A
    >v - 2 - <A
    ^A - 2 - >A
    ^< - 3 - v<A
    ^> - 3 - v>A
    ^v - 2 - vA
    vA - 3 - ^>A
    v> - 2 - >A
    v< - 2 - <A
    v^ - 2 - ^A


## Problem

The problem with mine is this...   Going from the '3' to the '7' has a shortest path:

    P0  A<vA<AA>>^AAvA<^A>AAvA^A
    C0  A  v <<   AA >  ^ AA > A
    C1  A         <<      ^^   A
    K0  3                      7

But my stupid just wing it solution gives a longer one:

    P0  Av<<A>>^AAv<A<A>>^AAvAA<^A>A
    C0  A   <   AA  v <   AA >>  ^ A
    C1  A       ^^        <<       A
    K0  3                          7

So K0 needs to find the shortest `^^<<A` which might be either `^^<<A` or `<<^^A`.
I try `^^<<A` (because it is safe going up then left) which is longer than `<<^^A`.

THOUGHT: Is it possible that there may be a way that changing directions could be
faster if you have multiple?   I don't see how.  Pressing A twice is a SINGLE
extra keypress through the whole pipeline.  Doing `^^<<A` requires moving the top
level 3 times and pressing A 5 times.   Doing `^<^<A` requiresing moving the top
level 5 times and pressing A 5 times.  CERTAINLY that will take more presses, RIGHT?

## Finding 'best' arrangement

So when I'm looking for the best arrangement


