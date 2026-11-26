package textutil

import "testing"

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: empty input
		{"", false},
		
		// Acceptance criteria: letters with accents
		{"café", true},
		{"naïve", true},
		{"résumé", true},
		{"Ñoño", true},
		
		// Acceptance criteria: a digit
		{"abc1", false},
		{"123", false},
		{"a1b", false},
		
		// Acceptance criteria: a space
		{"hello world", false},
		{" ", false},
		{"a b", false},
		
		// Additional test cases
		{"abc", true},
		{"ABC", true},
		{"AbCdEf", true},
		{"hello", true},
		{"HELLO", true},
		{"a", true},
		{"Z", true},
		{"hello-world", false}, // contains dash
		{"hello_world", false}, // contains underscore
		{"hello!", false}, // contains punctuation
		{"hello\n", false}, // contains newline
		{"hello\t", false}, // contains tab
	}

	for _, tt := range tests {
		if got := IsAlpha(tt.in); got != tt.want {
			t.Errorf("IsAlpha(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
