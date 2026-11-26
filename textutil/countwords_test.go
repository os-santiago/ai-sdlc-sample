package textutil

import "testing"

func TestCountWords(t *testing.T) {
	tests := []struct {
		s    string
		min  int
		want int
	}{
		// Acceptance criteria: empty input
		{"", 0, 0},
		{"", 1, 0},
		{"", 5, 0},

		// Acceptance criteria: min 0 counts all words
		{"hello world", 0, 2},
		{"one two three", 0, 3},
		{"a b c d e", 0, 5},

		// Acceptance criteria: min that excludes short words
		{"a bb ccc dddd", 3, 2},    // Only "ccc" and "dddd" have >= 3 runes
		{"hi hello world", 4, 2},   // Only "hello" and "world" have >= 4 runes
		{"I am a developer", 2, 2}, // Only "am" and "developer" have >= 2 runes

		// Additional test cases
		{"hello", 5, 1},
		{"hello", 6, 0},
		{"  hello   world  ", 1, 2},       // Leading/trailing spaces
		{"one two three four five", 4, 3}, // "three", "four", and "five"
		{"a bb ccc", 1, 3},
		{"a bb ccc", 2, 2},
		{"a bb ccc", 4, 0},
		{"café", 4, 1},  // Unicode handling
		{"こんにちは", 5, 1}, // Japanese characters (5 runes)
		{"こんにちは", 6, 0},
	}

	for _, tt := range tests {
		got := CountWords(tt.s, tt.min)
		if got != tt.want {
			t.Errorf("CountWords(%q, %d) = %d, want %d", tt.s, tt.min, got, tt.want)
		}
	}
}
