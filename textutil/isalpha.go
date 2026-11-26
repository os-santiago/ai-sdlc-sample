// Package textutil provides small text helpers.
package textutil

import "unicode"

// IsAlpha reports whether s is non-empty and every rune is a Unicode letter.
// It returns false for empty strings.
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
