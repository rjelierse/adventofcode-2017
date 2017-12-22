package day19

import "testing"

var maze = [][]byte{
	[]byte("     |          "),
	[]byte("     |  +--+    "),
	[]byte("     A  |  C    "),
	[]byte(" F---|----E|--+ "),
	[]byte("     |  |  |  D "),
	[]byte("     +B-+  +--+ "),
	[]byte("                "),
}

func TestTravel(t *testing.T) {
	sequence, steps := NewMaze(maze).Travel()

	if string(sequence) != "ABCDEF" {
		t.Error("Unexpected sequence:", string(sequence))
	}

	if steps != 38 {
		t.Error("Unexpected number of steps:", steps)
	}
}
