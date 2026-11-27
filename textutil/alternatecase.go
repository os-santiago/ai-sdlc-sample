package textutil

import "unicode"

// AlternateCase upper-cases runes at even indexes and lower-cases runes at odd indexes, counting runes.
func AlternateCase(s string) string {
	if s == "" {
		return ""
	}
	
	runes := []rune(s)
	for i, r := range runes {
		if i%2 == 0 {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}