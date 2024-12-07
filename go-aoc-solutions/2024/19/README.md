# Day 19

Largest input line is 60, 400 lines.   447 towels, largest towel is 8

5 different colors, so would need 3 bits, 21 per int64, probably just keep string (or bytes)

So what's an efficient way to keep track of what towel's we've used?  As we're
building the collection, we can't use the same towel twice.  Thoughts:

1. Array of bools, true if one's been used
2. Array of ints, fill with ones from start to finish we use
3. Array of indexes into main
4. Map of towel to position
5. Bitmask?   Too many towels for an int64

Another idea I had was to...

## Checking data

Ok, so I can create regular expressions for all 447 towels, run them all for
all 400 patterns, and get about 60-80 towels and 100-180 total matches
as slice indexes into the string all in 125ms.   I think I can use this then
to create a recursive call and memoize in decent time, even with string
manipulation being slow.   Let's try it...

## Ok, my string method in TOOSLOW works, but guess what?

It's too slow!   I think it's all the string manipulation since they are
immutable.   I thought about converting to byte arrays and writing my own
'find' functions and such, but I think I can just start by finding matches
at the start of the string first, and then search from past the last match
until the string is done.

## I had to use []byte

It was just so much faster...

## Part 2

This wouldn't work with my original method.   It took me a bit to figure out
the trick.   I kept trying to think of a way to memoize starting at the front.
I also thought of just combining the towels.   For example if I found a towel
of `buwrrw`, I could continue on and multiply by all the combinations that
could make up that from other towels.   For instance I might have `bu`, `b`, `u` `r`,
and `w` towels and could go bu-w-r-r-w or b-u-w-r-r-w.   That wouldn't save
much though because I would still be checking most of them individually, right?

So I start from the end where I only match one character and store the count
for how many ways I can make matches (just 1).   Then on the second from the end
I see how many ways I can match the rest of the string by testing each towel
and adding all combinations along with the stored values at the position of the
end of the string.   For example say I have `w`, `r`, and `rw`.   With `rwrwrw`
above I start with ints `[0 0 0 0 0 0]`.   Then I look at the last position
with a `w` and find one way for that to terminate the string, so I have
`[0 0 0 0 0 1]`.   Then I look at the second to last where there's the second
`r`.   I can use `rw` to terminate the string giving me `1` and I can use
`r` to get to the last position which has a `1` so that gives me two ways
to get to the end from there: `[0 0 0 0 2 1]`.  The next spot I have a `w`
and only the single `w` fits there and leaves me at the second to last position
which has a '2', so I put 2 there also: `[0 0 0 2 2 1]`.   How it gets a bit
interesting because I have `rwrw` I'm looking at.   I have `r` which points to
the next index which has a '2' and also I have `rw` which points to the second
to last index which also has a '2'.   Add those together and there are 4
ways to end the string from that position: `[0 0 4 2 2 1]`.  Next spot
I'm left with `wrwrw` and the only matching towel is `w` which points to the
next index with a '4', so I have `[0 4 4 2 2 1]`.   Then at the start there
is `rwrwrw` which matches `r` as well as `rw`, so I add the next index value
which is '4' and the one after that which is also '4' to get '8': `[8 4 4 2 2 1]`
Since the 8 is what I get from the start of the string, that is the answer.
