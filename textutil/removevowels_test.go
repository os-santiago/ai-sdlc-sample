package textutil

import "testing"

func TestRemoveVowels(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{
			name: "empty string",
			input: "",
			want: "",
		},
		{
			name: "mixed case",
			input: "Hello World",
			want: "Hll Wrld",
		},
		{
			name: "string without vowels",
			input: "bcdfg",
			want: "bcdfg",
		},
		{
			name: "all vowels",
			input: "aeiouAEIOU",
			want: "",
		},
		{
			name: "with numbers and punctuation",
			input: "Test123!",
			want: "Tst123!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveVowels(tt.input)
			if got != tt.want {
				t.Errorf("RemoveVowels(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
