package textfmt

import "strings"

// Indented prefixes every non-empty line of s with n spaces; n <= 0 returns s.
func Indented(s string, n int) string {
	if n <= 0 {
		return s
	}

	if s == "" {
		return ""
	}

	prefix := strings.Repeat(" ", n)

	var result strings.Builder
	atLineStart := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if atLineStart && ch != '\n' && ch != '\r' {
			result.WriteString(prefix)
			atLineStart = false
		}

		result.WriteByte(ch)

		if ch == '\n' {
			atLineStart = true
		}
	}

	return result.String()
}
