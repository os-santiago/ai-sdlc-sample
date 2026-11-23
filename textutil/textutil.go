// Package textutil provides small text helpers.
package textutil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Reverse returns s with its characters in reverse order.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// WordCount returns the number of whitespace-separated words in s,
// excluding tokens that contain no letters or digits.
func WordCount(s string) int {
	fields := strings.Fields(s)
	count := 0
	for _, field := range fields {
		for _, r := range field {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				count++
				break
			}
		}
	}
	return count
}

// IsPalindrome reports whether s reads the same forwards and backwards,
// ignoring case and spaces.
func IsPalindrome(s string) bool {
	// Normalize: remove spaces and convert to lowercase
	normalized := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	
	// Compare with reverse
	r := []rune(normalized)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

// Initials returns the upper-case initials of the words in s.
// Words are separated by whitespace, and extra whitespace is ignored.
func Initials(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}
	
	var initials []rune
	for _, word := range words {
		runes := []rune(word)
		initials = append(initials, unicode.ToUpper(runes[0]))
	}
	return string(initials)
}

// Abbreviate returns the upper-case first letter of each word followed by a dot.
// Words are separated by whitespace, and extra whitespace is ignored.
func Abbreviate(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}
	
	var result strings.Builder
	for _, word := range words {
		runes := []rune(word)
		result.WriteRune(unicode.ToUpper(runes[0]))
		result.WriteRune('.')
	}
	return result.String()
}

// Capitalize returns s with its first letter upper-cased.
// It is Unicode-aware and leaves the rest of the string unchanged.
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// CountVowels returns the number of vowels (a, e, i, o, u) in s,
// counting both upper and lower case.
func CountVowels(s string) int {
	count := 0
	for _, r := range s {
		lower := unicode.ToLower(r)
		if lower == 'a' || lower == 'e' || lower == 'i' || lower == 'o' || lower == 'u' {
			count++
		}
	}
	return count
}

// ReverseWords returns s with the order of its words reversed.
// Words are separated by any whitespace, and the result joins them with single spaces.
func ReverseWords(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}
	
	// Reverse the slice
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	
	return strings.Join(words, " ")
}

// IsBlank reports whether s is empty or contains only whitespace.
// It considers all Unicode whitespace characters.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// HasPrefixFold reports whether s starts with prefix, ignoring Unicode case.
// An empty prefix always matches.
func HasPrefixFold(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return strings.EqualFold(s[:len(prefix)], prefix)
}

// Squeeze replaces each run of Unicode whitespace with a single space
// and trims both ends.
func Squeeze(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields, " ")
}

// CountRunes returns the number of Unicode code points in s.
// It counts runes (characters), not bytes.
func CountRunes(s string) int {
	return utf8.RuneCountInString(s)
}

// CountFields returns the number of non-empty fields of s split on sep.
// Fields are separated by the separator rune sep, and consecutive separators
// do not produce empty fields (they are skipped).
func CountFields(s string, sep rune) int {
	count := 0
	for _, field := range strings.Split(s, string(sep)) {
		if field != "" {
			count++
		}
	}
	return count
}

// StartsWithVowel reports whether s starts with an ASCII vowel (a, e, i, o, u).
// The check is case-insensitive. An empty string returns false.
func StartsWithVowel(s string) bool {
	if s == "" {
		return false
	}
	
	firstRune := []rune(s)[0]
	lower := unicode.ToLower(firstRune)
	return lower == 'a' || lower == 'e' || lower == 'i' || lower == 'o' || lower == 'u'
}

// CountChar returns the number of occurrences of r in s.
// It counts multibyte runes correctly.
func CountChar(s string, r rune) int {
	count := 0
	for _, char := range s {
		if char == r {
			count++
		}
	}
	return count
}

// IsUpper reports whether s has at least one letter and all letters are uppercase.
// Non-letter characters are ignored.
func IsUpper(s string) bool {
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetter
}

// Acronym returns the uppercase first letter of each space-separated word.
// Words are separated by whitespace, and extra whitespace is ignored.
func Acronym(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}
	
	var result strings.Builder
	for _, word := range words {
		runes := []rune(word)
		result.WriteRune(unicode.ToUpper(runes[0]))
	}
	return result.String()
}

// Lines splits s into lines, splitting on \n and \r\n without including the endings.
// A trailing line ending does not produce an empty last line.
// Lines("") returns an empty, non-nil slice.
func Lines(s string) []string {
	if s == "" {
		return []string{}
	}
	
	var lines []string
	var current strings.Builder
	
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, current.String())
			current.Reset()
		} else if s[i] == '\r' && i+1 < len(s) && s[i+1] == '\n' {
			lines = append(lines, current.String())
			current.Reset()
			i++ // Skip the \n
		} else {
			current.WriteByte(s[i])
		}
	}
	
	// Add the last line if there's content remaining
	if current.Len() > 0 {
		lines = append(lines, current.String())
	}
	
	return lines
}
