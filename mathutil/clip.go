// Package mathutil provides small numeric helpers.
package mathutil

// Clip returns a new slice with each value clamped to the closed range [lo, hi].
// The input slice is not modified.
// For a nil or empty slice, it returns nil.
func Clip(values []int, lo, hi int) []int {
	if values == nil {
		return nil
	}
	if len(values) == 0 {
		return []int{}
	}

	result := make([]int, len(values))
	for i, v := range values {
		result[i] = Clamp(v, lo, hi)
	}
	return result
}
