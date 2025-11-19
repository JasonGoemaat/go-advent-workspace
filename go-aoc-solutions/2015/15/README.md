# 2015 Day 15

https://adventofcode.com/2015/day/15

My idea is to start with 100 teaspoons of EACH incredient.  Then start removing
teaspoons one at a time, removing the one that decreases the total by the
least amount.  Let's see how that works...

Pretty easy, but...

## Part 2

Shoot, need to have EXACTLY 500 calories.  So brute force recursion it is. I
put the code in the 'part2' subdirectory in a separate package to re-use the
same names and promote some local variables to package level to make it easy.
The first call is:

```go
// will have to try 94 million options
bestScore := recurse(0, 100)
```

This finishes in 35ms and returns the correct answer, how is it that fast?
Oh, I miscalculated the total calls as `100*99*98*97`.   In reality I start
at the highest value (greedy) of 100 for the first, so that is just 1 call.
When it gets to 99, there are three options to place that 1 remaining.
When it gets to 98, there are 6 options for placing the remaining.  What
is the formula for this?

I ran across [this question](https://math.stackexchange.com/questions/541790/counting-ways-to-partition-a-set-into-fixed-number-of-subsets)
on math.stackexchange and the wikipedia 
[Stirling_numbers_of_the_second_kind](https://en.wikipedia.org/wiki/Stirling_numbers_of_the_second_kind),
but that's for non-empty subsets.   Maybe I can just add 1 for each subset 
that CAN be empty?

Let's think about fewer ingredients.

* For ONE ingredient, there is ONE way
* For TWO ingredients, the number can be from 0-100, so 101 ways
* For THREE ingredients it gets difficult...

So 101 ways to set the value of A.   And for B,C then
there is 1 way for A=100.  For A=99 there are 2 ways.
for A=98 there are 3 ways (2-0,1-1,0-2).  For A=97 there
are 4 ways (3-0, 2-1, 1-2, 0-3)