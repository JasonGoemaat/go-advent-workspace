package part2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JasonGoemaat/go-aoc/aoc"
)

type (
	Keypad struct {
		Rows       [][]string
		CurrentPos KeypadPos
		InitialPos KeypadPos
		Depth      int
		Type       int
	}
	KeypadPos struct {
		X, Y int
	}
)

func NewNumericKeypad() Keypad {
	keypad := Keypad{
		// Layout:     []string{"789", "456", "123", " 0A"},
		Rows: [][]string{
			{"7", "8", "9"},
			{"4", "5", "6"},
			{"1", "2", "3"},
			{" ", "0", "A"},
		},
		CurrentPos: KeypadPos{X: 2, Y: 3},
		InitialPos: KeypadPos{X: 2, Y: 3},
		Type:       0,
	}
	return keypad
}

func NewControlKeypad(depth int) Keypad {
	keypad := Keypad{
		// Layout:     []string{" ^A", "<v>"},
		Rows: [][]string{
			{" ", "^", "A"},
			{"<", "v", ">"},
		},
		CurrentPos: KeypadPos{X: 2, Y: 0},
		InitialPos: KeypadPos{X: 2, Y: 0},
		Type:       1,
		Depth:      depth,
	}
	return keypad
}

func NewPlayerKeypad(depth int) Keypad {
	keypad := Keypad{
		// Layout:     []string{" ^A", "<v>"},
		Rows: [][]string{
			{" ", "^", "A"},
			{"<", "v", ">"},
		},
		CurrentPos: KeypadPos{X: 2, Y: 0},
		InitialPos: KeypadPos{X: 2, Y: 0},
		Type:       2,
		Depth:      depth,
	}
	return keypad
}

func (keypad Keypad) FindKey(key string) *KeypadPos {
	for r, row := range keypad.Rows {
		for c, value := range row {
			if value == key {
				return &KeypadPos{c, r}
			}
		}
	}
	return nil
}

var waysController = `A^ <
A> v
Av v< <v
A< v<<
^A >
^> >v v>
^v v
^< v<
>A ^
>^ <^ ^<
>v <
>< <<
v^ ^
v> >
vA >^ ^>
v< <
<v >
<> >>
<^ >^
<A >>^`

// 789
// 456
// 123
//  0A

var waysKeypad = `
A0 <
0A >
A1 ^<<
1A >>v
A2 <^ ^<
2A v> >v
A3 ^
3A v
A4 ^^<<
4A >>vv
A5 <^^ ^^<
5A >vv vv>
A6 ^^
6A vv
A7 ^^^<<
7A >>vvv
A8 <^^^ ^^^<
8A >vvv vvv>
A9 ^^^
9A vvv

01 ^<
10 >v
02 ^
20 v
03 >^ ^>
30 v< <v
04 ^^<
40 >vv
05 ^^
50 vv
06 >^^ ^^>
60 vv< <vv
07 ^^<
70 >vvv
08 ^^^
80 vvv
09 >^^^ ^^^>
90 vvv< <vvv

12 >
21 <
13 >>
31 <<
14 ^
41 v
15 >^ ^>
51 v< <v
16 >>^ ^>>
61 <<v v<<
17 ^^
71 vv
18 >^^ ^^>
81 <vv vv<
19 >>^^ ^^>>
91 <<vv vv<<

23 >
32 <
24 ^< <^
42 v> >v
25 ^
52 v
26 >^ ^>
62 <v v<
27 ^^< <^^
72 vv> >vv
28 ^^
82 vv
29 >^^ ^^>
92 <vv vv<

34 <<^ ^<<
43 v>> >>v
35 <^ ^<
53 v> >v
36 ^
63 v
37 <<^^ ^^<<
73 >>vv vv>>
38 <^^ ^^<
83 vv> >vv
39 ^^
93 vv

45 >
54 <
46 >>
64 <<
47 ^
74 v
48 ^> >^
84 <v v<
49 ^>> >>^
94 v<< <<v

56 >
65 <<
57 <^ ^<
75 v> >v
58 ^
85 v
59 >^ ^>
95 v< <v

67 <<^ ^<<
76 >>v v>>
68 <^ ^<
86 v> >v
69 ^
96 v

78 >
87 <
79 >>
97 <<

89 >
98 <
`

func findWays(s string) map[string][]string {
	ways := map[string][]string{}
	for _, line := range aoc.ParseLines(s) {
		if len(line) > 0 {
			parts := strings.Split(line, " ")
			ways[parts[0]] = parts[1:len(parts)]
		}
	}
	return ways
}

var (
	KEYPAD_DIGITS     = 0
	KEYPAD_CONTROLLER = 1
	KEYPAD_PLAYER     = 2
)

var (
	keypadWays     = findWays(waysKeypad)
	controllerWays = findWays(waysController)
)

type State struct {
	Keypads []Keypad
	Memo    map[int]map[string]int // depth and move (i.e. A7 to move from A to 7) - cost just for move
}

// 'keypresses' is a sequence of keypresses, starting at 'A' is assumed
// and we will end on 'A' to submit the keypress.
func (state *State) CalculateCost(depth int, keypresses string) int {
	kp := state.Keypads[depth]

	// player keypad, cost is the length of keypresses
	if kp.Type == KEYPAD_PLAYER {
		return len(keypresses)
	}

	// if we've already calculated the cheapest way to do it at this depth, return it
	if cost, exists := state.Memo[depth][keypresses]; exists {
		return cost
	}

	// we need to calculate the costs and pick the cheapest
	lastKey := "A"
	cost := 0
	var ways map[string][]string
	if kp.Type == KEYPAD_DIGITS {
		ways = keypadWays
	} else {
		ways = controllerWays
	}
	for _, ch := range keypresses {
		newKey := string(ch)
		if newKey == lastKey {
			// any secondary press should just cost 1
			cost++
		} else {
			if moves, exists := ways[lastKey+newKey]; exists {
				minCost := state.CalculateCost(depth+1, moves[0]+"A")
				if len(moves) > 1 {
					minCost = min(minCost, state.CalculateCost(depth+1, moves[1]+"A"))
				}
				cost += minCost
			} else {
				panic("UNKNOWN MOVES")
			}
		}
		lastKey = newKey
	}
	state.Memo[depth][keypresses] = cost
	return cost
}

func CreateState(robotControllers int) *State {
	keypads := make([]Keypad, 0, robotControllers+2)
	keypads = append(keypads, NewNumericKeypad())
	for i := range robotControllers {
		keypad := NewControlKeypad(i + 1)
		keypads = append(keypads, keypad)
	}
	keypads = append(keypads, NewPlayerKeypad(robotControllers+1))
	memo := map[int]map[string]int{}
	for i := range len(keypads) {
		memo[i] = map[string]int{}
	}
	return &State{keypads, memo}
}

func solve(contents string, robotControllers int) interface{} {
	lines := aoc.ParseLines(contents)
	state := CreateState(robotControllers) // ONE robot controller
	total := 0
	for _, line := range lines {
		cost := state.CalculateCost(0, line)
		if value, err := strconv.Atoi(line[0:3]); err != nil {
			panic("ERROR GETTING INTEGER FROM " + line[0:3])
		} else {
			total += cost * value
		}
	}
	return total
}

func Part1(contents string) interface{} {
	return solve(contents, 2)
}

func Part2(contents string) interface{} {
	return solve(contents, 25)
}

// FOR DEBUGGING - problems with:
// 980A (59, should be 60, maybe missing an A?)

func (state *State) CalculateCostDisplay(depth int, keypresses string) int {
	kp := state.Keypads[depth]

	// player keypad, cost is the length of keypresses
	if kp.Type == KEYPAD_PLAYER {
		return len(keypresses)
	}

	// if we've already calculated the cheapest way to do it at this depth, return it
	if cost, exists := state.Memo[depth][keypresses]; exists {
		return cost
	}

	// we need to calculate the costs and pick the cheapest
	lastKey := "A"
	cost := 0
	var ways map[string][]string
	if kp.Type == KEYPAD_DIGITS {
		ways = keypadWays
	} else {
		ways = controllerWays
	}
	for _, ch := range keypresses {
		newKey := string(ch)
		if newKey == lastKey {
			// any secondary press should just cost 1
			cost++
		} else {
			move := lastKey + newKey
			if moves, exists := ways[move]; exists {
				minCost := state.CalculateCostDisplay(depth+1, moves[0]+"A")
				if len(moves) > 1 {
					minCost = min(minCost, state.CalculateCostDisplay(depth+1, moves[1]+"A"))
				}
				cost += minCost
			} else {
				panic("UNKNOWN MOVES")
			}
		}
		lastKey = newKey
	}
	state.Memo[depth][keypresses] = cost
	fmt.Printf("%sDepth %d keys '%s' cost %d\n", strings.Repeat("  ", depth), depth, keypresses, cost)
	return cost
}
