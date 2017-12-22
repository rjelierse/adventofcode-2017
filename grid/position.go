package grid

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

func (p Position) InBounds(xMin, xMax, yMin, yMax int) bool {
	return xMin <= p[0] && p[0] <= xMax && yMin <= p[1] && p[1] <= yMax
}
