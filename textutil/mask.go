package textutil

// Mask replaces every rune except the last keep runes with *.
// If keep <= 0, all runes are masked.
// If keep >= the number of runes in s, the original string is returned.
func Mask(s string, keep int) string {
	runes := []rune(s)
	length := len(runes)

	// If keep >= length, return the original string
	if keep >= length {
		return s
	}

	// Calculate how many characters to mask
	maskCount := length - keep
	if keep <= 0 {
		maskCount = length
	}

	// Replace the first maskCount runes with *
	for i := 0; i < maskCount; i++ {
		runes[i] = '*'
	}

	return string(runes)
}
