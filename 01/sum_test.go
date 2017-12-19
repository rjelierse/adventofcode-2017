package day01

import "testing"

type expectation struct {
	input  []byte
	output int
}

var expectOffsetNext = []expectation{
	{[]byte{'1', '1', '2', '2'}, 3},
	{[]byte{'1', '1', '1', '1'}, 4},
	{[]byte{'1', '2', '3', '4'}, 0},
	{[]byte{'9', '1', '2', '1', '2', '1', '2', '9'}, 9},
}

var expectOffsetSplit = []expectation{
	{[]byte{'1', '2', '1', '2'}, 6},
	{[]byte{'1', '2', '2', '1'}, 0},
	{[]byte{'1', '2', '3', '4', '2', '5'}, 4},
	{[]byte{'1', '2', '3', '1', '2', '3'}, 12},
	{[]byte{'1', '2', '1', '3', '1', '4', '1', '5'}, 4},
}

func TestOffsetNext(t *testing.T) {
	for _, expect := range expectOffsetNext {
		sum := calculateSum(expect.input, 1)

		if sum != expect.output {
			t.Error("For", expect.input, "expected", expect.output, "got", sum)
		}
	}
}

func TestOffsetSplit(t *testing.T) {
	for _, expect := range expectOffsetSplit {
		sum := calculateSum(expect.input, len(expect.input)/2)

		if sum != expect.output {
			t.Error("For", expect.input, "expected", expect.output, "got", sum)
		}
	}
}
