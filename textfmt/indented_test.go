package textfmt

import "testing"

func TestIndentedZero(t *testing.T) {
	input := "hello\nworld"
	want := "hello\nworld"
	if got := Indented(input, 0); got != want {
		t.Errorf("Indented(%q, 0) = %q; want %q", input, got, want)
	}
}

func TestIndentedNegative(t *testing.T) {
	input := "hello\nworld"
	want := "hello\nworld"
	if got := Indented(input, -5); got != want {
		t.Errorf("Indented(%q, -5) = %q; want %q", input, got, want)
	}
}

func TestIndentedBlankLine(t *testing.T) {
	input := "first\n\nthird"
	want := "  first\n\n  third"
	if got := Indented(input, 2); got != want {
		t.Errorf("Indented(%q, 2) = %q; want %q", input, got, want)
	}
}

func TestIndentedSeveralLines(t *testing.T) {
	input := "first\nsecond\nthird"
	want := "    first\n    second\n    third"
	if got := Indented(input, 4); got != want {
		t.Errorf("Indented(%q, 4) = %q; want %q", input, got, want)
	}
}

func TestIndentedSingleLine(t *testing.T) {
	input := "hello"
	want := "   hello"
	if got := Indented(input, 3); got != want {
		t.Errorf("Indented(%q, 3) = %q; want %q", input, got, want)
	}
}

func TestIndentedEmptyString(t *testing.T) {
	if got := Indented("", 5); got != "" {
		t.Errorf("Indented(%q, 5) = %q; want empty", "", got)
	}
}

func TestIndentedTrailingNewline(t *testing.T) {
	input := "hello\n"
	want := "  hello\n"
	if got := Indented(input, 2); got != want {
		t.Errorf("Indented(%q, 2) = %q; want %q", input, got, want)
	}
}

func TestIndentedMultipleBlankLines(t *testing.T) {
	input := "a\n\n\nb"
	want := "  a\n\n\n  b"
	if got := Indented(input, 2); got != want {
		t.Errorf("Indented(%q, 2) = %q; want %q", input, got, want)
	}
}
