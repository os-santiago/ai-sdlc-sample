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
