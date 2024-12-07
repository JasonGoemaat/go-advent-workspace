# Day 20: Race!

I think I figured out a nice way to do this.   I can use my code doing A*
from day 18 and the visualization stuff too.   No falling tiles and there's
only one path to the finish line so that's super easy, just solve it once.

But then I need to 'cheat'.   The explanation I think is over-complicated.
it seems like removing any wall might create a shorter path.   You do need
to count all the ways you can do that though.

Luckily with my A* I have a path back to the start when I solved it.
I think what I can do is just start at the end from each node in my path and
look for a wall with another cheaper path node on the other side.   Then
I can subtract the difference in the cost (G) to get my savings if I cheated
through that wall.

My guess is that part 2 will let you cheat multiple times.   I think that
should be easy enough with the way I have it, but maybe not using Step()
for visualization, or I'd have to track the state of how many cheats are
used.

Easy-peasy, hardest part was adding the visualization stuff...

## Part 2

Ok a cheat can last for up to 20 picoseconds now.   At least you need to
use all the time at once.

Probably a better way, but I might just brute-force the 41x41 area around
each path node looking in my `path` map for other nodes with those 1,681
coordinates.  Hmmm.....  Let me check that path length...

Awww.   9356 nodes :(...  I thought about just checking the other nodes
back to the start so see if they were in range.   That might still
be viable, less than 80 million checks for sure