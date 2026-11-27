package mathutil

import "testing"

func TestCollatzSteps(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 0},
		{6, 8},
		{7, 16},
	}
	
	for _, tt := range tests {
		got := CollatzSteps(tt.n)
		if got != tt.want {
			t.Errorf("CollatzSteps(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
