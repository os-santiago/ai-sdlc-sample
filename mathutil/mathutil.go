// Package mathutil provides small numeric helpers.
package mathutil

import ("math"; "sort")

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
	n := len(sorted)
	if n%2 == 0 {
		// Even length: average of the two middle values
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	// Odd length: return the middle value
	return sorted[n/2]
}

// Mean returns the arithmetic mean of xs, or 0 for an empty slice.
func Mean(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	var sum float64
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

// Sum returns the sum of all values in xs, or 0 for an empty slice.
func Sum(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	var sum float64
	for _, x := range xs {
		sum += x
	}
	return sum
}

// Product returns the product of all values in xs, or 1 for an empty slice.
func Product(xs []int) int {
	if len(xs) == 0 {
		return 1
	}
	product := 1
	for _, x := range xs {
		product *= x
	}
	return product
}

// Mode returns the most frequent value in xs and true, or 0, false for an empty slice.
// On a tie, it returns the smallest of the tied values.
func Mode(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}

	// Count frequencies
	freq := make(map[int]int)
	for _, x := range xs {
		freq[x]++
	}

	// Find the maximum frequency
	maxFreq := 0
	for _, count := range freq {
		if count > maxFreq {
			maxFreq = count
		}
	}

	// Find the smallest value with maximum frequency
	var mode int
	first := true
	for val, count := range freq {
		if count == maxFreq {
			if first {
				mode = val
				first = false
			} else if val < mode {
				mode = val
			}
		}
	}

	return mode, true
}

// StdDev returns the population standard deviation of xs, or 0 for an empty slice.
// It uses the formula: sqrt(sum((x - mean)^2) / n)
func StdDev(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	mean := Mean(xs)
	var sumSquaredDiff float64
	for _, x := range xs {
		diff := x - mean
		sumSquaredDiff += diff * diff
	}
	variance := sumSquaredDiff / float64(len(xs))
	return math.Sqrt(variance)
}

// MinMax returns the smallest and largest value in xs and true,
// or 0, 0, false for an empty slice.
func MinMax(xs []int) (min, max int, ok bool) {
	if len(xs) == 0 {
		return 0, 0, false
	}

	min = xs[0]
	max = xs[0]

	for _, x := range xs[1:] {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	return min, max, true
}

// Max returns the largest value in xs and true,
// or 0, false for an empty slice.
func Max(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}

	max := xs[0]
	for _, x := range xs[1:] {
		if x > max {
			max = x
		}
	}
	return max, true
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Range returns a slice of integers from start up to but not including end.
// When end <= start, it returns an empty, non-nil slice.
func Range(start, end int) []int {
	if end <= start {
		return []int{}
	}
	result := make([]int, 0, end-start)
	for i := start; i < end; i++ {
		result = append(result, i)
	}
	return result
}

// GCD returns the greatest common divisor of the absolute values of a and b.
// GCD(0, 0) returns 0.
func GCD(a, b int) int {
	// Use absolute values
	a = Abs(a)
	b = Abs(b)
	
	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Sign returns the sign of x: -1 for negative, 0 for zero, 1 for positive.
func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

// Pow returns base raised to the power of exp for non-negative integer exponents.
// It returns the result and true for exp >= 0.
// For negative exponents, it returns 0, false.
func Pow(base, exp int) (int, bool) {
	if exp < 0 {
		return 0, false
	}
	
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	
	return result, true
}
