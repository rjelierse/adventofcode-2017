package grid

type Iterator struct {
	Position  Position
	Direction Direction
}

func (i *Iterator) TurnLeft() {
	i.Direction = i.Direction.TurnLeft()
}

func (i *Iterator) TurnRight() {
	i.Direction = i.Direction.TurnRight()
}

func (i *Iterator) Reverse() {
	i.Direction = i.Direction.Reverse()
}

func (i *Iterator) Move() {
	i.Position = i.Position.Move(i.Direction)
}
