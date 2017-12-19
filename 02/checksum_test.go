package day02

import "testing"

type expectation struct {
	input  []int
	output int
}

func TestDifference(t *testing.T) {
	expectations := []expectation{
		{[]int{5, 1, 9, 5}, 8},
		{[]int{7, 5, 3}, 4},
		{[]int{2, 4, 6, 8}, 6},
	}

	for _, test := range expectations {
		result := Difference(test.input)

		if test.output != result {
			t.Error("For", test.input, "expected", test.output, "got", result)
		}
	}
}

func TestDivision(t *testing.T) {
	expectations := []expectation{
		{[]int{5, 9, 2, 8}, 4},
		{[]int{9, 4, 7, 3}, 3},
		{[]int{3, 8, 6, 5}, 2},
	}

	for _, test := range expectations {
		result := Division(test.input)

		if test.output != result {
			t.Error("For", test.input, "expected", test.output, "got", result)
		}
	}
}
