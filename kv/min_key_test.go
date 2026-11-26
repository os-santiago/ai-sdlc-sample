package kv

import "testing"

func TestMinKeyNil(t *testing.T) {
	// Test nil map - should return empty string and false
	var nilMap map[string]string
	got, ok := MinKey(nilMap)
	if ok {
		t.Fatalf("MinKey(nil) returned true, want false")
	}
	if got != "" {
		t.Fatalf("MinKey(nil) = %q, want empty string", got)
	}
}

func TestMinKeyEmpty(t *testing.T) {
	// Test empty map - should return empty string and false
	m := map[string]string{}
	got, ok := MinKey(m)
	if ok {
		t.Fatalf("MinKey(empty) returned true, want false")
	}
	if got != "" {
		t.Fatalf("MinKey(empty) = %q, want empty string", got)
	}
}

func TestMinKeyOneKey(t *testing.T) {
	// Test one key - should return that key and true
	m := map[string]string{"single": "value"}
	got, ok := MinKey(m)
	if !ok {
		t.Fatalf("MinKey(one key) returned false, want true")
	}
	if got != "single" {
		t.Fatalf("MinKey(one key) = %q, want %q", got, "single")
	}
}

func TestMinKeySeveralKeys(t *testing.T) {
	// Test several keys - should return the lexically smallest
	m := map[string]string{
		"apple":  "1",
		"banana": "2",
		"cherry": "3",
		"date":   "4",
	}
	got, ok := MinKey(m)
	if !ok {
		t.Fatalf("MinKey(several keys) returned false, want true")
	}
	if got != "apple" {
		t.Fatalf("MinKey(several keys) = %q, want %q", got, "apple")
	}
}

func TestMinKeySeveralKeysNumeric(t *testing.T) {
	// Test lexical ordering with numeric-like keys
	m := map[string]string{
		"z":   "1",
		"a":   "2",
		"m":   "3",
		"aaa": "4",
	}
	got, ok := MinKey(m)
	if !ok {
		t.Fatalf("MinKey(numeric-like keys) returned false, want true")
	}
	if got != "a" {
		t.Fatalf("MinKey(numeric-like keys) = %q, want %q", got, "a")
	}
}
