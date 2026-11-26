// Package mathutil provides small numeric helpers.
package mathutil

// Average3 returns the mean of the three values.
func Average3(a, b, c int) float64 {
	return float64(a+b+c) / 3.0
}
