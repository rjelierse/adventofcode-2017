package generator

type Generator struct {
	Value      int
	Multiplier int
	Divider    int
}

const ValueMax = 2147483647

func NewGenerator(start int, multiplier int, divider int) *Generator {
	return &Generator{start, multiplier, divider}
}

func (g *Generator) Next(shouldDivide bool) int {
	v := g.nextValue()

	if shouldDivide {
		for v % g.Divider != 0 {
			v = g.nextValue()
		}
	}

	return v
}

func (g *Generator) nextValue() int {
	g.Value = (g.Value * g.Multiplier) % ValueMax
	return g.Value
}
