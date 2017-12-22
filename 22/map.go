package day22

import (
	"bytes"
)

type Map struct {
	grid map[Position]State
}

type State int

const (
	StateClean State = iota
	StateInfected
	StateWeakened
	StateFlagged
)

func (m *Map) GetState(p Position) State {
	return m.grid[p]
}

func (m *Map) SetState(p Position, s State) {
	m.grid[p] = s
}

func NewMap(input []byte) *Map {
	m := new(Map)
	m.grid = make(map[Position]State)

	lines := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	height, width := len(lines), len(lines[0])

	xMin := (width - 1) / -2
	yMin := (height - 1) / -2

	for y, line := range lines {
		y = yMin + y
		for x, char := range line {
			x = xMin + x
			p := Position{x, y}
			switch char {
			case '.':
				m.grid[p] = StateClean
			case '#':
				m.grid[p] = StateInfected
			default:
				panic("unknown character")
			}
		}
	}
	return m
}
