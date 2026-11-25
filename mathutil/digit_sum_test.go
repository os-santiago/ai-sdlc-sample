package mathutil

import "testing"

// Tests for DigitSum function
func TestDigitSumZero(t *testing.T) {
	if got := DigitSum(0); got != 0 {
		t.Errorf("DigitSum(0) = %d, want 0", got)
	}
}

func TestDigitSumPositive(t *testing.T) {
	if got := DigitSum(123); got != 6 {
		t.Errorf("DigitSum(123) = %d, want 6", got)
	}
}

func TestDigitSumNegative(t *testing.T) {
	if got := DigitSum(-123); got != 6 {
		t.Errorf("DigitSum(-123) = %d, want 6", got)
	}
}

func TestDigitSumSingleDigit(t *testing.T) {
	if got := DigitSum(7); got != 7 {
		t.Errorf("DigitSum(7) = %d, want 7", got)
	}
}

func TestDigitSumLargeNumber(t *testing.T) {
	if got := DigitSum(987654321); got != 45 {
		t.Errorf("DigitSum(987654321) = %d, want 45", got)
	}
}