// FilterValuePrefix returns a new map with the entries whose value starts with prefix.
package kv

import "strings"

// FilterValuePrefix returns a new map with the entries whose value starts with prefix.
func FilterValuePrefix(m map[string]string, prefix string) map[string]string {
	out := make(map[string]string)
	for k, v := range m {
		if strings.HasPrefix(v, prefix) {
			out[k] = v
		}
	}
	return out
}