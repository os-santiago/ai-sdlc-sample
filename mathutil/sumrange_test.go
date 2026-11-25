package mathutil

import "testing"

func TestSumRange(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 5, 15},   // 1+2+3+4+5 = 15
		{2, 4, 9},    // 2+3+4 = 9
		{3, 3, 3},    // 3 = 3
		{5, 2, 0},    // a > b, should return 0
		{0, 0, 0},    // 0 = 0
		{-2, 2, 0},   // -2 + -1 + 0 + 1 + 2 = 0
		{-3, -1, -6}, // -3 + -2 + -1 = -6
	}

	for _, tc := range tests {
		got := SumRange(tc.a, tc.b)
		if got != tc.expected {
			t.Errorf("SumRange(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.expected)
		}
	}
}