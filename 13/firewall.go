package firewall

type Firewall struct {
	layers map[int]int
	depth  int
}

func NewFirewall() *Firewall {
	f := Firewall{}
	f.layers = make(map[int]int)
	return &f
}

func (f *Firewall) AddLayer(depth int, rng int) {
	f.layers[depth] = rng

	if depth > f.depth {
		f.depth = depth
	}
}

func (f *Firewall) Traverse(delay int) int {
	t := delay
	penalty := 0

	for p := 0; p <= f.depth; p++ {
		r, exists := f.layers[p]

		if exists && Position(t, r) == 0 {
			penalty = penalty + (t * r)
		}

		t++
	}

	return penalty
}
