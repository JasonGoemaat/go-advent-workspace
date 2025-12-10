I'm having trouble figuring out an algorithm that will work in
all cases.   I see from posts that there seems to be a special
design to the puzzle that should make it more ameniable to an
visual/brute force approach.   I'd rather write something
generic that will work for all cases, but I'm having trouble.

What I tried to do for speed is to take any pair of points
and calculate the area, then sort the list of what I call 'Rectangle'
by this area descending so I start with the largest areas first.

Then I take the first 'valid' rectangle from the list and return that,
so my 'valid' function has to determine if the rectangle is valid.  To
do that I check if there are any points inside that rectangle
(`MinX < x < MaxX` and `MinY < y < MaxY`), or if any horizontal
segments intersect that rectangle.

The problem I have is that the rectangle formed by G,E below is
valid, while the rectangle formed by F,D is not.  I know I can
see if an individual point is within the area by looking at it
and counting how many line segments it intersects in one direction,
if that number is odd then it is inside a shape.  But how do I do
that for an entire area?

    ..............
    .......AXXXB..
    .......XXXXX..
    ..GXXXXHXXXX..
    ..XXXXXXXXXX..
    ..FXXXXXXEXX..
    .........XXX..
    .........DXC..
    ..............

Writing this I just had an idea.   Maybe I can just take the middle
of the area and check that, along with the intersecting tests of course.