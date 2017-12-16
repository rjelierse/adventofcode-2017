package dance

import "testing"

func TestGetExchangeInstruction(t *testing.T) {
	input := "x3/4"
	instruction, err := ParseExchangeInstruction(input)
	if err != nil {
		t.Error(err)
	} else if instruction.PositionA != 3 || instruction.PositionB != 4 {
		t.Error("Unexpected positions", instruction.PositionA, instruction.PositionB)
	}
}
