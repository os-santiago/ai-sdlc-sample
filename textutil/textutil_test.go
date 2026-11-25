package textutil

import "testing"

func TestReverseASCII(t *testing.T) {
	for in, want := range map[string]string{"" : "", "a": "a", "abc": "cba", "ai-sdlc": "clds-ia"} {
		if got := Reverse(in); got != want {
			t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestReverseMultibyte(t *testing.T) {
	for in, want := range map[string]string{
		"héllo":   "olléh",
		"日本語": "語本日",
	} {
		if got := Reverse(in); got != want {
			t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Existing test cases
		{"", 0},
		{"one", 1},
		{" two  words ", 2},
		// New test cases for punctuation-only tokens
		{"hello , world !", 2},
		{"one two three", 3},
		// Edge cases with punctuation
		{"!", 0},
		{"!!!", 0},
		{"   ,   ", 0},
		{"a , b , c", 3},
		{"word!", 1},
		{"!word", 1},
		{"wo!rd", 1},
		// Numbers count as words
		{"123", 1},
		{"abc123", 1},
		{"123abc", 1},
		// Mixed punctuation and words
		{"hello, world!", 2},
		{"   ,   !   ?   ", 0},
	}

	for _, tt := range tests {
		if got := WordCount(tt.in); got != tt.want {
			t.Errorf("WordCount(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Empty string
		{"", true},
		// Single character
		{"a", true},
		// Simple palindromes
		{"aba", true},
		{"abba", true},
		// Not palindromes
		{"abc", false},
		{"ab", false},
		// Case and space handling
		{"Never odd or even", true},
		{"A man a plan a canal Panama", true},
		// Unicode characters
		{"ñaña", false},
		{"añña", true},
	}

	for _, tt := range tests {
		if got := IsPalindrome(tt.in); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestInitials(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Empty string
		{"", ""},
		// Basic case
		{"ada lovelace", "AL"},
		// Extra whitespace
		{"  ada   lovelace  ", "AL"},
		{"one  two   three", "OTT"},
		// Unicode awareness
		{"élan vital", "ÉV"},
		{"über alles", "ÜA"},
		// Single word
		{"hello", "H"},
		// Mixed case input (should uppercase)
		{"hello World", "HW"},
		{"HELLO world", "HW"},
	}

	for _, tt := range tests {
		if got := Initials(tt.in); got != tt.want {
			t.Errorf("Initials(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestAbbreviate(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: basic case
		{"ada lovelace", "A.L."},
		// Acceptance criteria: empty string
		{"", ""},
		// Extra whitespace (acceptance criteria)
		{"  ada   lovelace  ", "A.L."},
		{"one  two   three", "O.T.T."},
		// Unicode awareness (acceptance criteria)
		{"élan vital", "É.V."},
		{"über alles", "Ü.A."},
		// Single word
		{"hello", "H."},
		// Mixed case input (should uppercase)
		{"hello World", "H.W."},
		{"HELLO world", "H.W."},
		// Tab and newline whitespace
		{"one\ttwo\nthree", "O.T.T."},
		// Only whitespace
		{"   ", ""},
		{"\t\n  ", ""},
		// Numbers (first character)
		{"123 abc", "1.A."},
	}

	for _, tt := range tests {
		if got := Abbreviate(tt.in); got != tt.want {
			t.Errorf("Abbreviate(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Empty string
		{"", ""},
		// Basic ASCII
		{"hello world", "Hello world"},
		{"hello", "Hello"},
		// Already capitalized
		{"Hello", "Hello"},
		// Unicode awareness
		{"élan", "Élan"},
		{"über", "Über"},
		// Single character
		{"a", "A"},
		{"é", "É"},
		// Mixed case - only first changes
		{"hELLO", "HELLO"},
	}

	for _, tt := range tests {
		if got := Capitalize(tt.in); got != tt.want {
			t.Errorf("Capitalize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Acceptance criteria cases
		{"Hello World", 3},
		{"", 0},
		// Additional test cases
		// All vowels lowercase
		{"aeiou", 5},
		// All vowels uppercase
		{"AEIOU", 5},
		// Mixed case
		{"AeIoU", 5},
		// No vowels
		{"bcdfg", 0},
		{"xyz", 0},
		// Single vowel
		{"a", 1},
		{"E", 1},
		// Repeated vowels
		{"aaaa", 4},
		{"eeeee", 5},
		// Real words
		{"beautiful", 5},
		{"rhythm", 0},
		{"education", 5},
	}

	for _, tt := range tests {
		if got := CountVowels(tt.in); got != tt.want {
			t.Errorf("CountVowels(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria cases
		{"one two three", "three two one"},
		{"", ""},
		// Additional test cases
		// Single word
		{"hello", "hello"},
		// Two words
		{"hello world", "world hello"},
		// Multiple spaces (should normalize to single spaces)
		{"one  two   three", "three two one"},
		{"  hello   world  ", "world hello"},
		// Tab and newline whitespace
		{"one\ttwo\nthree", "three two one"},
		// Only whitespace
		{"   ", ""},
		{"\t\n  ", ""},
		// Four words
		{"the quick brown fox", "fox brown quick the"},
		// Numbers and mixed characters
		{"123 abc 456", "456 abc 123"},
		// Punctuation (treated as part of words)
		{"hello, world!", "world! hello,"},
		// Unicode
		{"héllo wörld", "wörld héllo"},
	}

	for _, tt := range tests {
		if got := ReverseWords(tt.in); got != tt.want {
			t.Errorf("ReverseWords(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: empty string
		{"", true},
		// Acceptance criteria: only whitespace (ASCII)
		{" ", true},
		{"  ", true},
		{"   ", true},
		{"\t", true},
		{"\n", true},
		{"\r", true},
		{" \t\n ", true},
		// Acceptance criteria: Unicode whitespace
		{"\u00A0", true},     // Non-breaking space
		{"\u2000", true},     // En quad
		{"\u2003", true},     // Em space
		{"\u3000", true},     // Ideographic space
		{" \u00A0\t", true}, // Mixed whitespace
		// Acceptance criteria: false for strings with content
		{" a ", false},
		{"a", false},
		{"hello", false},
		{" hello ", false},
		{"\thello\n", false},
	}

	for _, tt := range tests {
		if got := IsBlank(tt.in); got != tt.want {
			t.Errorf("IsBlank(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestHasPrefixFold(t *testing.T) {
	tests := []struct {
		s      string
		prefix string
		want   bool
	}{
		// Acceptance criteria cases
		{"Hello", "he", true},
		{"Hi", "hello", false},
		// Empty prefix
		{"", "", true},
		{"Hello", "", true},
		// Exact match (case-insensitive)
		{"hello", "HELLO", true},
		{"HELLO", "hello", true},
		// Partial match
		{"Hello World", "hello", true},
		{"Hello World", "HELLO", true},
		// No match
		{"Hello", "world", false},
		{"abc", "xyz", false},
		// Prefix longer than string
		{"Hi", "Hello", false},
		// Unicode awareness
		{"Élégant", "élé", true},
		{"ÉLÉGANT", "élé", true},
	}

	for _, tt := range tests {
		if got := HasPrefixFold(tt.s, tt.prefix); got != tt.want {
			t.Errorf("HasPrefixFold(%q, %q) = %v, want %v", tt.s, tt.prefix, got, tt.want)
		}
	}
}

func TestSqueeze(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: specific example
		{"  a   b  ", "a b"},
		// Acceptance criteria: empty string
		{"", ""},
		// Only whitespace (should return empty)
		{" ", ""},
		{"  ", ""},
		// No whitespace
		{"abc", "abc"},
		// Single space (already normalized)
		{"a b", "a b"},
		// Multiple spaces
		{"a  b", "a b"},
		{"a   b   c", "a b c"},
		// Leading/trailing whitespace
		{" a", "a"},
		{"a ", "a"},
		{" a ", "a"},
		// Tab characters
		{"a\tb", "a b"},
		{"a\t\tb", "a b"},
		// Newlines
		{"a\nb", "a b"},
		{"a\n\nb", "a b"},
		// Mixed whitespace
		{"  a\t\n  b  ", "a b"},
		{"one  \ttwo\n\nthree   ", "one two three"},
	}

	for _, tt := range tests {
		if got := Squeeze(tt.in); got != tt.want {
			t.Errorf("Squeeze(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCountRunes(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		// Acceptance criteria: empty input
		{"", 0},
		// Acceptance criteria: ASCII
		{"hello", 5},
		{"abc", 3},
		{"a", 1},
		// Acceptance criteria: multibyte Unicode ("héllo" returns 5)
		{"héllo", 5},
		// Additional multibyte test cases
		{"日本語", 3},
		{"こんにちは", 5},
		{"café", 4},
		{"naïve", 5},
		{"Zürich", 6},
		// Mixed ASCII and multibyte
		{"hello世界", 7},
		{"test😀emoji", 10},
		// Emojis (each emoji is one rune)
		{"😀😁😂", 3},
	}

	for _, tt := range tests {
		if got := CountRunes(tt.in); got != tt.want {
			t.Errorf("CountRunes(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestCountFields(t *testing.T) {
	tests := []struct {
		in   string
		sep  rune
		want int
	}{
		// Acceptance criteria: specific example
		{"a,,b,", ',', 2},
		// Acceptance criteria: empty string
		{"", ',', 0},
		// Single separator
		{"a,b", ',', 2},
		{"a,b,c", ',', 3},
		{"one,two,three", ',', 3},
		// Repeated separators (should skip empty fields)
		{"a,,b", ',', 2},
		{"a,,,b", ',', 2},
		{"a,,,,b", ',', 2},
		// Leading separator
		{",a", ',', 1},
		{",,a", ',', 1},
		// Trailing separator
		{"a,", ',', 1},
		{"a,,", ',', 1},
		// Leading and trailing separators
		{",a,", ',', 1},
		{",,a,,", ',', 1},
		// Only separators
		{",,", ',', 0},
		{",,,,", ',', 0},
		// No separators
		{"abc", ',', 1},
		{"hello", ',', 1},
		// Different separators
		{"a:b:c", ':', 3},
		{"a;b;c", ';', 3},
		{"a|b|c", '|', 3},
		{"a b c", ' ', 3},
		// Unicode separator
		{"a·b·c", '·', 3},
		// Mixed cases
		{"one,two,,three,", ',', 3},
		{",one,,two,,,three", ',', 3},
	}

	for _, tt := range tests {
		if got := CountFields(tt.in, tt.sep); got != tt.want {
			t.Errorf("CountFields(%q, %q) = %d, want %d", tt.in, tt.sep, got, tt.want)
		}
	}
}

func TestStartsWithVowel(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: empty string returns false
		{"", false},
		// Acceptance criteria: lowercase vowels
		{"apple", true},
		{"elephant", true},
		{"igloo", true},
		{"orange", true},
		{"umbrella", true},
		// Acceptance criteria: uppercase vowels
		{"Apple", true},
		{"Elephant", true},
		{"Igloo", true},
		{"Orange", true},
		{"Umbrella", true},
		// Acceptance criteria: consonants
		{"banana", false},
		{"cherry", false},
		{"date", false},
		{"grape", false},
		{"kiwi", false},
		{"Banana", false},
		{"CHERRY", false},
		// Single vowel
		{"a", true},
		{"e", true},
		{"i", true},
		{"o", true},
		{"u", true},
		{"A", true},
		{"E", true},
		{"I", true},
		{"O", true},
		{"U", true},
	}

	for _, tt := range tests {
		if got := StartsWithVowel(tt.in); got != tt.want {
			t.Errorf("StartsWithVowel(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestCountChar(t *testing.T) {
	tests := []struct {
		in   string
		r    rune
		want int
	}{
		// Acceptance criteria: empty input
		{"", 'a', 0},
		{"", '日', 0},
		// Acceptance criteria: absent runes
		{"hello", 'z', 0},
		{"world", 'x', 0},
		{"abc", 'd', 0},
		// Acceptance criteria: multibyte runes
		{"日本語", '日', 1},
		{"日本語日", '日', 2},
		{"こんにちは", 'ん', 1},
		{"héllo", 'é', 1},
		{"café", 'é', 1},
		{"naïve", 'ï', 1},
		// ASCII characters
		{"hello", 'l', 2},
		{"hello", 'h', 1},
		{"hello", 'o', 1},
		{"hello", 'e', 1},
		// Single character
		{"a", 'a', 1},
		{"a", 'b', 0},
		// Multiple occurrences
		{"aaa", 'a', 3},
		{"aaabbbccc", 'b', 3},
		{"banana", 'a', 3},
		{"banana", 'n', 2},
		// Case sensitivity
		{"Hello", 'H', 1},
		{"Hello", 'h', 0},
		{"Hello", 'l', 2},
		{"Hello", 'L', 0},
		// Emojis (multibyte)
		{"😀😁😀", '😀', 2},
		{"😀😁😂", '😀', 1},
		{"test😀emoji", '😀', 1},
		{"test😀emoji", 'e', 2},
	}

	for _, tt := range tests {
		if got := CountChar(tt.in, tt.r); got != tt.want {
			t.Errorf("CountChar(%q, %q) = %d, want %d", tt.in, tt.r, got, tt.want)
		}
	}
}

func TestIsUpper(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: ABC-1 is true
		{"ABC-1", true},
		// Acceptance criteria: 123 is false (no letters)
		{"123", false},
		// Acceptance criteria: AbC is false (mixed case)
		{"AbC", false},
		// Empty string (no letters)
		{"", false},
		// All uppercase letters
		{"ABC", true},
		{"HELLO", true},
		{"WORLD", true},
		// All lowercase letters
		{"abc", false},
		{"hello", false},
		// Mixed case
		{"Hello", false},
		{"HeLLo", false},
		{"ABC123abc", false},
		// Only non-letters (should return false)
		{"!@#$%", false},
		{"   ", false},
		{"---", false},
		// Uppercase with non-letters (should ignore non-letters)
		{"ABC-DEF", true},
		{"HELLO WORLD", true},
		{"A B C", true},
		{"X-Y-Z", true},
		{"A1B2C3", true},
		// Single character
		{"A", true},
		{"a", false},
		{"1", false},
		// Unicode uppercase
		{"ÉLAN", true},
		{"ÜBER", true},
		// Unicode mixed case
		{"Élan", false},
		{"über", false},
	}

	for _, tt := range tests {
		if got := IsUpper(tt.in); got != tt.want {
			t.Errorf("IsUpper(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestAcronym(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: empty string returns ""
		{"", ""},
		// Acceptance criteria: "portable network graphics" returns "PNG"
		{"portable network graphics", "PNG"},
		// Acceptance criteria: one word
		{"hello", "H"},
		// Acceptance criteria: several words
		{"hello world", "HW"},
		{"one two three", "OTT"},
		// Extra whitespace
		{"  hello   world  ", "HW"},
		{"one  two   three", "OTT"},
		// Unicode awareness
		{"élan vital", "ÉV"},
		{"über alles", "ÜA"},
		// Mixed case input (should uppercase)
		{"hello World", "HW"},
		{"HELLO world", "HW"},
		// Tab and newline whitespace
		{"one\ttwo\nthree", "OTT"},
	}

	for _, tt := range tests {
		if got := Acronym(tt.in); got != tt.want {
			t.Errorf("Acronym(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestLines(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		// Acceptance criteria: empty string returns empty, non-nil slice
		{"", []string{}},
		
		// Acceptance criteria: splits on \n without the ending
		{"one\ntwo\nthree", []string{"one", "two", "three"}},
		{"line1\nline2", []string{"line1", "line2"}},
		
		// Acceptance criteria: splits on \r\n without the ending
		{"one\r\ntwo\r\nthree", []string{"one", "two", "three"}},
		{"line1\r\nline2", []string{"line1", "line2"}},
		
		// Acceptance criteria: trailing line ending does not produce empty last line
		{"one\ntwo\n", []string{"one", "two"}},
		{"one\r\ntwo\r\n", []string{"one", "two"}},
		{"line\n", []string{"line"}},
		{"line\r\n", []string{"line"}},
		
		// Mixed line endings
		{"one\ntwo\r\nthree", []string{"one", "two", "three"}},
		{"one\r\ntwo\nthree", []string{"one", "two", "three"}},
		{"a\nb\r\nc\nd", []string{"a", "b", "c", "d"}},
		
		// Single line (no line endings)
		{"single line", []string{"single line"}},
		{"a", []string{"a"}},
		
		// Empty lines
		{"\n", []string{""}},
		{"\r\n", []string{""}},
		{"one\n\ntwo", []string{"one", "", "two"}},
		{"one\r\n\r\ntwo", []string{"one", "", "two"}},
		{"\n\n", []string{"", ""}},
		{"\r\n\r\n", []string{"", ""}},
		
		// Standalone \r (not followed by \n) - should be treated as content
		{"one\rtwo", []string{"one\rtwo"}},
		{"a\rb\nc", []string{"a\rb", "c"}},
		
		// Content with special characters
		{"hello world\ngoodbye world", []string{"hello world", "goodbye world"}},
		{"line with spaces   \nline 2", []string{"line with spaces   ", "line 2"}},
		{"tab\there\nnext", []string{"tab\there", "next"}},
		
		// Multiple consecutive line endings
		{"a\n\n\nb", []string{"a", "", "", "b"}},
		{"a\r\n\r\n\r\nb", []string{"a", "", "", "b"}},
	}

	for _, tt := range tests {
		got := Lines(tt.in)
		
		// Check nil vs non-nil
		if got == nil {
			t.Errorf("Lines(%q) returned nil, want non-nil slice", tt.in)
			continue
		}
		
		// Check length
		if len(got) != len(tt.want) {
			t.Errorf("Lines(%q) returned %d lines, want %d\nGot: %#v\nWant: %#v", 
				tt.in, len(got), len(tt.want), got, tt.want)
			continue
		}
		
		// Check each line
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("Lines(%q)[%d] = %q, want %q", tt.in, i, got[i], tt.want[i])
			}
		}
	}
}

func TestHasDigit(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: empty string
		{"", false},
		
		// Acceptance criteria: letters-only (should return false)
		{"abc", false},
		{"hello", false},
		{"WORLD", false},
		{"AbCdEf", false},
		
		// Acceptance criteria: mixed input (letters and digits)
		{"abc123", true},
		{"hello1world", true},
		{"test9", true},
		{"1test", true},
		{"a1b2c3", true},
		
		// Digits only
		{"123", true},
		{"0", true},
		{"9", true},
		{"0123456789", true},
		
		// Special characters without digits
		{"!@#$%", false},
		{"hello-world", false},
		{"test_case", false},
		
		// Special characters with digits
		{"test-123", true},
		{"abc@456", true},
		{"#1winner", true},
	}

	for _, tt := range tests {
		if got := HasDigit(tt.in); got != tt.want {
			t.Errorf("HasDigit(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestIsLower(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: lowercase text (should return true)
		{"hello", true},
		{"world", true},
		{"abc", true},
		{"lowercase", true},
		// Unicode lowercase
		{"über", true},
		{"café", true},
		
		// Acceptance criteria: mixed case (should return false)
		{"Hello", false},
		{"WORLD", false},
		{"MixedCase", false},
		{"hELLO", false},
		{"aBc", false},
		// Mixed with Unicode
		{"Über", false},
		{"Café", false},
		
		// Acceptance criteria: no letters (should return false)
		{"", false},
		{"123", false},
		{"!@#", false},
		{"   ", false},
		{"123!@#", false},
		{"\t\n", false},
		
		// Lowercase with numbers/special characters (should return true)
		{"hello123", true},
		{"abc!@#", true},
		{"test123world", true},
		{"a1b2c3", true},
		{"hello-world", true},
		{"test_case", true},
		{"abc 123 def", true},
		{"hello, world!", true},
	}

	for _, tt := range tests {
		if got := IsLower(tt.in); got != tt.want {
			t.Errorf("IsLower(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestHasUpper(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		// Acceptance criteria: empty string
		{"", false},
		
		// Acceptance criteria: lowercase-only (should return false)
		{"hello", false},
		{"world", false},
		{"abc", false},
		{"lowercase", false},
		// Unicode lowercase
		{"über", false},
		{"café", false},
		
		// Acceptance criteria: mixed input (should return true)
		{"Hello", true},
		{"WORLD", true},
		{"MixedCase", true},
		{"hELLO", true},
		{"aBc", true},
		// Mixed with Unicode
		{"Über", true},
		{"Café", true},
		
		// Acceptance criteria: non-ASCII uppercase input
		{"HELLO", true},
		{"ÜBER", true},
		{"CaféMüller", true},
		
		// Edge cases with numbers and special characters
		{"123", false},
		{"!@#", false},
		{"   ", false},
		{"123!@#", false},
		{"\t\n", false},
		{"hello123", false},
		{"abc!@#", false},
		{"test123world", false},
		{"a1b2c3", false},
		{"hello-world", false},
		{"test_case", false},
		{"abc 123 def", false},
		{"hello, world!", false},
		
		// Cases with uppercase letters
		{"hello WORLD", true},
		{"test CASE", true},
		{"aB1c2d3", true},
		{"!@#ABC$%^", true},
		{"123ABC", true},
		{"ABC123", true},
		{"aBcDeF", true},
	}

	for _, tt := range tests {
		if got := HasUpper(tt.in); got != tt.want {
			t.Errorf("HasUpper(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}