// Package kv parses simple "key=value" lists.
package kv

import "strings"

// CountPrefix counts the keys starting with the given prefix.
// An empty prefix counts every key.
func CountPrefix(m map[string]string, prefix string) int {
	count := 0
	for k := range m {
		if strings.HasPrefix(k, prefix) {
			count++
		}
	}
	return count
}
