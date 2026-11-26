// Package textutil provides small text helpers.
package textutil

import "unicode"

// CountPunct returns the number of punctuation marks in s.
// It counts runes for which unicode.IsPunct is true.
func CountPunct(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsPunct(r) {
			count++
		}
	}
	return count
}
