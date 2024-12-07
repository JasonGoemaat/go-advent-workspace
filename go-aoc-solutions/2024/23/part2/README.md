# 2024 Day 23 (Part 2)

I think some kind of actual graph will be the best way to solve this.
I need to find largest groups that are all interconnected (and I think
one member at least starts with a 't').  I can start with one node to
check all of it's possibilities.

The first step will be to get all connected nodes at depth 1.  This will form a
new group where each node is connected to the first.

Starting with each of those, count all remaining nodes in that group that are
connected to it.  That is one group.

Then remove that node from the group and try all the remaining nodes.

Say A is connected to BCDE.   I want to then check BCDE, BCD, BCE, BDE, BD, BE.  

Hmm...   Maybe I want 