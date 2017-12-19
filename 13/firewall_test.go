package day13

import "testing"

func TestNewFirewall(t *testing.T) {
	f := NewFirewall()
	if f.depth != 0 {
		t.Error("Unexpected firewall depth", f.depth)
	}
}

func TestFirewall_AddLayer(t *testing.T) {
	f := NewFirewall()
	f.AddLayer(2, 3)

	if f.depth != 2 {
		t.Error("Unexpected firewall depth", f.depth)
	}

	if f.layers[2] != 3 {
		t.Error("Unexpected layer range", f.layers[2])
	}
}

func TestFirewall_Traverse(t *testing.T) {
	f := NewFirewall()
	f.AddLayer(0, 3)
	f.AddLayer(1, 2)
	f.AddLayer(4, 4)
	f.AddLayer(6, 4)
	p := f.Traverse(0)

	if p != 24 {
		t.Error("Unexpected firewall penalty", p)
	}
}

func TestFirewall_TraverseWithDelay(t *testing.T) {
	f := NewFirewall()
	f.AddLayer(0, 3)
	f.AddLayer(1, 2)
	f.AddLayer(4, 4)
	f.AddLayer(6, 4)
	p := f.Traverse(10)

	if p != 0 {
		t.Error("Unexpected firewall penalty", p)
	}
}
