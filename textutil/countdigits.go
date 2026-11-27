// Package textutil provides small text helpers.
package textutil

import "unicode"

// CountDigits returns the number of digits in s.
// It counts runes for which unicode.IsDigit is true.
func CountDigits(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			count++
		}
	}
	return count
}