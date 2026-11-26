package textutil

import "testing"

func TestLongestWord(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: blank input
		{"", ""},
		{"  ", ""}, // Only spaces
		{"\t\n  ", ""}, // Only whitespace
		
		// Acceptance criteria: a tie (the first wins)
		{"abc def", "abc"}, // Equal length, first wins
		{"one two six", "one"}, // Three equal length, first wins
		{"aa bb cc", "aa"}, // All equal, first wins
		
		// Acceptance criteria: multi-byte runes
		{"hello 世界", "hello"}, // ASCII vs 2 multi-byte runes (hello is 5, 世界 is 2)
		{"世界 hello", "hello"}, // 2 multi-byte runes vs 5 ASCII
		{"a 世界", "世界"}, // 1 ASCII vs 2 multi-byte runes
		{"世 世界 界", "世界"}, // All multi-byte, middle is longest
		{"café naïve", "naïve"}, // Both have accents, naïve is longer
		
		// Additional test cases
		{"hello", "hello"}, // Single word
		{"a bb ccc", "ccc"}, // Clear winner
		{"short verylongword tiny", "verylongword"}, // Clear winner in middle
		{"   hello   world  ", "hello"}, // Tie with spaces
		{"the quick brown fox", "quick"}, // Real sentence
		{"a b c d e", "a"}, // All single char, first wins
		{"!@# hello", "hello"}, // Special chars vs word
		{"123 abcdef", "abcdef"}, // Numbers vs letters
	}

	for _, tt := range tests {
		if got := LongestWord(tt.in); got != tt.want {
			t.Errorf("LongestWord(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
