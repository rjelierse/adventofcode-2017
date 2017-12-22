package day22

type Position [2]int

type Virus struct {
	Direction  Direction
	Position   Position
	Infections int
}

func NewVirus() *Virus {
	v := new(Virus)
	v.Position = Position{0, 0}
	v.Direction = DirectionUp
	return v
}

func (v *Virus) TurnLeft() {
	v.Direction = v.Direction.TurnLeft()
}

func (v *Virus) TurnRight() {
	v.Direction = v.Direction.TurnRight()
}

func (v *Virus) Reverse() {
	v.Direction = v.Direction.Reverse()
}

func (v *Virus) Move() {
	x, y := v.Position[0], v.Position[1]
	switch v.Direction {
	case DirectionUp:
		v.Position = Position{x, y - 1}
	case DirectionLeft:
		v.Position = Position{x - 1, y}
	case DirectionDown:
		v.Position = Position{x, y + 1}
	case DirectionRight:
		v.Position = Position{x + 1, y}
	default:
		panic("unknown direction")
	}
}

func (v *Virus) AddInfection() {
	v.Infections++
}
