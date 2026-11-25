// Package textutil provides small text helpers.
package textutil

// CommonPrefix returns the longest common prefix of a and b, compared by rune.
func CommonPrefix(a, b string) string {
	// Convert strings to rune slices for proper Unicode handling
	runesA := []rune(a)
	runesB := []rune(b)
	
	// Find the minimum length to avoid index out of bounds
	minLen := len(runesA)
	if len(runesB) < minLen {
		minLen = len(runesB)
	}
	
	// Find the common prefix
	var prefix []rune
	for i := 0; i < minLen; i++ {
		if runesA[i] == runesB[i] {
			prefix = append(prefix, runesA[i])
		} else {
			break
		}
	}
	
	return string(prefix)
}