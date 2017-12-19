package day19

import "fmt"

type Position [2]int
type Direction [2]int

var (
	DirUp    = Direction{0, -1}
	DirDown  = Direction{0, 1}
	DirLeft  = Direction{-1, 0}
	DirRight = Direction{1, 0}
	DirNone  = Direction{0, 0}
)

const (
	NodeHorizontal = '-'
	NodeVertical   = '|'
	NodeTurn       = '+'
	NodeVoid       = ' '
)

func (p Position) Move(d Direction) Position {
	x, y := p[0], p[1]
	dx, dy := d[0], d[1]
	return Position{x + dx, y + dy}
}

func Travel(grid [][]byte) (sequence []byte, steps int) {
	sequence = make([]byte, 0)

	width := len(grid[0])
	height := len(grid)
	fmt.Printf("Working on a %dx%d grid.\n", width, height)

	get := func(pos Position) byte {
		x, y := pos[0], pos[1]
		return grid[y][x]
	}

	next := func(p Position, d Direction) Position {
		p = p.Move(d)
		if p[0] < 0 || p[0] > (width-1) {
			panic("x out of bounds")
		}
		if p[1] < 0 || p[1] > (height-1) {
			panic("y out of bounds")
		}
		return p
	}

	change := func(p Position, d Direction) Direction {
		if d != DirDown && get(p.Move(DirUp)) != NodeVoid {
			return DirUp
		}
		if d != DirUp && get(p.Move(DirDown)) != NodeVoid {
			return DirDown
		}
		if d != DirRight && get(p.Move(DirLeft)) != NodeVoid {
			return DirLeft
		}
		if d != DirLeft && get(p.Move(DirRight)) != NodeVoid {
			return DirRight
		}
		return DirNone
	}

	var position Position
	var direction Direction
	for i, b := range grid[0] {
		if b == NodeVertical {
			position = Position{i, 0}
			direction = DirDown
			steps++
		}
	}

	for {
		position = next(position, direction)
		node := get(position)

		switch node {
		case NodeVoid:
			return
		case NodeTurn:
			direction = change(position, direction)
		case NodeVertical:
			// Business as usual
		case NodeHorizontal:
			// Business as usual
		default:
			sequence = append(sequence, node)
		}

		steps++

		if direction == DirNone {
			return
		}
	}
}
