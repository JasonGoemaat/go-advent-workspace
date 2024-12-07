# 2015 Day 8

This one threw me a bit.  I had to keep reading this instruction over and
over:

> For example, given the four strings above, the total number of characters of
string code (2 + 5 + 10 + 6 = 23) minus the total number of characters in
memory for string values (0 + 3 + 7 + 1 = 11) is 23 - 11 = 12.

This really means "The number of characters in the given encoded string minus
the number of characters if it were decoded into memory."

That's just the line length minus the length of the string after you decode it.
I think I was doubly messed up because of the way I decided to do things.
I had a regexp to capture any matches that represented a character: `\\`, `\"`,
and `\x[0-9a-f][0-9a-f]`