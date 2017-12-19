package day19

type Position [2]int

func (p Position) Move(d Direction) Position {
	x, y := p[0], p[1]
	switch d {
	case DirectionUp:
		return Position{x, y - 1}
	case DirectionDown:
		return Position{x, y + 1}
	case DirectionLeft:
		return Position{x - 1, y}
	case DirectionRight:
		return Position{x + 1, y}
	default:
		panic("unknown direction")
	}
}

func (p Position) WithinBounds(width, height int) bool {
	switch {
	case p[0] < 0 || p[0] > (width-1):
		return false
	case p[1] < 0 || p[1] > (height-1):
		return false
	default:
		return true
	}
}
