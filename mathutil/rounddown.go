// Package mathutil provides small numeric helpers.
package mathutil

// RoundDown returns the largest multiple of step that is <= n.
// If step <= 0, returns n unchanged.
func RoundDown(n, step int) int {
	if step <= 0 {
		return n
	}
	q := n / step
	if n < 0 && n%step != 0 {
		q--
	}
	return q * step
}
