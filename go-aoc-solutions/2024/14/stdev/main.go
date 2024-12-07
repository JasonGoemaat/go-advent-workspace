package main

//--------------------------------------------------------------------------------
// THIS is to produce images for part 2
//--------------------------------------------------------------------------------

import (
	"cmp"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"slices"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

var (
	_heights [][]int
	_widths  [][]int
)

func main() {
	firstCmp := func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	}
	if bytes, err := os.ReadFile(aoc.GetSubPath("_hs.txt")); err == nil {
		groups := aoc.ParseGroups(string(bytes))
		_heights = aoc.ParseLinesToInts(aoc.ParseLines(groups[0]))
		_widths = aoc.ParseLinesToInts(aoc.ParseLines(groups[1]))
		slices.SortFunc(_heights, firstCmp)
		slices.SortFunc(_widths, firstCmp)
	} else {
		panic(err)
	}

	// https://adventofcode.com/2015/day/14
	// aoc.Local(testStdDev, "solveMe", "input.aoc", 0) // try some options...
	// aoc.Local(bestStdDevsHEIGHT, "solveMe", "unused.aoc", 0) // find 11-299
	// aoc.Local(bestStdDevsWIDTH, "solveMe", "unused.aoc", 0) // find 11-299
	aoc.Local(drawMatches, "drawMatches", "unused.aoc", 0) // use matches above
}

// use _heights and _widths to make one big picture
func drawMatches(content string) interface{} {
	// numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))

	lefts := []int{}
	next := 0
	for _, v := range _widths {
		lefts = append(lefts, next)
		next += (v[0] + 1) // 1 for border, draw at lefts[col]-1
	}
	imageWidth := next

	tops := []int{}
	next = 0
	for _, v := range _heights {
		tops = append(tops, next)
		next += (v[0] + 1)
	}
	imageHeight := next

	fmt.Printf("Will generate image %d x %d\n", imageWidth, imageHeight)
	return 0
}

func testStdDev(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 101, 103

	// sd1.txt is full over w*h, which has the exact same value
	// repeating every h!
	file, err := os.Create(aoc.GetSubPath("sd2.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for sec := range h { // we ONLY need h!   since it repeats, they are always the same!
		robots, _ := getRobots(numLists, w, h, sec)
		fmt.Fprintf(file, "%d\t%f\n", sec, getStandardDeviationHeight(robots))
	}

	return 0
}

func bestStdDevsWIDTH(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	h := 103 // doesn't matter

	// sd1.txt is full over w*h, which has the exact same value
	// repeating every h!  sd2.txt shows original input.aoc has
	// MASSIVE low at step 76 (457, next is 631, avg is 663)
	// this one will go through a ton of heights and try to find
	// low standard deviations after stripping the original christmas tree's points

	type STD struct {
		std   float64
		w, s  int
		score float64
	}

	stds := []STD{}

	for w := 11; w < 121; w++ {
		addedSeconds := h * w
		totalSTD := float64(0)
		mystds := []STD{}
		for sec := range w {
			robots, _ := getRobots(numLists, w, h, sec+addedSeconds) // adding seconds for high numbers where robots start clustered to the left/top
			// that doesn't solve it, dummy!   every loop will get us back to the crappy starting position
			std := getStandardDeviationWidth(robots)
			totalSTD += std
			mystds = append(mystds, STD{std, w, sec, 0})
			// fmt.Fprintf(file, "%f\t%d\t%d\n", getStandardDeviationHeight(robots), h, sec)
		}
		averageSTD := totalSTD / float64(w)
		for _, s := range mystds {
			stds = append(stds, STD{s.std, s.w, s.s, s.std / averageSTD})
		}
	}

	file, err := os.Create(aoc.GetSubPath("sd3w.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "SCORE\tStdDev\twidth\tsecond\n")
	for _, s := range stds {
		fmt.Fprintf(file, "%f\t%f\t%d\t%d\n", s.score, s.std, s.w, s.s)
	}

	return 0
}

func bestStdDevsHEIGHT(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w := 101 // doesn't matter

	// sd1.txt is full over w*h, which has the exact same value
	// repeating every h!  sd2.txt shows original input.aoc has
	// MASSIVE low at step 76 (457, next is 631, avg is 663)
	// this one will go through a ton of heights and try to find
	// low standard deviations after stripping the original christmas tree's points

	type STD struct {
		std   float64
		h, s  int
		score float64
	}

	stds := []STD{}

	for h := 11; h < 121; h++ {
		addedSeconds := h * w // for higher values like 299 they are grouped in the original 103
		totalSTD := float64(0)
		mystds := []STD{}
		for sec := range h { // we ONLY need h!   since it repeats, they are always the same!
			robots, _ := getRobots(numLists, w, h, sec+addedSeconds) // adding seconds for high numbers where robots start clustered to the left/top
			// that doesn't solve it, dummy!   every loop will get us back to the crappy starting position
			std := getStandardDeviationHeight(robots)
			totalSTD += std
			mystds = append(mystds, STD{std, h, sec, 0})
			// fmt.Fprintf(file, "%f\t%d\t%d\n", getStandardDeviationHeight(robots), h, sec)
		}
		averageSTD := totalSTD / float64(h)
		for _, s := range mystds {
			stds = append(stds, STD{s.std, s.h, s.s, s.std / averageSTD})
		}
	}

	file, err := os.Create(aoc.GetSubPath("sd3.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "SCORE\tStdDev\theight\tsecond\n")
	for _, s := range stds {
		fmt.Fprintf(file, "%f\t%f\t%d\t%d\n", s.score, s.std, s.h, s.s)
	}

	return 0
}

// calculate standard deviation in y value, low value will
// mean there should be a horizontal band
func getStandardDeviationHeight(robots Robots) float64 {
	totalHeight := 0
	totalCount := 0
	for pos, count := range robots {
		totalHeight += pos.y * count
		totalCount += count
	}
	mean := float64(totalHeight) / float64(totalCount)
	variance := float64(0)
	for pos, count := range robots {
		y := math.Abs(float64(pos.y) - mean)
		variance += float64(count) * (y * y)
	}
	return math.Sqrt(variance)
}

// calculate standard deviation in x value, low value will
// mean there should be a vertical band
func getStandardDeviationWidth(robots Robots) float64 {
	totalWidth := 0
	totalCount := 0
	for pos, count := range robots {
		totalWidth += pos.x * count
		totalCount += count
	}
	mean := float64(totalWidth) / float64(totalCount)
	variance := float64(0)
	for pos, count := range robots {
		x := math.Abs(float64(pos.x) - mean)
		variance += float64(count) * (x * x)
	}
	return math.Sqrt(variance)
}

type (
	Position     struct{ x, y int }
	Robots       map[Position]int
	RobotIndexes map[Position][]int
)

func getRobots(numLists [][]int, w, h, seconds int) (Robots, RobotIndexes) {
	robots := Robots{}
	indexes := RobotIndexes{}
	for i, nums := range numLists {
		p := Position{(nums[0] + nums[2]*seconds) % w, (nums[1] + nums[3]*seconds) % h}
		if p.x < 0 {
			p.x += w
		}
		if p.y < 0 {
			p.y += h
		}
		robots[p] += 1
		if robots[p] == 1 {
			indexes[p] = []int{i}
		} else {
			indexes[p] = append(indexes[p], i)
		}
	}
	return robots, indexes
}

func findMissing(contents string) interface{} {
	// on xmas tree image, find 28,28 inside frame
	// move up to hit top of frame
	// find rectangle
	// find all the points that end up in the rectangle and remove them
	// output remaining original points
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	w, h := 101, 103 // actual input size

	seconds := 7286 // my answer for part 2 with the tree
	robots, _ := getRobots(numLists, w, h, seconds)

	// find exact bounding box of tree frame, starting inside I know at 28,28
	x1, y1 := 28, 28
	for y1 > 0 {
		// fmt.Printf("y checking %d, %d\n", x1, y1)
		if _, ok := robots[Position{x1, y1}]; ok {
			// fmt.Printf("	FOUND %d, %d\n", x1, y1)
			break
		}
		y1--
	}
	for x1 > 1 {
		// fmt.Printf("x checking %d, %d\n", x1, y1)
		if _, ok := robots[Position{x1 - 1, y1}]; !ok {
			// fmt.Printf("	FOUND %d, %d\n", x1, y1)
			break
		}
		x1--
	}
	// have top-left, find bottom and right
	x2 := x1 + 1
	for x2 < 1000 {
		if _, ok := robots[Position{x2 + 1, y1}]; !ok {
			break
		}
		x2++
	}
	// have top-left, find bottom and right
	y2 := y1 + 1
	for y2 < 1000 {
		if _, ok := robots[Position{x1, y2 + 1}]; !ok {
			break
		}
		y2++
	}
	fmt.Printf("%d box at %d, %d - %d, %d\n", seconds, x1, y1, x2, y2)
	// 23,21 - 53,53

	return 0
}

func removeTree(numLists [][]int) [][]int {
	// remove any points that wil end up in tree frame at 7286
	w, h := 101, 103                 // actual input size
	x1, y1, x2, y2 := 23, 21, 53, 53 // found using findMissing()
	results := [][]int{}
	for _, l := range numLists {
		x := l[0] + (l[2]*7286)%w
		y := l[1] + (l[3]*7286)%h
		if x < 0 {
			x += w
		}
		if y < 0 {
			y += h
		}
		if x < x1 || x > x2 || y < y1 || y > y2 {
			results = append(results, l)
		}
	}
	return results
}

var (
	green = color.RGBA{128, 255, 128, 255}
	white = color.RGBA{255, 255, 255, 255}
	black = color.RGBA{0, 0, 0, 255}
)

func outputUnused(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	result := removeTree(numLists)
	for _, r := range result {
		fmt.Printf("p=%d,%d v=%d,%d\n", r[0], r[1], r[2], r[3])
	}
	return 0
}

func createGraph(numLists [][]int, w, h, start, step int) *image.RGBA {
	rows, cols := 10, 20
	imageHeight, imageWidth := h*rows+(rows-1), w*cols+(cols-1)
	m := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for y := range imageHeight {
		for x := range imageWidth {
			m.SetRGBA(x, y, black)
		}
	}
	for row := 1; row < rows; row++ {
		for x := range imageWidth {
			m.SetRGBA(x, row*(h+1), white)
		}
	}
	for col := 1; col < cols; col++ {
		for y := range imageHeight {
			m.SetRGBA(col*(w+1), y, white)
		}
	}

	// now solve
	second := start
	for row := range rows {
		oy := row * (h + 1)
		for col := range cols {
			ox := col * (w + 1)
			robots, _ := getRobots(numLists, w, h, second)

			for pos, _ := range robots {
				m.Set(pos.x+ox, pos.y+oy, green)
			}

			second += step
		}
	}
	return m
}

type Combo struct {
	w, sw, h, sh int // width value, height value, and starting seconds for those
}

// return second owhen they will cross
func (c Combo) Second() int {
	aw, ah := c.sw, c.sh
	max := (c.w * c.h) + c.sw + c.sh
	max = max * max
	for aw < max {
		if aw == ah {
			return aw
		}
		diff := ah - aw
		if diff < -1000 {
			max++
		}
		if aw < ah {
			aw += c.w
		} else if ah < aw {
			ah += c.h
		}
	}
	fmt.Printf("No result: %v (last checked %d, %d)\n", c, aw, ah)
	return -1
}

func createCombos(numLists, horiz, vert [][]int) *image.RGBA {
	// horiz are different values for WIDTH, and show vertical patterns
	maxW, maxH := 0, 0
	combos := []Combo{}
	for _, x := range horiz {
		for _, y := range vert {
			maxW = max(x[0], maxW)
			maxH = max(y[0], maxH)
			c := Combo{x[0], x[1], y[0], y[1]}
			combos = append(combos, c)
		}
	}
	rows, cols := 10, 14
	imageHeight, imageWidth := maxH*rows+(rows-1), maxW*cols+(cols-1)
	m := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for y := range imageHeight {
		for x := range imageWidth {
			m.SetRGBA(x, y, black)
		}
	}
	for row := 1; row < rows; row++ {
		for x := range imageWidth {
			m.SetRGBA(x, row*(maxH+1), white)
		}
	}
	for col := 1; col < cols; col++ {
		for y := range imageHeight {
			m.SetRGBA(col*(maxW+1), y, white)
		}
	}

	// now display combos
	for i, combo := range combos {
		row := i / cols
		col := i % cols
		second := combo.Second()
		if second < 0 {
			continue
			// panic("BAD COMBO")
		}
		robots, _ := getRobots(numLists, combo.w, combo.h, second)
		ox := col * (maxW + 1)
		oy := row * (maxH + 1)
		for pos, _ := range robots {
			m.Set(pos.x+ox, pos.y+oy, green)
		}
	}
	return m
}

func save(subPath string, image *image.RGBA) {
	pngPath := aoc.GetSubPath(subPath)
	file, err := os.Create(pngPath)
	if err != nil {
		panic(fmt.Sprintf("ERROR CREATING FILE: %s", pngPath))
	}
	defer file.Close()
	png.Encode(file, image)
}

func solveMe(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	// image := createGraph(numLists, 107, 109, 0, 1) // 15 has vertical spike at 107, don't see 109 though
	// image := createGraph(numLists, 107, 109, 15, 107) // 15 has spike at 107, don't see 109 though -- DEFINITELY SOMETHING AT 107
	// image := createGraph(numLists, 107, 107, 0, 1) // horizontal at 79, vertical kinda at 15
	// image := createGraph(numLists, 105, 107, 79, 107) // something at 99th iteration?
	// image := createGraph(numLists, 105, 107, 0, 1) //??

	// image := createGraph(numLists, 29, 31, 0, 1) // well, tree was 29x31 inside the frame...  this is too small, way filled

	for x := 61; x < 127; x += 2 {
		// ----- IGNORE THESE, I was using full data set, not unused, so they are artifacts of the other one I think
		// Horizontal at 76 repeating at 103
		// 79 has vertical for sure at 11, repeating at 79
		// 87 has vertical at 12
		// 93 at 13, 95 kind at 13 also
		// 101 of course at 14
		// 107 kinda at 15
		// 109 kinda at 15
		// 115 at 16
		// 121, 123 kinda at 17
		// -----
		// 87@12, 93@13?, 107@15?
		image := createGraph(numLists, x, 103, 0, 1)
		save(fmt.Sprintf("widths/%03d.png", x), image)
	}

	for y := 61; y < 127; y += 2 {
		// 95@70?  99@73?  107@79?, 111@83?, 115@85?
		image := createGraph(numLists, 101, y, 0, 1)
		save(fmt.Sprintf("heights/%03d.png", y), image)
	}

	return 0
}

// using information visually gotten from results of solveMe
func useKnown(contents string) interface{} {
	numLists := aoc.ParseLinesToInts(aoc.ParseLines(contents))
	horiz := [][]int{{87, 12}, {93, 13}, {107, 15}}
	vert := [][]int{{95, 70}, {99, 73}, {107, 79}, {111, 83}, {115, 85}}
	image := createCombos(numLists, horiz, vert)
	save("useKnown.png", image)
	return 0
}
