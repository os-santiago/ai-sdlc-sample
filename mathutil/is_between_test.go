package mathutil

import "testing"

func TestIsBetweenInside(t *testing.T) {
	if got := IsBetween(5, 1, 10); !got {
		t.Errorf("IsBetween(5, 1, 10) = %v, want true", got)
	}
}

func TestIsBetweenAtLowerBound(t *testing.T) {
	if got := IsBetween(1, 1, 10); !got {
		t.Errorf("IsBetween(1, 1, 10) = %v, want true", got)
	}
}

func TestIsBetweenAtUpperBound(t *testing.T) {
	if got := IsBetween(10, 1, 10); !got {
		t.Errorf("IsBetween(10, 1, 10) = %v, want true", got)
	}
}

func TestIsBetweenOutsideBelow(t *testing.T) {
	if got := IsBetween(0, 1, 10); got {
		t.Errorf("IsBetween(0, 1, 10) = %v, want false", got)
	}
}

func TestIsBetweenOutsideAbove(t *testing.T) {
	if got := IsBetween(11, 1, 10); got {
		t.Errorf("IsBetween(11, 1, 10) = %v, want false", got)
	}
}

func TestIsBetweenLoGreaterThanHi(t *testing.T) {
	if got := IsBetween(5, 10, 1); got {
		t.Errorf("IsBetween(5, 10, 1) = %v, want false", got)
	}
}
