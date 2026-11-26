package textutil

import "testing"

func TestCollapseDashes(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		// Acceptance criteria: no dashes
		{"", ""},
		{"hello", "hello"},
		{"hello world", "hello world"},

		// Acceptance criteria: a single dash
		{"hello-world", "hello-world"},
		{"-", "-"},
		{"a-b-c", "a-b-c"},

		// Acceptance criteria: a run of dashes
		{"hello--world", "hello-world"},
		{"hello---world", "hello-world"},
		{"hello----world", "hello-world"},
		{"a-----b", "a-b"},

		// Acceptance criteria: dashes at both ends
		{"--start", "-start"},
		{"end--", "end-"},
		{"--both--", "-both-"},
		{"---start-and-end---", "-start-and-end-"},

		// Additional test cases
		{"multiple--runs--here", "multiple-runs-here"},
		{"a--b---c----d", "a-b-c-d"},
		{"------", "-"},
	}

	for _, tt := range tests {
		if got := CollapseDashes(tt.in); got != tt.want {
			t.Errorf("CollapseDashes(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
