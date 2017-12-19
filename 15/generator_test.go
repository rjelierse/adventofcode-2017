package day15

import "testing"

func TestGenerator_Next(t *testing.T) {
	g := NewGenerator(startA, multiplierA, 4)
	values := []int{
		1092455,
		1181022009,
		245556042,
		1744312007,
		1352636452,
	}
	for _, v := range values {
		r := g.Next(false)
		if v != r {
			t.Error("Expected", v, "got", r)
		}
	}
}

func TestGenerator_Next1(t *testing.T) {
	g := NewGenerator(startA, multiplierA, 4)
	values := []int{
		1352636452,
		1992081072,
		530830436,
		1980017072,
		740335192,
	}
	for _, v := range values {
		r := g.Next(true)
		if v != r {
			t.Error("Expected", v, "got", r)
		}
	}
}
