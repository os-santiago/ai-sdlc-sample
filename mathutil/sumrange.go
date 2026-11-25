// Package mathutil provides small numeric helpers.
package mathutil

// SumRange returns the sum of the integers from a to b inclusive.
// When a > b, it returns 0.
func SumRange(a, b int) int {
	if a > b {
		return 0
	}
	
	// Use the arithmetic series formula: sum = n * (first + last) / 2
	// where n is the count of numbers: (b - a + 1)
	n := b - a + 1
	return n * (a + b) / 2
}