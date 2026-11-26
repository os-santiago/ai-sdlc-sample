package textutil

import "testing"

func TestTitleWords(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: empty input
		{"", ""},

		// Acceptance criteria: several words
		{"hello world", "Hello World"},
		{"one two three", "One Two Three"},
		{"the quick brown fox", "The Quick Brown Fox"},

		// Acceptance criteria: extra spaces
		{"  hello   world  ", "Hello World"},
		{"   one  two   three   ", "One Two Three"},
		{"\thello\tworld", "Hello World"},
		{"  \t  hello   \n  world  ", "Hello World"},

		// Additional edge cases
		{"hello", "Hello"},
		{"HELLO WORLD", "HELLO WORLD"}, // Already uppercase
		{"hELLO wORLD", "HELLO WORLD"}, // Mixed case - only first letter uppercased
		{"a b c", "A B C"},
		{"123 abc", "123 Abc"},
		{"ñoño español", "Ñoño Español"}, // Unicode support
	}

	for _, tt := range tests {
		if got := TitleWords(tt.in); got != tt.want {
			t.Errorf("TitleWords(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
