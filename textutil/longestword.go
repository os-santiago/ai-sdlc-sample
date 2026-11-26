package textutil

import "strings"

// LongestWord returns the first whitespace-separated word in s with the most runes,
// or the empty string when there are none.
func LongestWord(s string) string {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return ""
	}
	
	longest := fields[0]
	maxLen := len([]rune(longest))
	
	for i := 1; i < len(fields); i++ {
		currentLen := len([]rune(fields[i]))
		if currentLen > maxLen {
			longest = fields[i]
			maxLen = currentLen
		}
	}
	
	return longest
}
