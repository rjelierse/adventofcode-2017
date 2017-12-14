package defrag

import "testing"

func TestGrid_Count(t *testing.T) {
	count := 8108
	g := NewGrid("flqrgnkx")
	result := g.Count()

	if result != count {
		t.Error("Expected", count, "set squares, got", result)
	}
}

func TestGrid_GroupCount(t *testing.T) {
	count := 1242
	g := NewGrid("flqrgnkx")
	result := g.GroupCount()

	if result != count {
		t.Error("Expected", count, "groups of squares, got", result)
	}
}

type bitTest struct {
	Input   int
	Outcome []int
}

func TestBitSet(t *testing.T) {
	tests := []bitTest{
		{0x0, []int{0, 0, 0, 0}},
		{0x1, []int{0, 0, 0, 1}},
		{0xe, []int{1, 1, 1, 0}},
		{0xf, []int{1, 1, 1, 1}},
	}

	for _, test := range tests {
		for i, b := range test.Outcome {
			result := isSquareUsed(test.Input, uint(i), byteSize)
			if result != b {
				t.Errorf("For input 0x%x, expected bit %d to be %d, got %d", test.Input, i, b, result)
			}
		}
	}
}
