package mathutil

import "testing"

func TestClampBelowRange(t *testing.T) {
	if got := Clamp(-5, 0, 10); got != 0 {
		t.Errorf("Clamp(-5, 0, 10) = %d, want 0", got)
	}
}

func TestClampAboveRange(t *testing.T) {
	if got := Clamp(15, 0, 10); got != 10 {
		t.Errorf("Clamp(15, 0, 10) = %d, want 10", got)
	}
}

func TestClampInRange(t *testing.T) {
	if got := Clamp(5, 0, 10); got != 5 {
		t.Errorf("Clamp(5, 0, 10) = %d, want 5", got)
	}
}

func TestMedianOdd(t *testing.T) {
	if got := Median([]float64{3, 1, 2}); got != 2 {
		t.Errorf("Median = %v, want 2", got)
	}
}

func TestMedianEven(t *testing.T) {
	if got := Median([]float64{1, 2, 3, 4}); got != 2.5 {
		t.Errorf("Median([1, 2, 3, 4]) = %v, want 2.5", got)
	}
}

func TestMean(t *testing.T) {
	if got := Mean([]float64{1, 2, 3}); got != 2 {
		t.Errorf("Mean = %v, want 2", got)
	}
}
