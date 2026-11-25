// Package kv parses simple "key=[REDACTED]" lists.
package kv

import "strings"

// FilterPrefix returns a new map with only the keys starting with the given prefix.
// The input map is not modified.
func FilterPrefix(m map[string]string, prefix string) map[string]string {
	out := make(map[string]string)
	for k, v := range m {
		if strings.HasPrefix(k, prefix) {
			out[k] = v
		}
	}
	return out
}