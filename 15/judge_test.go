package generator

import "testing"

func TestJudge_FindMatches(t *testing.T) {
	j := NewJudge(startA, startB)
	v := j.FindMatches(sampleSize)
	if v != 588 {
		t.Error("Expected 588 matches, got", v)
	}
}
