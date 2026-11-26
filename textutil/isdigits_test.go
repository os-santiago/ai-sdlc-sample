package textutil

import "testing"

// IsDigits reports whether s is non-empty and every rune is an ASCII digit.
// It returns false for empty strings.
func TestIsDigits(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: empty input
		{"", false},
		
		// Acceptance criteria: digits only
		{"12345", true},
		
		// Acceptance criteria: string with a letter
		{"123a456", false},
		
		// Additional test cases
		{"0", true},
		{"9", true},
		{"1234567890", true},
		{"123 456", false}, // contains space
		{"123-456", false}, // contains dash
		{"123.456", false}, // contains dot
		{"abc", false},
		{"123abc", false},
		{"abc123", false},
		{"", false}, // empty string
		{" ", false}, // space only
		{"\t", false}, // tab only
		{"\n", false}, // newline only
	}

	for _, tt := range tests {
		if got := IsDigits(tt.in); got != tt.want {
			t.Errorf("IsDigits(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}