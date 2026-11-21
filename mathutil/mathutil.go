// Package mathutil provides small numeric helpers.
package mathutil

import "sort"

// Clamp limits x to the closed range [lo, hi].
func Clamp(x, lo, hi int) int {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}

// Median returns the median of xs, or 0 for an empty slice. xs is not modified.
func Median(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	sorted := append([]float64(nil), xs...)
	sort.Float64s(sorted)
	return sorted[len(sorted)/2]
}

// Mean returns the arithmetic mean of xs, or 0 for an empty slice.
func Mean(xs []float64) float64 {
	var sum float64
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}
