// Package mathutil provides small numeric helpers.
package mathutil

// IsBetween reports whether lo <= n <= hi.
// When lo > hi it returns false.
func IsBetween(n, lo, hi int) bool {
	if lo > hi {
		return false
	}
	return lo <= n && n <= hi
}
