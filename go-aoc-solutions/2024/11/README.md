# 2024 Day 11

Thoughts:

* Brute force *might* work.  After 25 times for Part 1 that is just 55312
* Could be improved with memoization, but how much time does it really take to figure out one?
* Could follow numbers a certain number of steps for the memoization.
    * Blocks don't mingle, they just expand.
    * Could follow any number 5 into the future which is an entirely reasonable amount
    * Then when I see a number, I skip ahead 5 steps and use the previous result
* Ah, I think duplication removal is the best option
    * I think there will be A LOT of duplication
    * Things will get to 0s, 0s always become 1s, 1s become 2024, become 20/24, become 2/0/2/4, etc.
* Hmmm..   They mention that 'order is preserved' which is ominous for part 2
    * Part 1 doesn't seem to care about order

If order didn't matter, I think a combination of memoization and duplication
removal is near optimal.  But the duplication messes with the order.  In that
case maybe looking into the future is the best option.

Ah, maybe I do it as a tree.  Each number exists once in a map, pointing to it's node in a list.
That node has 1 or 2 nodes pointing to children, left and right.  Unless it is split, the right
node will be nil.   We then traverse the tree  until we reach a depth of 25.

Ok, let's see how that will work in GO.   Actually I think a 
`map[string]TreeNode` will work, where TreeNode has the strings used to access
other nodes in the map, or "" for no node:

```go
type TreeNode struct {
    Left, Right string
}
```

When we see a new number not in the map, we create a new node with Left and Right.

Doing a traversal should work, but could be slow with a lot of values.  Memoizing
based on depth in the traversal should speed it up, but shouldn't be needed
for part 1.  That way `[]int{0 1}` would figure out 0 along with a bunch of
others on the way down.   The first iteration will convert it to 1 and we
will have a solution for 1 at depth 24, so by the time we get to the 1 in the list
we will have only one step.   Many more along the way would be sped up too I'm sure.

That way we would create a new array for the target depth minus current depth, or
maybe of the entire depth so when we get to the 1 we can finish, or if the
depth increases for a future iteration we can finish from where we left off.

## Part 2

> NOTE: This still doesn't seem to depend on order, so maybe I could use de-duping
to speed it up

Yeah, going from a depth of 25 to 75 is going to be an issue without memoization.

At what point do I do that though?  What I think would work best would be to
maintain an array for each value and track every level.   I'm having trouble
picturing how to make that work though.  We couldn't track counts I don't think,
we'd have to track the actual list.   Or each node would have not only the number,
but the depth.   That's in reverse though, isn't it?   I was thinking this:

    Value: 0
        Depth 1 - [1] - Count 1, new value, add to queue
        Depth 2 - [10] - Count 1, new value, add to queue
        Depth 3 - [1 0] - Count 2, no new values, though one is 0, so we are in a loop

At this point of processing 0, we no longer have new unique values, we have
offshoots of 10 and 1 we need to process, then the remainder is a loop back to 0.

So to start with say '0' alone, we maybe maintain the full set until
we haven't seen anything new?

So if I keep track of depth, it could actually look like this:

    [{0,0}]
    [{1,0}]
    [{10,0}]
    [{1,0},{0,0}] -- both of these exist

I still think de-duping may be the fastest way.   I ran a little test to see how many
unique numbers were at later depths:

    Depth 20 created 466 keys
    2ms part1("input.aoc") = 26972 (BAD - Expected 216996)
    Depth 25 created 801 keys
    16ms part1("input.aoc") = 216996 (GOOD)
    PS C:\git\go\advent-workspace\go-aoc-solutions> go run 2024/11/main.go
    Depth 40 created 2516 keys
    7169ms part1("input.aoc") = 113789113 (BAD - Expected 216996)

So depth 40 took quite a while (7s), but only had 2516 unique numbers.  Doubling
the depth multiplied the number of keys by 6.   So for a depth of 75 it should
be under 15,000....    Let's do that, though it's a complete re-write :)


