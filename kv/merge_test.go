package kv

import "testing"

// Merge tests - Issue #44 Acceptance Criteria

func TestMergeOverlappingKeys(t *testing.T) {
	// Acceptance Criteria 2: When a key is in both, the value from b wins
	a := map[string]string{"x": "a-value", "y": "1"}
	b := map[string]string{"x": "b-value", "z": "2"}
	got := Merge(a, b)

	// Check right-hand precedence
	if got["x"] != "b-value" {
		t.Fatalf("Merge(%v, %v)[\"x\"] = %q, want \"b-value\"", a, b, got["x"])
	}
	// Check all keys present
	if got["y"] != "1" || got["z"] != "2" {
		t.Fatalf("Merge(%v, %v) = %v, missing keys", a, b, got)
	}
	if len(got) != 3 {
		t.Fatalf("Merge(%v, %v) length = %d, want 3", a, b, len(got))
	}
}

func TestMergeAllKeys(t *testing.T) {
	// Acceptance Criteria 1: Returns a new map with every key of a and b
	a := map[string]string{"a": "1", "b": "2"}
	b := map[string]string{"c": "3", "d": "4"}
	got := Merge(a, b)

	if len(got) != 4 {
		t.Fatalf("Merge(%v, %v) length = %d, want 4", a, b, len(got))
	}
	if got["a"] != "1" || got["b"] != "2" || got["c"] != "3" || got["d"] != "4" {
		t.Fatalf("Merge(%v, %v) = %v, want all keys", a, b, got)
	}
}

func TestMergeInputsUnchanged(t *testing.T) {
	// Acceptance Criteria 3 & 4: Neither input map is modified
	a := map[string]string{"x": "original-a", "y": "1"}
	b := map[string]string{"x": "original-b", "z": "2"}

	// Store original values
	aOrigX, bOrigX := a["x"], b["x"]
	aLen, bLen := len(a), len(b)

	_ = Merge(a, b)

	// Verify a is unchanged
	if a["x"] != aOrigX || len(a) != aLen {
		t.Fatalf("Merge modified input a: got %v, original had x=%q, len=%d", a, aOrigX, aLen)
	}
	// Verify b is unchanged
	if b["x"] != bOrigX || len(b) != bLen {
		t.Fatalf("Merge modified input b: got %v, original had x=%q, len=%d", b, bOrigX, bLen)
	}
}

func TestMergeEmptyMaps(t *testing.T) {
	// Edge case: empty maps
	a := map[string]string{}
	b := map[string]string{"key": "value"}
	got := Merge(a, b)

	if len(got) != 1 || got["key"] != "value" {
		t.Fatalf("Merge(empty, %v) = %v, want {key: value}", b, got)
	}
}

func TestMergeNilMaps(t *testing.T) {
	// Edge case: nil maps
	var a map[string]string
	b := map[string]string{"key": "value"}
	got := Merge(a, b)

	if len(got) != 1 || got["key"] != "value" {
		t.Fatalf("Merge(nil, %v) = %v, want {key: value}", b, got)
	}
}
