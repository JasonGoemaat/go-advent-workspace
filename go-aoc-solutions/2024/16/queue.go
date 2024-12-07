package main

type Queue[T any] struct {
	backing    []T
	head, tail int
}

func NewQueue[T any](capacity int) *Queue[T] {
	var queue = Queue[T]{backing: make([]T, capacity), head: 0, tail: 0}
	return &queue
}

func (q *Queue[T]) Enqueue(t T) {
	next := (q.tail + 1) % len(q.backing)
	if next == q.head {
		// need to resize
		newBacking := make([]T, len(q.backing)*2)

		// need to copy 1-2 parts depending on if it spans the end of the array
		if q.tail < q.head {
			copy(newBacking, q.backing[q.head:len(q.backing)])
			hsize := len(q.backing) - q.head
			copy(newBacking[hsize:len(newBacking)-hsize:len(newBacking)-hsize], q.backing[0:q.tail])
			q.tail = hsize + q.tail
			q.head = 0
		} else {
			// easier, they're in order
			copy(newBacking, q.backing[q.head:q.tail])
			q.tail = q.tail - q.head
			q.head = 0
		}
		q.backing = newBacking
	}
	q.backing[q.tail] = t
	q.tail = (q.tail + 1) % len(q.backing)
}

func (q *Queue[T]) Dequeue() T {
	if q.head == q.tail {
		return q.backing[q.head] // not the best, but I hate multiple return values, best to check IsEmpty, eh?
	}

	current := q.head
	q.head = (q.head + 1) % len(q.backing)
	return q.backing[current]
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue[T]) Count() int {
	if q.head > q.tail {
		return (len(q.backing) - q.head) + q.tail
	}
	return q.tail - q.head
}
