// Package mathutil provides small numeric helpers.
package mathutil

// SumOfEvens returns the sum of the even values in xs.
// It handles nil slices by returning 0.
func SumOfEvens(xs []int) int {
	if xs == nil {
		return 0
	}
	
	sum := 0
	for _, x := range xs {
		if IsEven(x) {
			sum += x
		}
	}
	return sum
}
