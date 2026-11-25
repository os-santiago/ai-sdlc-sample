package mathutil

import "testing"

func TestCountEven(t *testing.T) {
	// Test nil slice
	if got := CountEven(nil); got != 0 {
		t.Errorf("CountEven(nil) = %d, want 0", got)
	}

	// Test empty slice
	if got := CountEven([]int{}); got != 0 {
		t.Errorf("CountEven([]int{}) = %d, want 0", got)
	}

	// Test all odd numbers
	if got := CountEven([]int{1, 3, 5, 7}); got != 0 {
		t.Errorf("CountEven([]int{1, 3, 5, 7}) = %d, want 0", got)
	}

	// Test all even numbers
	if got := CountEven([]int{2, 4, 6, 8}); got != 4 {
		t.Errorf("CountEven([]int{2, 4, 6, 8}) = %d, want 4", got)
	}

	// Test mix including negatives and zero
	if got := CountEven([]int{-4, -3, -2, -1, 0, 1, 2, 3, 4}); got != 5 {
		t.Errorf("CountEven([]int{-4, -3, -2, -1, 0, 1, 2, 3, 4}) = %d, want 5", got)
	}

	// Test single even number
	if got := CountEven([]int{5}); got != 0 {
		t.Errorf("CountEven([]int{5}) = %d, want 0", got)
	}

	// Test single odd number
	if got := CountEven([]int{6}); got != 1 {
		t.Errorf("CountEven([]int{6}) = %d, want 1", got)
	}

	// Test single zero
	if got := CountEven([]int{0}); got != 1 {
		t.Errorf("CountEven([]int{0}) = %d, want 1", got)
	}
}