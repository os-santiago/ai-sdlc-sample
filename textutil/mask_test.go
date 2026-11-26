package textutil

import "testing"

func TestMask(t *testing.T) {
	tests := []struct {
		name string
		s    string
		keep int
		want string
	}{
		{
			name: "keep 0 masks all",
			s:    "password",
			keep: 0,
			want: "********",
		},
		{
			name: "partial mask",
			s:    "password",
			keep: 4,
			want: "****word",
		},
		{
			name: "keep larger than length",
			s:    "password",
			keep: 20,
			want: "password",
		},
		{
			name: "multibyte string",
			s:    "hello世界",
			keep: 2,
			want: "*****世界",
		},
		{
			name: "negative keep masks all",
			s:    "test",
			keep: -5,
			want: "****",
		},
		{
			name: "empty string",
			s:    "",
			keep: 0,
			want: "",
		},
		{
			name: "keep equals length",
			s:    "test",
			keep: 4,
			want: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mask(tt.s, tt.keep)
			if got != tt.want {
				t.Errorf("Mask(%q, %d) = %q, want %q", tt.s, tt.keep, got, tt.want)
			}
		})
	}
}
