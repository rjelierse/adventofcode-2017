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

// Sync updates the particle positions and returns the particle closest to <0,0,0>
func (b *Buffer) Sync() *Particle {
	for _, p := range b.particles {
		p.Update()
	}

	sort.Slice(b.particles, func(i, j int) bool {
		return b.particles[i].Distance() < b.particles[j].Distance()
	})

	return b.particles[0]
}
