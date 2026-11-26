package kv

import "sort"

// KeysWithValue returns the sorted keys whose value equals the given value.
// Returns an empty, non-nil slice when no keys match or when the map is nil.
func KeysWithValue(m map[string]string, value string) []string {
	if m == nil {
		return []string{}
	}

	keys := []string{}
	for k, v := range m {
		if v == value {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	return keys
}
