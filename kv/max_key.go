package kv

// MaxKey returns the lexically largest key and true,
// or empty and false for an empty or nil map.
func MaxKey(m map[string]string) (string, bool) {
	if len(m) == 0 {
		return "", false
	}

	maxKey := ""
	first := true
	for k := range m {
		if first || k > maxKey {
			maxKey = k
			first = false
		}
	}

	return maxKey, true
}
