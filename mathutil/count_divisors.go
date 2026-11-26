// Package mathutil provides small numeric helpers.
package mathutil

import "math"

// CountDivisors returns the number of positive divisors of n.
// It returns 0 for n <= 0.
func CountDivisors(n int) int {
	if n <= 0 {
		return 0
	}
	
	if n == 1 {
		return 1
	}
	
	count := 0
	sqrtN := int(math.Sqrt(float64(n)))
	
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			count++
			if i != n/i {
				count++
			}
		}
	}
	
	return count
}
