package defrag

import "testing"

func TestGrid(t *testing.T) {
	key := "flqrgnkx"
	count := 8108
	g := NewGrid(key)
	result := g.Count()

	if result != count {
		t.Error("For key", key, "expected", count, "set bits, got", result)
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
