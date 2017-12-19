package day03

import (
	"math"
	"testing"
)

type expectation struct {
	input  int
	result int
}

var scenarios = []expectation{
	{1, 0},
	{12, 3},
	{23, 2},
	{1024, 31},
	{277678, 12},
}

func TestPosition(t *testing.T) {
	for _, scenario := range scenarios {
		x, y := position(scenario.input)
		carry := math.Abs(float64(x)) + math.Abs(float64(y))

		if int(carry) != scenario.result {
			t.Error("For", scenario.input, "expected", scenario.result, "got", carry)
		}
	}
}
