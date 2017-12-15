package generator

import "testing"

func TestJudge_FindMatches(t *testing.T) {
	j := NewJudge(startA, startB)
	v := j.FindMatches(sampleSize, false)
	if v != 588 {
		t.Error("Expected 588 matches, got", v)
	}
}

func TestJudge_FindMatches2(t *testing.T) {
	j := NewJudge(startA, startB)
	v := j.FindMatches(5000000, true)
	if v != 309 {
		t.Error("Expected 309 matches, got", v)
	}
}
