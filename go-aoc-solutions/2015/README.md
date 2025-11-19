# Advent of Code 2015

https://adventofcode.com/2015

## Day 11

I started to program something, then realized doing it manually was easiest
after looking at the samples, see comments in the code.  All criteria can
be satisfied using just the last 5 letters as one of the samples given has
`aabcc` at the end which is the lowest way to update the last 5 characters.
We do need to worry because if the first 4 are a double and straight for
example we only need to look at the last 2.  My code probably isn't
perfect, but works for now.

## Day 12

Seems super-simple.   No strings containing numbers, so just use regex to
find them?  Yep, used `aoc.ParseInts()` and summed the int array.

Part 2 ignores objects with `"red"` properties, I think I can regexp replace
those out.  Hmm..   Can't figure out regexp for it.   Looking at GO's JSON
parsing I figured this trick out, I should be able to create a recursive
function to return the same object, or nil if it's an object with a property
that has the string value 'red':

```go
var data interface{} // try fully generic - Works for both, either get []interface{} for an array or map[string]interface{} for object
contentBytes := []byte(contents)
_ := json.Unmarshal(contentBytes, &data)
switch t := data.(type) {
case map[string]interface{}:
    fmt.Printf("It's an object!\n")
case []interface{}:
    fmt.Printf("It's an array!\n")
default:
    fmt.Printf("It's an unknown type!  (%T)\n", t)
}
```

See [12/README.md](12/README.md) for more info, it was actually easy to do...
