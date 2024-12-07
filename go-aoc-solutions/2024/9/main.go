package main

import (
	aoc "github.com/JasonGoemaat/go-aoc/aoc"
)

// input.aoc is 20,000 characters, so max 200,000 values

func main() {
	// https://adventofcode.com/2024/day/9
	aoc.Local(part1, "part1", "sample.aoc", 1928)
	aoc.Local(part1, "part1", "input.aoc", 6349606724455)
	aoc.Local(part2, "part2", "sample.aoc", 2858)
	aoc.Local(part2, "part2", "input.aoc", 6376648986651)
}

// Chars are FBFBFBFBFB... where F are files and B are blank areas in a disk.
// Each 'F' or 'B' is a number of blocks from 0-9
// Each 'F' has an id that is a 0-based index of the Files only as the start,
//
//	skipping the 'B's.
//
// The point is to move the files from the end onen block at a time into the
//
//	blank spaces.
//
// The result is the sum of (FileId * BlockId) for every file block after
//
//	there are no more blank spaces between file blocks.
//
// Thought 1: make a huge array of int representing the file id for each block
//
//	and use -1 for empty space.  Then move head and tail, looking for
//	empty spaces at the head and file blocks at the tail and swap them.
//	Stop when head crosses tail.   This is easy, let's do it.
func part1(contents string) interface{} {
	var bytes = []byte(contents) // MUST have no CRLF, one big string
	// first count total to make int array
	totalBlocks := 0
	for i, _ := range bytes {
		bytes[i] = bytes[i] - '0'
		totalBlocks += int(bytes[i])
	}
	blocks := make([]int, totalBlocks)
	head := 0
	for i, b := range bytes {
		for _ = range b {
			if (i & 1) == 0 {
				// file, set to file id
				blocks[head] = (i >> 1)
			} else {
				// for blank, set to -1
				blocks[head] = -1
			}
			head++
		}
	}
	head = 0
	tail := totalBlocks - 1
	for tail > head {
		// position head over empty space and tail over a file
		for ; tail > head && blocks[head] != -1; head++ {
		}
		for ; tail > head && blocks[tail] == -1; tail-- {
		}
		blocks[tail], blocks[head] = blocks[head], blocks[tail]
		head++
		tail--
	}
	// now compute checksum
	checksum := 0
	for i := 0; (i < totalBlocks) && (blocks[i] >= 0); i++ {
		checksum += (blocks[i] * i)
	}
	return checksum
}

// Of course part 2 deals with fragmentation in some manner...
// Attempt to move each file exactly once then ignore it.
// gotta be a better way here.  Well, maybe not.   We're not
// necessarily filling the empty space, so we could leave
// earlier gaps that future smaller files may still fit into.
// I think we need structs of some sort of struct.
// Hmmmm...   Maybe go slices are the way to go?  We never
// actually split blanks, we just shrink them.  So we could
// start with arrays of blanks areas and files.  Head and Tail
// are then indexes into the Blanks and reverse indexes into
// the Files.  We will need to move the file array forward and
// place the file from the end earlier, but that's not a big
// deal with 10k files.

// OOOHHH, the criteria make it even easier.   We don't actually
// have to move anything since we're tracking offset.   When we
// 'move' a file to a blank spot, we just change the index and
// then decrease the size of the blank spot and shift its index
// forward by the file size.   At this point we COULD remove
// the empty blanks to make future lookups faster.  In this case
// I think keeping files and blanks separate makes sense

type section struct {
	id     int // -1 for blank
	length int
	offset int
}

func part2(contents string) interface{} {
	bytes := []byte(contents)
	sections := make([]section, len(bytes))
	offset := 0 // actual block index on disk
	for i, b := range bytes {
		id := -1 // identifies empty
		size := int(b - '0')
		if (i & 1) == 0 {
			id = i >> 1
		}
		sections[i] = section{id, size, offset}
		offset += size
	}
	// 0xfffffffe removes the last bit, ignoring blank if it's last
	for tail := (len(sections) - 1) & 0xfffffffe; tail > 0; tail -= 2 {
		file := &sections[tail]
		for head := 1; head < tail; head += 2 {
			blank := &sections[head]
			if blank.offset > file.offset {
				// can't move
				break
			}
			if blank.length >= file.length {
				// move file offset, shrink blank from the front
				file.offset = blank.offset
				blank.offset += file.length
				blank.length -= file.length
				break
			}
		}
	}
	checksum := 0
	for fileIndex := 0; fileIndex < len(sections); fileIndex += 2 {
		file := &sections[fileIndex]
		for i := range file.length {
			checksum += (i + file.offset) * file.id
		}
	}
	return checksum
}
