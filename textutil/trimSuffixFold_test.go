package textutil

import "testing"

func TestTrimSuffixFold(t *testing.T) {
	tests := []struct {
		s      string
		suffix string
		want   string
	}{
		// Basic match with different case
		{"hello", "o", "hell"},
		{"hello", "O", "hell"},
		{"HELLO", "o", "HELL"},
		{"HELLO", "O", "HELL"},
		
		// Different case combinations
		{"Hello World", "world", "Hello "},
		{"Hello World", "WORLD", "Hello "},
		{"Hello World", "WoRlD", "Hello "},
		
		// No match
		{"hello", "world", "hello"},
		{"hello", "loW", "hello"},
		
		// Empty suffix
		{"hello", "", "hello"},
		{"", "", ""},
		
		// Edge cases
		{"", "hello", ""},
		{"a", "a", ""},
		{"a", "A", ""},
		{"test", "test", ""},
		{"test", "TEST", ""},
		{"Test", "test", ""},
		{"Test", "TEST", ""},
		
		// Unicode
		{"héllo wörld", "wörld", "héllo "},
		{"héllo wörld", "WÖRLD", "héllo "},
		{"héllo wörld", "Wörld", "héllo "},
		
		// Suffix longer than string
		{"hello", "hello world", "hello"},
		
		// Case where suffix is not at the end
		{"hello world", "orld", "hello w"},
		{"hello world", "ORLD", "hello w"},
	}

	for _, tt := range tests {
		if got := TrimSuffixFold(tt.s, tt.suffix); got != tt.want {
			t.Errorf("TrimSuffixFold(%q, %q) = %q, want %q", tt.s, tt.suffix, got, tt.want)
		}
	}
}