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
	for in, want := range map[string]int{"" : 0, "one": 1, " two  words ": 2} {
		if got := WordCount(in); got != want {
			t.Errorf("WordCount(%q) = %d, want %d", in, got, want)
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
