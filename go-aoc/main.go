package aoc

import "fmt"

func SayHello() {
	fmt.Println("Hello, World! (from go-aoc/main.go SayHello())")
}

func init() {
	fmt.Println("go-aoc/main.go init() running")
}
