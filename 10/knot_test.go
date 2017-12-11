package knot

import "testing"

func TestKnot_Tie(t *testing.T) {
	k := NewKnot(5)
	k.Transform(3)

	if k.currentPosition != 3 {
		t.Error("Invalid position", k.currentPosition)
	}

	if k.skip != 1 {
		t.Error("Invalid skip value", k.skip)
	}

	if k.sparseHash[0] != 2 || k.sparseHash[1] != 1 || k.sparseHash[2] != 0 {
		t.Error("Invalid hash transformation", k.sparseHash)
	}
}

func TestKnot_Hash(t *testing.T) {
	lengths := []int{3, 4, 1, 5}
	k := NewKnot(5)
	k.Round(lengths)
	s := k.Checksum()

	if s != 12 {
		t.Error("Invalid checksum", s)
	}
}

func TestXorAll(t *testing.T) {
	input := []int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}
	output := XorAll(input)

	if output != 64 {
		t.Error("For", input, "expected 64 got", output)
	}
}