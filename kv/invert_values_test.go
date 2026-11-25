package kv

import "testing"

func TestInvertValues(t *testing.T) {
	// Test nil map returns empty map
	var nilMap map[string]string
	got := InvertValues(nilMap)
	if got == nil {
		t.Fatalf("InvertValues(nil) returned nil, want empty map")
	}
	if len(got) != 0 {
		t.Fatalf("InvertValues(nil) returned map with %d elements, want 0", len(got))
	}

	// Test unique values
	m1 := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	got = InvertValues(m1)
	want := map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
	}
	if len(got) != len(want) {
		t.Fatalf("InvertValues(unique) length mismatch: got %d, want %d", len(got), len(want))
	}
	for k, v := range want {
		if got[k] != v {
			t.Fatalf("InvertValues(unique)[%q] = %q, want %q", k, got[k], v)
		}
	}

	// Test repeated values - smallest key should win
	m2 := map[string]string{
		"a": "1",
		"b": "2",
		"c": "1", // duplicate value
		"d": "3",
	}
	got = InvertValues(m2)
	want = map[string]string{
		"1": "a", // "a" comes before "c" lexicographically
		"2": "b",
		"3": "d",
	}
	if len(got) != len(want) {
		t.Fatalf("InvertValues(repeated) length mismatch: got %d, want %d", len(got), len(want))
	}
	for k, v := range want {
		if got[k] != v {
			t.Fatalf("InvertValues(repeated)[%q] = %q, want %q", k, got[k], v)
		}
	}

	// Test empty map
	emptyMap := map[string]string{}
	got = InvertValues(emptyMap)
	if len(got) != 0 {
		t.Fatalf("InvertValues(empty) returned map with %d elements, want 0", len(got))
	}
}