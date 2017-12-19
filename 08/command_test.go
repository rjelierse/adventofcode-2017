package day08

import "testing"

func TestProgram_Run(t *testing.T) {
	var lines = []string{
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}

	p := &Program{lines: lines}
	_, value := p.Run()
	if value != 1 {
		t.Error("Expected largest register at 1 got", value)
	}
}
