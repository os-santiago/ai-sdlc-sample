package textfmt

import "testing"

func TestTruncateShort(t *testing.T) {
	if got := Truncate("hi", 10); got != "hi" {
		t.Errorf("Truncate = %q", got)
	}
}

func TestTruncateMultibyte(t *testing.T) {
	if got := Truncate("héllo wörld", 8); got != "héllo..." {
		t.Errorf("Truncate = %q, want %q", got, "héllo...")
	}
}

func TestTitleASCII(t *testing.T) {
	if got := Title("hello world"); got != "Hello World" {
		t.Errorf("Title = %q", got)
	}
}

func TestTitleMultibyte(t *testing.T) {
	if got := Title("élan vital"); got != "Élan Vital" {
		t.Errorf("Title = %q, want %q", got, "Élan Vital")
	}
}

func TestPadLeftASCII(t *testing.T) {
	if got := PadLeft("hi", 5, '.'); got != "...hi" {
		t.Errorf("PadLeft = %q, want %q", got, "...hi")
	}
}

func TestPadLeftMultibyte(t *testing.T) {
	if got := PadLeft("é", 3, '.'); got != "..é" {
		t.Errorf("PadLeft = %q, want %q", got, "..é")
	}
}

func TestPadLeftLonger(t *testing.T) {
	if got := PadLeft("hello", 3, '.'); got != "hello" {
		t.Errorf("PadLeft = %q, want %q", got, "hello")
	}
}

func TestCenterOddPadding(t *testing.T) {
	if got := Center("ab", 5, '*'); got != "*ab**" {
		t.Errorf("Center = %q, want %q", got, "*ab**")
	}
}

func TestCenterEvenPadding(t *testing.T) {
	if got := Center("hi", 6, '.'); got != "..hi.." {
		t.Errorf("Center = %q, want %q", got, "..hi..")
	}
}

func TestCenterMultibyte(t *testing.T) {
	if got := Center("é", 5, '-'); got != "--é--" {
		t.Errorf("Center = %q, want %q", got, "--é--")
	}
}

func TestCenterAlreadyWide(t *testing.T) {
	if got := Center("hello", 3, '*'); got != "hello" {
		t.Errorf("Center = %q, want %q", got, "hello")
	}
}

func TestCenterExactWidth(t *testing.T) {
	if got := Center("abc", 3, '.'); got != "abc" {
		t.Errorf("Center = %q, want %q", got, "abc")
	}
}

// Tests for Repeat function

func TestRepeatBasic(t *testing.T) {
	// Main acceptance criteria: Repeat("ab", "-", 3) returns "ab-ab-ab"
	if got := Repeat("ab", "-", 3); got != "ab-ab-ab" {
		t.Errorf("Repeat(\"ab\", \"-\", 3) = %q, want %q", got, "ab-ab-ab")
	}
}

func TestRepeatZero(t *testing.T) {
	// n of 0 returns empty string
	if got := Repeat("hello", ",", 0); got != "" {
		t.Errorf("Repeat(\"hello\", \",\", 0) = %q, want %q", got, "")
	}
}

func TestRepeatNegative(t *testing.T) {
	// n less than 0 returns empty string
	if got := Repeat("test", " ", -1); got != "" {
		t.Errorf("Repeat(\"test\", \" \", -1) = %q, want %q", got, "")
	}
}

func TestRepeatOne(t *testing.T) {
	// n of 1 returns just the string without separator
	if got := Repeat("single", "|", 1); got != "single" {
		t.Errorf("Repeat(\"single\", \"|\", 1) = %q, want %q", got, "single")
	}
}

func TestRepeatEmptyString(t *testing.T) {
	// Repeating empty string
	if got := Repeat("", "-", 3); got != "--" {
		t.Errorf("Repeat(\"\", \"-\", 3) = %q, want %q", got, "--")
	}
}

func TestRepeatEmptySeparator(t *testing.T) {
	// Repeating with empty separator
	if got := Repeat("ab", "", 3); got != "ababab" {
		t.Errorf("Repeat(\"ab\", \"\", 3) = %q, want %q", got, "ababab")
	}
}

func TestRepeatMultibyte(t *testing.T) {
	// Repeating multibyte characters
	if got := Repeat("é", "•", 2); got != "é•é" {
		t.Errorf("Repeat(\"é\", \"•\", 2) = %q, want %q", got, "é•é")
	}
}

// Tests for Ellipsis function

func TestEllipsisShortASCII(t *testing.T) {
	// String with at most n characters is returned unchanged
	if got := Ellipsis("hello", 10); got != "hello" {
		t.Errorf("Ellipsis(\"hello\", 10) = %q, want %q", got, "hello")
	}
}

func TestEllipsisExactLength(t *testing.T) {
	// String with exactly n characters is returned unchanged
	if got := Ellipsis("hello", 5); got != "hello" {
		t.Errorf("Ellipsis(\"hello\", 5) = %q, want %q", got, "hello")
	}
}

func TestEllipsisTruncateASCII(t *testing.T) {
	// String longer than n characters gets truncated with ellipsis
	if got := Ellipsis("hello world", 8); got != "hello w…" {
		t.Errorf("Ellipsis(\"hello world\", 8) = %q, want %q", got, "hello w…")
	}
}

func TestEllipsisMultibyte(t *testing.T) {
	// Multibyte characters are counted as single runes
	if got := Ellipsis("héllo wörld", 8); got != "héllo w…" {
		t.Errorf("Ellipsis(\"héllo wörld\", 8) = %q, want %q", got, "héllo w…")
	}
}

func TestEllipsisMultibyteShort(t *testing.T) {
	// Short multibyte string is returned unchanged
	if got := Ellipsis("café", 10); got != "café" {
		t.Errorf("Ellipsis(\"café\", 10) = %q, want %q", got, "café")
	}
}

func TestEllipsisZero(t *testing.T) {
	// n of 0 returns empty string
	if got := Ellipsis("hello", 0); got != "" {
		t.Errorf("Ellipsis(\"hello\", 0) = %q, want %q", got, "")
	}
}

func TestEllipsisNegative(t *testing.T) {
	// n less than 0 returns empty string
	if got := Ellipsis("hello", -1); got != "" {
		t.Errorf("Ellipsis(\"hello\", -1) = %q, want %q", got, "")
	}
}

func TestEllipsisEmptyString(t *testing.T) {
	// Empty string with positive n returns empty string
	if got := Ellipsis("", 5); got != "" {
		t.Errorf("Ellipsis(\"\", 5) = %q, want %q", got, "")
	}
}

func TestEllipsisOne(t *testing.T) {
	// n of 1 with longer string returns just ellipsis
	if got := Ellipsis("hello", 1); got != "…" {
		t.Errorf("Ellipsis(\"hello\", 1) = %q, want %q", got, "…")
	}
}

func TestEllipsisTwo(t *testing.T) {
	// n of 2 with longer string returns first character + ellipsis
	if got := Ellipsis("hello", 2); got != "h…" {
		t.Errorf("Ellipsis(\"hello\", 2) = %q, want %q", got, "h…")
	}
}
