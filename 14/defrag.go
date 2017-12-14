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
)

type Grid struct {
	bits [gridSize][gridSize]int
}

func NewGrid(key string) *Grid {
	g := new(Grid)

	for i := 0; i < gridSize; i++ {
		lengths := KeyToLengths(key, i)
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

func KeyToLengths(key string, row int) []int {
	input := fmt.Sprintf("%s-%d", key, row)
	lengths := make([]int, len(input))
	for i, b := range input {
		lengths[i] = int(b)
	}
	return lengths
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
