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

// Quote tests

func TestQuotePlain(t *testing.T) {
	if got := Quote("hello"); got != `"hello"` {
		t.Errorf("Quote = %q, want %q", got, `"hello"`)
	}
}

func TestQuoteWithDoubleQuote(t *testing.T) {
	if got := Quote(`a"b`); got != `"a\"b"` {
		t.Errorf("Quote = %q, want %q", got, `"a\"b"`)
	}
}

func TestQuoteWithBackslash(t *testing.T) {
	if got := Quote(`a\b`); got != `"a\\b"` {
		t.Errorf("Quote = %q, want %q", got, `"a\\b"`)
	}
}

func TestQuoteWithBoth(t *testing.T) {
	if got := Quote(`a\"b`); got != `"a\\\"b"` {
		t.Errorf("Quote = %q, want %q", got, `"a\\\"b"`)
	}
}

func TestQuoteEmpty(t *testing.T) {
	if got := Quote(""); got != `""` {
		t.Errorf("Quote = %q, want %q", got, `""`)
	}
}

func TestQuoteMultipleQuotes(t *testing.T) {
	if got := Quote(`a"b"c`); got != `"a\"b\"c"` {
		t.Errorf("Quote = %q, want %q", got, `"a\"b\"c"`)
	}
}

func TestQuoteMultipleBackslashes(t *testing.T) {
	if got := Quote(`a\\b`); got != `"a\\\\b"` {
		t.Errorf("Quote = %q, want %q", got, `"a\\\\b"`)
	}
}

// Surround tests

func TestSurroundBasic(t *testing.T) {
	if got := Surround("x", "**"); got != "**x**" {
		t.Errorf("Surround = %q, want %q", got, "**x**")
	}
}

func TestSurroundEmptyText(t *testing.T) {
	if got := Surround("", "**"); got != "****" {
		t.Errorf("Surround = %q, want %q", got, "****")
	}
}

func TestSurroundEmptyMarker(t *testing.T) {
	if got := Surround("hello", ""); got != "hello" {
		t.Errorf("Surround = %q, want %q", got, "hello")
	}
}

func TestSurroundMultibyteMarker(t *testing.T) {
	if got := Surround("text", "🔥"); got != "🔥text🔥" {
		t.Errorf("Surround = %q, want %q", got, "🔥text🔥")
	}
}

func TestSurroundMultibyteText(t *testing.T) {
	if got := Surround("café", "["); got != "[café[" {
		t.Errorf("Surround = %q, want %q", got, "[café[")
	}
}

func TestSurroundMultibyteMarkerAndText(t *testing.T) {
	if got := Surround("émoji", "✨"); got != "✨émoji✨" {
		t.Errorf("Surround = %q, want %q", got, "✨émoji✨")
	}
}

// Bullet tests

func TestBulletEmpty(t *testing.T) {
	if got := Bullet([]string{}); got != "" {
		t.Errorf("Bullet = %q, want empty", got)
	}
}

func TestBulletOne(t *testing.T) {
	if got := Bullet([]string{"item"}); got != "- item" {
		t.Errorf("Bullet = %q, want %q", got, "- item")
	}
}

func TestBulletSeveral(t *testing.T) {
	want := "- first\n- second\n- third"
	if got := Bullet([]string{"first", "second", "third"}); got != want {
		t.Errorf("Bullet = %q, want %q", got, want)
	}
}

func TestBulletMultibyte(t *testing.T) {
	want := "- café\n- thé"
	if got := Bullet([]string{"café", "thé"}); got != want {
		t.Errorf("Bullet = %q, want %q", got, want)
	}
}

func TestBulletNil(t *testing.T) {
	if got := Bullet(nil); got != "" {
		t.Errorf("Bullet = %q, want empty", got)
	}
}

// Dedent tests

func TestDedentNoIndentation(t *testing.T) {
	input := "hello\nworld"
	want := "hello\nworld"
	if got := Dedent(input); got != want {
		t.Errorf("Dedent = %q, want %q", got, want)
	}
}

func TestDedentUniformIndentation(t *testing.T) {
	input := "    hello\n    world"
	want := "hello\nworld"
	if got := Dedent(input); got != want {
		t.Errorf("Dedent = %q, want %q", got, want)
	}
}

func TestDedentMixedIndentation(t *testing.T) {
	input := "    hello\n      world"
	want := "hello\n  world"
	if got := Dedent(input); got != want {
		t.Errorf("Dedent = %q, want %q", got, want)
	}
}

func TestDedentWithBlankLines(t *testing.T) {
	input := "    hello\n\n    world"
	want := "hello\n\nworld"
	if got := Dedent(input); got != want {
		t.Errorf("Dedent = %q, want %q", got, want)
	}
}

func TestDedentBlankLineDoesNotAffectIndent(t *testing.T) {
	input := "    hello\n  \n    world"
	want := "hello\n  \nworld"
	if got := Dedent(input); got != want {
		t.Errorf("Dedent = %q, want %q", got, want)
	}
}

func TestDedentEmptyString(t *testing.T) {
	if got := Dedent(""); got != "" {
		t.Errorf("Dedent = %q, want empty", got)
	}
}

func TestDedentSingleLine(t *testing.T) {
	input := "    hello"
	want := "hello"
	if got := Dedent(input); got != want {
		t.Errorf("Dedent = %q, want %q", got, want)
	}
}

func TestDedentOnlyBlankLines(t *testing.T) {
	input := "\n  \n"
	want := "\n  \n"
	if got := Dedent(input); got != want {
		t.Errorf("Dedent = %q, want %q", got, want)
	}
}

// TrimRunes tests

func TestTrimRunesEmptyCutset(t *testing.T) {
	if got := TrimRunes("hello", ""); got != "hello" {
		t.Errorf("TrimRunes = %q, want %q", got, "hello")
	}
}

func TestTrimRunesASCII(t *testing.T) {
	if got := TrimRunes("--hello--", "-"); got != "hello" {
		t.Errorf("TrimRunes = %q, want %q", got, "hello")
	}
}

func TestTrimRunesASCIIMultipleChars(t *testing.T) {
	if got := TrimRunes(".-*hello*-.", ".-*"); got != "hello" {
		t.Errorf("TrimRunes = %q, want %q", got, "hello")
	}
}

func TestTrimRunesMultibyte(t *testing.T) {
	if got := TrimRunes("★café★", "★"); got != "café" {
		t.Errorf("TrimRunes = %q, want %q", got, "café")
	}
}

func TestTrimRunesMultibyteMultiple(t *testing.T) {
	if got := TrimRunes("🔥✨hello✨🔥", "🔥✨"); got != "hello" {
		t.Errorf("TrimRunes = %q, want %q", got, "hello")
	}
}

func TestTrimRunesMixedASCIIAndMultibyte(t *testing.T) {
	if got := TrimRunes("★-café-★", "★-"); got != "café" {
		t.Errorf("TrimRunes = %q, want %q", got, "café")
	}
}

func TestTrimRunesEmptyString(t *testing.T) {
	if got := TrimRunes("", "abc"); got != "" {
		t.Errorf("TrimRunes = %q, want empty", got)
	}
}

func TestTrimRunesNothingToTrim(t *testing.T) {
	if got := TrimRunes("hello", "xyz"); got != "hello" {
		t.Errorf("TrimRunes = %q, want %q", got, "hello")
	}
}

func TestTrimRunesLeadingOnly(t *testing.T) {
	if got := TrimRunes("--hello", "-"); got != "hello" {
		t.Errorf("TrimRunes = %q, want %q", got, "hello")
	}
}

func TestTrimRunesTrailingOnly(t *testing.T) {
	if got := TrimRunes("hello--", "-"); got != "hello" {
		t.Errorf("TrimRunes = %q, want %q", got, "hello")
	}
}

func TestTrimRunesAllCharactersInCutset(t *testing.T) {
	if got := TrimRunes("---", "-"); got != "" {
		t.Errorf("TrimRunes = %q, want empty", got)
	}
}

func TestTrimRunesMiddleNotTrimmed(t *testing.T) {
	if got := TrimRunes("--he-llo--", "-"); got != "he-llo" {
		t.Errorf("TrimRunes = %q, want %q", got, "he-llo")
	}
}

func TestTrimRunesSpaces(t *testing.T) {
	if got := TrimRunes("  hello  ", " "); got != "hello" {
		t.Errorf("TrimRunes = %q, want %q", got, "hello")
	}
}

// TruncateDots tests

func TestTruncateDotsShortInput(t *testing.T) {
	if got := TruncateDots("hello", 10); got != "hello" {
		t.Errorf("TruncateDots = %q, want %q", got, "hello")
	}
}

func TestTruncateDotsTruncation(t *testing.T) {
	if got := TruncateDots("hello world", 8); got != "hello w…" {
		t.Errorf("TruncateDots = %q, want %q", got, "hello w…")
	}
}

func TestTruncateDotsMaxZero(t *testing.T) {
	if got := TruncateDots("hello", 0); got != "" {
		t.Errorf("TruncateDots = %q, want empty", got)
	}
}

func TestTruncateDotsMaxNegative(t *testing.T) {
	if got := TruncateDots("hello", -1); got != "" {
		t.Errorf("TruncateDots = %q, want empty", got)
	}
}

// RepeatJoin tests

func TestRepeatJoinBasic(t *testing.T) {
	if got := RepeatJoin("ab", "-", 3); got != "ab-ab-ab" {
		t.Errorf("RepeatJoin = %q, want %q", got, "ab-ab-ab")
	}
}

func TestRepeatJoinZero(t *testing.T) {
	if got := RepeatJoin("hello", ",", 0); got != "" {
		t.Errorf("RepeatJoin = %q, want empty", got)
	}
}

func TestRepeatJoinNegative(t *testing.T) {
	if got := RepeatJoin("test", " ", -1); got != "" {
		t.Errorf("RepeatJoin = %q, want empty", got)
	}
}

func TestRepeatJoinOne(t *testing.T) {
	if got := RepeatJoin("single", "|", 1); got != "single" {
		t.Errorf("RepeatJoin = %q, want %q", got, "single")
	}
}

func TestRepeatJoinEmptyString(t *testing.T) {
	if got := RepeatJoin("", "-", 3); got != "--" {
		t.Errorf("RepeatJoin = %q, want %q", got, "--")
	}
}

func TestRepeatJoinEmptySeparator(t *testing.T) {
	if got := RepeatJoin("ab", "", 3); got != "ababab" {
		t.Errorf("RepeatJoin = %q, want %q", got, "ababab")
	}
}

func TestRepeatJoinMultibyte(t *testing.T) {
	if got := RepeatJoin("é", "•", 2); got != "é•é" {
		t.Errorf("RepeatJoin = %q, want %q", got, "é•é")
	}
}

// Indent2 tests

func TestIndent2SingleLine(t *testing.T) {
	if got := Indent2("hello"); got != "  hello" {
		t.Errorf("Indent2 = %q, want %q", got, "  hello")
	}
}

func TestIndent2MultipleLines(t *testing.T) {
	input := "first\nsecond\nthird"
	expected := "  first\n  second\n  third"
	if got := Indent2(input); got != expected {
		t.Errorf("Indent2 = %q, want %q", got, expected)
	}
}

func TestIndent2EmptyLineBetween(t *testing.T) {
	input := "first\n\nthird"
	expected := "  first\n\n  third"
	if got := Indent2(input); got != expected {
		t.Errorf("Indent2 = %q, want %q", got, expected)
	}
}

func TestIndent2EmptyString(t *testing.T) {
	if got := Indent2(""); got != "" {
		t.Errorf("Indent2 = %q, want empty", got)
	}
}

// Underline tests

func TestUnderlineASCII(t *testing.T) {
	expected := "heading\n-------"
	if got := Underline("heading"); got != expected {
		t.Errorf("Underline = %q, want %q", got, expected)
	}
}

func TestUnderlineMultibyte(t *testing.T) {
	expected := "héllo wörld\n-----------"
	if got := Underline("héllo wörld"); got != expected {
		t.Errorf("Underline = %q, want %q", got, expected)
	}
}

func TestUnderlineEmpty(t *testing.T) {
	expected := "\n"
	if got := Underline(""); got != expected {
		t.Errorf("Underline = %q, want %q", got, expected)
	}
}
func TestIndent2OnlyEmptyLines(t *testing.T) {
	input := "\n\n"
	expected := "\n\n"
	if got := Indent2(input); got != expected {
		t.Errorf("Indent2 = %q, want %q", got, expected)
	}
}

// Quoted tests

func TestQuotedEmpty(t *testing.T) {
	if got := Quoted(""); got != "" {
		t.Errorf("Quoted = %q, want empty", got)
	}
}

func TestQuotedOneLine(t *testing.T) {
	if got := Quoted("hello"); got != "> hello" {
		t.Errorf("Quoted = %q, want %q", got, "> hello")
	}
}

func TestQuotedSeveralLines(t *testing.T) {
	input := "first\nsecond\nthird"
	want := "> first\n> second\n> third"
	if got := Quoted(input); got != want {
		t.Errorf("Quoted = %q, want %q", got, want)
	}
}

func TestQuotedTrailingNewline(t *testing.T) {
	input := "hello\n"
	want := "> hello\n"
	if got := Quoted(input); got != want {
		t.Errorf("Quoted = %q, want %q", got, want)
	}
}

func TestQuotedEmptyLinesBetween(t *testing.T) {
	input := "first\n\nthird"
	want := "> first\n> \n> third"
	if got := Quoted(input); got != want {
		t.Errorf("Quoted = %q, want %q", got, want)
	}
}

// Box tests

func TestBoxASCII(t *testing.T) {
	expected := "+-------+\n| hello |\n+-------+"
	if got := Box("hello"); got != expected {
		t.Errorf("Box = %q, want %q", got, expected)
	}
}

func TestBoxMultibyte(t *testing.T) {
	expected := "+---+\n| 🔥 |\n+---+"
	if got := Box("🔥"); got != expected {
		t.Errorf("Box = %q, want %q", got, expected)
	}
}

func TestBoxEmpty(t *testing.T) {
	expected := "+--+\n|  |\n+--+"
	if got := Box(""); got != expected {
		t.Errorf("Box = %q, want %q", got, expected)
	}
}
