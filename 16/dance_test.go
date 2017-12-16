package dance

import "testing"

func TestNewDancefloor(t *testing.T) {
	d := NewDancefloor(5)
	if d.Positions[0] != 'a' || d.Positions[1] != 'b' || d.Positions[2] != 'c' || d.Positions[3] != 'd' || d.Positions[4] != 'e' {
		t.Error("Unexpected dancefloor positioning", string(d.Positions))
	}
}

func TestDancefloor_Spin(t *testing.T) {
	d := NewDancefloor(5)
	d.Spin(1)
	if d.Positions[0] != 'e' || d.Positions[1] != 'a' || d.Positions[2] != 'b' || d.Positions[3] != 'c' || d.Positions[4] != 'd' {
		t.Error("Unexpected dancefloor positioning", string(d.Positions))
	}
}

func TestDancefloor_Exchange(t *testing.T) {
	d := NewDancefloor(5)
	d.Exchange(3, 4)
	if d.Positions[0] != 'a' || d.Positions[1] != 'b' || d.Positions[2] != 'c' || d.Positions[3] != 'e' || d.Positions[4] != 'd' {
		t.Error("Unexpected dancefloor positioning", string(d.Positions))
	}
}

func TestDancefloor_Partner(t *testing.T) {
	d := NewDancefloor(5)
	d.Partner('e', 'b')
	if d.Positions[0] != 'a' || d.Positions[1] != 'e' || d.Positions[2] != 'c' || d.Positions[3] != 'd' || d.Positions[4] != 'b' {
		t.Error("Unexpected dancefloor positioning", string(d.Positions))
	}
}
