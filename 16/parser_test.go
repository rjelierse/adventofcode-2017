package dance

import (
	"testing"
)

func TestSplitInstructions(t *testing.T) {
	input := "s1,x3/4,pe/b"
	output := []string{"s1", "x3/4", "pe/b"}

	result := splitInstructions(input)
	if result[0] != output[0] || result[1] != output[1] || result[2] != output[2] {
		t.Error("Unexpected instructions sequence", result)
	}
}
