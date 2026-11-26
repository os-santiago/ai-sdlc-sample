package mathutil

import "testing"

func TestStepsOne(t *testing.T) {
	got := Steps(1)
	want := 0
	if got != want {
		t.Errorf("Steps(1) = %d, want %d", got, want)
	}
}

func TestStepsSix(t *testing.T) {
	got := Steps(6)
	want := 8
	if got != want {
		t.Errorf("Steps(6) = %d, want %d", got, want)
	}
}

func TestStepsTwentySeven(t *testing.T) {
	got := Steps(27)
	want := 111
	if got != want {
		t.Errorf("Steps(27) = %d, want %d", got, want)
	}
}

func TestStepsZero(t *testing.T) {
	got := Steps(0)
	want := -1
	if got != want {
		t.Errorf("Steps(0) = %d, want %d", got, want)
	}
}
