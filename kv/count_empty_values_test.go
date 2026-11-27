package kv

import "testing"

func TestCountEmptyValues(t *testing.T) {
	// Test case: nil map
	got := CountEmptyValues(nil)
	want := 0
	if got != want {
		t.Fatalf("CountEmptyValues(nil) = %d, want %d", got, want)
	}

	// Test case: no empty values
	m := map[string]string{
		"foo": "1",
		"bar": "2",
		"baz": "3",
	}
	got = CountEmptyValues(m)
	want = 0
	if got != want {
		t.Fatalf("CountEmptyValues(%v) = %d, want %d", m, got, want)
	}

	// Test case: several empty values
	m = map[string]string{
		"foo": "",
		"bar": "value",
		"baz": "",
		"qux": "another",
		"quux": "",
	}
	got = CountEmptyValues(m)
	want = 3
	if got != want {
		t.Fatalf("CountEmptyValues(%v) = %d, want %d", m, got, want)
	}

	// Test case: all empty values
	m = map[string]string{
		"a": "",
		"b": "",
		"c": "",
	}
	got = CountEmptyValues(m)
	want = 3
	if got != want {
		t.Fatalf("CountEmptyValues(%v) = %d, want %d", m, got, want)
	}

	// Test case: empty map
	got = CountEmptyValues(map[string]string{})
	want = 0
	if got != want {
		t.Fatalf("CountEmptyValues(empty map) = %d, want %d", got, want)
	}
}
