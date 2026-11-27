package kv

import "testing"

func TestFilterValuePrefix(t *testing.T) {
	// Test case: nil map
	got := FilterValuePrefix(nil, "prefix")
	want := map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("FilterValuePrefix(nil, \"prefix\") = %v, want %v", got, want)
	}

	// Test case: no match
	m := map[string]string{
		"foo": "1",
		"bar": "2",
	}
	got = FilterValuePrefix(m, "baz")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("FilterValuePrefix(%v, \"baz\") = %v, want %v", m, got, want)
	}

	// Test case: some matches
	m = map[string]string{
		"foo": "prefix1",
		"foobar": "prefix2",
		"bar": "2",
		"fooqux": "prefix3",
	}
	got = FilterValuePrefix(m, "prefix")
	want = map[string]string{
		"foo": "prefix1",
		"foobar": "prefix2",
		"fooqux": "prefix3",
	}
	if !Equal(got, want) {
		t.Fatalf("FilterValuePrefix(%v, \"prefix\") = %v, want %v", m, got, want)
	}

	// Test case: empty prefix
	m = map[string]string{
		"foo": "1",
		"bar": "2",
	}
	got = FilterValuePrefix(m, "")
	want = map[string]string{
		"foo": "1",
		"bar": "2",
	}
	if !Equal(got, want) {
		t.Fatalf("FilterValuePrefix(%v, \"\") = %v, want %v", m, got, want)
	}

	// Test case: empty map
	got = FilterValuePrefix(map[string]string{}, "prefix")
	want = map[string]string{}
	if !Equal(got, want) {
		t.Fatalf("FilterValuePrefix(empty map, \"prefix\") = %v, want %v", got, want)
	}
}