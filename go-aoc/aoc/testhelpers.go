package aoc

import (
	"encoding/json"
	"fmt"
	"testing"
)

// nope
// type AocTB testing.TB

type AocTB interface {
	testing.TB
}

// type MyRouter mux.Router
// func (m *MyRouter) F() { ... }

// ExpectJson will convert the expected and actual arguments to JSON, or by
// using fmt.Sprintf("%v", xxx) if that fails for some reason like there is
// a cycle.   If they are different, it will fail the text and log both
// values.
func ExpectJson(t *testing.T, expected, actual interface{}) {
	t.Helper()
	actualString := ""
	expectedString := ""

	bytes, err := json.Marshal(actual)
	if err != nil {
		actualString = fmt.Sprintf("%v", actual)
	} else {
		actualString = string(bytes)
	}

	bytes, err = json.Marshal(expected)
	if err != nil {
		expectedString = fmt.Sprintf("%v", expected)
	} else {
		expectedString = string(bytes)
	}

	if actualString != expectedString {
		t.Logf("Expected: %s", expectedString)
		t.Logf("  Actual: %s", actualString)
		t.Fail()
	}
}
