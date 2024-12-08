package aoc

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func loadString(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func solveFile(puzzlePath string, parts ...func(string) interface{}) {
	contents, err := loadString(puzzlePath)
	fileName := filepath.Base(puzzlePath)
	if err != nil {
		fmt.Println("COULD NOT LOAD:", puzzlePath)
		return
	}
	fmt.Printf("Using: %s\n", fileName)
	for i, part := range parts {
		startTime := time.Now()
		result := part(contents)
		endTime := time.Now()
		ms := endTime.Sub(startTime).Milliseconds()
		fmt.Printf("  Part %d in %d ms: %v\n", i+1, ms, result)
	}
}

// SolveLocal takes an array of solver functions and calls them with the
// contents of 'sample.txt' and 'input.txt' files in the same diretory as
// the `.go` file making the call.   It reports the results and how much
// time it took for the solving function to run.
func SolveLocal(parts ...func(string) interface{}) {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)

	samplePath := filepath.Join(dir, "sample.txt")
	solveFile(samplePath, parts...)
	inputPath := filepath.Join(dir, "input.txt")
	solveFile(inputPath, parts...)
}
