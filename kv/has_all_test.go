package kv

import "testing"

func TestHasAll(t *testing.T) {
	// Test all keys present
	m := map[string]string{"a": "1", "b": "2", "c": "3"}
	if !HasAll(m, "a", "b", "c") {
		t.Fatalf("HasAll(m, \"a\", \"b\", \"c\") = false, want true")
	}
	
	// Test one key missing
	if HasAll(m, "a", "b", "d") {
		t.Fatalf("HasAll(m, \"a\", \"b\", \"d\") = true, want false")
	}
	
	// Test no keys provided
	if !HasAll(m) {
		t.Fatalf("HasAll(m) with no keys = false, want true")
	}
	
	// Test with nil map
	var nilMap map[string]string
	if HasAll(nilMap, "a") {
		t.Fatalf("HasAll(nilMap, \"a\") = true, want false")
	}
	
	// Test with empty map
	emptyMap := map[string]string{}
	if HasAll(emptyMap, "a") {
		t.Fatalf("HasAll(emptyMap, \"a\") = true, want false")
	}
	
	// Test with empty map and no keys
	if !HasAll(emptyMap) {
		t.Fatalf("HasAll(emptyMap) with no keys = false, want true")
	}
}