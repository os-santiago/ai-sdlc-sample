// Package textutil provides small text helpers.
package textutil

// HasDuplicateRune reports whether any rune appears more than once in s.
func HasDuplicateRune(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		if seen[r] {
			return true
		}
		seen[r] = true
	}
	return false
}
