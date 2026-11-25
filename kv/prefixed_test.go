package kv

import "testing"

func TestPrefixed(t *testing.T) {
	// Test with empty map
	empty := map[string]string{}
	got := Prefixed(empty, "prefix_")
	if got == nil {
		t.Fatalf("Prefixed(empty) returned nil, want empty map")
	}
	if len(got) != 0 {
		t.Fatalf("Prefixed(empty) = %v, want empty map", got)
	}

	// Test with several keys
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	got = Prefixed(m, "prefix_")
	want := map[string]string{
		"prefix_key1": "value1",
		"prefix_key2": "value2",
		"prefix_key3": "value3",
	}
	if len(got) != len(want) {
		t.Fatalf("Prefixed(m) length = %d, want %d", len(got), len(want))
	}
	for k, v := range want {
		if got[k] != v {
			t.Fatalf("Prefixed(m)[%q] = %q, want %q", k, got[k], v)
		}
	}

	// Test that input map is not modified
	expectedOriginal := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	for k, v := range expectedOriginal {
		if m[k] != v {
			t.Fatalf("Input map was modified: m[%q] = %q, want %q", k, m[k], v)
		}
	}

	// Test with nil map
	var nilMap map[string]string
	got = Prefixed(nilMap, "prefix_")
	if got != nil {
		t.Fatalf("Prefixed(nil) = %v, want nil", got)
	}
}