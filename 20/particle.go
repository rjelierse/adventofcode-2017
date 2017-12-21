package day20

import (
	"fmt"
)

type Particle struct {
	Id           int
	Position     Vector
	Velocity     Vector
	Acceleration Vector
}

func ParticleFromInput(id int, input string) (*Particle, error) {
	var p, v, a Vector
	_, err := fmt.Sscanf(input, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &p.X, &p.Y, &p.Z, &v.X, &v.Y, &v.Z, &a.X, &a.Y, &a.Z)
	if err != nil {
		return nil, err
	}

	return &Particle{id, p, v, a}, nil
}

func (p *Particle) Distance() int {
	return p.Position.Sum()
}

func (p *Particle) Update() {
	p.Velocity = p.Velocity.Add(p.Acceleration)
	p.Position = p.Position.Add(p.Velocity)
}
