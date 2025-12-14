package main

import (
	"regexp"

	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

func main() {
	aoc.Local(part1, "part1", "sample.aoc", 5)
	aoc.Local(part1, "part1", "input.aoc", 772)
	aoc.Local(part2, "part2", "sample2.aoc", 2)
	aoc.Local(part2, "part2", "input.aoc", 423227545768872)
}

func part1(content string) any {
	lines := aoc.ParseLines(content)
	rxIds := regexp.MustCompile("[a-z]+")
	paths := make(map[string][]string)
	for _, line := range lines {
		ids := rxIds.FindAllString(line, -1)
		paths[ids[0]] = ids[1:]
	}

	total := 0
	var goFrom func(string)
	goFrom = func(id string) {
		if id == "out" {
			total++
			return
		}
		for _, nextId := range paths[id] {
			goFrom(nextId)
		}
	}
	goFrom("you")
	return total
}

func part2(content string) any {
	lines := aoc.ParseLines(content)
	rxIds := regexp.MustCompile("[a-z]+")
	paths := make(map[string][]string)
	backpaths := make(map[string][]string)
	for _, line := range lines {
		ids := rxIds.FindAllString(line, -1)
		from := ids[0]
		to := ids[1:]
		paths[from] = to
		for _, child := range to {
			parents, exists := backpaths[child]
			if exists {
				parents = append(parents, from)
			} else {
				parents = []string{from}
			}
			backpaths[child] = parents
		}
	}

	queue := make([]string, 0)
	visited := make(map[string]bool)
	queue = append(queue, "out")
	visited["out"] = true
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, parent := range backpaths[current] {
			if !visited[parent] {
				queue = append(queue, parent)
				visited[parent] = true
			}
		}
	}

	// first ensure both maps have all devices, even
	// if the list of parents/children are empty
	for _, value := range paths {
		for _, id := range value {
			_, ok := paths[id]
			if !ok {
				paths[id] = make([]string, 0)
			}
		}
	}
	for _, value := range backpaths {
		for _, id := range value {
			_, ok := backpaths[id]
			if !ok {
				backpaths[id] = make([]string, 0)
			}
		}
	}

	idList := make([]string, 0)
	for key := range paths {
		idList = append(idList, key)
	}

	moved := make(map[string]bool)
	sortedIdList := make([]string, 0, len(idList))

	var hasParent = func(id string) bool {
		for _, parent := range backpaths[id] {
			if !moved[parent] {
				return true
			}
		}
		return false
	}

	for len(idList) > 0 {
		for _, id := range idList {
			if !hasParent(id) {
				sortedIdList = append(sortedIdList, id)
				moved[id] = true
			}
		}
		newIndex := 0
		for i := range idList {
			id := idList[i]
			if !moved[id] {
				if newIndex != i {
					idList[newIndex] = idList[i]
				}
				newIndex++
			}
		}
		// resize slice, chopping off end
		idList = idList[:newIndex]
	}

	// fmt.Println(sortedIdList)

	indexForDevice := make(map[string]int)
	for i, id := range sortedIdList {
		indexForDevice[id] = i
	}

	counts := make([]int, len(sortedIdList))

	// first loop, find first of 'dac' or 'fft'
	// counts is total ways to get to the device
	start := indexForDevice["svr"]
	counts[start] = 1 // visiting svr first, need to find 'dac' or 'fft'
	for i := start; i < len(sortedIdList); i++ {
		id := sortedIdList[i]
		if id == "dac" || id == "fft" {
			// clear counts past, so we only count paths flowing
			// through this device
			for j := i + 1; j < len(sortedIdList); j++ {
				counts[j] = 0
			}
			start = i
			break
		}

		// add current counts to children
		for _, child := range paths[id] {
			childIndex := indexForDevice[child]
			counts[childIndex] += counts[i]
		}
	}

	// second loop, find other of 'dac' or 'fft'
	// counts is total ways to get to the device
	for i := start; i < len(sortedIdList); i++ {
		id := sortedIdList[i]
		if i != start && (id == "dac" || id == "fft") {
			// clear counts past, so we only count paths flowing
			// through this device
			for j := i + 1; j < len(sortedIdList); j++ {
				counts[j] = 0
			}
			start = i
			break
		}

		// add current counts to children
		for _, child := range paths[id] {
			childIndex := indexForDevice[child]
			counts[childIndex] += counts[i]
		}
	}

	// third loop, just looking for end, counts should only include
	// paths going through both fft and dac
	result := -1
	for i := start; i < len(sortedIdList); i++ {
		id := sortedIdList[i]
		if id == "out" {
			result = counts[i]
			break
		}

		// add current counts to children
		for _, child := range paths[id] {
			childIndex := indexForDevice[child]
			counts[childIndex] += counts[i]
		}
	}

	return result
}
