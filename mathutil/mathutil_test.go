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

func TestSumNil(t *testing.T) {
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum(nil) = %v, want 0", got)
	}
}

func TestSumEmpty(t *testing.T) {
	if got := Sum([]float64{}); got != 0 {
		t.Errorf("Sum([]float64{}) = %v, want 0", got)
	}
}

func TestSumExample(t *testing.T) {
	if got := Sum([]float64{1.5, 2.5, -1}); got != 3 {
		t.Errorf("Sum([1.5, 2.5, -1]) = %v, want 3", got)
	}
}

func TestSumPositive(t *testing.T) {
	if got := Sum([]float64{1, 2, 3, 4}); got != 10 {
		t.Errorf("Sum([1, 2, 3, 4]) = %v, want 10", got)
	}
}

func TestSumNegative(t *testing.T) {
	if got := Sum([]float64{-1, -2, -3}); got != -6 {
		t.Errorf("Sum([-1, -2, -3]) = %v, want -6", got)
	}
}

func TestSumSingleValue(t *testing.T) {
	if got := Sum([]float64{42}); got != 42 {
		t.Errorf("Sum([42]) = %v, want 42", got)
	}
}

func TestProductNil(t *testing.T) {
	if got := Product(nil); got != 1 {
		t.Errorf("Product(nil) = %v, want 1", got)
	}
}

func TestProductEmpty(t *testing.T) {
	if got := Product([]int{}); got != 1 {
		t.Errorf("Product([]int{}) = %v, want 1", got)
	}
}

func TestProductExample(t *testing.T) {
	if got := Product([]int{2, 3, 4}); got != 24 {
		t.Errorf("Product([2, 3, 4]) = %v, want 24", got)
	}
}

func TestProductSingleValue(t *testing.T) {
	if got := Product([]int{7}); got != 7 {
		t.Errorf("Product([7]) = %v, want 7", got)
	}
}

func TestProductWithZero(t *testing.T) {
	if got := Product([]int{2, 0, 5}); got != 0 {
		t.Errorf("Product([2, 0, 5]) = %v, want 0", got)
	}
}

func TestProductNegative(t *testing.T) {
	if got := Product([]int{-2, 3, -4}); got != 24 {
		t.Errorf("Product([-2, 3, -4]) = %v, want 24", got)
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

func TestMinMaxNil(t *testing.T) {
	min, max, ok := MinMax(nil)
	if ok {
		t.Error("MinMax(nil) returned true, want false")
	}
	if min != 0 || max != 0 {
		t.Errorf("MinMax(nil) = %d, %d, want 0, 0", min, max)
	}
}

func TestMinMaxEmpty(t *testing.T) {
	min, max, ok := MinMax([]int{})
	if ok {
		t.Error("MinMax([]int{}) returned true, want false")
	}
	if min != 0 || max != 0 {
		t.Errorf("MinMax([]int{}) = %d, %d, want 0, 0", min, max)
	}
}

func TestMinMaxExample(t *testing.T) {
	min, max, ok := MinMax([]int{3, -1, 7})
	if !ok {
		t.Error("MinMax returned false for non-empty slice")
	}
	if min != -1 {
		t.Errorf("MinMax([3, -1, 7]) min = %d, want -1", min)
	}
	if max != 7 {
		t.Errorf("MinMax([3, -1, 7]) max = %d, want 7", max)
	}
}

func TestMinMaxSingleValue(t *testing.T) {
	min, max, ok := MinMax([]int{42})
	if !ok {
		t.Error("MinMax returned false for non-empty slice")
	}
	if min != 42 || max != 42 {
		t.Errorf("MinMax([42]) = %d, %d, want 42, 42", min, max)
	}
}

func TestMinMaxAllSame(t *testing.T) {
	min, max, ok := MinMax([]int{5, 5, 5, 5})
	if !ok {
		t.Error("MinMax returned false for non-empty slice")
	}
	if min != 5 || max != 5 {
		t.Errorf("MinMax([5, 5, 5, 5]) = %d, %d, want 5, 5", min, max)
	}
}

func TestMinMaxNegativeValues(t *testing.T) {
	min, max, ok := MinMax([]int{-10, -20, -5, -15})
	if !ok {
		t.Error("MinMax returned false for non-empty slice")
	}
	if min != -20 {
		t.Errorf("MinMax([-10, -20, -5, -15]) min = %d, want -20", min)
	}
	if max != -5 {
		t.Errorf("MinMax([-10, -20, -5, -15]) max = %d, want -5", max)
	}
}

func TestAbsNegative(t *testing.T) {
	if got := Abs(-3); got != 3 {
		t.Errorf("Abs(-3) = %d, want 3", got)
	}
}

func TestAbsZero(t *testing.T) {
	if got := Abs(0); got != 0 {
		t.Errorf("Abs(0) = %d, want 0", got)
	}
}

func TestAbsPositive(t *testing.T) {
	if got := Abs(5); got != 5 {
		t.Errorf("Abs(5) = %d, want 5", got)
	}
}
