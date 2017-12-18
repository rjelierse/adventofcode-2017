package sound

import "testing"

func TestRunQueue(t *testing.T) {
	instructions := []string{
		"snd 1",
		"snd 2",
		"snd p",
		"rcv a",
		"rcv b",
		"rcv c",
		"rcv d",
	}
	sent := RunQueue(instructions)
	if sent != 3 {
		t.Error("Unexpected number of sent values:", sent)
	}
}
