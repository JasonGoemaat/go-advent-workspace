# Day 13

Pretty straight-forward to brute-force, though I had a mistake in my 
calculating 'if' statement that wasn't apparent and had to do some debugging
and debug printing.

Going to try a better solution.   I could represent each person as a vertex in
a graph and each link as an edge.  I want to visit EVERY vertex once and end
going back to the starting vertex.   The starting vertex shouldn't matter.

The problem is I don't know how to traverse it any more quickly than doing a
brute force.  Each time we go down a node we also have to pass the whole list
of where we've been and avoid going there again.  We end when there are no
more paths to take and calculate the total cost.

Well, reading about graphs it is NP-hard.  I can decrease the cost a bit I think
because I can pick two people for left and right?   No, that's no faster because
I still have to swap one, or if I keep that one, swap the other.   Then try any
other two people also.
