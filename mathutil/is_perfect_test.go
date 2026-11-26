package mathutil

import "testing"

func TestIsPerfect(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, false},   // n <= 1 is not perfect
		{1, false},   // n <= 1 is not perfect
		{6, true},    // 6 = 1 + 2 + 3
		{28, true},   // 28 = 1 + 2 + 4 + 7 + 14
		{12, false},  // 12 != 1 + 2 + 3 + 4 + 6 = 16
	}

	for _, tt := range tests {
		got := IsPerfect(tt.n)
		if got != tt.want {
			t.Errorf("IsPerfect(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
