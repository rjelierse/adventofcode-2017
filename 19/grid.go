package day19

type Grid [][]byte

func (g Grid) Start() Position {
	p := Position{0, 0}
	d := DirectionRight
	for g.Get(p).IsVoid() {
		p = p.Move(d)
	}
	return p
}

func (g Grid) Get(p Position) Node {
	return Node(g[p[1]][p[0]])
}

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) InGrid(p Position) bool {
	return p.WithinBounds(g.Width(), g.Height())
}
