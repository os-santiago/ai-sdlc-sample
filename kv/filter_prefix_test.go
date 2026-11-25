package kv

import "testing"

func TestFilterPrefix(t *testing.T) {
	// Test case: no match
	m := map[string]string{
		"foo": "1",
		"bar": "2",
	}
	got := FilterPrefix(m, "baz")
	want := map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("FilterPrefix(%v, \"baz\") = %v, want %v", m, got, want)
	}

	// Test case: some matches
	m = map[string]string{
		"foo": "1",
		"foobar": "2",
		"bar": "3",
		"fooqux": "4",
	}
	got = FilterPrefix(m, "foo")
	want = map[string]string{
		"foo": "1",
		"foobar": "2",
		"fooqux": "4",
	}
	if !Equal(got, want) {
		t.Fatalf("FilterPrefix(%v, \"foo\") = %v, want %v", m, got, want)
	}

	// Test case: empty prefix
	m = map[string]string{
		"foo": "1",
		"bar": "2",
	}
	got = FilterPrefix(m, "")
	want = map[string]string{
		"foo": "1",
		"bar": "2",
	}
	if !Equal(got, want) {
		t.Fatalf("FilterPrefix(%v, \"\") = %v, want %v", m, got, want)
	}

	// Test case: empty map
	got = FilterPrefix(map[string]string{}, "prefix")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("FilterPrefix(empty map, \"prefix\") = %v, want %v", got, want)
	}
}