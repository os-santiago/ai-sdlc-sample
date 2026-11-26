package textutil

import "unicode"

// Initial returns the first rune of s upper-cased, or the empty string for empty input.
func Initial(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	return string(unicode.ToUpper(r[0]))
}
