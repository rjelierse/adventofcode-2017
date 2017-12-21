package day21

import (
	"fmt"
	"os"
)

type Pattern string

func ParseEnhancements(lines []string) map[Pattern]Pattern {
	enhancements := make(map[Pattern]Pattern)
	for _, line := range lines {
		var input, output Pattern
		_, err := fmt.Sscanf(line, "%s => %s", &input, &output)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot parse line:", err)
			continue
		}
		for _, pattern := range input.Permutations() {
			if _, exists := enhancements[pattern]; !exists {
				enhancements[pattern] = output
			}
		}
	}
	return enhancements
}

func (p Pattern) Permutations() (patterns []Pattern) {
	var r, f Grid
	r = p.ToGrid()
	f = r.Flip()
	for i := 0; i <= 360; i += 90 {
		patterns = append(patterns, r.ToPattern(), f.ToPattern())
		r = r.Rotate()
		f = r.Flip()
	}
	return
}

func (p Pattern) ToGrid() Grid {
	return NewGridFromPattern(p)
}
