// Package mathutil provides small numeric helpers.
package mathutil

// CountEven returns the number of even values in values.
// It handles nil slices by returning 0.
func CountEven(values []int) int {
	if values == nil {
		return 0
	}
	
	count := 0
	for _, v := range values {
		if IsEven(v) {
			count++
		}
	}
	return count
}