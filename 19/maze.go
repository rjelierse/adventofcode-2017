package day19

import "github.com/rjelierse/adventofcode-2017/grid"

type Maze struct {
	Grid   map[grid.Position]Node
	Width  int
	Height int
}

func (m *Maze) FindStartPosition() grid.Position {
	p := grid.Position{0, 0}
	for m.Get(p).IsVoid() {
		p = p.Move(grid.DirectionRight)
	}
	return p
}

func (m *Maze) Get(p grid.Position) Node {
	return m.Grid[p]
}

func (m *Maze) Travel() (sequence string, steps int) {
	i := &grid.Iterator{
		Position: m.FindStartPosition(),
		Direction: grid.DirectionDown,
	}
	peekLeft := func() Node {
		d := i.Direction.TurnLeft()
		p := i.Position.Move(d)
		return m.Get(p)
	}
	peekRight := func() Node {
		d := i.Direction.TurnRight()
		p := i.Position.Move(d)
		return m.Get(p)
	}
	steps = 1
	for {
		i.Move()
		if !m.InBounds(i) {
			return
		}

		n := m.Get(i.Position)
		switch {
		case n.IsVoid():
			return
		case n.IsAlpha():
			sequence += string(n)
		case n.IsTurn():
			switch {
			case peekLeft().IsNotVoid():
				i.TurnLeft()
			case peekRight().IsNotVoid():
				i.TurnRight()
			default:
				steps++
				return
			}
		}
		steps++
	}
	return
}

func (m *Maze) InBounds(i *grid.Iterator) bool {
	return i.Position.InBounds(0, m.Width - 1, 0, m.Height - 1)
}

func NewMaze(data [][]byte) *Maze {
	m := new(Maze)
	m.Width = len(data[0])
	m.Height = len(data)
	m.Grid = make(map[grid.Position]Node)
	for y, row := range data {
		for x, char := range row {
			m.Grid[grid.Position{x, y}] = Node(char)
		}
	}
	return m
}
