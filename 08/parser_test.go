package registers

import "testing"

func TestParseLine(t *testing.T) {
	action, condition := ParseLine("b inc 5 if a > 1")

	if action.Register != "b" {
		t.Error("Expected register 'b' got", action.Register)
	}

	if action.Operator != OperatorIncrease {
		t.Error("Expected operator", OperatorIncrease, "got", action.Operator)
	}

	if action.Value != 5 {
		t.Error("Expected value 5 got", action.Value)
	}

	if condition.Register != "a" {
		t.Error("Expected register 'a' got", condition.Register)
	}

	if condition.Comparator != CompareGreater {
		t.Error("Expected comparator", CompareGreater, "go", condition.Comparator)
	}

	if condition.Value != 1 {
		t.Error("Expected value 1 got", condition.Value)
	}
}
