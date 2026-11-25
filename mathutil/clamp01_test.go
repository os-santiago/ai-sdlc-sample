package mathutil

import "testing"

func TestClamp01BelowZero(t *testing.T) {
	got := Clamp01(-0.5)
	if got != 0 {
		t.Errorf("Clamp01(-0.5) = %v, want 0", got)
	}
}

func TestClamp01AboveOne(t *testing.T) {
	got := Clamp01(1.5)
	if got != 1 {
		t.Errorf("Clamp01(1.5) = %v, want 1", got)
	}
}

func TestClamp01InRange(t *testing.T) {
	got := Clamp01(0.5)
	if got != 0.5 {
		t.Errorf("Clamp01(0.5) = %v, want 0.5", got)
	}
}