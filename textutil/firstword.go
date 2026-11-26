package textutil

import "strings"

// FirstWord returns the first whitespace-separated word in s,
// or the empty string when there are none.
func FirstWord(s string) string {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return ""
	}
	return fields[0]
}