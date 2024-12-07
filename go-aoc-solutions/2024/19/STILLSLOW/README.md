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

