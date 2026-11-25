package mathutil

import "testing"

func TestDistance(t *testing.T) {
	// Test a < b
	if got := Distance(3, 7); got != 4 {
		t.Errorf("Distance(3, 7) = %d, want 4", got)
	}

	// Test a > b
	if got := Distance(7, 3); got != 4 {
		t.Errorf("Distance(7, 3) = %d, want 4", got)
	}

	// Test a == b
	if got := Distance(5, 5); got != 0 {
		t.Errorf("Distance(5, 5) = %d, want 0", got)
	}
}