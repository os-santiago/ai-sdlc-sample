// Package mathutil provides small numeric helpers.
package mathutil

// NextEven returns the smallest even number strictly greater than n.
// Examples:
//   NextEven(0) = 2
//   NextEven(1) = 2
//   NextEven(2) = 4
//   NextEven(-1) = 0
//   NextEven(-2) = 0
func NextEven(n int) int {
	if n%2 == 0 {
		// n is even, return n + 2
		return n + 2
	}
	// n is odd, return n + 1
	return n + 1
}