package garbage

import "testing"

type test struct {
	Stream string
	Score int
}

var tests = []test{
	{"{}", 1},
	{"{{{}}}", 6},
	{"{{},{}}", 5},
	{"{{{},{},{{}}}}", 16},
	{"{<a>,<a>,<a>,<a>}", 1},
	{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
	{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
	{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
}

func TestReadStream(t *testing.T) {
	for _, i := range tests {
		s, _ := ReadStream(i.Stream)
		if s != i.Score {
			t.Error("For input", i.Stream, "expected score", i.Score, "got", s)
		}
	}
}
