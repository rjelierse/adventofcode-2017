package sound

import (
	"log"
	"testing"
)

func TestSoundBoard(t *testing.T) {
	instructions := []string{
		"set a 1",
		"add a 2",
		"mul a a",
		"mod a 5",
		"snd a",
		"set a 0",
		"rcv a",
		"jgz a -1",
		"set a 1",
		"jgz a -2",
	}
	lastPlayed, err := SoundBoard(instructions)
	if err != nil {
		log.Fatal("Cannot calculate last played:", err)
	}
	if lastPlayed != 4 {
		t.Error("Unexpected last played value:", lastPlayed)
	}
}
