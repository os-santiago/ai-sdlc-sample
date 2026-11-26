package kv

import "testing"

func TestCountPrefix(t *testing.T) {
	// Test case: nil map
	got := CountPrefix(nil, "foo")
	want := 0
	if got != want {
		t.Fatalf("CountPrefix(nil, \"foo\") = %d, want %d", got, want)
	}

	// Test case: empty prefix counts every key
	m := map[string]string{
		"foo": "1",
		"bar": "2",
		"baz": "3",
	}
	got = CountPrefix(m, "")
	want = 3
	if got != want {
		t.Fatalf("CountPrefix(%v, \"\") = %d, want %d", m, got, want)
	}

	// Test case: prefix matching some keys
	m = map[string]string{
		"foo": "1",
		"foobar": "2",
		"bar": "3",
		"fooqux": "4",
	}
	got = CountPrefix(m, "foo")
	want = 3
	if got != want {
		t.Fatalf("CountPrefix(%v, \"foo\") = %d, want %d", m, got, want)
	}

	// Test case: no match
	m = map[string]string{
		"foo": "1",
		"bar": "2",
	}
	got = CountPrefix(m, "baz")
	want = 0
	if got != want {
		t.Fatalf("CountPrefix(%v, \"baz\") = %d, want %d", m, got, want)
	}

	// Test case: empty map
	got = CountPrefix(map[string]string{}, "prefix")
	want = 0
	if got != want {
		t.Fatalf("CountPrefix(empty map, \"prefix\") = %d, want %d", got, want)
	}
}
