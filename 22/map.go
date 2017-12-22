package day22

import (
	"bytes"
	"github.com/rjelierse/adventofcode-2017/grid"
)

type Map struct {
	Grid map[grid.Position]State
}

type State int

const (
	StateClean State = iota
	StateInfected
	StateWeakened
	StateFlagged
)

func (m *Map) GetState(p grid.Position) State {
	return m.Grid[p]
}

func (m *Map) SetState(p grid.Position, s State) {
	m.Grid[p] = s
}

func NewMap(input []byte) *Map {
	m := new(Map)
	m.Grid = make(map[grid.Position]State)

	lines := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	height, width := len(lines), len(lines[0])

	xMin := (width - 1) / -2
	yMin := (height - 1) / -2

	for y, line := range lines {
		y = yMin + y
		for x, char := range line {
			x = xMin + x
			p := grid.Position{x, y}
			switch char {
			case '.':
				m.Grid[p] = StateClean
			case '#':
				m.Grid[p] = StateInfected
			default:
				panic("unknown character")
			}
		}
	}
	return m
}
