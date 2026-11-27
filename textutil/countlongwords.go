package textutil

import (
	"strings"
	"unicode/utf8"
)

// CountLongWords counts whitespace-separated words in s with at least n runes.
func CountLongWords(s string, n int) int {
	fields := strings.Fields(s)
	count := 0
	for _, word := range fields {
		if utf8.RuneCountInString(word) >= n {
			count++
		}
	}
	return count
}
