package textutil

import "testing"

// CommonPrefix tests
func TestCommonPrefix(t *testing.T) {
	tests := []struct {
		a, b   string
		want   string
	}{
		// Acceptance criteria: no common prefix
		{"hello", "world", ""},
		
		// Acceptance criteria: partial prefix
		{"hello", "help", "hel"},
		
		// Acceptance criteria: identical strings
		{"hello", "hello", "hello"},
		
		// Additional test cases
		{"", "test", ""},
		{"test", "", ""},
		{"same", "same", "same"},
		{"abc", "abcde", "abc"},
		{"abcde", "abc", "abc"},
		{"test", "testing", "test"},
		{"café", "café", "café"},
		{"héllo", "héllo", "héllo"},
		{"héllo", "hello", "h"},
		{"日本語", "日本語", "日本語"},
		{"日本語", "日本", "日本"},
		{"a", "a", "a"},
		{"a", "b", ""},
		{"abc", "def", ""},
		{"123", "12345", "123"},
		{"12345", "123", "123"},
		{"test123", "test456", "test"},
	}
	
	for _, tt := range tests {
		if got := CommonPrefix(tt.a, tt.b); got != tt.want {
			t.Errorf("CommonPrefix(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
		}
	}
}