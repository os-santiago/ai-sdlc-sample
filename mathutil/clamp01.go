// Package mathutil provides small numeric helpers.
package mathutil

// Clamp01 limits x to the closed range [0, 1].
// It returns 0 for x < 0, 1 for x > 1, and x otherwise.
func Clamp01(x float64) float64 {
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return x
}