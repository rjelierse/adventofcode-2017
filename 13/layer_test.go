package day13

import (
	"testing"
)

type test struct {
	Range     int
	Positions []int
}

var tests = []test{
	{1, []int{0, 0, 0, 0, 0, 0, 0}},
	{2, []int{0, 1, 0, 1, 0, 1, 0}},
	{3, []int{0, 1, 2, 1, 0, 1, 2}},
	{4, []int{0, 1, 2, 3, 2, 1, 0}},
}

func TestPosition(t *testing.T) {
	for _, s := range tests {
		for time, pos := range s.Positions {
			result := Position(time, s.Range)
			if result != pos {
				t.Error("For range", s.Range, "at time", time, "expected position at", pos, "got", result)
			}
		}
	}
}
