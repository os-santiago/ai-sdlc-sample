package kv

// MinKey returns the lexically smallest key and true,
// or empty and false for an empty or nil map.
func MinKey(m map[string]string) (string, bool) {
	if len(m) == 0 {
		return "", false
	}

	minKey := ""
	first := true
	for k := range m {
		if first || k < minKey {
			minKey = k
			first = false
		}
	}

	return minKey, true
}
