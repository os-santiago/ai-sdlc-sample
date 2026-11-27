package mathutil

import "testing"

func TestSumOfEvens(t *testing.T) {
	// Test nil slice
	if got := SumOfEvens(nil); got != 0 {
		t.Errorf("SumOfEvens(nil) = %d, want 0", got)
	}

	// Test empty slice
	if got := SumOfEvens([]int{}); got != 0 {
		t.Errorf("SumOfEvens([]int{}) = %d, want 0", got)
	}

	// Test only odd values
	if got := SumOfEvens([]int{1, 3, 5, 7}); got != 0 {
		t.Errorf("SumOfEvens([]int{1, 3, 5, 7}) = %d, want 0", got)
	}

	// Test all even numbers
	if got := SumOfEvens([]int{2, 4, 6, 8}); got != 20 {
		t.Errorf("SumOfEvens([]int{2, 4, 6, 8}) = %d, want 20", got)
	}

	// Test mix including negatives
	if got := SumOfEvens([]int{-4, -3, -2, -1, 0, 1, 2, 3, 4}); got != 0 {
		t.Errorf("SumOfEvens([]int{-4, -3, -2, -1, 0, 1, 2, 3, 4}) = %d, want 0", got)
	}

	// Test single even number
	if got := SumOfEvens([]int{6}); got != 6 {
		t.Errorf("SumOfEvens([]int{6}) = %d, want 6", got)
	}

	// Test single odd number
	if got := SumOfEvens([]int{5}); got != 0 {
		t.Errorf("SumOfEvens([]int{5}) = %d, want 0", got)
	}

	// Test with zero
	if got := SumOfEvens([]int{0}); got != 0 {
		t.Errorf("SumOfEvens([]int{0}) = %d, want 0", got)
	}

	// Test negative evens only
	if got := SumOfEvens([]int{-2, -4, -6}); got != -12 {
		t.Errorf("SumOfEvens([]int{-2, -4, -6}) = %d, want -12", got)
	}
}
