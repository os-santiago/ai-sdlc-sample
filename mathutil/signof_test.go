package mathutil

import "testing"

func TestSignOf(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: -5, want: -1},
		{n: -1, want: -1},
		{n: 0, want: 0},
		{n: 1, want: 1},
		{n: 5, want: 1},
	}

	for _, tc := range tests {
		if got := SignOf(tc.n); got != tc.want {
			t.Errorf("SignOf(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}