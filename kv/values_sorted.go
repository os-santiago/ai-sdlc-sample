// Package kv parses simple "key=[REDACTED]" lists.
package kv

import "sort"

// ValuesSorted returns the distinct values of m in ascending order.
// An empty or nil map returns an empty, non-nil slice.
func ValuesSorted(m map[string]string) []string {
	if m == nil || len(m) == 0 {
		return []string{}
	}

	// Use a map to track distinct values
	seen := make(map[string]bool)
	for _, v := range m {
		seen[v] = true
	}

	// Extract unique values into a slice
	values := make([]string, 0, len(seen))
	for v := range seen {
		values = append(values, v)
	}

	// Sort and return
	sort.Strings(values)
	return values
}
