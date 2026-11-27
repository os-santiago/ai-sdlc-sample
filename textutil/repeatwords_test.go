package textutil

import "testing"

func TestRepeatWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    int
		want string
	}{
		// n = 0 cases
		{
			name: "n=0 with single word",
			s:    "hello",
			n:    0,
			want: "",
		},
		{
			name: "n=0 with multiple words",
			s:    "hello world",
			n:    0,
			want: "",
		},
		{
			name: "n=-1 with multiple words",
			s:    "foo bar baz",
			n:    -1,
			want: "",
		},
		// n = 1 cases
		{
			name: "n=1 with single word",
			s:    "hello",
			n:    1,
			want: "hello",
		},
		{
			name: "n=1 with multiple words",
			s:    "hello world",
			n:    1,
			want: "hello world",
		},
		{
			name: "n=1 with extra whitespace",
			s:    "  hello   world  ",
			n:    1,
			want: "hello world",
		},
		{
			name: "n=1 with three words",
			s:    "foo bar baz",
			n:    1,
			want: "foo bar baz",
		},
		// n = 2 cases
		{
			name: "n=2 with single word",
			s:    "hello",
			n:    2,
			want: "hello hello",
		},
		{
			name: "n=2 with two words",
			s:    "hello world",
			n:    2,
			want: "hello hello world world",
		},
		{
			name: "n=2 with three words",
			s:    "foo bar baz",
			n:    2,
			want: "foo foo bar bar baz baz",
		},
		{
			name: "n=2 with multiple words and whitespace",
			s:    "  one  two  three  ",
			n:    2,
			want: "one one two two three three",
		},
		// Edge cases
		{
			name: "empty string with n=1",
			s:    "",
			n:    1,
			want: "",
		},
		{
			name: "empty string with n=2",
			s:    "",
			n:    2,
			want: "",
		},
		{
			name: "whitespace only with n=2",
			s:    "   ",
			n:    2,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RepeatWords(tt.s, tt.n)
			if got != tt.want {
				t.Errorf("RepeatWords(%q, %d) = %q, want %q", tt.s, tt.n, got, tt.want)
			}
		})
	}
}
