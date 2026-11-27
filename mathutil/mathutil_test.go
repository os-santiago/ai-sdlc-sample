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

func TestLCMExample1(t *testing.T) {
	got := LCM(4, 6)
	if got != 12 {
		t.Errorf("LCM(4, 6) = %d, want 12", got)
	}
}

func TestLCMZeroFirst(t *testing.T) {
	got := LCM(0, 5)
	if got != 0 {
		t.Errorf("LCM(0, 5) = %d, want 0", got)
	}
}

func TestLCMZeroSecond(t *testing.T) {
	got := LCM(5, 0)
	if got != 0 {
		t.Errorf("LCM(5, 0) = %d, want 0", got)
	}
}

func TestLCMBothZero(t *testing.T) {
	got := LCM(0, 0)
	if got != 0 {
		t.Errorf("LCM(0, 0) = %d, want 0", got)
	}
}

func TestLCMPositive(t *testing.T) {
	got := LCM(12, 18)
	if got != 36 {
		t.Errorf("LCM(12, 18) = %d, want 36", got)
	}
}

func TestLCMNegativeFirst(t *testing.T) {
	got := LCM(-4, 6)
	if got != 12 {
		t.Errorf("LCM(-4, 6) = %d, want 12 (using absolute values)", got)
	}
}

func TestLCMNegativeSecond(t *testing.T) {
	got := LCM(4, -6)
	if got != 12 {
		t.Errorf("LCM(4, -6) = %d, want 12 (using absolute values)", got)
	}
}

func TestLCMBothNegative(t *testing.T) {
	got := LCM(-4, -6)
	if got != 12 {
		t.Errorf("LCM(-4, -6) = %d, want 12 (using absolute values)", got)
	}
}

func TestLCMIdenticalValues(t *testing.T) {
	got := LCM(7, 7)
	if got != 7 {
		t.Errorf("LCM(7, 7) = %d, want 7", got)
	}
}

func TestLCMCoprime(t *testing.T) {
	got := LCM(5, 7)
	if got != 35 {
		t.Errorf("LCM(5, 7) = %d, want 35 (coprime numbers)", got)
	}
}

func TestIsPrimeNegative(t *testing.T) {
	if got := IsPrime(-5); got {
		t.Error("IsPrime(-5) = true, want false (negative numbers are not prime)")
	}
	if got := IsPrime(-2); got {
		t.Error("IsPrime(-2) = true, want false (negative numbers are not prime)")
	}
}

func TestIsPrimeZero(t *testing.T) {
	if got := IsPrime(0); got {
		t.Error("IsPrime(0) = true, want false (0 is not prime)")
	}
}

func TestIsPrimeOne(t *testing.T) {
	if got := IsPrime(1); got {
		t.Error("IsPrime(1) = true, want false (1 is not prime)")
	}
}

func TestIsPrimeTwo(t *testing.T) {
	if got := IsPrime(2); !got {
		t.Error("IsPrime(2) = false, want true (2 is prime)")
	}
}

func TestIsPrimeThirteen(t *testing.T) {
	if got := IsPrime(13); !got {
		t.Error("IsPrime(13) = false, want true (13 is prime)")
	}
}

func TestIsPrimeFifteen(t *testing.T) {
	if got := IsPrime(15); got {
		t.Error("IsPrime(15) = true, want false (15 = 3 × 5 is composite)")
	}
}

func TestIsPrimeSmallPrimes(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	for _, p := range primes {
		if got := IsPrime(p); !got {
			t.Errorf("IsPrime(%d) = false, want true", p)
		}
	}
}

func TestIsPrimeSmallComposites(t *testing.T) {
	composites := []int{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28}
	for _, c := range composites {
		if got := IsPrime(c); got {
			t.Errorf("IsPrime(%d) = true, want false", c)
		}
	}
}

func TestIsPrimeFour(t *testing.T) {
	if got := IsPrime(4); got {
		t.Error("IsPrime(4) = true, want false (4 = 2 × 2 is composite)")
	}
}

func TestIsPrimeNine(t *testing.T) {
	if got := IsPrime(9); got {
		t.Error("IsPrime(9) = true, want false (9 = 3 × 3 is composite)")
	}
}

func TestIsPrimeLargePrime(t *testing.T) {
	if got := IsPrime(97); !got {
		t.Error("IsPrime(97) = false, want true (97 is prime)")
	}
}

func TestIsPrimeLargeComposite(t *testing.T) {
	if got := IsPrime(100); got {
		t.Error("IsPrime(100) = true, want false (100 = 10 × 10 is composite)")
	}
}

func TestIsPrimeEvenComposite(t *testing.T) {
	if got := IsPrime(50); got {
		t.Error("IsPrime(50) = true, want false (50 = 2 × 25 is composite)")
	}
}

func TestFactorialZero(t *testing.T) {
	got, ok := Factorial(0)
	if !ok {
		t.Error("Factorial(0) returned false, want true")
	}
	if got != 1 {
		t.Errorf("Factorial(0) = %d, want 1", got)
	}
}

func TestFactorialOne(t *testing.T) {
	got, ok := Factorial(1)
	if !ok {
		t.Error("Factorial(1) returned false, want true")
	}
	if got != 1 {
		t.Errorf("Factorial(1) = %d, want 1", got)
	}
}

func TestFactorialFive(t *testing.T) {
	got, ok := Factorial(5)
	if !ok {
		t.Error("Factorial(5) returned false, want true")
	}
	if got != 120 {
		t.Errorf("Factorial(5) = %d, want 120", got)
	}
}

func TestFactorialTwenty(t *testing.T) {
	got, ok := Factorial(20)
	if !ok {
		t.Error("Factorial(20) returned false, want true")
	}
	// 20! = 2432902008176640000
	if got != 2432902008176640000 {
		t.Errorf("Factorial(20) = %d, want 2432902008176640000", got)
	}
}

func TestFactorialTwentyOne(t *testing.T) {
	got, ok := Factorial(21)
	if ok {
		t.Error("Factorial(21) returned true, want false (overflow guard)")
	}
	if got != 0 {
		t.Errorf("Factorial(21) = %d, want 0", got)
	}
}

func TestFactorialNegative(t *testing.T) {
	got, ok := Factorial(-1)
	if ok {
		t.Error("Factorial(-1) returned true, want false (negative input)")
	}
	if got != 0 {
		t.Errorf("Factorial(-1) = %d, want 0", got)
	}
}

func TestFactorialNegativeLarge(t *testing.T) {
	got, ok := Factorial(-10)
	if ok {
		t.Error("Factorial(-10) returned true, want false (negative input)")
	}
	if got != 0 {
		t.Errorf("Factorial(-10) = %d, want 0", got)
	}
}

func TestFactorialOverflow(t *testing.T) {
	got, ok := Factorial(25)
	if ok {
		t.Error("Factorial(25) returned true, want false (overflow guard)")
	}
	if got != 0 {
		t.Errorf("Factorial(25) = %d, want 0", got)
	}
}

func TestDigitsZero(t *testing.T) {
	got := Digits(0)
	if got != 1 {
		t.Errorf("Digits(0) = %d, want 1", got)
	}
}

func TestDigitsPositiveSingle(t *testing.T) {
	got := Digits(5)
	if got != 1 {
		t.Errorf("Digits(5) = %d, want 1", got)
	}
}

func TestDigitsPositiveThree(t *testing.T) {
	got := Digits(123)
	if got != 3 {
		t.Errorf("Digits(123) = %d, want 3", got)
	}
}

func TestDigitsNegativeSingle(t *testing.T) {
	got := Digits(-7)
	if got != 1 {
		t.Errorf("Digits(-7) = %d, want 1 (ignoring sign)", got)
	}
}

func TestDigitsNegativeThree(t *testing.T) {
	got := Digits(-123)
	if got != 3 {
		t.Errorf("Digits(-123) = %d, want 3 (ignoring sign)", got)
	}
}

func TestDigitsLargePositive(t *testing.T) {
	got := Digits(9999)
	if got != 4 {
		t.Errorf("Digits(9999) = %d, want 4", got)
	}
}

func TestDigitsLargeNegative(t *testing.T) {
	got := Digits(-9999)
	if got != 4 {
		t.Errorf("Digits(-9999) = %d, want 4 (ignoring sign)", got)
	}
}

func TestSumOfSquaresEmpty(t *testing.T) {
	if got := SumOfSquares([]int{}); got != 0 {
		t.Errorf("SumOfSquares([]int{}) = %d, want 0", got)
	}
}

func TestSumOfSquaresNil(t *testing.T) {
	if got := SumOfSquares(nil); got != 0 {
		t.Errorf("SumOfSquares(nil) = %d, want 0", got)
	}
}

func TestSumOfSquaresPositive(t *testing.T) {
	if got := SumOfSquares([]int{1, 2, 3}); got != 14 {
		t.Errorf("SumOfSquares([1, 2, 3]) = %d, want 14 (1 + 4 + 9)", got)
	}
}

func TestSumOfSquaresNegative(t *testing.T) {
	if got := SumOfSquares([]int{-1, -2, -3}); got != 14 {
		t.Errorf("SumOfSquares([-1, -2, -3]) = %d, want 14 (1 + 4 + 9)", got)
	}
}

func TestSumOfSquaresMixed(t *testing.T) {
	if got := SumOfSquares([]int{-2, 3}); got != 13 {
		t.Errorf("SumOfSquares([-2, 3]) = %d, want 13 (4 + 9)", got)
	}
}

func TestSumOfSquaresSingleValue(t *testing.T) {
	if got := SumOfSquares([]int{5}); got != 25 {
		t.Errorf("SumOfSquares([5]) = %d, want 25", got)
	}
}

func TestSumOfSquaresZero(t *testing.T) {
	if got := SumOfSquares([]int{0}); got != 0 {
		t.Errorf("SumOfSquares([0]) = %d, want 0", got)
	}
}

func TestSumOfSquaresMultipleValues(t *testing.T) {
	if got := SumOfSquares([]int{1, 2, 3, 4, 5}); got != 55 {
		t.Errorf("SumOfSquares([1, 2, 3, 4, 5]) = %d, want 55 (1 + 4 + 9 + 16 + 25)", got)
	}
}

func TestAverageFloatEmpty(t *testing.T) {
	got, ok := AverageFloat([]float64{})
	if ok {
		t.Errorf("AverageFloat([]) returned ok=true, want false")
	}
	if got != 0 {
		t.Errorf("AverageFloat([]) = %v, want 0", got)
	}
}

func TestAverageFloatSingleValue(t *testing.T) {
	got, ok := AverageFloat([]float64{5.5})
	if !ok {
		t.Errorf("AverageFloat([5.5]) returned ok=false, want true")
	}
	if got != 5.5 {
		t.Errorf("AverageFloat([5.5]) = %v, want 5.5", got)
	}
}

func TestAverageFloatMultipleValues(t *testing.T) {
	got, ok := AverageFloat([]float64{1.0, 2.0, 3.0, 4.0, 5.0})
	if !ok {
		t.Errorf("AverageFloat([1.0, 2.0, 3.0, 4.0, 5.0]) returned ok=false, want true")
	}
	if got != 3.0 {
		t.Errorf("AverageFloat([1.0, 2.0, 3.0, 4.0, 5.0]) = %v, want 3.0", got)
	}
}

func TestIsEvenZero(t *testing.T) {
	if got := IsEven(0); !got {
		t.Errorf("IsEven(0) = %v, want true", got)
	}
}

func TestIsEvenPositiveEven(t *testing.T) {
	if got := IsEven(2); !got {
		t.Errorf("IsEven(2) = %v, want true", got)
	}
}

func TestIsEvenPositiveOdd(t *testing.T) {
	if got := IsEven(1); got {
		t.Errorf("IsEven(1) = %v, want false", got)
	}
}

func TestIsEvenNegativeEven(t *testing.T) {
	if got := IsEven(-2); !got {
		t.Errorf("IsEven(-2) = %v, want true", got)
	}
}

func TestIsEvenNegativeOdd(t *testing.T) {
	if got := IsEven(-1); got {
		t.Errorf("IsEven(-1) = %v, want false", got)
	}
}

func TestIsEvenLargePositiveEven(t *testing.T) {
	if got := IsEven(1000); !got {
		t.Errorf("IsEven(1000) = %v, want true", got)
	}
}

func TestIsEvenLargeNegativeOdd(t *testing.T) {
	if got := IsEven(-999); got {
		t.Errorf("IsEven(-999) = %v, want false", got)
	}
}

func TestIsOddZero(t *testing.T) {
	if got := IsOdd(0); got {
		t.Errorf("IsOdd(0) = %v, want false", got)
	}
}

func TestIsOddPositiveOdd(t *testing.T) {
	if got := IsOdd(1); !got {
		t.Errorf("IsOdd(1) = %v, want true", got)
	}
}

func TestIsOddPositiveEven(t *testing.T) {
	if got := IsOdd(2); got {
		t.Errorf("IsOdd(2) = %v, want false", got)
	}
}

func TestIsOddNegativeOdd(t *testing.T) {
	if got := IsOdd(-1); !got {
		t.Errorf("IsOdd(-1) = %v, want true", got)
	}
}

func TestIsOddNegativeEven(t *testing.T) {
	if got := IsOdd(-2); got {
		t.Errorf("IsOdd(-2) = %v, want false", got)
	}
}

func TestIsOddLargeValues(t *testing.T) {
	if got := IsOdd(999); !got {
		t.Errorf("IsOdd(999) = %v, want true", got)
	}
	if got := IsOdd(-1000); got {
		t.Errorf("IsOdd(-1000) = %v, want false", got)
	}
}

func TestSquarePositive(t *testing.T) {
	if got := Square(5); got != 25 {
		t.Errorf("Square(5) = %d, want 25", got)
	}
}

func TestSquareNegative(t *testing.T) {
	if got := Square(-3); got != 9 {
		t.Errorf("Square(-3) = %d, want 9", got)
	}
}

func TestSquareZero(t *testing.T) {
	if got := Square(0); got != 0 {
		t.Errorf("Square(0) = %d, want 0", got)
	}
}

func TestSquareLargePositive(t *testing.T) {
	if got := Square(100); got != 10000 {
		t.Errorf("Square(100) = %d, want 10000", got)
	}
}

func TestSquareLargeNegative(t *testing.T) {
	if got := Square(-10); got != 100 {
		t.Errorf("Square(-10) = %d, want 100", got)
	}
}

func TestHalfZero(t *testing.T) {
	if got := Half(0); got != 0 {
		t.Errorf("Half(0) = %d, want 0", got)
	}
}

func TestHalfPositiveEven(t *testing.T) {
	if got := Half(4); got != 2 {
		t.Errorf("Half(4) = %d, want 2", got)
	}
}

func TestHalfPositiveOdd(t *testing.T) {
	if got := Half(5); got != 2 {
		t.Errorf("Half(5) = %d, want 2", got)
	}
}

func TestHalfNegativeEven(t *testing.T) {
	if got := Half(-4); got != -2 {
		t.Errorf("Half(-4) = %d, want -2", got)
	}
}

func TestHalfNegativeOdd(t *testing.T) {
	if got := Half(-5); got != -2 {
		t.Errorf("Half(-5) = %d, want -2", got)
	}
}

func TestHalfLargePositiveOdd(t *testing.T) {
	if got := Half(99); got != 49 {
		t.Errorf("Half(99) = %d, want 49", got)
	}
}

func TestHalfLargeNegativeOdd(t *testing.T) {
	if got := Half(-99); got != -49 {
		t.Errorf("Half(-99) = %d, want -49", got)
	}
}

func TestClipNil(t *testing.T) {
	got := Clip(nil, 0, 10)
	if got != nil {
		t.Errorf("Clip(nil, 0, 10) = %v, want nil", got)
	}
}

func TestClipEmpty(t *testing.T) {
	got := Clip([]int{}, 0, 10)
	if got == nil {
		t.Errorf("Clip([]int{}, 0, 10) = nil, want empty slice")
	}
	if len(got) != 0 {
		t.Errorf("Clip([]int{}, 0, 10) = %v, want []", got)
	}
}

func TestClipValuesInside(t *testing.T) {
	input := []int{1, 5, 9}
	got := Clip(input, 0, 10)
	want := []int{1, 5, 9}

	if len(got) != len(want) {
		t.Fatalf("Clip(%v, 0, 10) length = %d, want %d", input, len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Clip(%v, 0, 10)[%d] = %d, want %d", input, i, got[i], want[i])
		}
	}
}

func TestClipValuesBelow(t *testing.T) {
	input := []int{-5, -10, -1}
	got := Clip(input, 0, 10)
	want := []int{0, 0, 0}

	if len(got) != len(want) {
		t.Fatalf("Clip(%v, 0, 10) length = %d, want %d", input, len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Clip(%v, 0, 10)[%d] = %d, want %d", input, i, got[i], want[i])
		}
	}
}

func TestClipValuesAbove(t *testing.T) {
	input := []int{15, 20, 11}
	got := Clip(input, 0, 10)
	want := []int{10, 10, 10}

	if len(got) != len(want) {
		t.Fatalf("Clip(%v, 0, 10) length = %d, want %d", input, len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Clip(%v, 0, 10)[%d] = %d, want %d", input, i, got[i], want[i])
		}
	}
}

func TestClipMixedValues(t *testing.T) {
	input := []int{-5, 5, 15}
	got := Clip(input, 0, 10)
	want := []int{0, 5, 10}

	if len(got) != len(want) {
		t.Fatalf("Clip(%v, 0, 10) length = %d, want %d", input, len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Clip(%v, 0, 10)[%d] = %d, want %d", input, i, got[i], want[i])
		}
	}
}

func TestClipInputUnchanged(t *testing.T) {
	input := []int{-5, 5, 15}
	original := []int{-5, 5, 15}

	Clip(input, 0, 10)

	// Check that input was not modified
	if len(input) != len(original) {
		t.Fatalf("input length changed from %d to %d", len(original), len(input))
	}

	for i := range original {
		if input[i] != original[i] {
			t.Errorf("input[%d] = %d, want %d (input was modified)", i, input[i], original[i])
		}
	}
}

func TestClipBoundaryValues(t *testing.T) {
	input := []int{0, 10}
	got := Clip(input, 0, 10)
	want := []int{0, 10}

	if len(got) != len(want) {
		t.Fatalf("Clip(%v, 0, 10) length = %d, want %d", input, len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Clip(%v, 0, 10)[%d] = %d, want %d", input, i, got[i], want[i])
		}
	}
}

func TestAbsDiffSumNil(t *testing.T) {
	got := AbsDiffSum(nil)
	if got != 0 {
		t.Errorf("AbsDiffSum(nil) = %d, want 0", got)
	}
}

func TestAbsDiffSumSingleValue(t *testing.T) {
	got := AbsDiffSum([]int{42})
	if got != 0 {
		t.Errorf("AbsDiffSum([]int{42}) = %d, want 0", got)
	}
}

func TestAbsDiffSumUpsAndDowns(t *testing.T) {
	// Test a slice with ups and downs
	got := AbsDiffSum([]int{1, 3, 2, 5, 4})
	// Differences: |3-1|=2, |2-3|=1, |5-2|=3, |4-5|=1
	// Sum: 2+1+3+1 = 7
	if got != 7 {
		t.Errorf("AbsDiffSum([]int{1, 3, 2, 5, 4}) = %d, want 7", got)
	}
}

func TestAbsDiffSumExample(t *testing.T) {
	// Test a simple case
	got := AbsDiffSum([]int{10, 15, 12})
	// Differences: |15-10|=5, |12-15|=3
	// Sum: 5+3 = 8
	if got != 8 {
		t.Errorf("AbsDiffSum([]int{10, 15, 12}) = %d, want 8", got)
	}
}
