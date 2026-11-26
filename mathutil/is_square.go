// Package mathutil provides small numeric helpers.
package mathutil

import "math"

// IsSquare reports whether n is a perfect square.
// Negative values return false.
func IsSquare(n int) bool {
	if n < 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}
