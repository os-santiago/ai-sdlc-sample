package textutil

import "testing"

func TestReverseASCII(t *testing.T) {
	for in, want := range map[string]string{"" : "", "a": "a", "abc": "cba", "ai-sdlc": "clds-ia"} {
		if got := Reverse(in); got != want {
			t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestReverseMultibyte(t *testing.T) {
	for in, want := range map[string]string{
		"héllo":   "olléh",
		"日本語": "語本日",
	} {
		if got := Reverse(in); got != want {
			t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Existing test cases
		{"", 0},
		{"one", 1},
		{" two  words ", 2},
		// New test cases for punctuation-only tokens
		{"hello , world !", 2},
		{"one two three", 3},
		// Edge cases with punctuation
		{"!", 0},
		{"!!!", 0},
		{"   ,   ", 0},
		{"a , b , c", 3},
		{"word!", 1},
		{"!word", 1},
		{"wo!rd", 1},
		// Numbers count as words
		{"123", 1},
		{"abc123", 1},
		{"123abc", 1},
		// Mixed punctuation and words
		{"hello, world!", 2},
		{"   ,   !   ?   ", 0},
	}

	for _, tt := range tests {
		if got := WordCount(tt.in); got != tt.want {
			t.Errorf("WordCount(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Empty string
		{"", true},
		// Single character
		{"a", true},
		// Simple palindromes
		{"aba", true},
		{"abba", true},
		// Not palindromes
		{"abc", false},
		{"ab", false},
		// Case and space handling
		{"Never odd or even", true},
		{"A man a plan a canal Panama", true},
		// Unicode characters
		{"ñaña", false},
		{"añña", true},
	}

	for _, tt := range tests {
		if got := IsPalindrome(tt.in); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestInitials(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Empty string
		{"", ""},
		// Basic case
		{"ada lovelace", "AL"},
		// Extra whitespace
		{"  ada   lovelace  ", "AL"},
		{"one  two   three", "OTT"},
		// Unicode awareness
		{"élan vital", "ÉV"},
		{"über alles", "ÜA"},
		// Single word
		{"hello", "H"},
		// Mixed case input (should uppercase)
		{"hello World", "HW"},
		{"HELLO world", "HW"},
	}

	for _, tt := range tests {
		if got := Initials(tt.in); got != tt.want {
			t.Errorf("Initials(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Empty string
		{"", ""},
		// Basic ASCII
		{"hello world", "Hello world"},
		{"hello", "Hello"},
		// Already capitalized
		{"Hello", "Hello"},
		// Unicode awareness
		{"élan", "Élan"},
		{"über", "Über"},
		// Single character
		{"a", "A"},
		{"é", "É"},
		// Mixed case - only first changes
		{"hELLO", "HELLO"},
	}

	for _, tt := range tests {
		if got := Capitalize(tt.in); got != tt.want {
			t.Errorf("Capitalize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Acceptance criteria cases
		{"Hello World", 3},
		{"", 0},
		// Additional test cases
		// All vowels lowercase
		{"aeiou", 5},
		// All vowels uppercase
		{"AEIOU", 5},
		// Mixed case
		{"AeIoU", 5},
		// No vowels
		{"bcdfg", 0},
		{"xyz", 0},
		// Single vowel
		{"a", 1},
		{"E", 1},
		// Repeated vowels
		{"aaaa", 4},
		{"eeeee", 5},
		// Real words
		{"beautiful", 5},
		{"rhythm", 0},
		{"education", 5},
	}

	for _, tt := range tests {
		if got := CountVowels(tt.in); got != tt.want {
			t.Errorf("CountVowels(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria cases
		{"one two three", "three two one"},
		{"", ""},
		// Additional test cases
		// Single word
		{"hello", "hello"},
		// Two words
		{"hello world", "world hello"},
		// Multiple spaces (should normalize to single spaces)
		{"one  two   three", "three two one"},
		{"  hello   world  ", "world hello"},
		// Tab and newline whitespace
		{"one\ttwo\nthree", "three two one"},
		// Only whitespace
		{"   ", ""},
		{"\t\n  ", ""},
		// Four words
		{"the quick brown fox", "fox brown quick the"},
		// Numbers and mixed characters
		{"123 abc 456", "456 abc 123"},
		// Punctuation (treated as part of words)
		{"hello, world!", "world! hello,"},
		// Unicode
		{"héllo wörld", "wörld héllo"},
	}

	for _, tt := range tests {
		if got := ReverseWords(tt.in); got != tt.want {
			t.Errorf("ReverseWords(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: empty string
		{"", true},
		// Acceptance criteria: only whitespace (ASCII)
		{" ", true},
		{"  ", true},
		{"   ", true},
		{"\t", true},
		{"\n", true},
		{"\r", true},
		{" \t\n ", true},
		// Acceptance criteria: Unicode whitespace
		{"\u00A0", true},     // Non-breaking space
		{"\u2000", true},     // En quad
		{"\u2003", true},     // Em space
		{"\u3000", true},     // Ideographic space
		{" \u00A0\t", true}, // Mixed whitespace
		// Acceptance criteria: false for strings with content
		{" a ", false},
		{"a", false},
		{"hello", false},
		{" hello ", false},
		{"\thello\n", false},
	}

	for _, tt := range tests {
		if got := IsBlank(tt.in); got != tt.want {
			t.Errorf("IsBlank(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestHasPrefixFold(t *testing.T) {
	tests := []struct {
		s      string
		prefix string
		want   bool
	}{
		// Acceptance criteria cases
		{"Hello", "he", true},
		{"Hi", "hello", false},
		// Empty prefix
		{"", "", true},
		{"Hello", "", true},
		// Exact match (case-insensitive)
		{"hello", "HELLO", true},
		{"HELLO", "hello", true},
		// Partial match
		{"Hello World", "hello", true},
		{"Hello World", "HELLO", true},
		// No match
		{"Hello", "world", false},
		{"abc", "xyz", false},
		// Prefix longer than string
		{"Hi", "Hello", false},
		// Unicode awareness
		{"Élégant", "élé", true},
		{"ÉLÉGANT", "élé", true},
	}

	for _, tt := range tests {
		if got := HasPrefixFold(tt.s, tt.prefix); got != tt.want {
			t.Errorf("HasPrefixFold(%q, %q) = %v, want %v", tt.s, tt.prefix, got, tt.want)
		}
	}
}

func TestSqueeze(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: specific example
		{"  a   b  ", "a b"},
		// Acceptance criteria: empty string
		{"", ""},
		// Only whitespace (should return empty)
		{" ", ""},
		{"  ", ""},
		// No whitespace
		{"abc", "abc"},
		// Single space (already normalized)
		{"a b", "a b"},
		// Multiple spaces
		{"a  b", "a b"},
		{"a   b   c", "a b c"},
		// Leading/trailing whitespace
		{" a", "a"},
		{"a ", "a"},
		{" a ", "a"},
		// Tab characters
		{"a\tb", "a b"},
		{"a\t\tb", "a b"},
		// Newlines
		{"a\nb", "a b"},
		{"a\n\nb", "a b"},
		// Mixed whitespace
		{"  a\t\n  b  ", "a b"},
		{"one  \ttwo\n\nthree   ", "one two three"},
	}

	for _, tt := range tests {
		if got := Squeeze(tt.in); got != tt.want {
			t.Errorf("Squeeze(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCountRunes(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Acceptance criteria: empty input
		{"", 0},
		// Acceptance criteria: ASCII
		{"hello", 5},
		{"abc", 3},
		{"a", 1},
		// Acceptance criteria: multibyte Unicode ("héllo" returns 5)
		{"héllo", 5},
		// Additional multibyte test cases
		{"日本語", 3},
		{"こんにちは", 5},
		{"café", 4},
		{"naïve", 5},
		{"Zürich", 6},
		// Mixed ASCII and multibyte
		{"hello世界", 7},
		{"test😀emoji", 10},
		// Emojis (each emoji is one rune)
		{"😀😁😂", 3},
	}

	for _, tt := range tests {
		if got := CountRunes(tt.in); got != tt.want {
			t.Errorf("CountRunes(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}
