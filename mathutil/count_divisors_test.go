package mathutil

import "testing"

func TestCountDivisors(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"negative", -5, 0},
		{"one", 1, 1},
		{"prime", 7, 2},
		{"twelve", 12, 6},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountDivisors(tt.n)
			if got != tt.want {
				t.Errorf("CountDivisors(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
