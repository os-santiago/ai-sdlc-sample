package textutil

import "testing"

func TestInitial(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: empty input
		{"", ""},

		// Acceptance criteria: ASCII first rune
		{"hello", "H"},
		{"world", "W"},
		{"a", "A"},
		{"ABC", "A"},
		{"test", "T"},

		// Acceptance criteria: multibyte first rune
		{"élan", "É"},
		{"über", "Ü"},
		{"日本語", "日"},
		{"café", "C"},
		{"ñoño", "Ñ"},
	}

	for _, tt := range tests {
		if got := Initial(tt.in); got != tt.want {
			t.Errorf("Initial(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
