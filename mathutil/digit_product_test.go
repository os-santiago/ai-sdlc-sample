package mathutil

import "testing"

func TestDigitProduct(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "single digit",
			n:    5,
			want: 5,
		},
		{
			name: "negative number",
			n:    -234,
			want: 24,
		},
		{
			name: "234",
			n:    234,
			want: 24,
		},
		{
			name: "multiple digits with zero",
			n:    105,
			want: 0,
		},
		{
			name: "large number",
			n:    123,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DigitProduct(tt.n)
			if got != tt.want {
				t.Errorf("DigitProduct(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
