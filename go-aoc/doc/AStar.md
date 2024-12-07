# astar.AStar

Some puzzles require finding a path or paths.   There is an `astar` package to
help.  See 2024 days 18 and 20 in my solutions for how they are used.

### Structs

```go
AStar struct {
    Area       *aoc.Area
    Open       map[AStarPosition]*AStarNode
    Closed     map[AStarPosition]*AStarNode
    Start, End AStarPosition
    Last       *AStarNode
}

AStarPosition struct {
    X, Y int
}

AStarNode struct {
    F, G, H  int
    Position AStarPosition
    Parent   *AStarNode
}
```

## Methods

```go
func NewAStar(area *aoc.Area, start, end AStarPosition) *AStar
func (as *AStar) GetShortestPath() *AStarNode
func (as *AStar) StepShortestPath() (*AStarNode, bool)
func (as *AStar) Reset()
func (node *AStarNode) Length() int
```

### NewAStar

This works with the `Area` struct and functions above, so parse the area then
you can create an `*Astar` to find paths with this:

```go
func NewAStar(area *aoc.Area, start, end AStarPosition) *AStar
```

`start` and `end` need to be provided.   Day 18 uses the top-left and
bottom-right of the area.   The `AStarPosition` is created like so:

```go
start := astar.AStarPosition{X: 0, Y: 0}
```

### GetShortestPath

You can find the shortest path calling this, which returns the end node
where following nodes back using `.Parent` to get to the start will give
you the full path.

```go
func (as *AStar) GetShortestPath() *AStarNode {
	node, done := as.StepShortestPath()
	for !done {
		node, done = as.StepShortestPath()
	}
	return node
}
```

Note the `StepShortestPath()` method that can be called separately if you
want to do visualization and use the `Open`, `Closed`, and `Last` properties
to render the visualization.

### Reset

There's a `Reset()` method that will reset the path finder.   This was used
in 2024 Day 18.   In this puzzle for part 2 we would drop new ceiling tiles
onto the area and needed to find which one broke the path.   For this to be
used we found the initial path.   For each tile we checked if it fell on the
existing path, and if so we reset and calculated a new path and stopped when
no path was found.

### Length

Just a helper function you can call on the result of `GetShortestPath()` to
return the number of steps in the path.

```go
func (node *AStarNode) Length() int
```
