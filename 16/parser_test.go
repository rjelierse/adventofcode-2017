package dance

import (
	"testing"
)

func TestSplitInstructions(t *testing.T) {
	input := "s1,x3/4,pe/b"
	output := []string{"s1", "x3/4", "pe/b"}

	result := SplitInstructions(input)
	if result[0] != output[0] || result[1] != output[1] || result[2] != output[2] {
		t.Error("Unexpected instructions sequence", result)
	}
}

func TestParseSpinInstruction(t *testing.T) {
	input := "s1"
	i := GetInstruction(input)
	if i != InstructionSpin {
		t.Error("Unexpected instruction", string(i))
	}
	count := ParseSpinInstruction(input)
	if count != 1 {
		t.Error("Unexpected spin count")
	}
}

func TestGetExchangeInstruction(t *testing.T) {
	input := "x3/4"
	i := GetInstruction(input)
	if i != InstructionExchange {
		t.Error("Unexpected instruction", string(i))
	}
	a, b := ParseExchangeInstruction(input)
	if a != 3 || b != 4 {
		t.Error("Unexpected positions", a, b)
	}
}

func TestGetPartnerInstruction(t *testing.T) {
	input := "pe/b"
	i := GetInstruction(input)
	if i != InstructionPartner {
		t.Error("Unexpected instruction", string(i))
	}
	a, b := ParsePartnerInstruction(input)
	if a != 'e' || b != 'b' {
		t.Error("Unexpected partners", string(a), string(b))
	}
}
