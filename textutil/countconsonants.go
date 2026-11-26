// Package textutil provides small text helpers.
package textutil

// CountConsonants returns the number of ASCII consonants in s.
// It counts ASCII letters that are not vowels (a, e, i, o, u in either case).
func CountConsonants(s string) int {
	count := 0
	for _, r := range s {
		// Check if it's an ASCII letter (A-Z or a-z)
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			// Check if it's NOT a vowel
			if r != 'A' && r != 'E' && r != 'I' && r != 'O' && r != 'U' &&
				r != 'a' && r != 'e' && r != 'i' && r != 'o' && r != 'u' {
				count++
			}
		}
	}
	return count
}
