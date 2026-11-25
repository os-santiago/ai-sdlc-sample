package mathutil

import "testing"

// TestTriangleNumber tests the TriangleNumber function.
func TestTriangleNumber(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{4, 10}, // 1+2+3+4 = 10
		{-1, 0},
	}
	
	for _, tc := range tests {
		got := TriangleNumber(tc.n)
		if got != tc.expected {
			t.Errorf("TriangleNumber(%d) = %d, want %d", tc.n, got, tc.expected)
		}
	}
}