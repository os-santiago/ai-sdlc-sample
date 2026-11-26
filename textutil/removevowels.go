// Package textutil provides small text helpers.
package textutil

// RemoveVowels returns s without a, e, i, o, u in either case, keeping every other rune.
func RemoveVowels(s string) string {
	var result []rune
	for _, r := range s {
		// Check if it's NOT a vowel
		if r != 'A' && r != 'E' && r != 'I' && r != 'O' && r != 'U' &&
			r != 'a' && r != 'e' && r != 'i' && r != 'o' && r != 'u' {
			result = append(result, r)
		}
	}
	return string(result)
}
