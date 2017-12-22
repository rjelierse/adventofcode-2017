package day22

import "github.com/rjelierse/adventofcode-2017/grid"

type Virus struct {
	grid.Iterator
	Infections int
}

func NewVirus() *Virus {
	v := new(Virus)
	v.Position = grid.Position{0, 0}
	v.Direction = grid.DirectionUp
	return v
}

func (v *Virus) AddInfection() {
	v.Infections++
}
