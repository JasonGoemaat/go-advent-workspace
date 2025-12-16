/*
From: https://codeberg.org/cdd/aoc/raw/branch/main/2025/go/10.go
Reddit: https://www.reddit.com/r/adventofcode/comments/1pity70/comment/nu6v5fq/
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"sync"
)

func parseIndicator(b []byte) uint32 {
	var (
		res uint32
		i   int
		c   byte
	)
	b = b[1 : len(b)-1]

	for i, c = range b {
		if c == '#' {
			res |= (1 << i)
		}
	}
	return res
}

func parseButton(b []byte) uint32 {
	var (
		res     uint32
		i, next int
	)
	b = b[1 : len(b)-1]

	next = bytes.IndexByte(b, ',')
	for next != -1 {
		res |= 1 << a2i(b[i:i+next])
		i += next + 1
		next = bytes.IndexByte(b[i:], ',')
	}

	res |= 1 << a2i(b[i:])

	return res
}

type light struct {
	a uint32
	b uint32
}

type machine struct {
	targ uint32
	nums []uint32
}

func worker(in chan machine, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		m    machine
		ok   bool
		c, n [1024]light
		seen = make(map[uint32]struct{}, 1024)
	)
	for m, ok = <-in; ok; m, ok = <-in {
		out <- buttonBfs(m.targ, m.nums, c[:len(m.nums)], n[:0], seen)
	}
}

func buttonBfs(targ uint32, nums []uint32, cur, next []light, seen map[uint32]struct{}) int {
	var (
		res, i int
		t      light
		x      uint32
		ok     bool
	)

	for i, x = range nums {
		if targ == x {
			return 1
		}
		cur[i].a = x
		cur[i].b |= 1 << i
	}

	clear(seen)

	for len(cur) > 0 {
		res++
		for _, t = range cur {
			for i, x = range nums {
				if t.b&(1<<i) != 0 {
					continue
				}
				x ^= t.a
				if _, ok = seen[x]; ok {
					continue
				}
				if x == targ {
					return res + 1
				}
				seen[x] = struct{}{}
				next = append(next, light{a: x, b: t.b & (1 << i)})
			}
		}
		cur, next = next, cur[:0]
	}

	return res
}

func parseButtonRaw(b []byte, l int) []int {
	var (
		res     = make([]int, l)
		i, next int
	)
	b = b[1 : len(b)-1]

	next = bytes.IndexByte(b, ',')
	for next != -1 {
		res[a2i(b[i:i+next])] = 1
		i += next + 1
		next = bytes.IndexByte(b[i:], ',')
	}
	res[a2i(b[i:])] = 1

	return res
}

func parseJoltage(b []byte) []int {
	var (
		nums  []int
		i     int
		split [][]byte
	)
	i = bytes.IndexByte(b, '{')
	split = bytes.Split(b[i+1:len(b)-1], []byte{','})
	nums = make([]int, len(split))
	for i = range nums {
		nums[i] = a2i(split[i])
	}
	return nums
}

func solution10A(r *bytes.Reader) int {
	var (
		res, i    int
		b         []byte
		pieces    [][]byte
		indicator uint32
		wg        sync.WaitGroup
		in        = make(chan machine, 128)
		out       = make(chan int, 256)
		scanner   = bufio.NewScanner(r)
		buttons   = make([]uint32, 0, 64)
	)

	for range 12 {
		wg.Add(1)
		go worker(in, out, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for scanner.Scan() {
		b = scanner.Bytes()
		if len(b) == 0 {
			continue
		}
		pieces = bytes.Split(b, []byte{' '})
		indicator = parseIndicator(pieces[0])
		buttons = make([]uint32, 0, 8)
		for i = 1; i < len(pieces)-1; i++ {
			buttons = append(buttons, parseButton(pieces[i]))
		}
		in <- machine{targ: indicator, nums: buttons}
		// not necessary for the input, but good practice to avoid
		// deadlocking on inputs with more than 256 lines
		select {
		case i = <-out:
			res += i
		default:
			continue
		}
	}
	close(in)

	for i = range out {
		res += i
	}

	return res
}

func transpose[T number](m [][]T) [][]T {
	r := make([][]T, len(m[0]))
	for i := range r {
		r[i] = make([]T, len(m))
	}
	for i := range m {
		for j := range r {
			r[j][i] = m[i][j]
		}
	}
	return r
}

func swapRows[T number](m [][]T, col int) [][]T {
	firstCol := firstNonzero(m[col])
	if firstCol == col {
		return m
	}
	// otherwise, search for a column
	maxCol := firstCol
	maxRow := col
	for i := col + 1; i < len(m); i++ {
		t := firstNonzero(m[i])
		if t == col {
			maxRow = i
			break
		} else if t < maxCol {
			maxCol = t
			maxRow = i
		}
	}
	if maxRow == len(m) {
		return m
	}
	m[col], m[maxRow] = m[maxRow], m[col]
	return m
}

func eliminate[T number](m [][]T, beg, idx int) {
	for i := beg + 1; i < len(m); i++ {

		if m[i][idx] == 0 {
			continue
		}
		ratio := m[i][idx] / m[beg][idx]
		for j := range m[i] {
			m[i][j] -= ratio * m[beg][j]
		}
	}
}

func firstNonzero[T number](src []T) int {
	var i int
	for i < len(src) && (src[i] == 0 || absg(float64(src[i])) < .0001) {
		i++
	}
	return i
}

func cloneMatrix[T any](m [][]T) [][]T {
	n := make([][]T, len(m))
	for row := range m {
		n[row] = make([]T, len(m[row]))
		copy(n[row], m[row])
	}
	return n
}

type number interface{ int | float64 }

func minimize[T number](m [][]T) [][]T {
	// remove rows with all zero or 1 nonzero and zero in the result
	n := make([][]T, 0, len(m))
	for row := range m {
		zc := 0
		// neg := -1
		for col := range m[row][:len(m[row])-1] {
			zc += b2i(m[row][col] != 0)
		}
		if zc == 0 || zc == 1 && m[row][len(m[row])-1] == 0 {
			continue
		}
		n = append(n, make([]T, len(m[row])))
		copy(n[len(n)-1], m[row])
	}
	return n
}

func rowReduce[T number](m [][]T) [][]T {
	for i := range len(m) {
		if i >= len(m[i]) {
			break
		}
		m = swapRows(m, i)
		k := firstNonzero(m[i])

		if k >= len(m[i])-1 {
			break
		}
		// fmt.Println(k, m[i][k])

		t := m[i][k]
		for j := 0; j < len(m[i]); j++ {
			m[i][j] /= t
		}
		eliminate(m, i, k)
	}

	for i := len(m) - 1; i > -1; i-- {
		k := firstNonzero(m[i])
		if k >= len(m[i])-1 {
			continue
		}
		t := m[i][k]

		for j := i - 1; j > -1; j-- {
			ratio := m[j][k] / t
			for h := k; h < len(m[i]); h++ {
				m[j][h] -= ratio * m[i][h]
			}
		}
	}

	return m
}

type mat []float64

func (m mat) String() string {
	bldr := new(strings.Builder)
	bldr.WriteByte('[')
	for i, n := range m {
		bldr.WriteString(fmt.Sprintf("%.04f", float64(n)))
		if i < len(m)-1 {
			bldr.WriteByte(' ')
		}
	}
	bldr.WriteByte(']')
	return bldr.String()
}

func pmatrix(m any) {
	var f [][]float64
	var ok bool
	if f, ok = m.([][]float64); !ok {
		return
	}
	fmt.Println()
	for row := range f {
		fmt.Println(mat(f[row]))
	}
	fmt.Println()
}

type m2 []float64

func (m m2) String() string {
	bldr := new(strings.Builder)
	for i, n := range m {
		bldr.WriteString(fmt.Sprintf("%.04f", float64(n)))
		if i < len(m)-1 {
			bldr.WriteByte('\t')
		}
	}
	return bldr.String()
}

func emptyOrFree[T number](m [][]T) ([]int, []int) {
	var empty, free []int
	for col := range m[0][:len(m[0])-1] {
		count := 0
		for row := range m {
			count += b2i(m[row][col] != 0)
		}
		if count == 0 {
			empty = append(empty, col)
		} else if count > 1 {
			free = append(free, col)
		}
	}
	return empty, free
}

func removeColumn[T number](m [][]T, col int) {
	for row := range m {
		copy(m[row][col:], m[row][col+1:])
		m[row] = m[row][:len(m[row])-1]
	}
}

func sumFinal[T number](m [][]T) T {
	var n T
	for row := range m {
		n += m[row][len(m[row])-1]
	}
	return n
}

func isSolved[T number](m [][]T) bool {
	var nz int
	for row := range m {
		nz = 0
		for col := range m[row][:len(m[row])-1] {
			if m[row][col] != 0 {
				nz++
			}
		}
		if nz > 1 {
			return false
		}
		f0 := float64(m[row][len(m[row])-1])
		f1 := float64(int(m[row][len(m[row])-1]))
		if f0 != f1 {
			// now check if the absolute difference is below a threshold
			if absg(f0-f1) > 0.1 {
				return false
			}
		}
	}
	// now make sure the final column is all positive integers;
	// we assume the other values are already positive 1
	/*
	 * x0 + x3 - x5 = 2
	 * x1 + x5 = 5
	 * x2 + x3 - x5 = 1
	 * x4 + x5 = 3
	 *
	 * x0 = 2 + x5 - x3
	 * x1 = 5 - x5
	 * x2 = 1 + x5 - x3
	 * x4 = 3 - x5
	 *
	 * x3 = 0
	 * x5 = 0
	 * x0 = 2
	 * x1 = 5
	 * x2 = 1
	 * x4 = 3 --> total == 11
	 *
	 * x3 = 0
	 * x5 = 1
	 * x0 = 3
	 * x1 = 4
	 * x2 = 2
	 * x4 = 1 --> total = 11
	 *
	 * x3 = 3
	 * x5 = 2
	 * x0 = 1
	 * x1 = 3
	 * x2 = 0
	 * x4 = 1 --> total = 10
	 *
	 *
	 * we could start at 0 and then say --> we stop going once we hit an invalid number
	 * so this would be a depth first search
	 */

	return true
}

func evaluateRow[T number](row []T, freeVals []fv[T]) T {
	var res T
	i := firstNonzero(row)
	if i == len(row) {
		return 0
	}
	// we're not going to worry for now about invalid rows
	// i.e., ones without solutions
	res = row[len(row)-1]
	for j := i + 1; j < len(row)-1; j++ {
		res -= row[j] * freeVals[j].val
	}
	res /= row[i]

	return res
}

type fv[T number] struct {
	val   T
	upper T
}

func paramDFS[T number](m [][]T, freeVals []fv[T], freeIndices []bool, i int) (T, bool) {
	var res, t T
	res = 512000
	var ok, okt bool
	if i == len(freeVals) {
		res = 0
		for row := range m {
			t = evaluateRow(m[row], freeVals)
			if t < 0 {
				return 512000, false
			}
			f0 := float64(t)
			f1 := float64(int(t))
			if f0 != f1 {
				// fmt.Println(f0, f1)
				if absg(f0-f1) > 0.1 {
					// fmt.Println(f0, f1)
					return 512000, false
				}
			}
			res += t
		}
		for _, v := range freeVals {
			res += v.val
		}
		return res, true
	}
	if !freeIndices[i] {
		return paramDFS(m, freeVals, freeIndices, i+1)
	}

	// otherwise, while we have a valid solution, increase... freeVals[i] and try again
	for p := T(0); p <= freeVals[i].upper; p++ {
		t, okt = paramDFS(m, freeVals, freeIndices, i+1)
		if okt {
			res = min(res, t)
			ok = true
		}
		freeVals[i].val++
	}
	// fmt.Println(freeVals, res)
	freeVals[i].val = 0
	return res, ok
}

func nzmul[T number](row []T) T {
	res := T(1)
	for _, c := range row {
		if c != 0 {
			res *= c
		}
	}
	return absg(res)
}

func gcd(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func lcm(a, b int) int {
	return abs(abs(a*b) / gcd(a, b))
}

func llcm[T number](row []T) T {
	i := 0
	var res int
	for ; i < len(row) && row[i] == 0; i++ {
	}
	res = int(row[i])
	i++
	for ; i < len(row); i++ {
		if row[i] == 0 {
			continue
		}
		res = lcm(res, int(row[i]))
	}
	return T(res)
}

func upperBound[T number](row []T, cur T) T {
	// if all values in T are positive, we can set an upper bound of the final variable
	all := true
	for i := range row {
		if row[i] < 0 {
			all = false
			break
		}
	}
	if all {
		if cur == 0 {
			return row[len(row)-1]
		}
		return min(cur, row[len(row)-1])
	}
	if cur == 0 {
		return 4000
	}
	return cur
}

func solveFree[T number](m [][]T, cols []int) T {
	// make sure each row in the matrix is good

	for row := range m {
		for {
			t := getSmallestDecimal(m[row])
			if t == 1 || t == 0 {
				break
			}
			t = 1 / t
			for col := range m[row] {
				m[row][col] *= t
			}
		}
	}

	freeVals := make([]fv[T], len(m[0]))

	freeCols := make([]bool, len(m[0]))
	for col := range cols {
		freeCols[cols[col]] = true
		// find upper bound
		for row := range m {
			if m[row][cols[col]] != 0 {
				freeVals[cols[col]].upper = upperBound(m[row], freeVals[cols[col]].upper)
			}
		}
	}

	res, ok := paramDFS(m, freeVals, freeCols, 0)
	if !ok {
		panic(ok)
	}

	return res
}

func getSmallestDecimal[T number](row []T) T {
	res := T(1)

	for i := range row {
		f0 := float64(row[i])
		f1 := float64(int(row[i]))
		if f0 == f1 {
			continue
		}
		if absg(f0-f1) < 0.00001 {
			row[i] = T(f1)
			// fmt.Println(f0, f1)
			continue
		}
		res = min(res, absg(T(f0-f1)))
	}

	return res
}

func Solution10B(line string) int {
	return solution10B(line)
}

func solution10B(line string) int {
	r := bytes.NewReader([]byte(line))

	var (
		res, i int
		b      []byte
		pieces [][]byte
		// indicator uint32
		scanner  = bufio.NewScanner(r)
		buttons  = make([][]int, 0, 64)
		joltages []int
	)

	for scanner.Scan() {
		b = scanner.Bytes()
		if len(b) == 0 {
			continue
		}
		pieces = bytes.Split(b, []byte{' '})

		buttons = make([][]int, 0, 12)
		for i = 1; i < len(pieces)-1; i++ {
			buttons = append(buttons, parseButtonRaw(pieces[i], len(pieces[0])-2))
		}
		joltages = parseJoltage(pieces[len(pieces)-1])

		buttons = append(buttons, joltages)
		b2 := make([][]float64, len(buttons))
		for row := range buttons {
			b2[row] = make([]float64, len(buttons[row]))
			for col := range buttons[row] {
				b2[row][col] = float64(buttons[row][col])
			}
		}
		b2 = transpose(b2)
		rowReduce(b2)
		b2 = minimize(b2)
		_, free := emptyOrFree(b2)
		if len(free) > 0 {
			ans := solveFree(b2, free)
			res += int(ans)
		} else {
			res += int(sumFinal(b2))
		}

	}

	return res
}
