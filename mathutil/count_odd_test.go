package mathutil

import "testing"

func TestCountOdd(t *testing.T) {
	// Test nil slice
	if got := CountOdd(nil); got != 0 {
		t.Errorf("CountOdd(nil) = %d, want 0", got)
	}

	// Test empty slice
	if got := CountOdd([]int{}); got != 0 {
		t.Errorf("CountOdd([]int{}) = %d, want 0", got)
	}

	// Test all even numbers
	if got := CountOdd([]int{2, 4, 6, 8}); got != 0 {
		t.Errorf("CountOdd([]int{2, 4, 6, 8}) = %d, want 0", got)
	}

	// Test all odd numbers
	if got := CountOdd([]int{1, 3, 5, 7}); got != 4 {
		t.Errorf("CountOdd([]int{1, 3, 5, 7}) = %d, want 4", got)
	}

	// Test mix including negatives and zero
	if got := CountOdd([]int{-4, -3, -2, -1, 0, 1, 2, 3, 4}); got != 4 {
		t.Errorf("CountOdd([]int{-4, -3, -2, -1, 0, 1, 2, 3, 4}) = %d, want 4", got)
	}

	// Test single odd number
	if got := CountOdd([]int{5}); got != 1 {
		t.Errorf("CountOdd([]int{5}) = %d, want 1", got)
	}

	// Test single even number
	if got := CountOdd([]int{6}); got != 0 {
		t.Errorf("CountOdd([]int{6}) = %d, want 0", got)
	}

	// Test single zero
	if got := CountOdd([]int{0}); got != 0 {
		t.Errorf("CountOdd([]int{0}) = %d, want 0", got)
	}

	// Test negative odd values
	if got := CountOdd([]int{-1, -3, -5}); got != 3 {
		t.Errorf("CountOdd([]int{-1, -3, -5}) = %d, want 3", got)
	}
}
