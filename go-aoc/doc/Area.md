## aoc.Area

A lot of puzzles involve some sort of map or 2d area.  There is a helper
structure `aoc.Area` that can help with this:

    ##########
    #S.......#
    #.######.#
    ##.......#
    ##.##.####
    #.......E#
    ##########

There is a parser function to create a structure which just contains the width
and height and the data as a byte array.   I thought this would be especially
efficient for cloning after making changes.

```go
func ParseArea(content string) *Area

type Area struct {
    Width  int
    Height int
    Data   []byte
}
```

Sometimes it is easier to operate with row and column and sometimes with the
index into the array.   These functions let you do things one or both ways
and convert between them.

```go
func (area Area) Get(row, col int) byte
func (area Area) GetIndex(index int) byte
func (area Area) IndexToRowCol(index int) (int, int)
func (area Area) RowColToIndex(row, col int) int
func (area Area) Inside(row, col int) bool
func (area Area) InsideIndex(index int) bool
func (area Area) Is(row, col int, b byte) bool
func (area Area) Clone() *Area
```