package mathutil

import "testing"

func TestIsSquareZero(t *testing.T) {
	if got := IsSquare(0); !got {
		t.Errorf("IsSquare(0) = %v, want true", got)
	}
}

func TestIsSquareOne(t *testing.T) {
	if got := IsSquare(1); !got {
		t.Errorf("IsSquare(1) = %v, want true", got)
	}
}

func TestIsSquare16(t *testing.T) {
	if got := IsSquare(16); !got {
		t.Errorf("IsSquare(16) = %v, want true", got)
	}
}

func TestIsSquare15(t *testing.T) {
	if got := IsSquare(15); got {
		t.Errorf("IsSquare(15) = %v, want false", got)
	}
}

func TestIsSquareNegative(t *testing.T) {
	if got := IsSquare(-1); got {
		t.Errorf("IsSquare(-1) = %v, want false", got)
	}
}
