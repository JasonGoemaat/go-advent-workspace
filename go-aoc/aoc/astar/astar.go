package astar

import "github.com/JasonGoemaat/go-aoc/aoc"

// as of go 1.18, can use generic
type (
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
)

func (node *AStarNode) Length() int {
	if node == nil {
		return 0
	}
	length := 0
	for n := node; n.Parent != nil; n = n.Parent {
		length++
	}
	return length
}

func NewAStar(area *aoc.Area, start, end AStarPosition) *AStar {
	result := AStar{
		Area:   area,
		Open:   map[AStarPosition]*AStarNode{},
		Closed: map[AStarPosition]*AStarNode{},
		Start:  start,
		End:    end,
	}
	result.Open[start] = result.CreateOrUpdateNode(nil, start, 0)
	return &result
}

func (as *AStar) Reset() {
	as.Open = map[AStarPosition]*AStarNode{}
	as.Closed = map[AStarPosition]*AStarNode{}
	as.Last = nil
	as.Open[as.Start] = as.CreateOrUpdateNode(nil, as.Start, 0)
}

// Create new node or update existing node if it exists and the cost is less
// than the existing cost of an open node.   Will return nil if the node is
// already closed or outside the bounds of the area.
func (as *AStar) CreateOrUpdateNode(parent *AStarNode, pos AStarPosition, cost int) *AStarNode {
	if pos.X < 0 || pos.X >= as.Area.Width || pos.Y < 0 || pos.Y >= as.Area.Height {
		return nil
	}

	// check that area has a valid spot to move
	// TODO: check ValidBytes?  For now just disallow walls
	if as.Area.Get(pos.Y, pos.X) == '#' {
		return nil
	}

	if as.Closed[pos] != nil {
		return nil
	}
	open := as.Open[pos]
	if open != nil {
		if cost < open.G {
			open.G = cost
			open.F = open.H + open.G
			open.Parent = parent
		}
		return open
	}
	newOpen := AStarNode{Position: pos, Parent: parent}
	newOpen.H = as.CalculateH(pos)
	newOpen.G = cost
	newOpen.F = newOpen.G + newOpen.H
	as.Open[pos] = &newOpen
	return &newOpen
}

func (as *AStar) CalculateH(pos AStarPosition) int {
	dx := as.End.X - pos.X
	dy := as.End.Y - pos.Y
	if dx < 0 {
		dx = 0 - dx
	}
	if dy < 0 {
		dy = 0 - dy
	}

	// // full diagonal:
	// l, h := min(dx, dy), max(dx, dy)
	// return l*14 + (h-l)*10 // if we want to do diagonals, 1.4 vs. 1.0 is good approximation

	// simple taxicab
	return (dx + dy) * 10

	// NOTE: This doesn't actually work, duh!
	// // slightly prefer diagonal to straight...  In 100x100 open grid from one corner to another,
	// // taxicab will fill in the whole grid with diagonals to find the shortest path.  This is
	// // because right 100 and down 100 is the same distance as down 100 and right 100, or down 1, right 100, down 99.
	// // This sill prefer diagonals as if we could move diagonally.   So right 1 and down 1
	// // as a cost of 14, while right 2 has a cost of 20, looking at the heuristic.   I think this
	// // makes some sense as if the top or left of the end but not both were blocked, being at the
	// // top-left of the end gives you 2 paths to choose from with the same cost
	// l, h := min(dx, dy), max(dx, dy)
	// return l*4 + h*10 // if we want to do diagonals, 1.4 vs. 1.0 is good approximation.
}

func (as *AStar) GetShortestPath() *AStarNode {
	node, done := as.StepShortestPath()
	for !done {
		node, done = as.StepShortestPath()
	}
	return node
}

// Perform single-step in A* algorithm, returning the node that was checked (or nil if none)
// and a flag that is true if the algorithm is done.   When returning true for the flag, the
// node will be the shortest path (check .G or .F for final cost, follow .Parent for entire
// path), or nil if no path was found.
func (as *AStar) StepShortestPath() (*AStarNode, bool) {
	if as.Closed[as.End] != nil {
		// shortest path end node and true if solution found
		as.Last = as.Closed[as.End]
		return as.Closed[as.End], true
	}
	var smallest *AStarNode
	for _, node := range as.Open {
		if smallest == nil || node.F < smallest.F {
			smallest = node
		} else if node.F == smallest.F {
			if node.H < smallest.H {
				// same F, but smaller H, prefer closer to target
				smallest = node
			} else if node.H == smallest.H {
				// no reason to pick one over the other, except we want it to be
				// deterministic so we use dx first, then dy
				if smallest.Position.X < node.Position.X {
					smallest = node
				} else if smallest.Position.X == node.Position.X {
					if smallest.Position.Y < node.Position.Y {
						smallest = node
					}
				}
			}
		}
	}
	as.Last = smallest
	if smallest == nil {
		// nil and true if there are no solutions
		return nil, true
	}
	delete(as.Open, smallest.Position)
	as.Closed[smallest.Position] = smallest
	if smallest.Position == as.End {
		// this was the final end node, return true
		return smallest, true
	}

	// not at the end, add or update surrounding nodes
	as.CreateOrUpdateNode(smallest, AStarPosition{smallest.Position.X - 1, smallest.Position.Y}, smallest.G+10)
	as.CreateOrUpdateNode(smallest, AStarPosition{smallest.Position.X + 1, smallest.Position.Y}, smallest.G+10)
	as.CreateOrUpdateNode(smallest, AStarPosition{smallest.Position.X, smallest.Position.Y - 1}, smallest.G+10)
	as.CreateOrUpdateNode(smallest, AStarPosition{smallest.Position.X, smallest.Position.Y + 1}, smallest.G+10)
	return smallest, false
}
