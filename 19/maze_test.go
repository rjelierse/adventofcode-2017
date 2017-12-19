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
	result := Travel(maze)

	if string(result) != "ABCDEF" {
		t.Error("Unexpected result:", string(result))
	}
}