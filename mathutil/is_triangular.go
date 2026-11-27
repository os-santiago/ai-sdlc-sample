// Package mathutil provides small numeric helpers.
package mathutil

import "math"

// IsTriangular reports whether n equals k*(k+1)/2 for some integer k >= 1.
// Triangular numbers are 1, 3, 6, 10, 15, 21, ...
func IsTriangular(n int) bool {
	if n < 1 {
		return false
	}
	
	// For a triangular number n = k*(k+1)/2, we can solve for k:
	// k^2 + k - 2*n = 0
	// Using the quadratic formula: k = (-1 + sqrt(1 + 8*n)) / 2
	// We only need to check if k is a positive integer
	
	// Calculate discriminant
	discriminant := 1 + 8*n
	
	// Check if discriminant is a perfect square
	sqrtDiscriminant := int(math.Sqrt(float64(discriminant)))
	if sqrtDiscriminant*sqrtDiscriminant != discriminant {
		return false
	}
	
	// Check if k = (-1 + sqrtDiscriminant) / 2 is a positive integer
	k := (-1 + sqrtDiscriminant) / 2
	return k >= 1 && (-1+sqrtDiscriminant)%2 == 0
}