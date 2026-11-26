package textutil

import "testing"

func TestFirstWord(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: empty input
		{"", ""},
		
		// Acceptance criteria: leading spaces
		{"  hello world", "hello"},
		{"   hello", "hello"},
		
		// Acceptance criteria: several words
		{"hello world", "hello"},
		{"one two three", "one"},
		
		// Additional test cases
		{"  ", ""}, // Only spaces
		{"\t\n  ", ""}, // Only whitespace
		{"hello", "hello"},
		{"  hello   world  ", "hello"},
		{"a b c d e", "a"},
		{"   a   ", "a"},
		{"!@# hello", "!@#"},
		{"123 abc", "123"},
		{"   123   abc   ", "123"},
	}

	for _, tt := range tests {
		if got := FirstWord(tt.in); got != tt.want {
			t.Errorf("FirstWord(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}