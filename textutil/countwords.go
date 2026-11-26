package textutil

import (
	"strings"
	"unicode/utf8"
)

// CountWords counts whitespace-separated words in s with at least min runes.
func CountWords(s string, min int) int {
	fields := strings.Fields(s)
	count := 0
	for _, word := range fields {
		if utf8.RuneCountInString(word) >= min {
			count++
		}
	}
	return count
}
