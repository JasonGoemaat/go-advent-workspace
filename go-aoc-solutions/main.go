package main

import (
	"fmt"

	"github.com/JasonGoemaat/go-aoc"
)

func init() {
	fmt.Println("go-aoc-2024 init() running")
}

func main() {
	fmt.Println("go-aoc-2024 main() running")
	fmt.Println("calling aoc.SayHello() from other module in workspace")
	aoc.SayHello()
}
