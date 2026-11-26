package textutil

import (
	"strings"
	"unicode"
)

// TitleWords returns s with the first rune of each whitespace-separated word
// upper-cased and the rest unchanged. Words are joined with single spaces.
func TitleWords(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}

	for i, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}

	return strings.Join(words, " ")
}
