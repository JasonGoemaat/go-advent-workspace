Matrix:

    1a +    1b +    1c +    0d  = 10    (joltages 0,4)
    1a +    1b +    0c +    0d  = 5     (joltage 3)
    1a +    0b +    1c +    1d  = 11    (joltages 1,2)
    0a +    0b +    1c +    0d  = 5     (joltage 5)

0 - |1 1 1 0 10|
1 - |1 1 0 0  5|
2 - |1 0 1 1 11|
3 - |0 0 1 0  5|

I'm using 0 indexes for computer's sake, math usually starts
with index (row or column) of 0.

What if I make it so that rows 1 and 2 have 0 for the first column?   I can do that by subtracting row 0 from them since
row 0 has a 1:

0 - | 1  1  1  0  10|
1 - | 0  0 -1  0  -5|
2 - | 0 -1  0  1   1|
3 - | 0  0  1  0   5|

Then I swap rows 2 and 1 to get back to row echelon form, i.e.
row 1 column 1 is non-zero and everything below is 0:

0 - | 1  1  1  0  10|
1 - | 0 -1  0  1   1|
2 - | 0  0 -1  0  -5|
3 - | 0  0  1  0   5|

Rows 2 and 3 both have non-zero values now in column 2.  But
I can add row 3 to row 2 to make it zero:

0 - | 1  1  1  0  10|
1 - | 0 -1  0  1   1|
2 - | 0  0  0  0   0|
3 - | 0  0  1  0   5|

And re-order again, swapping rows 2 and 3:

0 - | 1  1  1  0  10|
1 - | 0 -1  0  1   1|
2 - | 0  0  1  0   5|
3 - | 0  0  0  0   0|

Now the -1 in row 1 has non-zero above it, so I add row 1 to row 0 to make that a zero:

0 - | 1  0  1  1  11|
1 - | 0 -1  0  1   1|
2 - | 0  0  1  0   5|
3 - | 0  0  0  0   0|

Now the 1 in row 2 has non-zero in row 0 above it, so subtract that from row 0:

0 - | 1  0  0  1   6|
1 - | 0 -1  0  1   1|
2 - | 0  0  1  0   5| = 5
3 - | 0  0  0  0   0|

# NOT Reduced

From row 40 above:

And re-order again, swapping rows 2 and 3:

0 - | 1  1  1  0  10|
1 - | 0 -1  0  1   1| 
2 - | 0  0  1  0   5|
3 - | 0  0  0  0   0|

So I should be able to figure out the solution now?
This short make it look easy: 
https://www.youtube.com/shorts/frqILDunV50

Hmm...  Can I multiply a row?  negate row 1 for instance?

0 - | 1  1  1  0  10|
1 - | 0  1  0 -1  -1| 
2 - | 0  0  1  0   5|
3 - | 0  0  0  0   0|

Now subtract 1 from 0 so column 1 only has a single 1

0 - | 1  0  1  1  11|
1 - | 0  1  0 -1  -1| 
2 - | 0  0  1  0   5|
3 - | 0  0  0  0   0|

Now subtract row 2 from row 0 so column 2 only has a single 1

0 - | 1  0  0  1   6|
1 - | 0  1  0 -1  -1| 
2 - | 0  0  1  0   5|
3 - | 0  0  0  0   0|

so row 2 has to be 