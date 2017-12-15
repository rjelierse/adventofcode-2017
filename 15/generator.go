package generator

type Generator struct {
	Value      int
	Multiplier int
}

const ValueMax = 2147483647

func NewGenerator(start int, multiplier int) *Generator {
	return &Generator{start, multiplier}
}

func (g *Generator) Next() int {
	g.Value = (g.Value * g.Multiplier) % ValueMax
	return g.Value
}
