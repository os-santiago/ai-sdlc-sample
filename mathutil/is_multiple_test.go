package mathutil

import "testing"

func TestIsMultipleMultiple(t *testing.T) {
	if got := IsMultiple(10, 2); !got {
		t.Errorf("IsMultiple(10, 2) = %v, want true", got)
	}
}

func TestIsMultipleNonMultiple(t *testing.T) {
	if got := IsMultiple(10, 3); got {
		t.Errorf("IsMultiple(10, 3) = %v, want false", got)
	}
}

func TestIsMultipleNegativeN(t *testing.T) {
	if got := IsMultiple(-10, 2); !got {
		t.Errorf("IsMultiple(-10, 2) = %v, want true", got)
	}
}

func TestIsMultipleZeroDivisor(t *testing.T) {
	if got := IsMultiple(10, 0); got {
		t.Errorf("IsMultiple(10, 0) = %v, want false", got)
	}
}