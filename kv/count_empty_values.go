package kv

// CountEmptyValues counts the keys whose value is an empty string.
func CountEmptyValues(m map[string]string) int {
	count := 0
	for _, v := range m {
		if v == "" {
			count++
		}
	}
	return count
}
