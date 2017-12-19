package day16

import "testing"

func TestParseSpinInstruction(t *testing.T) {
	input := "s1"
	instruction, err := ParseSpinInstruction(input)
	if err != nil {
		t.Error(err)
	} else if instruction.Count != 1 {
		t.Error("Unexpected spin count")
	}
}
