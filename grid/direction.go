package grid

type Direction int

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

func (d Direction) TurnLeft() Direction {
	switch d {
	case DirectionUp:
		return DirectionLeft
	case DirectionLeft:
		return DirectionDown
	case DirectionDown:
		return DirectionRight
	case DirectionRight:
		return DirectionUp
	default:
		panic("unknown direction")
	}
}

func (d Direction) TurnRight() Direction {
	switch d {
	case DirectionUp:
		return DirectionRight
	case DirectionLeft:
		return DirectionUp
	case DirectionDown:
		return DirectionLeft
	case DirectionRight:
		return DirectionDown
	default:
		panic("unknown direction")
	}
}

func (d Direction) Reverse() Direction {
	switch d {
	case DirectionUp:
		return DirectionDown
	case DirectionLeft:
		return DirectionRight
	case DirectionDown:
		return DirectionUp
	case DirectionRight:
		return DirectionLeft
	default:
		panic("unknown direction")
	}
}
