package mathutil

import "testing"

func TestIsTriangular(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{0, false}, // Not triangular (k must be >= 1)
		{1, true},  // 1 = 1*(1+1)/2
		{3, true},  // 3 = 2*(2+1)/2
		{10, true}, // 10 = 4*(4+1)/2
		{11, false}, // Not triangular
	}

	for _, tc := range tests {
		got := IsTriangular(tc.n)
		if got != tc.expected {
			t.Errorf("IsTriangular(%d) = %v, want %v", tc.n, got, tc.expected)
		}
	}
}