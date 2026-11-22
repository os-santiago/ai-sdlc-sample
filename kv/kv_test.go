package kv

import "testing"

func TestParseBasic(t *testing.T) {
	got, err := Parse("a=1;b=2")
	if err != nil || got["a"] != "1" || got["b"] != "2" || len(got) != 2 {
		t.Fatalf("Parse = %v, %v", got, err)
	}
}

func TestParseTrailingSeparator(t *testing.T) {
	// Trailing separator should be ignored
	got, err := Parse("a=1;b=2;")
	if err != nil {
		t.Fatalf("Parse with trailing separator returned error: %v", err)
	}
	if got["a"] != "1" || got["b"] != "2" || len(got) != 2 {
		t.Fatalf("Parse = %v, want {a:1, b:2}", got)
	}
}

func TestParseEmpty(t *testing.T) {
	// Empty input should return empty map
	got, err := Parse("")
	if err != nil {
		t.Fatalf("Parse empty string returned error: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("Parse = %v, want empty map", got)
	}
}

func TestParseMalformedPair(t *testing.T) {
	// Pair without '=' should return ErrSyntax
	_, err := Parse("a")
	if err != ErrSyntax {
		t.Fatalf("Parse malformed pair = %v, want ErrSyntax", err)
	}
}

func TestParseInts(t *testing.T) {
	// Valid input with integers
	got, err := ParseInts("retries=3;timeout=30")
	if err != nil {
		t.Fatalf("ParseInts returned error: %v", err)
	}
	if got["retries"] != 3 || got["timeout"] != 30 || len(got) != 2 {
		t.Fatalf("ParseInts = %v, want {retries:3, timeout:30}", got)
	}
}

func TestParseIntsNonInteger(t *testing.T) {
	// Non-integer value should return ErrSyntax
	_, err := ParseInts("count=abc")
	if err != ErrSyntax {
		t.Fatalf("ParseInts with non-integer = %v, want ErrSyntax", err)
	}
}

func TestParseIntsMalformed(t *testing.T) {
	// Malformed pair (no '=') should return ErrSyntax
	_, err := ParseInts("invalid")
	if err != ErrSyntax {
		t.Fatalf("ParseInts with malformed pair = %v, want ErrSyntax", err)
	}
}

func TestParseIntsTrailingSeparator(t *testing.T) {
	// Trailing separator should be ignored for ParseInts too
	got, err := ParseInts("a=1;b=2;")
	if err != nil {
		t.Fatalf("ParseInts with trailing separator returned error: %v", err)
	}
	if got["a"] != 1 || got["b"] != 2 || len(got) != 2 {
		t.Fatalf("ParseInts = %v, want {a:1, b:2}", got)
	}
}

func TestParseIntsEmpty(t *testing.T) {
	// Empty input should return empty map for ParseInts too
	got, err := ParseInts("")
	if err != nil {
		t.Fatalf("ParseInts empty string returned error: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("ParseInts = %v, want empty map", got)
	}
}

// Format tests - Acceptance Criteria

func TestFormatSorted(t *testing.T) {
	// Acceptance Criteria 1: Format with keys sorted
	m := map[string]string{"b": "2", "a": "1"}
	got := Format(m)
	want := "a=1;b=2"
	if got != want {
		t.Fatalf("Format(%v) = %q, want %q", m, got, want)
	}
}

func TestFormatEmpty(t *testing.T) {
	// Acceptance Criteria 2: Format of empty map returns empty string
	m := map[string]string{}
	got := Format(m)
	if got != "" {
		t.Fatalf("Format(empty map) = %q, want empty string", got)
	}
}

func TestFormatParseRoundTrip(t *testing.T) {
	// Acceptance Criteria 3: Parse(Format(m)) returns m
	tests := []map[string]string{
		{"a": "1", "b": "2"},
		{"foo": "bar", "baz": "qux", "x": "y"},
		{"single": "value"},
		{},
	}

	for _, m := range tests {
		formatted := Format(m)
		parsed, err := Parse(formatted)
		if err != nil {
			t.Fatalf("Parse(Format(%v)) returned error: %v", m, err)
		}
		if len(parsed) != len(m) {
			t.Fatalf("Parse(Format(%v)) = %v, length mismatch", m, parsed)
		}
		for k, v := range m {
			if parsed[k] != v {
				t.Fatalf("Parse(Format(%v))[%q] = %q, want %q", m, k, parsed[k], v)
			}
		}
	}
}

func TestFormatMultipleKeys(t *testing.T) {
	// Test with multiple keys to verify sorting
	m := map[string]string{
		"z": "last",
		"a": "first",
		"m": "middle",
	}
	got := Format(m)
	want := "a=first;m=middle;z=last"
	if got != want {
		t.Fatalf("Format(%v) = %q, want %q", m, got, want)
	}
}

func TestFormatSinglePair(t *testing.T) {
	// Test with single pair
	m := map[string]string{"key": "value"}
	got := Format(m)
	want := "key=value"
	if got != want {
		t.Fatalf("Format(%v) = %q, want %q", m, got, want)
	}
}
