// Package textutil provides small text helpers.
package textutil

import (
	"strings"
	"unicode"
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