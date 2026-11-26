package mathutil

// Steps returns the number of Collatz steps to reach 1 from n.
// For n <= 0, it returns -1.
func Steps(n int) int {
	if n <= 0 {
		return -1
	}

	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}

	return steps
}
