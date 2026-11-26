package kv

import (
	"reflect"
	"testing"
)

func TestKeysWithValue(t *testing.T) {
	// Test no match
	m := map[string]string{"a": "1", "b": "2", "c": "3"}
	got := KeysWithValue(m, "4")
	want := []string{}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("KeysWithValue(m, \"4\") = %v, want %v", got, want)
	}

	// Verify it returns non-nil slice
	if got == nil {
		t.Fatalf("KeysWithValue(m, \"4\") returned nil, want non-nil empty slice")
	}

	// Test one match
	got = KeysWithValue(m, "2")
	want = []string{"b"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("KeysWithValue(m, \"2\") = %v, want %v", got, want)
	}

	// Test several matches
	m2 := map[string]string{"a": "x", "b": "y", "c": "x", "d": "z", "e": "x"}
	got = KeysWithValue(m2, "x")
	want = []string{"a", "c", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("KeysWithValue(m2, \"x\") = %v, want %v", got, want)
	}

	// Test nil map
	var nilMap map[string]string
	got = KeysWithValue(nilMap, "any")
	want = []string{}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("KeysWithValue(nilMap, \"any\") = %v, want %v", got, want)
	}

	// Verify nil map returns non-nil slice
	if got == nil {
		t.Fatalf("KeysWithValue(nilMap, \"any\") returned nil, want non-nil empty slice")
	}
}
