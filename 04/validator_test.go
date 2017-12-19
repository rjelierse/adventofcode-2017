package day04

import (
	"testing"
)

type expectation struct {
	phrase string
	result bool
}

var expectations = []expectation{
	{"aa bb cc dd ee", true},
	{"aa bb cc dd aa", false},
	{"aa bb cc dd aaa", true},
}

func TestIsValid(t *testing.T) {
	for _, test := range expectations {
		isValid := IsUnique(test.phrase)

		if isValid != test.result {
			t.Error("For", test.phrase, "expected", test.result, "got", isValid)
		}
	}
}
