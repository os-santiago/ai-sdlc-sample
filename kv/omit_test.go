package kv

import "testing"

func TestOmit(t *testing.T) {
	// Test case: no keys - should return a copy of the entire map
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	got := Omit(m)
	want := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	if !Equal(got, want) {
		t.Fatalf("Omit(%v) with no keys = %v, want %v", m, got, want)
	}

	// Test case: omitting present keys
	got = Omit(m, "a", "c")
	want = map[string]string{
		"b": "2",
	}
	if !Equal(got, want) {
		t.Fatalf("Omit(%v, \"a\", \"c\") = %v, want %v", m, got, want)
	}

	// Test case: omitting missing keys
	got = Omit(m, "x", "y", "z")
	want = map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	if !Equal(got, want) {
		t.Fatalf("Omit(%v, \"x\", \"y\", \"z\") = %v, want %v", m, got, want)
	}

	// Test case: mix of present and missing keys
	got = Omit(m, "a", "x", "b")
	want = map[string]string{
		"c": "3",
	}
	if !Equal(got, want) {
		t.Fatalf("Omit(%v, \"a\", \"x\", \"b\") = %v, want %v", m, got, want)
	}

	// Test case: input map is unchanged
	original := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	_ = Omit(original, "a", "b")
	if !Equal(original, m) {
		t.Fatalf("Omit modified input map: got %v, want %v", original, m)
	}

	// Test case: nil map
	var nilMap map[string]string
	got = Omit(nilMap, "a", "b")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Omit(nil, \"a\", \"b\") = %v, want %v", got, want)
	}

	// Test case: nil map with no keys
	got = Omit(nilMap)
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Omit(nil) with no keys = %v, want %v", got, want)
	}

	// Test case: empty map
	emptyMap := map[string]string{}
	got = Omit(emptyMap, "a")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Omit(empty map, \"a\") = %v, want %v", got, want)
	}

	// Test case: omitting all keys
	got = Omit(m, "a", "b", "c")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Omit(%v, \"a\", \"b\", \"c\") = %v, want %v", m, got, want)
	}
}
