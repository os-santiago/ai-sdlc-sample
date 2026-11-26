package mathutil

import "testing"

func TestAverage3(t *testing.T) {
	tests := []struct {
		a, b, c  int
		expected float64
	}{
		// Equal values
		{5, 5, 5, 5.0},
		{0, 0, 0, 0.0},
		{-3, -3, -3, -3.0},

		// Mixed signs
		{-10, 5, 2, -1.0},
		{10, -5, -2, 1.0},
		{-1, 1, 0, 0.0},

		// Non-integer mean
		{1, 2, 3, 2.0},
		{1, 1, 2, 4.0 / 3.0},
		{10, 20, 31, 61.0 / 3.0},
	}

	for _, tt := range tests {
		got := Average3(tt.a, tt.b, tt.c)
		if got != tt.expected {
			t.Errorf("Average3(%d, %d, %d) = %v, want %v", tt.a, tt.b, tt.c, got, tt.expected)
		}
	}
}
