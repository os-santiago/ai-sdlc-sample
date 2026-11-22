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

func TestPadRightASCII(t *testing.T) {
	if got := PadRight("hi", 5, '.'); got != "hi..." {
		t.Errorf("PadRight = %q, want %q", got, "hi...")
	}
}

func TestPadRightMultibyte(t *testing.T) {
	if got := PadRight("é", 3, '.'); got != "é.." {
		t.Errorf("PadRight = %q, want %q", got, "é..")
	}
}

func TestPadRightLonger(t *testing.T) {
	if got := PadRight("hello", 3, '.'); got != "hello" {
		t.Errorf("PadRight = %q, want %q", got, "hello")
	}
}

func TestPadRightExactWidth(t *testing.T) {
	if got := PadRight("abc", 3, '.'); got != "abc" {
		t.Errorf("PadRight = %q, want %q", got, "abc")
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

func TestRepeatBasic(t *testing.T) {
	if got := Repeat("ab", "-", 3); got != "ab-ab-ab" {
		t.Errorf("Repeat = %q, want %q", got, "ab-ab-ab")
	}
}

func TestRepeatZero(t *testing.T) {
	if got := Repeat("hello", ",", 0); got != "" {
		t.Errorf("Repeat = %q, want empty", got)
	}
}

func TestRepeatNegative(t *testing.T) {
	if got := Repeat("test", " ", -1); got != "" {
		t.Errorf("Repeat = %q, want empty", got)
	}
}

func TestRepeatOne(t *testing.T) {
	if got := Repeat("single", "|", 1); got != "single" {
		t.Errorf("Repeat = %q, want %q", got, "single")
	}
}

func TestRepeatEmptyString(t *testing.T) {
	if got := Repeat("", "-", 3); got != "--" {
		t.Errorf("Repeat = %q, want %q", got, "--")
	}
}

func TestRepeatEmptySeparator(t *testing.T) {
	if got := Repeat("ab", "", 3); got != "ababab" {
		t.Errorf("Repeat = %q, want %q", got, "ababab")
	}
}

func TestRepeatMultibyte(t *testing.T) {
	if got := Repeat("é", "•", 2); got != "é•é" {
		t.Errorf("Repeat = %q, want %q", got, "é•é")
	}
}

func TestEllipsisShortASCII(t *testing.T) {
	if got := Ellipsis("hello", 10); got != "hello" {
		t.Errorf("Ellipsis = %q, want %q", got, "hello")
	}
}

func TestEllipsisExactLength(t *testing.T) {
	if got := Ellipsis("hello", 5); got != "hello" {
		t.Errorf("Ellipsis = %q, want %q", got, "hello")
	}
}

func TestEllipsisTruncateASCII(t *testing.T) {
	if got := Ellipsis("hello world", 8); got != "hello w…" {
		t.Errorf("Ellipsis = %q, want %q", got, "hello w…")
	}
}

func TestEllipsisMultibyte(t *testing.T) {
	if got := Ellipsis("héllo wörld", 8); got != "héllo w…" {
		t.Errorf("Ellipsis = %q, want %q", got, "héllo w…")
	}
}

func TestEllipsisMultibyteShort(t *testing.T) {
	if got := Ellipsis("café", 10); got != "café" {
		t.Errorf("Ellipsis = %q, want %q", got, "café")
	}
}

func TestEllipsisZero(t *testing.T) {
	if got := Ellipsis("hello", 0); got != "" {
		t.Errorf("Ellipsis = %q, want empty", got)
	}
}

func TestEllipsisNegative(t *testing.T) {
	if got := Ellipsis("hello", -1); got != "" {
		t.Errorf("Ellipsis = %q, want empty", got)
	}
}

func TestEllipsisEmptyString(t *testing.T) {
	if got := Ellipsis("", 5); got != "" {
		t.Errorf("Ellipsis = %q, want empty", got)
	}
}

func TestEllipsisOne(t *testing.T) {
	if got := Ellipsis("hello", 1); got != "…" {
		t.Errorf("Ellipsis = %q, want %q", got, "…")
	}
}

func TestEllipsisTwo(t *testing.T) {
	if got := Ellipsis("hello", 2); got != "h…" {
		t.Errorf("Ellipsis = %q, want %q", got, "h…")
	}
}

func TestIndentSingleLine(t *testing.T) {
	if got := Indent("hello", "  "); got != "  hello" {
		t.Errorf("Indent = %q, want %q", got, "  hello")
	}
}

func TestIndentMultiLine(t *testing.T) {
	input := "hello\nworld"
	want := "  hello\n  world"
	if got := Indent(input, "  "); got != want {
		t.Errorf("Indent = %q, want %q", got, want)
	}
}

func TestIndentEmptyLine(t *testing.T) {
	input := "hello\n\nworld"
	want := "  hello\n\n  world"
	if got := Indent(input, "  "); got != want {
		t.Errorf("Indent = %q, want %q", got, want)
	}
}

func TestIndentTrailingNewline(t *testing.T) {
	input := "hello\n"
	want := "  hello\n"
	if got := Indent(input, "  "); got != want {
		t.Errorf("Indent = %q, want %q", got, want)
	}
}

func TestIndentLeadingNewline(t *testing.T) {
	input := "\nhello"
	want := "\n  hello"
	if got := Indent(input, "  "); got != want {
		t.Errorf("Indent = %q, want %q", got, want)
	}
}

func TestIndentEmptyString(t *testing.T) {
	if got := Indent("", "  "); got != "" {
		t.Errorf("Indent = %q, want empty", got)
	}
}

func TestIndentDifferentPrefix(t *testing.T) {
	input := "line1\nline2"
	want := ">>line1\n>>line2"
	if got := Indent(input, ">>"); got != want {
		t.Errorf("Indent = %q, want %q", got, want)
	}
}

func TestIndentMultipleEmptyLines(t *testing.T) {
	input := "a\n\n\nb"
	want := "  a\n\n\n  b"
	if got := Indent(input, "  "); got != want {
		t.Errorf("Indent = %q, want %q", got, want)
	}
}

// Wrap tests

func TestWrapBasic(t *testing.T) {
	input := "hello world this is a test"
	want := "hello world\nthis is a\ntest"
	if got := Wrap(input, 12); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapSingleWord(t *testing.T) {
	input := "hello"
	want := "hello"
	if got := Wrap(input, 10); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapLongWord(t *testing.T) {
	input := "supercalifragilisticexpialidocious"
	want := "supercalifragilisticexpialidocious"
	if got := Wrap(input, 10); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapLongWordWithOthers(t *testing.T) {
	input := "hello supercalifragilisticexpialidocious world"
	want := "hello\nsupercalifragilisticexpialidocious\nworld"
	if got := Wrap(input, 10); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapZeroWidth(t *testing.T) {
	input := "hello world"
	want := "hello world"
	if got := Wrap(input, 0); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapNegativeWidth(t *testing.T) {
	input := "hello world"
	want := "hello world"
	if got := Wrap(input, -5); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapEmptyString(t *testing.T) {
	input := ""
	want := ""
	if got := Wrap(input, 10); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapMultipleSpaces(t *testing.T) {
	input := "hello    world"
	want := "hello\nworld"
	if got := Wrap(input, 7); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapMultibyte(t *testing.T) {
	input := "héllo wörld café"
	want := "héllo\nwörld café"
	if got := Wrap(input, 10); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}

func TestWrapNoSpaces(t *testing.T) {
	input := "helloworld"
	want := "helloworld"
	if got := Wrap(input, 5); got != want {
		t.Errorf("Wrap = %q, want %q", got, want)
	}
}
