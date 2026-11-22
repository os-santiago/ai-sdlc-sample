package textfmt

import "testing"

func TestTruncateShort(t *testing.T) {
	if got := Truncate("hi", 10); got != "hi" {
		t.Errorf("Truncate = %q", got)
	}
}

func TestTruncateMultibyte(t *testing.T) {
	if got := Truncate("héllo wörld", 8); got != "héllo..." {
		t.Errorf("Truncate = %q, want %q", got, "héllo...")
	}
}

func TestTitleASCII(t *testing.T) {
	if got := Title("hello world"); got != "Hello World" {
		t.Errorf("Title = %q", got)
	}
}

func TestTitleMultibyte(t *testing.T) {
	if got := Title("élan vital"); got != "Élan Vital" {
		t.Errorf("Title = %q, want %q", got, "Élan Vital")
	}
}

func TestPadLeftASCII(t *testing.T) {
	if got := PadLeft("hi", 5, '.'); got != "...hi" {
		t.Errorf("PadLeft = %q, want %q", got, "...hi")
	}
}

func TestPadLeftMultibyte(t *testing.T) {
	if got := PadLeft("é", 3, '.'); got != "..é" {
		t.Errorf("PadLeft = %q, want %q", got, "..é")
	}
}

func TestPadLeftLonger(t *testing.T) {
	if got := PadLeft("hello", 3, '.'); got != "hello" {
		t.Errorf("PadLeft = %q, want %q", got, "hello")
	}
}
