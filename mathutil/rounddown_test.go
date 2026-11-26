package mathutil_test

import (
	"testing"

	"github.com/os-santiago/ai-sdlc-sample/mathutil"
)

func TestRoundDown(t *testing.T) {
	tests := []struct {
		name string
		n    int
		step int
		want int
	}{
		{
			name: "step zero returns n unchanged",
			n:    10,
			step: 0,
			want: 10,
		},
		{
			name: "step negative returns n unchanged",
			n:    10,
			step: -3,
			want: 10,
		},
		{
			name: "negative n with positive step",
			n:    -7,
			step: 3,
			want: -9,
		},
		{
			name: "exact multiple",
			n:    6,
			step: 3,
			want: 6,
		},
		{
			name: "value between multiples",
			n:    7,
			step: 3,
			want: 6,
		},
		{
			name: "positive n and step",
			n:    10,
			step: 3,
			want: 9,
		},
		{
			name: "negative exact multiple",
			n:    -6,
			step: 3,
			want: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mathutil.RoundDown(tt.n, tt.step)
			if got != tt.want {
				t.Errorf("RoundDown(%d, %d) = %d, want %d", tt.n, tt.step, got, tt.want)
			}
		})
	}
}
