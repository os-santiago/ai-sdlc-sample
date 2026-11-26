package mathutil

import "math"

// IsPerfect reports whether n equals the sum of its positive divisors excluding itself.
// n <= 1 is not perfect.
func IsPerfect(n int) bool {
	if n <= 1 {
		return false
	}
	
	sum := 0
	sqrtN := int(math.Sqrt(float64(n)))
	
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			sum += i
			if i != 1 && i != n/i {
				sum += n / i
			}
		}
	}
	
	return sum == n
}
