# General Library

## aoc.LogF()

You can enable logging and use the LogF function when investigating or
debugging a puzzle solution.  Then you can leave the statements in and
just don't enable logging and it won't output

``` go
var LoggingEnabled = false
func LogF(format string, args ...interface{})
```

## aoc.Local()

This is the simple function called for each day's solvers in their `main` func.
It lets you run a sample/input through your function, loading the file and calling
the solver function and reporting the results nicely on one line along with the
expected output

```go
func Local(solver func(string) interface{}, name string, fileName string, expected interface{})
```

## JSON

These functions are quick and dirty for converting GO structs or values into
JSON and comparing them.

```go
func JsonString(value interface{}) string
func JsonEquals(value1, value2 interface{}) bool
```

## Path

These functions let you get the path relative to the `.go` file calling them,
and to ensure a subdirectory of that relative path exists.   This can help
if you want to create output, or read separate input.

```go
func GetSubPath(subPath string) string
func GetOrCreateSubDirectory(subPath string) string
```

## aoc.ExpectJson()

This function gives an easy way to write tests and compare complex results.
The expected and actual can be any type and the test will fail if their
JSON representations are not identical.

```go
func ExpectJson(t *testing.T, expected, actual interface{})
```

An example is in [parsing_test.go](../aoc/parsing_test.go) where I want to
ensure an array of strings is correctly parsed into a two-dimensional
array of integers.   The package name isn't needed on this call because the
test is in the same package, but you'd be calling like `aoc.ExpectJson()`.

```go
func TestParseLinesToInts(t *testing.T) {
	lines := []string{"13|7", "8    5 3   12", "47,999,1234,9999"}
	expected := [][]int{{13, 7}, {8, 5, 3, 12}, {47, 999, 1234, 9999}}
	actual := ParseLinesToInts(lines)
	ExpectJson(t, expected, actual)
}
```
