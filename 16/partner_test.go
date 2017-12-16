package dance

import "testing"

func TestGetPartnerInstruction(t *testing.T) {
	input := "pe/b"
	instruction := ParsePartnerInstruction(input)
	if instruction.PartnerA != 'e' || instruction.PartnerB != 'b' {
		t.Error("Unexpected partners", string(instruction.PartnerA), string(instruction.PartnerB))
	}
}
