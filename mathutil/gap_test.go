package mathutil

import "testing"

func TestGapNil(t *testing.T) {
	got := Gap(nil)
	if got != 0 {
		t.Errorf("Gap(nil) = %v, want 0", got)
	}
}

func TestGapOneValue(t *testing.T) {
	got := Gap([]int{5})
	if got != 0 {
		t.Errorf("Gap([]int{5}) = %v, want 0", got)
	}
}

func TestGapIncreasingSequence(t *testing.T) {
	got := Gap([]int{1, 3, 5, 10})
	want := 5 // gap between 5 and 10
	if got != want {
		t.Errorf("Gap([]int{1, 3, 5, 10}) = %v, want %v", got, want)
	}
}

func TestGapDecreasingSequence(t *testing.T) {
	got := Gap([]int{10, 5, 3, 1})
	want := 5 // gap between 10 and 5
	if got != want {
		t.Errorf("Gap([]int{10, 5, 3, 1}) = %v, want %v", got, want)
	}
}

func TestGapMixedSequence(t *testing.T) {
	got := Gap([]int{5, 2, 8, 3})
	want := 6 // gap between 2 and 8
	if got != want {
		t.Errorf("Gap([]int{5, 2, 8, 3}) = %v, want %v", got, want)
	}
}

func TestGapAllSameValues(t *testing.T) {
	got := Gap([]int{5, 5, 5, 5})
	want := 0
	if got != want {
		t.Errorf("Gap([]int{5, 5, 5, 5}) = %v, want %v", got, want)
	}
}

func TestGapTwoValues(t *testing.T) {
	got := Gap([]int{1, 10})
	want := 9
	if got != want {
		t.Errorf("Gap([]int{1, 10}) = %v, want %v", got, want)
	}
}
