package day21

import "testing"

func TestPattern_Size(t *testing.T) {
	if size := Pattern("../..").Size(); size != 2 {
		t.Error("Expected size 2, got", size)
	}
	if size := Pattern(".../.../...").Size(); size != 3 {
		t.Error("Expected size 3, got", size)
	}
	if size := Pattern("..../..../..../....").Size(); size != 4 {
		t.Error("Expected size 4, got", size)
	}
}

func TestPattern_Split(t *testing.T) {
	patterns := Pattern("#..#/..../..../#..#").Split()
	if patterns[0][0] != "#./.." {
		t.Error("Pattern at (0,0) mismatch", patterns)
	}
	if patterns[0][1] != ".#/.." {
		t.Error("Pattern at (1,0) mismatch", patterns)
	}
	if patterns[1][0] != "../#." {
		t.Error("Pattern at (0,1) mismatch", patterns)
	}
	if patterns[1][1] != "../.#" {
		t.Error("Pattern at (1,1) mismatch", patterns)
	}
}
