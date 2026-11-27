package textutil

import "testing"

func TestAlternateCase(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},                    // empty input
		{"a", "A"},                  // one rune
		{"ab", "Ab"},                // two runes
		{"abc", "AbC"},              // three runes
		{"hello", "HeLlO"},          // mixed word
		{"world", "WoRlD"},          // another mixed word
		{"Go", "Go"},                // already mixed case
		{"golang", "GoLaNg"},        // longer mixed word
	}

	for _, tt := range tests {
		if got := AlternateCase(tt.in); got != tt.want {
			t.Errorf("AlternateCase(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}