package mathutil

import "testing"

func TestCountInRange(t *testing.T) {
	// Test nil slice
	if got := CountInRange(nil, 1, 10); got != 0 {
		t.Errorf("CountInRange(nil, 1, 10) = %d, want 0", got)
	}

	// Test lo > hi
	if got := CountInRange([]int{1, 2, 3}, 10, 1); got != 0 {
		t.Errorf("CountInRange([]int{1, 2, 3}, 10, 1) = %d, want 0", got)
	}

	// Test empty slice
	if got := CountInRange([]int{}, 1, 10); got != 0 {
		t.Errorf("CountInRange([]int{}, 1, 10) = %d, want 0", got)
	}

	// Test mix of values inside and outside range
	if got := CountInRange([]int{0, 1, 2, 5, 10, 11, 15}, 2, 10); got != 3 {
		t.Errorf("CountInRange([]int{0, 1, 2, 5, 10, 11, 15}, 2, 10) = %d, want 3", got)
	}

	// Test all values inside range
	if got := CountInRange([]int{3, 4, 5, 6}, 2, 10); got != 4 {
		t.Errorf("CountInRange([]int{3, 4, 5, 6}, 2, 10) = %d, want 4", got)
	}

	// Test all values outside range
	if got := CountInRange([]int{1, 2, 11, 12}, 3, 10); got != 0 {
		t.Errorf("CountInRange([]int{1, 2, 11, 12}, 3, 10) = %d, want 0", got)
	}

	// Test single value in range
	if got := CountInRange([]int{5}, 1, 10); got != 1 {
		t.Errorf("CountInRange([]int{5}, 1, 10) = %d, want 1", got)
	}

	// Test single value out of range
	if got := CountInRange([]int{15}, 1, 10); got != 0 {
		t.Errorf("CountInRange([]int{15}, 1, 10) = %d, want 0", got)
	}

	// Test boundary values included
	if got := CountInRange([]int{1, 5, 10}, 1, 10); got != 3 {
		t.Errorf("CountInRange([]int{1, 5, 10}, 1, 10) = %d, want 3", got)
	}

	// Test negative values
	if got := CountInRange([]int{-5, -1, 0, 1, 5}, -2, 2); got != 3 {
		t.Errorf("CountInRange([]int{-5, -1, 0, 1, 5}, -2, 2) = %d, want 3", got)
	}
}
