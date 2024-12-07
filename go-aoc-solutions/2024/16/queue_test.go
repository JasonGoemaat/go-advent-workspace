package main

import (
	"fmt"
	"testing"

	"github.com/JasonGoemaat/go-aoc/aoc"
)

func TestQueue1(t *testing.T) {
	queue := NewQueue[int](4)
	for i := range 8 {
		for j := range 5 {
			k := i*5 + j
			queue.Enqueue(k)
			t.Logf("Enqueue: %02d: head %d, tail %d, queue size: %d, ptr: %v", k, queue.head, queue.tail, len(queue.backing), &queue.backing[0])
			t.Logf("    %v", queue.backing)
		}
		for _ = range 3 {
			v := queue.Dequeue()
			t.Logf("Dequeue: %02d: head %d, tail %d, queue size: %d, ptr: %v", v, queue.head, queue.tail, len(queue.backing), &queue.backing[0])
			t.Logf("    %v", queue.backing)
		}
	}
	aoc.ExpectJson(t, "[18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 0 0 0 0 0 0 0 0 0 0]", fmt.Sprintf("%v", queue.backing))
}
