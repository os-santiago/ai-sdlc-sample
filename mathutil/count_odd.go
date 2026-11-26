// Package mathutil provides small numeric helpers.
package mathutil

// CountOdd returns the number of odd values in values.
// It handles nil slices by returning 0.
func CountOdd(values []int) int {
	if values == nil {
		return 0
	}

	count := 0
	for _, v := range values {
		if IsOdd(v) {
			count++
		}
	}
	return count
}
