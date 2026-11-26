package kv

// Pick returns a new map with only the given keys that exist in m.
// Keys that don't exist in m are ignored.
func Pick(m map[string]string, keys ...string) map[string]string {
	out := make(map[string]string)
	for _, key := range keys {
		if v, ok := m[key]; ok {
			out[key] = v
		}
	}
	return out
}
