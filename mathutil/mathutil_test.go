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

func TestMaxMixed(t *testing.T) {
	got, ok := Max([]int{3, -1, 7, 2})
	if !ok {
		t.Error("Max returned false for non-empty slice")
	}
	if got != 7 {
		t.Errorf("Max([3, -1, 7, 2]) = %d, want 7", got)
	}
}

func TestMaxSingleValue(t *testing.T) {
	got, ok := Max([]int{42})
	if !ok {
		t.Error("Max returned false for non-empty slice")
	}
	if got != 42 {
		t.Errorf("Max([42]) = %d, want 42", got)
	}
}

func TestMaxEmpty(t *testing.T) {
	if _, ok := Max(nil); ok {
		t.Error("Max(nil) returned true, want false")
	}
	if got, ok := Max(nil); ok || got != 0 {
		t.Errorf("Max(nil) = %d, %v, want 0, false", got, ok)
	}
	if _, ok := Max([]int{}); ok {
		t.Error("Max([]int{}) returned true, want false")
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

func TestRangeExample(t *testing.T) {
	got := Range(2, 5)
	want := []int{2, 3, 4}
	if len(got) != len(want) {
		t.Fatalf("Range(2, 5) length = %d, want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Range(2, 5)[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestRangePositive(t *testing.T) {
	got := Range(0, 3)
	want := []int{0, 1, 2}
	if len(got) != len(want) {
		t.Fatalf("Range(0, 3) length = %d, want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Range(0, 3)[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestRangeEqual(t *testing.T) {
	got := Range(5, 5)
	if got == nil {
		t.Error("Range(5, 5) returned nil, want empty non-nil slice")
	}
	if len(got) != 0 {
		t.Errorf("Range(5, 5) length = %d, want 0", len(got))
	}
}

func TestRangeReversed(t *testing.T) {
	got := Range(10, 5)
	if got == nil {
		t.Error("Range(10, 5) returned nil, want empty non-nil slice")
	}
	if len(got) != 0 {
		t.Errorf("Range(10, 5) length = %d, want 0", len(got))
	}
}

func TestRangeSingleValue(t *testing.T) {
	got := Range(7, 8)
	if len(got) != 1 || got[0] != 7 {
		t.Errorf("Range(7, 8) = %v, want [7]", got)
	}
}

func TestGCDExample1(t *testing.T) {
	if got := GCD(12, 18); got != 6 {
		t.Errorf("GCD(12, 18) = %d, want 6", got)
	}
}

func TestGCDExample2(t *testing.T) {
	if got := GCD(0, 5); got != 5 {
		t.Errorf("GCD(0, 5) = %d, want 5", got)
	}
}

func TestGCDExample3(t *testing.T) {
	if got := GCD(0, 0); got != 0 {
		t.Errorf("GCD(0, 0) = %d, want 0", got)
	}
}

func TestGCDNegativeFirst(t *testing.T) {
	if got := GCD(-12, 18); got != 6 {
		t.Errorf("GCD(-12, 18) = %d, want 6", got)
	}
}

func TestGCDNegativeSecond(t *testing.T) {
	if got := GCD(12, -18); got != 6 {
		t.Errorf("GCD(12, -18) = %d, want 6", got)
	}
}

func TestGCDBothNegative(t *testing.T) {
	if got := GCD(-12, -18); got != 6 {
		t.Errorf("GCD(-12, -18) = %d, want 6", got)
	}
}

func TestGCDCommutative(t *testing.T) {
	if got := GCD(18, 12); got != 6 {
		t.Errorf("GCD(18, 12) = %d, want 6", got)
	}
}

func TestGCDZeroFirst(t *testing.T) {
	if got := GCD(0, 10); got != 10 {
		t.Errorf("GCD(0, 10) = %d, want 10", got)
	}
}

func TestGCDLargeNumbers(t *testing.T) {
	if got := GCD(1071, 462); got != 21 {
		t.Errorf("GCD(1071, 462) = %d, want 21", got)
	}
}

func TestSignNegative(t *testing.T) {
	if got := Sign(-5); got != -1 {
		t.Errorf("Sign(-5) = %d, want -1", got)
	}
}

func TestSignZero(t *testing.T) {
	if got := Sign(0); got != 0 {
		t.Errorf("Sign(0) = %d, want 0", got)
	}
}

func TestSignPositive(t *testing.T) {
	if got := Sign(5); got != 1 {
		t.Errorf("Sign(5) = %d, want 1", got)
	}
}
func TestPowBasicExample(t *testing.T) {
	got, ok := Pow(2, 10)
	if !ok {
		t.Error("Pow(2, 10) returned false, want true")
	}
	if got != 1024 {
		t.Errorf("Pow(2, 10) = %d, want 1024", got)
	}
}

func TestPowZeroExponent(t *testing.T) {
	got, ok := Pow(5, 0)
	if !ok {
		t.Error("Pow(5, 0) returned false, want true")
	}
	if got != 1 {
		t.Errorf("Pow(5, 0) = %d, want 1", got)
	}
}

func TestPowNegativeExponent(t *testing.T) {
	got, ok := Pow(2, -1)
	if ok {
		t.Error("Pow(2, -1) returned true, want false")
	}
	if got != 0 {
		t.Errorf("Pow(2, -1) = %d, want 0", got)
	}
}

func TestPowOne(t *testing.T) {
	got, ok := Pow(7, 1)
	if !ok {
		t.Error("Pow(7, 1) returned false, want true")
	}
	if got != 7 {
		t.Errorf("Pow(7, 1) = %d, want 7", got)
	}
}

func TestPowZeroBase(t *testing.T) {
	got, ok := Pow(0, 5)
	if !ok {
		t.Error("Pow(0, 5) returned false, want true")
	}
	if got != 0 {
		t.Errorf("Pow(0, 5) = %d, want 0", got)
	}
}

func TestPowNegativeBase(t *testing.T) {
	got, ok := Pow(-2, 3)
	if !ok {
		t.Error("Pow(-2, 3) returned false, want true")
	}
	if got != -8 {
		t.Errorf("Pow(-2, 3) = %d, want -8", got)
	}
}

func TestPowLargerExponent(t *testing.T) {
	got, ok := Pow(3, 4)
	if !ok {
		t.Error("Pow(3, 4) returned false, want true")
	}
	if got != 81 {
		t.Errorf("Pow(3, 4) = %d, want 81", got)
	}
}

func TestAverageIntEmpty(t *testing.T) {
	got, ok := AverageInt(nil)
	if ok {
		t.Error("AverageInt(nil) returned true, want false")
	}
	if got != 0 {
		t.Errorf("AverageInt(nil) = %d, want 0", got)
	}
	
	got, ok = AverageInt([]int{})
	if ok {
		t.Error("AverageInt([]int{}) returned true, want false")
	}
	if got != 0 {
		t.Errorf("AverageInt([]int{}) = %d, want 0", got)
	}
}

func TestAverageIntExample1(t *testing.T) {
	got, ok := AverageInt([]int{1, 2, 4})
	if !ok {
		t.Error("AverageInt([1, 2, 4]) returned false, want true")
	}
	if got != 2 {
		t.Errorf("AverageInt([1, 2, 4]) = %d, want 2 (floor of 2.333...)", got)
	}
}

func TestAverageIntExample2(t *testing.T) {
	got, ok := AverageInt([]int{-3, -4})
	if !ok {
		t.Error("AverageInt([-3, -4]) returned false, want true")
	}
	if got != -4 {
		t.Errorf("AverageInt([-3, -4]) = %d, want -4 (floor of -3.5)", got)
	}
}

func TestAverageIntPositiveExact(t *testing.T) {
	got, ok := AverageInt([]int{2, 4, 6})
	if !ok {
		t.Error("AverageInt([2, 4, 6]) returned false, want true")
	}
	if got != 4 {
		t.Errorf("AverageInt([2, 4, 6]) = %d, want 4", got)
	}
}

func TestAverageIntNegativeExact(t *testing.T) {
	got, ok := AverageInt([]int{-2, -4, -6})
	if !ok {
		t.Error("AverageInt([-2, -4, -6]) returned false, want true")
	}
	if got != -4 {
		t.Errorf("AverageInt([-2, -4, -6]) = %d, want -4", got)
	}
}

func TestAverageIntSingleValue(t *testing.T) {
	got, ok := AverageInt([]int{7})
	if !ok {
		t.Error("AverageInt([7]) returned false, want true")
	}
	if got != 7 {
		t.Errorf("AverageInt([7]) = %d, want 7", got)
	}
}

func TestAverageIntMixed(t *testing.T) {
	got, ok := AverageInt([]int{-5, 10, -2})
	if !ok {
		t.Error("AverageInt([-5, 10, -2]) returned false, want true")
	}
	if got != 1 {
		t.Errorf("AverageInt([-5, 10, -2]) = %d, want 1 (floor of 1.0)", got)
	}
}
