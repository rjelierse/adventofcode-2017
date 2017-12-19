package day19

type Direction int

const (
	DirectionNone Direction = iota
	DirectionUp
	DirectionDown
	DirectionLeft
	DirectionRight
)

func (d Direction) IsHorizontal() bool {
	return d == DirectionLeft || d == DirectionRight
}

func (d Direction) IsVertical() bool {
	return d == DirectionUp || d == DirectionDown
}

func (d Direction) Change(g Grid, p Position) Direction {
	switch {
	case d.IsHorizontal():
		switch {
		case g.Get(p.Move(DirectionUp)).IsNotVoid():
			return DirectionUp
		case g.Get(p.Move(DirectionDown)).IsNotVoid():
			return DirectionDown
		}
	case d.IsVertical():
		switch {
		case g.Get(p.Move(DirectionLeft)).IsNotVoid():
			return DirectionLeft
		case g.Get(p.Move(DirectionRight)).IsNotVoid():
			return DirectionRight
		}
	}
	return DirectionNone
}
