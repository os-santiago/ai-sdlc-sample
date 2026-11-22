// Package textfmt formats short texts for display.
package textfmt

import "strings"

// Truncate shortens s to at most n characters, ending with "..." when it
// had to cut. n smaller than 4 returns s cut to n characters.
func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	if n < 4 {
		return string(runes[:n])
	}
	return string(runes[:n-3]) + "..."
}

// Title upper-cases the first letter of each space-separated word.
func Title(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		if w != "" {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}
