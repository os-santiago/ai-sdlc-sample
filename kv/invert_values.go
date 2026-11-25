// Package kv parses simple "key=[REDACTED]" lists.
package kv

// InvertValues returns a map from each value to its key.
// When values repeat, the smallest key (in lexicographic order) wins.
// A nil map returns an empty map.
func InvertValues(m map[string]string) map[string]string {
	if m == nil {
		return map[string]string{}
	}

	// Create a map to track which key has the smallest value for each value
	result := make(map[string]string)
	
	// For each key-value pair in the input map
	for k, v := range m {
		// If this value hasn't been seen before, or if this key is smaller than the existing one
		if existingKey, exists := result[v]; !exists || k < existingKey {
			result[v] = k
		}
	}
	
	return result
}