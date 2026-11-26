package kv

// Omit returns a new map without the given keys.
// The input map m is not modified.
func Omit(m map[string]string, keys ...string) map[string]string {
	// Create a set of keys to omit for efficient lookup
	omitSet := make(map[string]bool)
	for _, key := range keys {
		omitSet[key] = true
	}

	// Create output map with all keys except the ones to omit
	out := make(map[string]string)
	for k, v := range m {
		if !omitSet[k] {
			out[k] = v
		}
	}
	return out
}
