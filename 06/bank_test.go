package day06

import "testing"

var steps = []Configuration{
	{2, 4, 1, 2},
	{3, 1, 2, 3},
	{0, 2, 3, 4},
	{1, 3, 4, 1},
	{2, 4, 1, 2},
}

func TestNewBank(t *testing.T) {
	config := Configuration{0, 2, 7, 0}
	bank := NewBank(config)

	if bank.NumSlots != 4 {
		t.Error("Expected bank to have 4 slots, got", bank.NumSlots)
	}

	if !bank.Configuration.Is(config) {
		t.Error("At step", 0, "expected configuration", config, "received", bank.Configuration, "instead.")
	}

	for i, step := range steps {
		bank.Redistribute()

		if !bank.Configuration.Is(step) {
			t.Error("At step", i + 1, "expected configuration", step, "received", bank.Configuration, "instead.")
		}

		seen, _ := bank.ConfigurationSeenBefore()
		if seen && i != 4 {
			t.Error("Expected configuration to have been seend before at step 5, got", i + 1, "instead.")
		}
	}
}
