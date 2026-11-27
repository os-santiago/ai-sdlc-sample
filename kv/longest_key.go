package kv

// LongestKey returns the longest key (the lexically smallest on ties) and true,
// or empty and false for an empty or nil map.
func LongestKey(m map[string]string) (string, bool) {
	if len(m) == 0 {
		return "", false
	}

	longestKey := ""
	maxLen := -1
	for k := range m {
		kLen := len(k)
		if kLen > maxLen || (kLen == maxLen && k < longestKey) {
			longestKey = k
			maxLen = kLen
		}
	}

	return longestKey, true
}
