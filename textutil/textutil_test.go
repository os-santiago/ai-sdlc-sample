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
