package main

import "strconv"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func absg[T number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func sum(nums []int) int {
	var res, n int
	for _, n = range nums {
		res += n
	}
	return res
}

func a2i(a []byte) int {
	i, err := strconv.Atoi(string(a))
	if err != nil {
		panic(err)
	}
	return i
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
