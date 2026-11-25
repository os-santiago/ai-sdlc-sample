// Package mathutil provides small numeric helpers.
package mathutil

// IsMultiple reports whether n is a multiple of d.
// d == 0 returns false.
func IsMultiple(n, d int) bool {
	if d == 0 {
		return false
	}
	return n%d == 0
}