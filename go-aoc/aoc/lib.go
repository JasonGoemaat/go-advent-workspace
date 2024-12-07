package aoc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var LoggingEnabled = false

func LogF(format string, args ...interface{}) {
	if !LoggingEnabled {
		return
	}
	fmt.Printf(format, args...)
}

func loadString(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func solveFile(puzzlePath string, parts ...func(string) interface{}) {
	contents, err := loadString(puzzlePath)
	if err != nil {
		fmt.Println("COULD NOT LOAD:", puzzlePath)
		return
	}
	if len(contents) == 0 {
		fmt.Printf("NO INPUT IN FILE: %s\n", puzzlePath)
	}

	fileName := filepath.Base(puzzlePath)
	fmt.Printf("Using: %s\n", fileName)
	for i, part := range parts {
		// check for existing file to override for this test
		// Example: 2024 Day 3 has a different sample for part 2,
		// so this looks for a file called 'sample.aoc' when
		// processing index 1 representing part 2
		startTime := time.Now()
		result := part(contents)
		endTime := time.Now()
		ms := endTime.Sub(startTime).Milliseconds()
		if len(parts) > 1 {
			fmt.Printf("  Part %d in %d ms: %v\n", i+1, ms, result)
		} else {
			fmt.Printf("  Solved in %d ms: %v\n", ms, result)
		}
	}
}

// GetSubPath returns a full directory path relative to the calling `.go`
// file.
func GetSubPath(subPath string) string {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, subPath)
}

// Return full path to passed relative path, creating it as a directory if
// it doesn't already exist.
func GetOrCreateSubDirectory(subPath string) string {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	fullDir := filepath.Join(dir, subPath)
	_ = os.MkdirAll(fullDir, 0755)
	return fullDir
}

// SolveLocal takes an array of solver functions and calls them with the
// contents of 'sample.aoc' and 'input.aoc' files in the same diretory as
// the `.go` file making the call.   It reports the results and how much
// time it took for the solving function to run.
func SolveLocal(parts ...func(string) interface{}) {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)

	samplePath := filepath.Join(dir, "sample.aoc")
	solveFile(samplePath, parts...)
	inputPath := filepath.Join(dir, "input.aoc")
	solveFile(inputPath, parts...)
}

// sample call: aoc.Local(part1, "Part1", "sample.aoc", 14)
func Local(solver func(string) interface{}, name string, fileName string, expected interface{}) {
	_, caller, _, _ := runtime.Caller(1)
	callerDir := filepath.Dir(caller)
	inputPath := filepath.Join(callerDir, fileName)
	contents, _ := loadString(inputPath)
	if len(contents) == 0 {
		fmt.Printf("EMPTY FILE: %s\n", fileName)
	}
	startTime := time.Now()
	result := solver(contents)
	endTime := time.Now()
	ms := endTime.Sub(startTime).Milliseconds()
	if JsonEquals(result, expected) {
		fmt.Printf("%dms %s(%q) = %v (GOOD)\n", ms, name, fileName, result)
	} else {
		fmt.Printf("%dms %s(%q) = %v (BAD - Expected %v)\n", ms, name, fileName, result, expected)
	}
}

func SolveALocal(name string, solver func(string) interface{}, expected interface{}) {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	puzzlePath := filepath.Join(dir, name)
	solveFile(puzzlePath, solver)
	fmt.Printf("        Expected: %v\n", expected)
}

func JsonString(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	return string(bytes)
}

func JsonEquals(value1, value2 interface{}) bool {
	string1 := JsonString(value1)
	string2 := JsonString(value2)
	return string1 == string2
}
