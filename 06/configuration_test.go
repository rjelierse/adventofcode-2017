package day06

import (
	"testing"
)

func TestConfiguration_High(t *testing.T) {
	config := Configuration{0, 2, 7, 0}
	position, value := config.High()

	if position != 2 {
		t.Errorf("Expected position of highest value at 2, got %d instead.", position)
	}

	if value != 7 {
		t.Errorf("Expected highest value at 7, got %d instead.", value)
	}
}

func TestConfiguration_Is(t *testing.T) {
	a := Configuration{0, 2, 7, 0}
	b := Configuration{0, 2, 7, 0}
	c := Configuration{2, 4, 1, 2}
	d := Configuration{2, 4, 1}

	if !a.Is(b) {
		t.Errorf("Expected %v to equal %v.", a, b)
	}

	if a.Is(c) {
		t.Errorf("Expected %v not to equal %v.", a, c)
	}

	if c.Is(d) {
		t.Errorf("Expected %v not to equal %v.", c, d)
	}
}

func TestConfiguration_Clone(t *testing.T) {
	a := Configuration{0, 2, 7, 0}
	b := a.Clone()

	if !a.Is(b) {
		t.Errorf("Expected %v to equal %v.", a, b)
	}

	a[0] = 2
	if a[0] == b[0] {
		t.Errorf("Expected cloned value not to change.")
	}
}
