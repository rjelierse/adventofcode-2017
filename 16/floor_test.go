package day16

import "testing"

func TestNewDancefloor(t *testing.T) {
	d := NewFloor(5)
	if d.positions[0] != 'a' || d.positions[1] != 'b' || d.positions[2] != 'c' || d.positions[3] != 'd' || d.positions[4] != 'e' {
		t.Error("Unexpected dancefloor positioning", string(d.positions))
	}
}

func TestDancefloor_Spin(t *testing.T) {
	d := NewFloor(5)
	d.Spin(1)
	if d.positions[0] != 'e' || d.positions[1] != 'a' || d.positions[2] != 'b' || d.positions[3] != 'c' || d.positions[4] != 'd' {
		t.Error("Unexpected dancefloor positioning", string(d.positions))
	}
}

func TestDancefloor_Exchange(t *testing.T) {
	d := NewFloor(5)
	d.Exchange(3, 4)
	if d.positions[0] != 'a' || d.positions[1] != 'b' || d.positions[2] != 'c' || d.positions[3] != 'e' || d.positions[4] != 'd' {
		t.Error("Unexpected dancefloor positioning", string(d.positions))
	}
}

func TestDancefloor_Partner(t *testing.T) {
	d := NewFloor(5)
	d.Partner('e', 'b')
	if d.positions[0] != 'a' || d.positions[1] != 'e' || d.positions[2] != 'c' || d.positions[3] != 'd' || d.positions[4] != 'b' {
		t.Error("Unexpected dancefloor positioning", string(d.positions))
	}
}
