package textfmt

import "testing"

func TestTruncateShort(t *testing.T) {
	if got := Truncate("hi", 10); got != "hi" {
		t.Errorf("Truncate = %q", got)
	}
}

func TestTitleASCII(t *testing.T) {
	if got := Title("hello world"); got != "Hello World" {
		t.Errorf("Title = %q", got)
	}
}
