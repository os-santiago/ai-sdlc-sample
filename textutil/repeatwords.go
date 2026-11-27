package textutil

import "strings"

// RepeatWords repeats every whitespace-separated word n times, joined by single spaces.
// If n <= 0, it returns an empty string.
func RepeatWords(s string, n int) string {
	if n <= 0 {
		return ""
	}
	
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}
	
	var result []string
	for _, word := range words {
		for i := 0; i < n; i++ {
			result = append(result, word)
		}
	}
	
	return strings.Join(result, " ")
}
