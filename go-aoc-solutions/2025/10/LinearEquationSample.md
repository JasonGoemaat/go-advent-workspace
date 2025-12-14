# Linear Equation Sample

    [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}

So there are 6 buttons and 4 joltages.  Here's the matrix,
with buttons labled B0-B5 and Joltages J0-J3, TJ for total joltage.

         B0  B1  B2  B3  B4  B5  TJ
    J0 |  0   0   0   0   1   1   3|
    J1 |  0   1   0   0   0   1   5|
    J2 |  0   0   1   1   1   0   4|
    J3 |  1   1   0   1   0   0   7|

So maybe to start we sort so that 1s in each column are on the top:


          A   B   C   D   E   F
         B0  B1  B2  B3  B4  B5  TJ
    J3 |  1   1   0   1   0   0   7|
    J1 |  0   1   0   0   0   1   5|
    J2 |  0   0   1   1   1   0   4|
    J0 |  0   0   0   0   1   1   3|

So it doesn't give us a single value...   When I see for solving
is people looking at the bottom row and there being a single
value.   For example if B4/J0 were a 0, then we could use only 
B5 to get the total joltage of 3, so B5 would be pressed 3 times.

How it is now, We just know that B4 and B5 must be pressed 
a total of 3 times, or B4 + B5 = 3.   That diesn't help so much.
The row J2 has 3 values still.   Maybe we can make it somewhat better by trying REDUCED echelon form.   That is making sure the pivot is the ONLY value in that column.

Here the Pivot in B0J3 is the only 1 in that column.   But
the pivot in column B in B1J1 has a 1 above it.   So I'll
subtract J1 from J3:

          A   B   C   D   E   F
         B0  B1  B2  B3  B4  B5  TJ
    J3 |  1   0   0   1   0  -1   2|
    J1 |  0   1   0   0   0   1   5|
    J2 |  0   0   1   1   1   0   4|
    J0 |  0   0   0   0   1   1   3|

Column C has only the pivot

Column D does NOT contain a pivot

Column E has a pivot in B4J0, but B4J2 has a non-zero, so we can subtract J2 = J2-J0


          A   B   C   D   E   F
         B0  B1  B2  B3  B4  B5  TJ
    J3 |  1   0   0   1   0  -1   2|
    J1 |  0   1   0   0   0   1   5|
    J2 |  0   0   1   1   0  -1   1|
    J0 |  0   0   0   0   1   1   3|


What if we re-order the columns so that the pivots form
a diagonal line:

          A   B   C   E   D   F
         B0  B1  B2  B4  B3  B5  TJ
    J3 |  1   0   0   0   1  -1   2|
    J1 |  0   1   0   0   0   1   5|
    J2 |  0   0   1   0   1  -1   1|
    J0 |  0   0   0   1   0   1   3|

Thinking...   So D and F have no pivots.   For every time I press F,
I have to press either D or A+C to compensate for the -1.  Hmm...
Actually D perfectly cancels out F, but F adds to J0 and J1 also.  But
I can add to J0 and J1 using E and B without having to cancel.
Maybe this situation is just a happy coincidence.  