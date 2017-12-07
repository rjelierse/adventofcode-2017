package tower

import "testing"

type expectation struct {
	Line    string
	Program Program
}

var inputs = []expectation{
	{"pbga (66)", Program{Name: "pbga", SelfWeight: 66}},
	{"fwft (72) -> ktlj, cntj, xhth", Program{Name: "fwft", SelfWeight: 72, Children: []string{"ktlj", "cntj", "xhth"}}},
}

func TestParseLine(t *testing.T) {
	for _, test := range inputs {
		program, err := ParseLine(test.Line)
		if err != nil {
			t.Error("Parsing failed", err)
		}

		if program.Name != test.Program.Name {
			t.Errorf("Expected program name %s, got %s instead.", test.Program.Name, program.Name)
		}

		if program.SelfWeight != test.Program.SelfWeight {
			t.Errorf("Expected program weight %d, got %d instead.", test.Program.SelfWeight, program.SelfWeight)
		}

		if len(program.Children) != len(test.Program.Children) {
			t.Errorf("Expected program to have %d children, got %d instead.", len(test.Program.Children), len(program.Children))
		}
	}
}
