// Package mathutil provides small numeric helpers.
package mathutil

// CountInRange returns the number of values in xs that satisfy lo <= v <= hi.
// It returns 0 for nil slices or when lo > hi.
func CountInRange(xs []int, lo, hi int) int {
	if xs == nil || lo > hi {
		return 0
	}
	
	count := 0
	for _, v := range xs {
		if v >= lo && v <= hi {
			count++
		}
	}
	return count
}
