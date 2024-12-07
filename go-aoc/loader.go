package aoc

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetDir() string {
	fmt.Println("init() called in go-aoc/loader.go")
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	fmt.Println("    dir:", dir)
	return dir
}
