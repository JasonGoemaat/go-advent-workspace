# Say What??!?!?!??

Part 2 is just about the most complicated to understand puzzle I've ever seen!
It doesn't make any logical sense.   So you are trying to figure out the
best price to buy from each monkey.   But because of some language barrier
you have to have the Historian translate something into monkey for your
monkey agent to know when to sell.

For some unfathomable reason you need to sell to each of the monkeys after
they have had the same 4 consecutive price changes.   That's the only way
the Historian can describe to the monkey when to buy, and the monkey can
only remember 1 set of changes.   BUT the monkey agen then has to keep track
themselves of the last 4 prices of each other monkey agent and realize when
the right time to sell is.   Seems a lot easier to just give the monkey
one number for each buyer, then you would get the best price as well.

Oh, well.

## Failed Assumptions

Part 1 was just so easy, and I was trying to think what part 2 would be.

I thought "I bet they are going to ask for like 10 billion secret numbers"
and the trick would be to find a cycle in the secret number calculation
and then mod that.   That would have been too easy for day 22.  I'm glad
I didn't waste the time.

## Solving Part 2

This shouldn't take too long.   I was thinking of an efficient way to
store the 4 changes.   I can store any change in 5 bits by adding 9.   The
largest price change is +/- 9 so there are 19 possibilities, adding 9
will change -9 to +9 into 0 to 18 which can be stored in 5 bits.  Then
each new price will do

`lastPrices = ((lastPrices << 5) | (newPrice+9)) & 0xfffff`

And use that as a key.   If I use multiply, add, and modulo though
I could store it in even fewer possibilities, 130321 or 0x1fd11.
This would be easy array territory.  I only have to calculate 4*1997
values also, so modulo vs. bit arithmetic ain't gonna matter enough.

Or I could just use strings, would that be easier?

Screw it, each monkey with bit masking array of 1 million ints is just
8mb, so 32mb overhead for having arrays for 4 monkeys.   Then one final
loop to process easily and find the most bananas I can get for my hiding
spots.  

Oh, wait, they sell the first time, so I need a number to represent no sale.

## Maps it is

Oh, dummy...  There's actually 1629 monkeys.   Time to use maps because I
don't really want to deal with it.   Actually, using modulo it would be 1.6gb
and just 12gb with not.   Maybe that's feasible?