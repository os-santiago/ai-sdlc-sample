// Package kv parses simple "key=[REDACTED]" lists.
package kv

// Prefixed returns a new map with prefix added to every key of m.
// The input map m is not modified.
func Prefixed(m map[string]string, prefix string) map[string]string {
	if m == nil {
		return nil
	}
	
	out := make(map[string]string, len(m))
	for k, v := range m {
		out[prefix+k] = v
	}
	return out
}