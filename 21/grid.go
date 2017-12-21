package day21

import "bytes"

const PatternSeparator = "/"

type Grid [][]byte

func (src Grid) Rotate() Grid {
	size := len(src)
	dst := NewGrid(size)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dst[y][x] = src[size-x-1][y]
		}
	}
	return dst
}

func (src Grid) Flip() Grid {
	size := len(src)
	dst := NewGrid(size)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dst[y][x] = src[y][size-x-1]
		}
	}
	return dst
}

func (src Grid) ToPattern() Pattern {
	return Pattern(bytes.Join(src, []byte(PatternSeparator)))
}

func NewGrid(size int) Grid {
	g := make([][]byte, size)
	for i := 0; i < size; i++ {
		g[i] = make([]byte, size)
	}
	return g
}

func NewGridFromPattern(p Pattern) Grid {
	return bytes.Split([]byte(p), []byte(PatternSeparator))
}
