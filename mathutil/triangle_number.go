// Package mathutil provides small numeric helpers.
package mathutil

// TriangleNumber returns the triangular number 1+2+...+n, or 0 for n <= 0.
func TriangleNumber(n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}