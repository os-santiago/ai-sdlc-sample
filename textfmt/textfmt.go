// Package textfmt formats short texts for display.
package textfmt

import "strings"

// Truncate shortens s to at most n characters, ending with "..." when it
// had to cut. n smaller than 4 returns s cut to n characters.
func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	if n < 4 {
		return string(runes[:n])
	}
	return string(runes[:n-3]) + "..."
}

// Title upper-cases the first letter of each space-separated word.
func Title(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		if w != "" {
			runes := []rune(w)
			words[i] = strings.ToUpper(string(runes[:1])) + string(runes[1:])
		}
	}
	return strings.Join(words, " ")
}

// PadLeft prepends pad runes to s until it reaches width characters.
// Width counts characters (runes), not bytes. If s already has width
// or more characters, it is returned unchanged.
func PadLeft(s string, width int, pad rune) string {
	runes := []rune(s)
	if len(runes) >= width {
		return s
	}
	needed := width - len(runes)
	padding := strings.Repeat(string(pad), needed)
	return padding + s
}

// PadRight appends pad runes to s until it reaches width characters.
// Width counts characters (runes), not bytes. If s already has width
// or more characters, it is returned unchanged.
func PadRight(s string, width int, pad rune) string {
	runes := []rune(s)
	if len(runes) >= width {
		return s
	}
	needed := width - len(runes)
	padding := strings.Repeat(string(pad), needed)
	return s + padding
}

// Center pads s on both sides to width characters using pad. Width counts
// characters (runes), not bytes. When the total padding is odd, the extra
// character goes on the right. If s already has width or more characters,
// it is returned unchanged.
func Center(s string, width int, pad rune) string {
	runes := []rune(s)
	if len(runes) >= width {
		return s
	}
	needed := width - len(runes)
	leftPad := needed / 2
	rightPad := needed - leftPad
	return strings.Repeat(string(pad), leftPad) + s + strings.Repeat(string(pad), rightPad)
}

// Repeat joins n copies of s with sep as the separator between them.
// If n is 0 or less, it returns an empty string.
func Repeat(s, sep string, n int) string {
	if n <= 0 {
		return ""
	}
	if n == 1 {
		return s
	}
	var result strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			result.WriteString(sep)
		}
		result.WriteString(s)
	}
	return result.String()
}

// Ellipsis shortens s to at most n characters for display, adding an ellipsis
// when it had to cut. n of 0 or less returns an empty string. When s has at
// most n characters (runes), it is returned unchanged. Otherwise the first
// n-1 characters are returned followed by "…".
func Ellipsis(s string, n int) string {
	if n <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n-1]) + "…"
}

// TruncateDots shortens s to at most max runes, ending with "…" when it
// had to cut. max smaller than 1 returns an empty string. When s has at
// most max runes (runes), it is returned unchanged. Otherwise the first
// max-1 runes are returned followed by "…".
func TruncateDots(s string, max int) string {
	if max <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= max {
		return s
	}
	return string(runes[:max-1]) + "…"
}

// Indent adds prefix at the start of every non-empty line in s.
// Empty lines stay empty, and line endings are kept.
func Indent(s, prefix string) string {
	if s == "" {
		return ""
	}

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

// Wrap breaks s into lines of at most width characters at spaces, joined with \n.
// A single word longer than width stays on its own line. Width of 0 or less
// returns s unchanged.
func Wrap(s string, width int) string {
	if width <= 0 {
		return s
	}

	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}

	var result strings.Builder
	lineLen := 0

	for _, word := range words {
		wordLen := len([]rune(word))

		if lineLen == 0 {
			// Start of a new line
			result.WriteString(word)
			lineLen = wordLen
		} else if lineLen+1+wordLen <= width {
			// Word fits on current line
			result.WriteString(" ")
			result.WriteString(word)
			lineLen += 1 + wordLen
		} else {
			// Word doesn't fit, start a new line
			result.WriteString("\n")
			result.WriteString(word)
			lineLen = wordLen
		}
	}

	return result.String()
}

// RepeatJoin joins n copies of s with sep as the separator between them.
// If n is 0 or less, it returns an empty string.
func RepeatJoin(s, sep string, n int) string {
	if n <= 0 {
		return ""
	}
	if n == 1 {
		return s
	}
	var result strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			result.WriteString(sep)
		}
		result.WriteString(s)
	}
	return result.String()
}

// Quote wraps s in double quotes and escapes inner double quotes and backslashes
// with a backslash. The result is suitable for display in quoted contexts.
func Quote(s string) string {
	var result strings.Builder
	result.WriteByte('"')
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' || s[i] == '"' {
			result.WriteByte('\\')
		}
		result.WriteByte(s[i])
	}
	result.WriteByte('"')
	return result.String()
}

// Surround wraps s with marker on both sides. If marker is empty, s is
// returned unchanged.
func Surround(s, marker string) string {
	if marker == "" {
		return s
	}
	return marker + s + marker
}

// Bullet formats items as a bulleted list with each item on its own line
// prefixed with "- ", joined with "\n", with no trailing newline.
// An empty slice returns an empty string.
func Bullet(items []string) string {
	if len(items) == 0 {
		return ""
	}
	var result strings.Builder
	for i, item := range items {
		if i > 0 {
			result.WriteString("\n")
		}
		result.WriteString("- ")
		result.WriteString(item)
	}
	return result.String()
}

// Underline returns s, a newline, and one hyphen per rune of s.
func Underline(s string) string {
	runes := []rune(s)
	underline := make([]rune, len(runes))
	for i := range underline {
		underline[i] = '-'
	}
	return s + "\n" + string(underline)
}

// Dedent removes the longest run of leading spaces shared by all non-blank lines.
// Blank lines are kept and do not affect the shared indentation.
func Dedent(s string) string {
	if s == "" {
		return ""
	}

	lines := strings.Split(s, "\n")

	// Find minimum indentation among non-blank lines
	minIndent := -1
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			// Skip blank lines
			continue
		}

		// Count leading spaces
		indent := 0
		for _, ch := range line {
			if ch == ' ' {
				indent++
			} else {
				break
			}
		}

		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}

	// If no non-blank lines or no indentation, return original
	if minIndent <= 0 {
		return s
	}

	// Remove minIndent spaces from each line
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" && len(line) >= minIndent {
			lines[i] = line[minIndent:]
		}
	}

	return strings.Join(lines, "\n")
}

// TrimRunes removes leading and trailing runes from s that appear in cutset.
// If cutset is empty, s is returned unchanged. All runes in cutset are
// considered as a set, and any rune at the start or end of s that appears
// in cutset is removed.
func TrimRunes(s, cutset string) string {
	if cutset == "" {
		return s
	}

	runes := []rune(s)
	cutRunes := []rune(cutset)

	// Trim from start
	start := 0
	for start < len(runes) && contains(cutRunes, runes[start]) {
		start++
	}

	// Trim from end
	end := len(runes)
	for end > start && contains(cutRunes, runes[end-1]) {
		end--
	}

	return string(runes[start:end])
}

// contains checks if a rune is in a slice of runes.
func contains(runes []rune, r rune) bool {
	for _, cr := range runes {
		if cr == r {
			return true
		}
	}
	return false
}

// Indent2 prefixes every non-empty line of s with two spaces; empty lines stay empty.
// Line endings are preserved.
func Indent2(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	atLineStart := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if atLineStart && ch != '\n' && ch != '\r' {
			result.WriteString("  ")
			atLineStart = false
		}

		result.WriteByte(ch)

		if ch == '\n' {
			atLineStart = true
		}
	}

	return result.String()
}

// Quoted prefixes every line of s with "> "; an empty string returns an empty string.
func Quoted(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	result.WriteString("> ")

	for i := 0; i < len(s); i++ {
		ch := s[i]
		result.WriteByte(ch)

		if ch == '\n' && i < len(s)-1 {
			result.WriteString("> ")
		}
	}

	return result.String()
}
