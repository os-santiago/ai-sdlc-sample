package textfmt

import "testing"

func TestBannerEvenPadding(t *testing.T) {
	got := Banner("hi", 10)
	want := "=== hi ==="
	if got != want {
		t.Errorf("Banner(%q, %d) = %q; want %q", "hi", 10, got, want)
	}
}

func TestBannerOddPadding(t *testing.T) {
	got := Banner("ab", 11)
	want := "=== ab ===="
	if got != want {
		t.Errorf("Banner(%q, %d) = %q; want %q", "ab", 11, got, want)
	}
}

func TestBannerWidthTooSmall(t *testing.T) {
	got := Banner("hello", 5)
	want := "hello"
	if got != want {
		t.Errorf("Banner(%q, %d) = %q; want %q", "hello", 5, got, want)
	}
}

func TestBannerMultibyte(t *testing.T) {
	got := Banner("日本", 10)
	want := "=== 日本 ==="
	if got != want {
		t.Errorf("Banner(%q, %d) = %q; want %q", "日本", 10, got, want)
	}
}
