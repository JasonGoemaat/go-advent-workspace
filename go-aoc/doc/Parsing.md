## Parsing

These functions in the `aoc` package help with parsing:

### ParseGroups

```go
func ParseGroups(content string) []string
```

Used to split puzzle input with different sections separated by
empty lines, returns an array of strings representing each section.
For example if the input has an area and a set of moves:

    #####
    #...#
    #S.E#
    #####

    UP 2
    RIGHT 5
    DOWN 5

It would return:

```go
[]string{"#####\n#...#\n#S.E#\n#####","UP 2\nRIGHT 5\nDOWN 5"}
```

### ParseLines

This returns an array of strings for each line in the input:

```go
func ParseLines(content string) []string
```

### ParseInts

This finds any integers in the input and returns an array of them.

```go
func ParseInts(content string) []int 
```

An example of using this was 2024 day 22 where each line was a number,
it was just easier than splitting into lines and parsing each one as
an integer

### ParseIntsPerLine

This returns an array of int for each line consisting of the integers
on the line.

```go
func ParseIntsPerLine(content string) [][]int
```








The `aoc.Local` call handles loading the named file from the same directory
as the '.go' file and passes the text contents of the file as a string to
the solver function (`part1` and `part2` is what I named them above).

Using regular expressions to split strings in go is pretty easy.  As an
example, this is used in `aoc.ParseGroups()` to split into a string that
contains empty lines dividing groups of input:

    var reDoubleCRLF = regexp.MustCompile("[\r]?[\n][\r]?[\n]")
    var sections := reDoubleCRLF.Split(text)

The `aoc.ParseGroups()` function omits the last group if empty, I was thinking
there might be a blank line at the end of the file, but the last line doesn't
end in a newline so I don't think that's necessary.

See [parsing_test.go](https://github.com/JasonGoemaat/go-aoc/blob/main/aoc/parsing_test.go)
for some examples.
