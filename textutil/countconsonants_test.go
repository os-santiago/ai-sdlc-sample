package textutil

import "testing"

func TestCountConsonants(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Acceptance criteria: empty input
		{"", 0},

		// Acceptance criteria: mixed case
		{"Hello", 3}, // H, l, l
		{"WORLD", 4}, // W, R, L, D
		{"AaBbCc", 4}, // B, b, C, c (not counting A, a)

		// Acceptance criteria: non-letter characters
		{"Hello123", 3}, // H, l, l
		{"a1b2c3", 2}, // b, c
		{"Hello, World!", 7}, // H, l, l, W, r, l, d

		// Additional test cases
		{"aeiou", 0}, // All vowels
		{"AEIOU", 0}, // All vowels uppercase
		{"bcdfg", 5}, // All consonants
		{"BCDFG", 5}, // All consonants uppercase
		{"xyz", 3}, // x, y, z
		{"123!@#", 0}, // No letters
		{"The quick brown fox", 11}, // T, h, q, c, k, b, r, w, n, f, x
	}

	for _, tt := range tests {
		if got := CountConsonants(tt.in); got != tt.want {
			t.Errorf("CountConsonants(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
