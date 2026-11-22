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
			runes := []rune(w)
			words[i] = strings.ToUpper(string(runes[:1])) + string(runes[1:])
		}
	}
	return strings.Join(words, " ")
}

// PadLeft prepends pad runes to s until it reaches width characters.
// Width counts characters (runes), not bytes. If s already has width
// or more characters, it is returned unchanged.
func PadLeft(s string, width int, pad rune) string {
	runes := []rune(s)
	if len(runes) >= width {
		return s
	}
	needed := width - len(runes)
	padding := strings.Repeat(string(pad), needed)
	return padding + s
}
