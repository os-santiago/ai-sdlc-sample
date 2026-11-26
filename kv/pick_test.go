package kv

import "testing"

func TestPick(t *testing.T) {
	// Test case: present keys
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	got := Pick(m, "a", "c")
	want := map[string]string{
		"a": "1",
		"c": "3",
	}
	if !Equal(got, want) {
		t.Fatalf("Pick(%v, \"a\", \"c\") = %v, want %v", m, got, want)
	}

	// Test case: missing keys
	got = Pick(m, "x", "y", "z")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Pick(%v, \"x\", \"y\", \"z\") = %v, want %v", m, got, want)
	}

	// Test case: mix of present and missing keys
	got = Pick(m, "a", "x", "b")
	want = map[string]string{
		"a": "1",
		"b": "2",
	}
	if !Equal(got, want) {
		t.Fatalf("Pick(%v, \"a\", \"x\", \"b\") = %v, want %v", m, got, want)
	}

	// Test case: no keys
	got = Pick(m)
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Pick(%v) with no keys = %v, want %v", m, got, want)
	}

	// Test case: nil map
	var nilMap map[string]string
	got = Pick(nilMap, "a", "b")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Pick(nil, \"a\", \"b\") = %v, want %v", got, want)
	}

	// Test case: nil map with no keys
	got = Pick(nilMap)
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Pick(nil) with no keys = %v, want %v", got, want)
	}

	// Test case: empty map
	emptyMap := map[string]string{}
	got = Pick(emptyMap, "a")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("Pick(empty map, \"a\") = %v, want %v", got, want)
	}
}
