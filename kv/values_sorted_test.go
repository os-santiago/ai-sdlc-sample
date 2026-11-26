package kv

import (
	"reflect"
	"testing"
)

func TestValuesSorted(t *testing.T) {
	// Test nil map
	var nilMap map[string]string
	got := ValuesSorted(nilMap)
	if got == nil {
		t.Fatalf("ValuesSorted(nil) returned nil, want empty non-nil slice")
	}
	if len(got) != 0 {
		t.Fatalf("ValuesSorted(nil) = %v, want []", got)
	}

	// Test empty map
	emptyMap := map[string]string{}
	got = ValuesSorted(emptyMap)
	if got == nil {
		t.Fatalf("ValuesSorted(empty) returned nil, want empty non-nil slice")
	}
	if len(got) != 0 {
		t.Fatalf("ValuesSorted(empty) = %v, want []", got)
	}

	// Test duplicates
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "1", // duplicate value
		"d": "3",
		"e": "2", // another duplicate
	}
	got = ValuesSorted(m)
	want := []string{"1", "2", "3"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ValuesSorted(duplicates) = %v, want %v", got, want)
	}

	// Test several values (all unique)
	m2 := map[string]string{
		"z": "zebra",
		"a": "apple",
		"m": "mango",
		"b": "banana",
	}
	got = ValuesSorted(m2)
	want = []string{"apple", "banana", "mango", "zebra"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ValuesSorted(several) = %v, want %v", got, want)
	}

	// Test single value
	m3 := map[string]string{
		"only": "one",
	}
	got = ValuesSorted(m3)
	want = []string{"one"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ValuesSorted(single) = %v, want %v", got, want)
	}
}
