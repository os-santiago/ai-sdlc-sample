package mathutil

// NextMultiple returns the smallest multiple of step that is >= n.
// When step <= 0, it returns n unchanged.
//
// Examples:
//   NextMultiple(7, 3) = 9
//   NextMultiple(6, 3) = 6
//   NextMultiple(-7, 3) = -6
//   NextMultiple(5, 0) = 5
//   NextMultiple(5, -3) = 5
func NextMultiple(n, step int) int {
	if step <= 0 {
		return n
	}
	
	remainder := n % step
	if remainder == 0 {
		return n
	}
	
	// For negative remainder (when n is negative), add the absolute value
	// For positive remainder, move to the next multiple
	if remainder < 0 {
		return n - remainder
	}
	return n + step - remainder
}
