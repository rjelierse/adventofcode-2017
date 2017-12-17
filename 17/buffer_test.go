package spinlock

import "testing"

func TestNewRing(t *testing.T) {
	r := NewRing()
	if len(r.buffer) != 1 {
		t.Error("Unexpected buffer length", len(r.buffer))
	}
	if r.buffer[0] != 0 {
		t.Error("Unexpected buffer contents", r.buffer)
	}
	if r.position != 0 {
		t.Error("Unexpected buffer position", r.position)
	}
	if r.value != 0 {
		t.Error("Unexpected next value", r.value)
	}
}

func TestRing_Run(t *testing.T) {
	r := NewRing()
	r.Run(3, 2017)
	if len(r.buffer) != 2018 {
		t.Error("Unexpected buffer length", len(r.buffer))
	}
	if r.buffer[r.position] != 2017 {
		t.Error("Unexpected buffer value", r.buffer[r.position])
	}
	v := r.ValueAfter(2017)
	if v != 638 {
		t.Error("Unexpected next value", v)
	}
}

func TestRing_Step(t *testing.T) {
	l := 3
	r := NewRing()
	r.Step(l)
	if r.buffer[0] != 0 || r.buffer[1] != 1 {
		t.Error("Unexpected buffer content after first step", r.buffer)
	}
	r.Step(l)
	if r.buffer[0] != 0 || r.buffer[1] != 2 || r.buffer[2] != 1 {
		t.Error("Unexpected buffer content after second step", r.buffer)
	}
}

func TestRing_nextPosition(t *testing.T) {
	r := NewRing()
	p := r.nextPosition(3)
	if p != 1 {
		t.Error("Expected next position at 1, got", p)
	}
}

func TestRing_insertAfter(t *testing.T) {
	r := NewRing()
	r.insertAt(1, 1)
	if len(r.buffer) != 2 {
		t.Error("Unexpected buffer length", len(r.buffer))
	}
	if r.buffer[0] != 0 || r.buffer[1] != 1 {
		t.Error("Unexpected buffer contents", r.buffer)
	}
}
