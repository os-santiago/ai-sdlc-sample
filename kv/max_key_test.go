package kv

import "testing"

func TestMaxKeyNil(t *testing.T) {
	// Test nil map - should return empty string and false
	var nilMap map[string]string
	got, ok := MaxKey(nilMap)
	if ok {
		t.Fatalf("MaxKey(nil) returned true, want false")
	}
	if got != "" {
		t.Fatalf("MaxKey(nil) = %q, want empty string", got)
	}
}

func TestMaxKeyEmpty(t *testing.T) {
	// Test empty map - should return empty string and false
	m := map[string]string{}
	got, ok := MaxKey(m)
	if ok {
		t.Fatalf("MaxKey(empty) returned true, want false")
	}
	if got != "" {
		t.Fatalf("MaxKey(empty) = %q, want empty string", got)
	}
}

func TestMaxKeyOneKey(t *testing.T) {
	// Test one key - should return that key and true
	m := map[string]string{"single": "value"}
	got, ok := MaxKey(m)
	if !ok {
		t.Fatalf("MaxKey(one key) returned false, want true")
	}
	if got != "single" {
		t.Fatalf("MaxKey(one key) = %q, want %q", got, "single")
	}
}

func TestMaxKeySeveralKeys(t *testing.T) {
	// Test several keys - should return the lexically largest
	m := map[string]string{
		"apple":  "1",
		"banana": "2",
		"cherry": "3",
		"date":   "4",
	}
	got, ok := MaxKey(m)
	if !ok {
		t.Fatalf("MaxKey(several keys) returned false, want true")
	}
	if got != "date" {
		t.Fatalf("MaxKey(several keys) = %q, want %q", got, "date")
	}
}

func TestMaxKeySeveralKeysNumeric(t *testing.T) {
	// Test lexical ordering with numeric-like keys
	m := map[string]string{
		"z":   "1",
		"a":   "2",
		"m":   "3",
		"zzz": "4",
	}
	got, ok := MaxKey(m)
	if !ok {
		t.Fatalf("MaxKey(numeric-like keys) returned false, want true")
	}
	if got != "zzz" {
		t.Fatalf("MaxKey(numeric-like keys) = %q, want %q", got, "zzz")
	}
}
