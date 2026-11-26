package textutil

import "testing"

// CountPunct counts the runes for which unicode.IsPunct is true.
func TestCountPunct(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Acceptance criteria: empty input
		{"", 0},

		// Acceptance criteria: a sentence with punctuation
		{"Hello, world! How are you?", 3},

		// Acceptance criteria: a string without any
		{"Hello world", 0},

		// Additional test cases
		{"!!!", 3},
		{".", 1},
		{"one, two; three: four.", 4},
		{"abc123", 0},
		{"(parens)", 2},
		{"[brackets]", 2},
		{"{braces}", 2},
		{"a-b", 1},
		{"test@example.com", 2},
	}

	for _, tt := range tests {
		if got := CountPunct(tt.in); got != tt.want {
			t.Errorf("CountPunct(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
