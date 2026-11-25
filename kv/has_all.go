package kv

// HasAll reports whether every key in keys is present in m.
// Returns true when no keys are provided.
func HasAll(m map[string]string, keys ...string) bool {
	// Handle nil map case
	if m == nil {
		return false
	}
	
	// Check each key
	for _, key := range keys {
		if _, ok := m[key]; !ok {
			return false
		}
	}
	
	return true
}