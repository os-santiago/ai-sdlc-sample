package mathutil

import "testing"

func TestNextEven(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 2},
		{1, 2},
		{2, 4},
		{-1, 0},
		{-2, 0},
		{3, 4},
		{-3, -2},
		{100, 102},
		{-101, -100},
	}

	for _, tt := range tests {
		got := NextEven(tt.n)
		if got != tt.expected {
			t.Errorf("NextEven(%d) = %d, want %d", tt.n, got, tt.expected)
		}
	}
}