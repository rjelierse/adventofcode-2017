package tower

import "fmt"

type Program struct {
	Name          string
	SelfWeight    int
	Children      []string
	Parent        *Program
	ChildPrograms []*Program
}

// AddChild sets the references to the parent and child program on both nodes
func (p *Program) AddChild(c *Program) {
	p.ChildPrograms = append(p.ChildPrograms, c)

	c.Parent = p
}

func (p *Program) Weight() int {
	if len(p.ChildPrograms) == 0 {
		return p.SelfWeight
	}

	weight := 0
	weights := make(map[string]int)

	for _, c := range p.ChildPrograms {
		w := c.Weight()
		weights[c.Name] = w
		weight = weight + w
	}

	fmt.Printf("%s weighs %d (%v)\n", p.Name, p.SelfWeight, weights)

	return p.SelfWeight + weight
}
