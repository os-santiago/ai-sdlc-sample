package mathutil

// Gap returns the largest absolute difference between consecutive values.
// If there are fewer than two values, it returns 0.
func Gap(values []int) int {
	if len(values) < 2 {
		return 0
	}

	maxGap := 0
	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff < 0 {
			diff = -diff
		}
		if diff > maxGap {
			maxGap = diff
		}
	}

	return maxGap
}
