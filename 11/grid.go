package day11

type Grid struct {
	position map[string]int
}

// NewGrid creates a new grid of hexagons to iterate over
func NewGrid() *Grid {
	g := &Grid{}
	g.position = make(map[string]int)
	return g
}

// Move adjusts the position within the grid
func (g *Grid) Move(move string) {
	g.position[move] = g.position[move] + 1
	g.normalize()
}

// Distance calculates the distance from the center of the grid
func (g *Grid) Distance() int {
	d := 0

	for _, v := range g.position {
		d = d + v
	}

	return d
}

func (g *Grid) normalize() {
	// Normalize "nw" + "se" = 0
	g.adjustForDiff("nw", "se")
	// Normalize "ne" + "sw" = 0
	g.adjustForDiff("ne", "sw")

	// Normalize "ne" + "nw" = "n"
	g.adjustForSum("nw", "ne", "n")
	// Normalize "se" + "sw" = "s"
	g.adjustForSum("sw", "se", "s")

	// Normalize "n" + "s" = 0
	g.adjustForDiff("n", "s")

	// Normalize "ne" + "s" = "se"
	g.adjustForSum("ne", "s", "se")
	// Normalize "nw" + "s" = "sw"
	g.adjustForSum("nw", "s", "sw")
	// Normalize "se" + "n" = "ne"
	g.adjustForSum("se", "n", "ne")
	// Normalize "sw" + "n" = "nw"
	g.adjustForSum("sw", "n", "nw")
}

func (g *Grid) adjustForDiff(a string, b string) {
	if g.position[a] == 0 || g.position[b] == 0 {
		return
	}

	if g.position[a] < g.position[b] {
		g.position[b] = g.position[b] - g.position[a]
		g.position[a] = 0
	} else {
		g.position[a] = g.position[a] - g.position[b]
		g.position[b] = 0
	}
}

func (g *Grid) adjustForSum(a string, b string, c string) {
	if g.position[a] == 0 || g.position[b] == 0 {
		return
	}

	if g.position[a] < g.position[b] {
		g.position[c] = g.position[c] + g.position[a]
		g.position[b] = g.position[b] - g.position[a]
		g.position[a] = 0
	} else {
		g.position[c] = g.position[c] + g.position[b]
		g.position[a] = g.position[a] - g.position[b]
		g.position[b] = 0
	}
}
