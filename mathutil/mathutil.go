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

// LCM returns the least common multiple of the absolute values of a and b.
// LCM(0, x) returns 0 for any x.
func LCM(a, b int) int {
	// LCM(0, x) = 0 for any x
	if a == 0 || b == 0 {
		return 0
	}
	// LCM(a, b) = |a * b| / GCD(a, b)
	abs_a, abs_b := Abs(a), Abs(b)
	return (abs_a * abs_b) / GCD(a, b)
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

// AverageInt returns the floor of the arithmetic mean of xs and true,
// or 0, false for an empty slice.
// The floor function rounds towards negative infinity, so:
// AverageInt([]int{1, 2, 4}) returns 2 (floor of 2.333...)
// AverageInt([]int{-3, -4}) returns -4 (floor of -3.5)
func AverageInt(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	
	sum := 0
	for _, x := range xs {
		sum += x
	}
	
	// Floor division: adjust for negative remainders
	avg := sum / len(xs)
	if sum%len(xs) != 0 && sum < 0 {
		avg--
	}
	
	return avg, true
}

// AverageFloat returns the arithmetic mean of values and true,
// or 0, false for an empty slice.
func AverageFloat(values []float64) (float64, bool) {
	if len(values) == 0 {
		return 0, false
	}
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values)), true
}

// IsPrime returns true if n is a prime number, false otherwise.
// A prime number is a natural number greater than 1 that has no positive divisors other than 1 and itself.
// Values below 2 (including negative numbers, 0, and 1) are not prime.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Factorial returns n! (n factorial) for small non-negative integers.
// It returns the factorial and true for 0 <= n <= 20.
// For negative n or n > 20, it returns 0, false to prevent integer overflow.
//
// Examples:
//   Factorial(0) = 1, true
//   Factorial(5) = 120, true
//   Factorial(20) = 2432902008176640000, true
//   Factorial(21) = 0, false (overflow guard)
//   Factorial(-1) = 0, false (negative input)
func Factorial(n int) (int, bool) {
	if n < 0 || n > 20 {
		return 0, false
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, true
}

// Digits returns the number of decimal digits of n, ignoring the sign.
// Digits(0) returns 1; Digits(-123) returns 3.
func Digits(n int) int {
	if n == 0 {
		return 1
	}
	
	// Use absolute value to ignore sign
	if n < 0 {
		n = -n
	}
	
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	
	return count
}

// SumOfSquares returns the sum of each value squared in values, or 0 for an empty slice.
// Examples:
//   SumOfSquares([]int{1, 2, 3}) = 1 + 4 + 9 = 14
//   SumOfSquares([]int{-2, 3}) = 4 + 9 = 13
//   SumOfSquares([]int{}) = 0
func SumOfSquares(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sum := 0
	for _, v := range values {
		sum += v * v
	}
	return sum
}

// IsEven returns true if n is even, false otherwise.
// A number is even if it is divisible by 2.
// This includes zero and negative even numbers.
// Examples:
//   IsEven(0) = true
//   IsEven(2) = true
//   IsEven(-2) = true
//   IsEven(1) = false
//   IsEven(-1) = false
func IsEven(n int) bool {
	return n%2 == 0
}
// IsOdd returns true if n is odd, false otherwise.
// A number is odd if it is not divisible by 2.
// This includes negative odd numbers but not zero.
// Examples:
//   IsOdd(1) = true
//   IsOdd(-1) = true
//   IsOdd(0) = false
//   IsOdd(2) = false
func IsOdd(n int) bool {
	return n%2 != 0
}

// Square returns n squared (n*n).
// Examples:
//   Square(5) = 25
//   Square(-3) = 9
//   Square(0) = 0
func Square(n int) int {
	return n * n
}

// Half returns n/2 rounded toward zero.
// Examples:
//   Half(5) = 2
//   Half(-5) = -2
//   Half(4) = 2
//   Half(0) = 0
func Half(n int) int {
	return n / 2
}
// SignOf returns the sign of n: -1 for negative, 0 for zero, 1 for positive.
func SignOf(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}

// DigitSum returns the sum of the decimal digits of the absolute value of n.
// Examples:
//   DigitSum(0) = 0
//   DigitSum(123) = 6 (1 + 2 + 3)
//   DigitSum(-123) = 6 (1 + 2 + 3)
func DigitSum(n int) int {
	// Use absolute value to ignore sign
	if n < 0 {
		n = -n
	}
	
	// Special case for 0
	if n == 0 {
		return 0
	}
	
	// Sum the digits
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	
	return sum
}