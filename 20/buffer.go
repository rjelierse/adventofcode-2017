package day20

import (
	"sort"
)

type Buffer struct {
	particles []*Particle
}

func BufferFromInput(input []string) (*Buffer, error) {
	b := new(Buffer)
	b.particles = make([]*Particle, len(input))
	for i, line := range input {
		p, err := ParticleFromInput(i, line)
		if err != nil {
			return nil, err
		}
		b.particles[i] = p
	}
	return b, nil
}

// Closest orders the particles slice by distance at time t and returns the particle closest to <0,0,0>
func (b *Buffer) Closest(t int) *Particle {
	sort.Slice(b.particles, func(i, j int) bool {
		return b.particles[i].CalcPosition(t).Sum() < b.particles[j].CalcPosition(t).Sum()
	})

	return b.particles[0]
}
