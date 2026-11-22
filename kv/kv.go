// Package kv parses simple "key=value" lists.
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

// Format renders a map as a deterministic "key=value;key=value" string.
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