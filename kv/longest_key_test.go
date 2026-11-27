package kv

import "testing"

func TestLongestKeyNil(t *testing.T) {
	// Test nil map - should return empty string and false
	var nilMap map[string]string
	got, ok := LongestKey(nilMap)
	if ok {
		t.Fatalf("LongestKey(nil) returned true, want false")
	}
	if got != "" {
		t.Fatalf("LongestKey(nil) = %q, want empty string", got)
	}
}

func TestLongestKeyEmpty(t *testing.T) {
	// Test empty map - should return empty string and false
	m := map[string]string{}
	got, ok := LongestKey(m)
	if ok {
		t.Fatalf("LongestKey(empty) returned true, want false")
	}
	if got != "" {
		t.Fatalf("LongestKey(empty) = %q, want empty string", got)
	}
}

func TestLongestKeyTie(t *testing.T) {
	// Test tie - same length keys, should return lexically smallest
	m := map[string]string{
		"cat": "1",
		"dog": "2",
		"bat": "3",
	}
	got, ok := LongestKey(m)
	if !ok {
		t.Fatalf("LongestKey(tie) returned false, want true")
	}
	if got != "bat" {
		t.Fatalf("LongestKey(tie) = %q, want %q", got, "bat")
	}
}

func TestLongestKeyDistinctLengths(t *testing.T) {
	// Test distinct lengths - should return the longest
	m := map[string]string{
		"a":     "1",
		"ab":    "2",
		"abc":   "3",
		"abcde": "4",
	}
	got, ok := LongestKey(m)
	if !ok {
		t.Fatalf("LongestKey(distinct lengths) returned false, want true")
	}
	if got != "abcde" {
		t.Fatalf("LongestKey(distinct lengths) = %q, want %q", got, "abcde")
	}
}

func TestLongestKeyTieWithShorter(t *testing.T) {
	// Test tie among longest keys with shorter keys present
	m := map[string]string{
		"a":        "1",
		"elephant": "2",
		"giraffe":  "3",
		"ab":       "4",
	}
	got, ok := LongestKey(m)
	if !ok {
		t.Fatalf("LongestKey(tie with shorter) returned false, want true")
	}
	if got != "elephant" {
		t.Fatalf("LongestKey(tie with shorter) = %q, want %q", got, "elephant")
	}
}
