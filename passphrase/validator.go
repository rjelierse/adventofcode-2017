package passphrase

import (
	"sort"
	"strings"
)

func IsUnique(phrase string) bool {
	words := strings.Split(phrase, " ")

	return unique(words)
}

func IsNotAnagram(phrase string) bool {
	words := strings.Split(phrase, " ")
	grams := words[:0]

	for _, word := range words {
		glyphs := strings.Split(word, "")
		sort.Strings(glyphs)
		grams = append(grams, strings.Join(glyphs, ""))
	}

	return unique(grams)
}

func unique(words []string) bool {
	sort.Strings(words)

	for i := 1; i < len(words); i++ {
		if words[i-1] == words[i] {
			return false
		}
	}

	return true
}
