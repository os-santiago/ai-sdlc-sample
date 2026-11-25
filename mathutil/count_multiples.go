// Package mathutil provides small numeric helpers.
package mathutil

// CountMultiples returns how many integers in [lo, hi] are multiples of k.
// It returns 0 when k <= 0 or lo > hi.
func CountMultiples(lo, hi, k int) int {
	if k <= 0 || lo > hi {
		return 0
	}

	// Find the first multiple of k >= lo
	firstMultiple := ((lo + k - 1) / k) * k
	
	// If the first multiple is beyond hi, there are no multiples
	if firstMultiple > hi {
		return 0
	}

	// Count how many multiples of k exist in [firstMultiple, hi]
	count := (hi-firstMultiple)/k + 1
	return count
}