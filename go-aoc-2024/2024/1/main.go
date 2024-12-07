package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/JasonGoemaat/go-aoc"
)

func init() {
	fmt.Println("init() called in 2024/1/main.go")
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	fmt.Println(dir)
}

func main() {
	fmt.Println("2024 Day 1")
	dir := aoc.GetDir()
	fmt.Println("aoc.GetDir() returned:", dir)
}
