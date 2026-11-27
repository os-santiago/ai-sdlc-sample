package textutil

import "unicode"

// FirstNonSpace returns the first rune that is not Unicode white space and true,
// or 0 and false if there is none.
func FirstNonSpace(s string) (rune, bool) {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return r, true
		}
	}
	return 0, false
}
