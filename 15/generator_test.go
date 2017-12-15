package generator

import "testing"

func TestGenerator_Next(t *testing.T) {
	g := NewGenerator(startA, multiplierA)
	values := []int{
		1092455,
		1181022009,
		245556042,
		1744312007,
		1352636452,
	}
	for _, v := range values {
		r := g.Next()
		if v != r {
			t.Error("Expected", v, "got", r)
		}
	}
}
