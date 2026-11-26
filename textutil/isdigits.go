// Package textutil provides small text helpers.
package textutil

// IsDigits reports whether s is non-empty and every rune is an ASCII digit.
// It returns false for empty strings.
func IsDigits(s string) bool {
	if s == "" {
		return false
	}
	
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}