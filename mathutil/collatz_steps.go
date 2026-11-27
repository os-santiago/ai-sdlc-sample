package mathutil

// CollatzSteps returns how many steps the Collatz rule (halve if even, else 3n+1)
// takes to reach 1 from n, and 0 for n <= 1.
func CollatzSteps(n int) int {
	if n <= 1 {
		return 0
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
