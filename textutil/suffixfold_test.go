package textutil

import "testing"

func TestHasSuffixFold(t *testing.T) {
	tests := []struct {
		s      string
		suffix string
		want   bool
	}{
		// Basic cases
		{"hello", "o", true},
		{"hello", "O", true},
		{"HELLO", "o", true},
		{"HELLO", "O", true},
		
		// Different case combinations
		{"Hello World", "world", true},
		{"Hello World", "WORLD", true},
		{"Hello World", "WoRlD", true},
		
		// Non-matching cases
		{"hello", "world", false},
		{"hello", "loW", false},
		
		// Empty suffix
		{"hello", "", true},
		{"", "", true},
		
		// Edge cases
		{"", "hello", false},
		{"a", "a", true},
		{"a", "A", true},
		{"test", "test", true},
		{"test", "TEST", true},
		{"Test", "test", true},
		{"Test", "TEST", true},
		
		// Unicode
		{"héllo wörld", "wörld", true},
		{"héllo wörld", "WÖRLD", true},
		{"héllo wörld", "Wörld", true},
	}

	for _, tt := range tests {
		if got := HasSuffixFold(tt.s, tt.suffix); got != tt.want {
			t.Errorf("HasSuffixFold(%q, %q) = %v, want %v", tt.s, tt.suffix, got, tt.want)
		}
	}
}