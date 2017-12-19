package day07

import "testing"

var lines = []string{
	"pbga (66)",
	"xhth (57)",
	"ebii (61)",
	"havc (66)",
	"ktlj (57)",
	"fwft (72) -> ktlj, cntj, xhth",
	"qoyq (66)",
	"padx (45) -> pbga, havc, qoyq",
	"tknk (41) -> ugml, padx, fwft",
	"jptl (61)",
	"ugml (68) -> gyxo, ebii, jptl",
	"gyxo (61)",
	"cntj (57)",
}

func TestProgram(t *testing.T) {
	program := BuildTree(lines)
	if program == nil {
		t.Error("Cannot find parent node")
	}

	if program.Name != "tknk" {
		t.Errorf("Expected root node to have name tknk, got %d instead.", program.Name)
	}
}
