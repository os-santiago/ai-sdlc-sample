package textutil

import "testing"

func TestFirstNonSpace(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  rune
		wantOk bool
	}{
		{
			name:   "empty input",
			input:  "",
			want:   0,
			wantOk: false,
		},
		{
			name:   "only spaces",
			input:  "   ",
			want:   0,
			wantOk: false,
		},
		{
			name:   "leading spaces with ASCII char",
			input:  "   hello",
			want:   'h',
			wantOk: true,
		},
		{
			name:   "leading spaces with multi-byte rune",
			input:  "   世界",
			want:   '世',
			wantOk: true,
		},
		{
			name:   "no leading spaces",
			input:  "hello",
			want:   'h',
			wantOk: true,
		},
		{
			name:   "multi-byte first rune no spaces",
			input:  "世界",
			want:   '世',
			wantOk: true,
		},
		{
			name:   "tab and newline prefix",
			input:  "\t\n  test",
			want:   't',
			wantOk: true,
		},
		{
			name:   "only whitespace chars",
			input:  "\t\n\r  ",
			want:   0,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := FirstNonSpace(tt.input)
			if got != tt.want || gotOk != tt.wantOk {
				t.Errorf("FirstNonSpace(%q) = (%q, %v), want (%q, %v)", tt.input, got, gotOk, tt.want, tt.wantOk)
			}
		})
	}
}
