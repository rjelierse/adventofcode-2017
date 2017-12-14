package defrag

import (
	"fmt"
	"github.com/rjelierse/adventofcode-2017/10"
)

const gridSize = 128
const byteSize = 4
const divider = 16

const (
	squareFree = iota
	squareUsed
	squareVisited
)

type Grid struct {
	bits [gridSize][gridSize]int
}

func NewGrid(key string) *Grid {
	g := new(Grid)

	for i := 0; i < gridSize; i++ {
		input := fmt.Sprintf("%s-%d", key, i)
		lengths := knot.Length([]byte(input))
		for j, bits := range knot.Hash(lengths) {
			upper := bits / divider
			lower := bits % divider

			var b uint
			for b = 0; b < byteSize; b++ {
				upperIndex := 8*j + int(b)
				lowerIndex := 8*j + int(b) + 4
				g.bits[i][upperIndex] = isSquareUsed(upper, b, byteSize)
				g.bits[i][lowerIndex] = isSquareUsed(lower, b, byteSize)
			}
		}
	}

	return g
}

func (g *Grid) Count() int {
	count := 0

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if g.bits[i][j] == squareUsed {
				count++
			}
		}
	}

	return count
}

func (g *Grid) GroupCount() int {
	groups := 0

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if g.bits[i][j] == squareUsed {
				groups++
				g.fill(i, j)
			}
		}
	}

	return groups
}

func (g *Grid) fill(i int, j int) {
	if g.bits[i][j] != squareUsed {
		return
	}

	g.bits[i][j] = squareVisited

	if i > 0 {
		g.fill(i-1, j)
	}

	if j > 0 {
		g.fill(i, j-1)
	}

	if i < gridSize-1 {
		g.fill(i+1, j)
	}

	if j < gridSize-1 {
		g.fill(i, j+1)
	}
}

func isSquareUsed(value int, bit uint, bits uint) int {
	lsb := bits - bit - 1
	outcome := value & (1 << lsb)

	if outcome > 0 {
		return squareUsed
	} else {
		return squareFree
	}
}
