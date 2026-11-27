// Package kv parses simple "key=[REDACTED]" lists.
package kv

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// ErrSyntax reports a malformed pair.
var ErrSyntax = errors.New("kv: malformed pair")

// Parse reads pairs separated by ";" such as "a=1;b=2" into a map.
func Parse(s string) (map[string]string, error) {
	out := map[string]string{}
	for _, pair := range strings.Split(s, ";") {
		if pair == "" {
			continue
		}
		k, v, ok := strings.Cut(strings.TrimSpace(pair), "=")
		k, v = strings.TrimSpace(k), strings.TrimSpace(v)
		if !ok {
			return nil, ErrSyntax
		}
		out[k] = v
	}
	return out, nil
}

// Format renders a map as a deterministic "key=[REDACTED];key=[REDACTED]" string.
// Keys are sorted alphabetically. Returns empty string for empty map.
func Format(m map[string]string) string {
	if len(m) == 0 {
		return ""
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	pairs := make([]string, len(keys))
	for i, k := range keys {
		pairs[i] = k + "=" + m[k]
	}
	return strings.Join(pairs, ";")
}

// ParseInts reads pairs separated by ";" and converts values to integers.
func ParseInts(s string) (map[string]int, error) {
	out := map[string]int{}
	for _, pair := range strings.Split(s, ";") {
		if pair == "" {
			continue
		}
		k, v, ok := strings.Cut(strings.TrimSpace(pair), "=")
		k, v = strings.TrimSpace(k), strings.TrimSpace(v)
		if !ok {
			return nil, ErrSyntax
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, ErrSyntax
		}
		out[k] = n
	}
	return out, nil
}

// Keys returns the keys of the map sorted in ascending order.
// Returns an empty, non-nil slice for a nil map.
func Keys(m map[string]string) []string {
	if m == nil {
		return []string{}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Values returns the values of the map ordered by their keys ascending.
// Returns an empty, non-nil slice for a nil map.
func Values(m map[string]string) []string {
	if m == nil {
		return []string{}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	values := make([]string, len(keys))
	for i, k := range keys {
		values[i] = m[k]
	}
	return values
}

// Merge returns a new map with every key of a and b.
// When a key is in both, the value from b wins (right-hand precedence).
// Neither input map is modified.
func Merge(a, b map[string]string) map[string]string {
	out := make(map[string]string, len(a)+len(b))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

// MergeRight returns a new map with every key of a and b.
// When a key is in both, the value from b wins (right-hand precedence).
// Neither input map is modified.
func MergeRight(a, b map[string]string) map[string]string {
	out := make(map[string]string, len(a)+len(b))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

// Delete returns a new map with all entries of m except the given key.
// The input map is not modified. If key is not in m, returns an equal copy.
func Delete(m map[string]string, key string) map[string]string {
	out := make(map[string]string, len(m))
	for k, v := range m {
		if k != key {
			out[k] = v
		}
	}
	return out
}

// Equal returns true when both maps have the same keys and values.
// A nil map equals an empty map.
func Equal(a, b map[string]string) bool {
	// Treat nil maps as empty
	lenA := len(a)
	lenB := len(b)
	if lenA != lenB {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}
	return true
}

// EqualMaps returns true when both maps have the same keys and values.
// A nil map equals an empty map.
func EqualMaps(a, b map[string]string) bool {
	// Delegate to Equal for consistent behavior
	return Equal(a, b)
}

// Filter returns a new map with only the keys starting with the given prefix.
// The input map is not modified.
func Filter(m map[string]string, prefix string) map[string]string {
	out := make(map[string]string)
	for k, v := range m {
		if strings.HasPrefix(k, prefix) {
			out[k] = v
		}
	}
	return out
}

// Get returns the value for key in m, or fallback when the key is missing.
// A present key with an empty value returns the empty value, not the fallback.
func Get(m map[string]string, key, fallback string) string {
	if v, ok := m[key]; ok {
		return v
	}
	return fallback
}

// Has reports whether key exists in the map.
// Returns true even when the key has an empty value.
func Has(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}

// Invert returns a new map with keys and values swapped.
// When all values are unique, it returns the inverted map and true.
// When two or more keys share the same value, it returns nil, false.
func Invert(m map[string]string) (map[string]string, bool) {
	out := make(map[string]string, len(m))
	for k, v := range m {
		if _, exists := out[v]; exists {
			return nil, false
		}
		out[v] = k
	}
	return out, true
}

// Rename renames a key in the map from "from" to "to".
// It returns true when the rename succeeds.
// It returns false and leaves the map unchanged when:
// - the "from" key is missing, or
// - the "to" key already exists.
func Rename(m map[string]string, from, to string) bool {
	// Check if "from" key exists
	value, fromExists := m[from]
	if !fromExists {
		return false
	}
	
	// Check if "to" key already exists
	_, toExists := m[to]
	if toExists {
		return false
	}
	
	// Perform the rename
	delete(m, from)
	m[to] = value
	return true
}

// Count returns the number of keys whose value equals the given value.
func Count(m map[string]string, value string) int {
	count := 0
	for _, v := range m {
		if v == value {
			count++
		}
	}
	return count
}

// ByteSize returns the total number of bytes in all keys and values of the map.
// A nil map returns 0.
func ByteSize(m map[string]string) int {
	if m == nil {
		return 0
	}
	
	total := 0
	for k, v := range m {
		total += len(k) + len(v)
	}
	return total
}

// SortedKeys returns the keys of the map sorted in ascending order.
// Returns an empty, non-nil slice for a nil map.
func SortedKeys(m map[string]string) []string {
	if m == nil {
		return []string{}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// HasAny reports whether at least one of keys is in m.
func HasAny(m map[string]string, keys ...string) bool {
	if m == nil {
		return false
	}
	for _, key := range keys {
		if _, ok := m[key]; ok {
			return true
		}
	}
	return false
}