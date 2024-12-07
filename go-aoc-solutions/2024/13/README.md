
Sample puzzle:

    Button A: X+94, Y+34
    Button B: X+22, Y+67
    Prize: X=10000000008400, Y=10000000005400

For X:

    X/94 = 106382978813 remainder 72
    X/22 = 454545454927 remainder 6 // more than 3x the number of A presses

For Y:

    Y/34 = 294117647218 remainder 22
    Y/67 = 149253731424 remainder 59

Button A:

    X/94 = 106382978813 remainder 72
    Y/34 = 294117647218 remainder 22

Button B:

    X/22 = 454545454927 remainder 6 // more than 3x the number of A presses
    Y/67 = 149253731424 remainder 59


So the largest number of presses for A is 106382978813 or we shoot past the
X limit, and we have a lot of Y remaining.

So the largest number of presses for B is 149253731424 or we shoot past the
Y limit, and have a lot of X remaining.

## Eureka! (maybe)

What about a binary search through that space?  Yeah, that's the ticket!
Hmmm...   We can do two searches, once maximizing for A presses and once
maximizing for B presses

Start with `ALOW=0, AHIGH=MIN(GOALX/AX, GOALY/AY)`.  We search the space
using a binary search to try and find a value where we can add more B to
get to the end.  How do we determine up and down?

If it is equally divisible, we found a solution

## Combo of LCM?

LCM of 22,94 is `2*11*47` or 1034.  So we need 11 A to reach that or 47 B
to reach the same value.

LCM of 67,34 is `67*34` or 2368.   So we need 67 A to reach that or 34 B
to reach the same value.

Those two numbers have factors in common:

    1034 = 2 * 11 * 47
    2368 = 2 * 2 * 2 * 2 * 2 * 2 * 37
    Removing common factor of '2':
    LCM = 11 * 47 * 2 * 2 * 2 * 2 * 2 * 37
    LCM = 612128

Therefore it is possible to get to 612128, does that help or am I going
completely in the wrong direction?

## IS IT POSSIBLE?

Ok, I'm an idiot.   I don't think you can ever get to the same resulting
position by using different amounts of A and B.  Trying to minimize the
cost is a red herring.   The puzzle needs some way to combine the
number of each press into a single value for the answer, that's the only
reason it's there I think.  So maybe binary search is the answer, or
quadinary search :

Looking at the problem again:

    Button A:

        X/94 = 106382978813 remainder 72
        Y/34 = 294117647218 remainder 22

    Button B:

        X/22 = 454545454927 remainder 6 // more than 3x the number of A presses
        Y/67 = 149253731424 remainder 59

I need to press A between 0 and 106382978813 times.  Starting in the middle:

Ok, see here: https://www.sveltelab.dev/zyyy4ogk9m2l6r9

```json
{
    "puzzle": {
        "a": {
            "x": 94,
            "y": 34
        },
        "b": {
            "x": 22,
            "y": 67
        },
        "goal": {
            "x": 10000000008400,
            "y": 10000000005400
        }
    },
    "aHigh": 106382978812,
    "bHigh": 149253731423,
    "apresses": 52720768261,
    "ax": 4955752216534,
    "ay": 1792506120874,
    "bpressesx": 229283990539,
    "bpressesy": 122499908724
}
```

So I need two different numbers for b presses.   What happens if I decrease my A presses?
Somewhere around here the gap is the shortest between the 'b' presses needed for x and y
to equal the result:

{
  "apresses": 80964036972,
  "ax": 7610619475368,
  "ay": 2752777257048,
  "bpressesx": 108608206046,
  "bpressesy": 108167503706
}
```

How to I know to move right (ALOW = AMID) instead of left (AHIGH = AMID) when
doing my search?  Because bpressesx is larger than bpressesy!  In other words,
`(bpressesx - bpressesy) > 0`.  If it is equal to 0, I found the solution.

NOW, is there a situation where this would be reversed?  Maybe I just check both
ways.  So I want to move right if:

I want more A presses if:

    bpressesx is higher than bpressesy and 