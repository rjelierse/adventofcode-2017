package day11

import "testing"

type test struct {
	moves    []string
	distance int
}

var tests = []test{
	{[]string{"ne", "ne", "ne"}, 3},
	{[]string{"ne", "ne", "sw", "sw"}, 0},
	{[]string{"ne", "ne", "s", "s"}, 2},
	{[]string{"se", "sw", "se", "sw", "sw"}, 3},
}

func TestNewGrid(t *testing.T) {
	for _, s := range tests {
		g := NewGrid()
		for _, move := range s.moves {
			g.Move(move)
		}
		d := g.Distance()

		if s.distance != d {
			t.Error("For", s.moves, "expected distance", s.distance, "got", d)
		}
	}
}
