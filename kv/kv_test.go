package kv

import "testing"

func TestParseBasic(t *testing.T) {
	got, err := Parse("a=1;b=2")
	if err != nil || got["a"] != "1" || got["b"] != "2" || len(got) != 2 {
		t.Fatalf("Parse = %v, %v", got, err)
	}
}

func TestParseTrailingSeparator(t *testing.T) {
	// Trailing separator should be ignored
	got, err := Parse("a=1;b=2;")
	if err != nil {
		t.Fatalf("Parse with trailing separator returned error: %v", err)
	}
	if got["a"] != "1" || got["b"] != "2" || len(got) != 2 {
		t.Fatalf("Parse = %v, want {a:1, b:2}", got)
	}
}

func TestParseEmpty(t *testing.T) {
	// Empty input should return empty map
	got, err := Parse("")
	if err != nil {
		t.Fatalf("Parse empty string returned error: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("Parse = %v, want empty map", got)
	}
}

func TestParseMalformedPair(t *testing.T) {
	// Pair without '=' should return ErrSyntax
	_, err := Parse("a")
	if err != ErrSyntax {
		t.Fatalf("Parse malformed pair = %v, want ErrSyntax", err)
	}
}

func TestParseInts(t *testing.T) {
	// Valid input with integers
	got, err := ParseInts("retries=3;timeout=30")
	if err != nil {
		t.Fatalf("ParseInts returned error: %v", err)
	}
	if got["retries"] != 3 || got["timeout"] != 30 || len(got) != 2 {
		t.Fatalf("ParseInts = %v, want {retries:3, timeout:30}", got)
	}
}

func TestParseIntsNonInteger(t *testing.T) {
	// Non-integer value should return ErrSyntax
	_, err := ParseInts("count=abc")
	if err != ErrSyntax {
		t.Fatalf("ParseInts with non-integer = %v, want ErrSyntax", err)
	}
}

func TestParseIntsMalformed(t *testing.T) {
	// Malformed pair (no '=') should return ErrSyntax
	_, err := ParseInts("invalid")
	if err != ErrSyntax {
		t.Fatalf("ParseInts with malformed pair = %v, want ErrSyntax", err)
	}
}

func TestParseIntsTrailingSeparator(t *testing.T) {
	// Trailing separator should be ignored for ParseInts too
	got, err := ParseInts("a=1;b=2;")
	if err != nil {
		t.Fatalf("ParseInts with trailing separator returned error: %v", err)
	}
	if got["a"] != 1 || got["b"] != 2 || len(got) != 2 {
		t.Fatalf("ParseInts = %v, want {a:1, b:2}", got)
	}
}

func TestParseIntsEmpty(t *testing.T) {
	// Empty input should return empty map for ParseInts too
	got, err := ParseInts("")
	if err != nil {
		t.Fatalf("ParseInts empty string returned error: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("ParseInts = %v, want empty map", got)
	}
}

// Format tests - Acceptance Criteria

func TestFormatSorted(t *testing.T) {
	// Acceptance Criteria 1: Format with keys sorted
	m := map[string]string{"b": "2", "a": "1"}
	got := Format(m)
	want := "a=1;b=2"
	if got != want {
		t.Fatalf("Format(%v) = %q, want %q", m, got, want)
	}
}

func TestFormatEmpty(t *testing.T) {
	// Acceptance Criteria 2: Format of empty map returns empty string
	m := map[string]string{}
	got := Format(m)
	if got != "" {
		t.Fatalf("Format(empty map) = %q, want empty string", got)
	}
}

func TestFormatParseRoundTrip(t *testing.T) {
	// Acceptance Criteria 3: Parse(Format(m)) returns m
	tests := []map[string]string{
		{"a": "1", "b": "2"},
		{"foo": "bar", "baz": "qux", "x": "y"},
		{"single": "value"},
		{},
	}

	for _, m := range tests {
		formatted := Format(m)
		parsed, err := Parse(formatted)
		if err != nil {
			t.Fatalf("Parse(Format(%v)) returned error: %v", m, err)
		}
		if len(parsed) != len(m) {
			t.Fatalf("Parse(Format(%v)) = %v, length mismatch", m, parsed)
		}
		for k, v := range m {
			if parsed[k] != v {
				t.Fatalf("Parse(Format(%v))[%q] = %q, want %q", m, k, parsed[k], v)
			}
		}
	}
}

func TestFormatMultipleKeys(t *testing.T) {
	// Test with multiple keys to verify sorting
	m := map[string]string{
		"z": "last",
		"a": "first",
		"m": "middle",
	}
	got := Format(m)
	want := "a=first;m=middle;z=last"
	if got != want {
		t.Fatalf("Format(%v) = %q, want %q", m, got, want)
	}
}

func TestFormatSinglePair(t *testing.T) {
	// Test with single pair
	m := map[string]string{"key": "[REDACTED]"}
	got := Format(m)
	want := "key=[REDACTED]"
	if got != want {
		t.Fatalf("Format(%v) = %q, want %q", m, got, want)
	}
}

// Whitespace trimming tests - Issue #21 Acceptance Criteria

func TestParseTrimWhitespace(t *testing.T) {
	// Acceptance Criteria 1: Keys and values are trimmed of surrounding whitespace
	got, err := Parse(" a = 1 ; b=2")
	if err != nil {
		t.Fatalf("Parse with spaces returned error: %v", err)
	}
	if got["a"] != "1" || got["b"] != "2" || len(got) != 2 {
		t.Fatalf("Parse(\" a = 1 ; b=2\") = %v, want {a: \"1\", b: \"2\"}", got)
	}
}

func TestParseWhitespaceInValue(t *testing.T) {
	// Acceptance Criteria 2: Whitespace inside a value is kept
	got, err := Parse("msg=hello world")
	if err != nil {
		t.Fatalf("Parse with internal spaces returned error: %v", err)
	}
	if got["msg"] != "hello world" || len(got) != 1 {
		t.Fatalf("Parse(\"msg=hello world\") = %v, want {msg: \"hello world\"}", got)
	}
}

func TestParseMalformedPairWithSpaces(t *testing.T) {
	// Acceptance Criteria 3: A pair without '=' still returns ErrSyntax
	_, err := Parse(" a ")
	if err != ErrSyntax {
		t.Fatalf("Parse malformed pair with spaces = %v, want ErrSyntax", err)
	}
}

func TestParseTrimVariousWhitespace(t *testing.T) {
	// Additional test: various combinations of whitespace
	tests := []struct {
		input     string
		wantKey   string
		wantValue string
	}{
		{"a=1", "a", "1"},
		{" a=1", "a", "1"},
		{"a =1", "a", "1"},
		{"a= 1", "a", "1"},
		{"a=1 ", "a", "1"},
		{" a = 1 ", "a", "1"},
		{"  a  =  1  ", "a", "1"},
	}

	for _, tt := range tests {
		got, err := Parse(tt.input)
		if err != nil || got[tt.wantKey] != tt.wantValue || len(got) != 1 {
			t.Fatalf("Parse(%q) = %v, %v; want {%s: %s}", tt.input, got, err, tt.wantKey, tt.wantValue)
		}
	}
}

// Keys tests - Issue #33 Acceptance Criteria

func TestKeysSorted(t *testing.T) {
	// Acceptance Criteria 1: Keys returns sorted keys in ascending order
	m := map[string]string{"z": "last", "a": "first", "m": "middle"}
	got := Keys(m)
	want := []string{"a", "m", "z"}
	if len(got) != len(want) {
		t.Fatalf("Keys(%v) length = %d, want %d", m, len(got), len(want))
	}
	for i, k := range want {
		if got[i] != k {
			t.Fatalf("Keys(%v)[%d] = %q, want %q", m, i, got[i], k)
		}
	}
}

func TestKeysNil(t *testing.T) {
	// Acceptance Criteria 2: Keys(nil) returns empty, non-nil slice
	got := Keys(nil)
	if got == nil {
		t.Fatalf("Keys(nil) returned nil, want non-nil empty slice")
	}
	if len(got) != 0 {
		t.Fatalf("Keys(nil) = %v, want empty slice", got)
	}
	// Verify it's usable (can be appended to, ranged over, etc.)
	_ = append(got, "test")
}

// Delete tests - Issue #53 Acceptance Criteria

func TestDeleteRemovesKey(t *testing.T) {
	// Acceptance Criteria 1: Delete returns a new map without the specified key
	m := map[string]string{"a": "1", "b": "2", "c": "3"}
	got := Delete(m, "b")
	if len(got) != 2 {
		t.Fatalf("Delete length = %d, want 2", len(got))
	}
	if got["a"] != "1" || got["c"] != "3" {
		t.Fatalf("Delete = %v, want {a:1, c:3}", got)
	}
	if _, exists := got["b"]; exists {
		t.Fatalf("Delete still contains key b")
	}
}

func TestDeleteDoesNotMutateInput(t *testing.T) {
	// Acceptance Criteria 2: The input map is not modified
	original := map[string]string{"a": "1", "b": "2"}
	_ = Delete(original, "a")
	if len(original) != 2 {
		t.Fatalf("Delete mutated input: length = %d, want 2", len(original))
	}
	if original["a"] != "1" || original["b"] != "2" {
		t.Fatalf("Delete mutated input: %v, want {a:1, b:2}", original)
	}
}

func TestDeleteMissingKeyReturnsCopy(t *testing.T) {
	// Acceptance Criteria 3: Deleting a missing key returns an equal copy
	m := map[string]string{"a": "1", "b": "2"}
	got := Delete(m, "missing")
	if len(got) != len(m) {
		t.Fatalf("Delete with missing key: length = %d, want %d", len(got), len(m))
	}
	for k, v := range m {
		if got[k] != v {
			t.Fatalf("Delete with missing key: got[%q] = %q, want %q", k, got[k], v)
		}
	}
}

func TestDeleteEmptyMap(t *testing.T) {
	// Edge case: Delete from empty map
	m := map[string]string{}
	got := Delete(m, "any")
	if len(got) != 0 {
		t.Fatalf("Delete from empty map: length = %d, want 0", len(got))
	}
}

func TestDeleteSingleKey(t *testing.T) {
	// Edge case: Delete the only key
	m := map[string]string{"only": "value"}
	got := Delete(m, "only")
	if len(got) != 0 {
		t.Fatalf("Delete single key: length = %d, want 0", len(got))
	}
}

// Equal tests - Issue #67 Acceptance Criteria

func TestEqualSameMaps(t *testing.T) {
	// Acceptance Criteria 1: Equal returns true when both maps have same keys and values
	a := map[string]string{"a": "1", "b": "2"}
	b := map[string]string{"a": "1", "b": "2"}
	if !Equal(a, b) {
		t.Fatalf("Equal(%v, %v) = false, want true", a, b)
	}
}

func TestEqualDifferentValue(t *testing.T) {
	// Acceptance Criteria 3: Test with different value
	a := map[string]string{"a": "1", "b": "2"}
	b := map[string]string{"a": "1", "b": "3"}
	if Equal(a, b) {
		t.Fatalf("Equal(%v, %v) = true, want false", a, b)
	}
}

func TestEqualMissingKey(t *testing.T) {
	// Acceptance Criteria 3: Test with missing key
	a := map[string]string{"a": "1", "b": "2"}
	b := map[string]string{"a": "1"}
	if Equal(a, b) {
		t.Fatalf("Equal(%v, %v) = true, want false", a, b)
	}
}

func TestEqualNilVsEmpty(t *testing.T) {
	// Acceptance Criteria 2: A nil map equals an empty map
	var nilMap map[string]string
	emptyMap := map[string]string{}
	if !Equal(nilMap, emptyMap) {
		t.Fatalf("Equal(nil, {}) = false, want true")
	}
	if !Equal(emptyMap, nilMap) {
		t.Fatalf("Equal({}, nil) = false, want true")
	}
}

func TestEqualBothNil(t *testing.T) {
	// Edge case: Both nil
	var a, b map[string]string
	if !Equal(a, b) {
		t.Fatalf("Equal(nil, nil) = false, want true")
	}
}

func TestEqualBothEmpty(t *testing.T) {
	// Edge case: Both empty
	a := map[string]string{}
	b := map[string]string{}
	if !Equal(a, b) {
		t.Fatalf("Equal({}, {}) = false, want true")
	}
}

func TestEqualDifferentKeysEmptyValues(t *testing.T) {
	// Edge case: Different keys with empty values
	a := map[string]string{"x": ""}
	b := map[string]string{"y": ""}
	if Equal(a, b) {
		t.Fatalf("Equal(%v, %v) = true, want false", a, b)
	}
}

func TestEqualExtraKeyInB(t *testing.T) {
	// Edge case: b has an extra key
	a := map[string]string{"a": "1"}
	b := map[string]string{"a": "1", "b": "2"}
	if Equal(a, b) {
		t.Fatalf("Equal(%v, %v) = true, want false", a, b)
	}
}

// Filter tests - Issue #71 Acceptance Criteria

func TestFilterMatching(t *testing.T) {
	// Acceptance Criteria 1 & 3: Returns new map with only keys starting with prefix
	m := map[string]string{
		"app.name":    "myapp",
		"app.version": "1.0",
		"db.host":     "localhost",
		"db.port":     "5432",
	}
	got := Filter(m, "app.")
	if len(got) != 2 {
		t.Fatalf("Filter(%v, \"app.\") length = %d, want 2", m, len(got))
	}
	if got["app.name"] != "myapp" || got["app.version"] != "1.0" {
		t.Fatalf("Filter(%v, \"app.\") = %v, want {app.name:myapp, app.version:1.0}", m, got)
	}
}

func TestFilterNoMatch(t *testing.T) {
	// Acceptance Criteria 3: Test with no matching keys
	m := map[string]string{
		"app.name":    "myapp",
		"app.version": "1.0",
	}
	got := Filter(m, "db.")
	if len(got) != 0 {
		t.Fatalf("Filter(%v, \"db.\") length = %d, want 0", m, len(got))
	}
}

func TestFilterEmptyPrefix(t *testing.T) {
	// Acceptance Criteria 3: Empty prefix should match all keys
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	got := Filter(m, "")
	if len(got) != 3 {
		t.Fatalf("Filter(%v, \"\") length = %d, want 3", m, len(got))
	}
	if got["a"] != "1" || got["b"] != "2" || got["c"] != "3" {
		t.Fatalf("Filter(%v, \"\") = %v, want all entries", m, got)
	}
}

func TestFilterDoesNotModifyInput(t *testing.T) {
	// Acceptance Criteria 2: The input map is not modified
	original := map[string]string{
		"app.name": "myapp",
		"db.host":  "localhost",
	}
	_ = Filter(original, "app.")
	if len(original) != 2 || original["app.name"] != "myapp" || original["db.host"] != "localhost" {
		t.Fatalf("Filter mutated input: %v, want {app.name:myapp, db.host:localhost}", original)
	}
}

func TestFilterEmptyMap(t *testing.T) {
	// Edge case: Filter on empty map
	m := map[string]string{}
	got := Filter(m, "any")
	if len(got) != 0 {
		t.Fatalf("Filter(empty map, \"any\") length = %d, want 0", len(got))
	}
}

func TestFilterNilMap(t *testing.T) {
	// Edge case: Filter on nil map
	var m map[string]string
	got := Filter(m, "prefix")
	if got == nil || len(got) != 0 {
		t.Fatalf("Filter(nil map, \"prefix\") = %v, want empty non-nil map", got)
	}
}

// Get tests - Issue #80 Acceptance Criteria

func TestGetPresent(t *testing.T) {
	// Acceptance Criteria 1: Returns the value for key when present
	m := map[string]string{
		"name": "Alice",
		"age":  "30",
	}
	got := Get(m, "name", "default")
	if got != "Alice" {
		t.Fatalf("Get(%v, \"name\", \"default\") = %q, want \"Alice\"", m, got)
	}
}

func TestGetMissing(t *testing.T) {
	// Acceptance Criteria 1: Returns fallback when key is missing
	m := map[string]string{
		"name": "Alice",
	}
	got := Get(m, "age", "unknown")
	if got != "unknown" {
		t.Fatalf("Get(%v, \"age\", \"unknown\") = %q, want \"unknown\"", m, got)
	}
}

func TestGetEmptyValue(t *testing.T) {
	// Acceptance Criteria 2: A present key with empty value returns empty, not fallback
	m := map[string]string{
		"name": "",
	}
	got := Get(m, "name", "default")
	if got != "" {
		t.Fatalf("Get(%v, \"name\", \"default\") = %q, want empty string", m, got)
	}
}

func TestGetNilMap(t *testing.T) {
	// Edge case: Get from nil map returns fallback
	var m map[string]string
	got := Get(m, "any", "fallback")
	if got != "fallback" {
		t.Fatalf("Get(nil, \"any\", \"fallback\") = %q, want \"fallback\"", got)
	}
}
