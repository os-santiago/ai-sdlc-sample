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

func TestMeanEmpty(t *testing.T) {
	if got := Mean(nil); got != 0 {
		t.Errorf("Mean(nil) = %v, want 0", got)
	}
	if got := Mean([]float64{}); got != 0 {
		t.Errorf("Mean([]float64{}) = %v, want 0", got)
	}
}

func TestModeClearWinner(t *testing.T) {
	got, ok := Mode([]int{1, 2, 2, 3})
	if !ok {
		t.Error("Mode returned false for non-empty slice")
	}
	if got != 2 {
		t.Errorf("Mode([1, 2, 2, 3]) = %d, want 2", got)
	}
}

func TestModeTie(t *testing.T) {
	got, ok := Mode([]int{3, 1, 1, 3, 2})
	if !ok {
		t.Error("Mode returned false for non-empty slice")
	}
	if got != 1 {
		t.Errorf("Mode([3, 1, 1, 3, 2]) = %d, want 1 (smallest of tied values)", got)
	}
}

func TestModeEmpty(t *testing.T) {
	if _, ok := Mode(nil); ok {
		t.Error("Mode(nil) returned true, want false")
	}
	if _, ok := Mode([]int{}); ok {
		t.Error("Mode([]int{}) returned true, want false")
	}
}

func TestStdDevExample(t *testing.T) {
	got := StdDev([]float64{2, 4, 4, 4, 5, 5, 7, 9})
	if got != 2 {
		t.Errorf("StdDev([2, 4, 4, 4, 5, 5, 7, 9]) = %v, want 2", got)
	}
}

func TestStdDevEmpty(t *testing.T) {
	if got := StdDev(nil); got != 0 {
		t.Errorf("StdDev(nil) = %v, want 0", got)
	}
	if got := StdDev([]float64{}); got != 0 {
		t.Errorf("StdDev([]float64{}) = %v, want 0", got)
	}
}

func TestStdDevSingleValue(t *testing.T) {
	got := StdDev([]float64{5})
	if got != 0 {
		t.Errorf("StdDev([5]) = %v, want 0 (no deviation from mean)", got)
	}
}
