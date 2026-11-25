package mathutil

import "testing"

func TestCountMultiples(t *testing.T) {
	tests := []struct {
		name     string
		lo, hi, k int
		want     int
	}{
		{
			name: "normal range",
			lo:   1, hi: 10, k: 2,
			want: 5, // 2, 4, 6, 8, 10
		},
		{
			name: "k <= 0",
			lo:   1, hi: 10, k: 0,
			want: 0,
		},
		{
			name: "k <= 0 negative",
			lo:   1, hi: 10, k: -1,
			want: 0,
		},
		{
			name: "empty range",
			lo:   5, hi: 3, k: 2,
			want: 0,
		},
		{
			name: "no multiples",
			lo:   1, hi: 3, k: 5,
			want: 0, // 5 is greater than 3
		},
		{
			name: "start at multiple",
			lo:   2, hi: 10, k: 2,
			want: 5, // 2, 4, 6, 8, 10
		},
		{
			name: "single element range",
			lo:   5, hi: 5, k: 5,
			want: 1, // 5
		},
		{
			name: "single element range no multiple",
			lo:   3, hi: 3, k: 5,
			want: 0,
		},
		{
			name: "large range",
			lo:   10, hi: 100, k: 10,
			want: 10, // 10, 20, 30, 40, 50, 60, 70, 80, 90, 100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountMultiples(tt.lo, tt.hi, tt.k); got != tt.want {
				t.Errorf("CountMultiples(%d, %d, %d) = %d, want %d", tt.lo, tt.hi, tt.k, got, tt.want)
			}
		})
	}
}