package textutil

import "testing"

func TestCountDigits(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Acceptance criteria: empty input
		{"", 0},
		
		// Acceptance criteria: no digits
		{"hello world", 0},
		{"abc", 0},
		{"!@#$%", 0},
		{"   ", 0},
		
		// Acceptance criteria: mixed text
		{"hello123world", 3},
		{"abc123def456", 6},
		{"123", 3},
		{"a1b2c3", 3},
		{"test123test", 3},
		
		// Additional test cases
		{"1234567890", 10}, // All digits
		{"1a2b3c4d5e", 5},  // Alternating digits and letters
		{"!@#123$%^456", 6}, // Mixed with special characters
		{"Hello 123 World 456", 6}, // With spaces
		{"café123", 3}, // Unicode characters with digits
		{"αβγ123", 3}, // Greek letters with digits
		{"٠١٢٣٤٥٦٧٨٩", 10}, // Arabic numerals
		{"۰۱۲۳۴۵۶۷۸۹", 10}, // Persian numerals
		{"123abc456def789", 9}, // Multiple digit sequences
		{"a1b2c3d4e5", 5}, // Single digits
		{"", 0}, // Empty string
		{" ", 0}, // Single space
		{"\t\n", 0}, // Whitespace only
		{"!@#$%^&*()", 0}, // Special characters only
		{"12345678901234567890", 20}, // Large number of digits
	}

	for _, tt := range tests {
		if got := CountDigits(tt.in); got != tt.want {
			t.Errorf("CountDigits(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}