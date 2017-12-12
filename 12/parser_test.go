package pipes

import "testing"

var lines = []string{
	"0 <-> 2",
	"1 <-> 1",
	"2 <-> 0, 3, 4",
	"3 <-> 2, 4",
	"4 <-> 2, 3, 6",
	"5 <-> 6",
	"6 <-> 4, 5",
}


func TestParseGroup(t *testing.T) {
	g, err := ParseGroup(lines[2])
	if err != nil {
		t.Error(err)
	}
	if g.pid != 2 {
		t.Error("Unexpected pid", g.pid)
	}
	if len(g.pipes) != 3 || g.pipes[0] != 0 || g.pipes[1] != 3 || g.pipes[2] != 4 {
		t.Error("Unexpected pipes", g.pipes)
	}
}

func TestParseTree(t *testing.T) {
	tree, err := ParseTree(lines)
	if err != nil {
		t.Error(err)
	}
	if len(tree.programs) != 7 {
		t.Error("Unknown program count", len(tree.programs))
	}
}

func TestTree_Size(t *testing.T) {
	tree, _ := ParseTree(lines)
	size := tree.Size(0)

	if size != 6 {
		t.Error("Unknown group size", size)
	}
}
