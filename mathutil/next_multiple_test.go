package mathutil

import "testing"

func TestNextMultiple(t *testing.T) {
	tests := []struct {
		name string
		n    int
		step int
		want int
	}{
		// Step 0 edge case
		{"step 0", 5, 0, 5},
		{"step 0 negative n", -5, 0, -5},
		
		// Negative step edge case
		{"negative step", 5, -3, 5},
		{"negative step with negative n", -5, -3, -5},
		
		// Exact multiple
		{"exact multiple positive", 6, 3, 6},
		{"exact multiple zero", 0, 5, 0},
		{"exact multiple negative", -6, 3, -6},
		
		// Value between multiples
		{"between multiples", 7, 3, 9},
		{"between multiples small", 1, 5, 5},
		{"between multiples large step", 15, 10, 20},
		
		// Negative n
		{"negative n between multiples", -7, 3, -6},
		{"negative n small", -1, 5, 0},
		{"negative n large", -15, 10, -10},
		
		// Additional edge cases
		{"step 1", 5, 1, 5},
		{"large numbers", 1000, 7, 1001},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NextMultiple(test.n, test.step)
			if got != test.want {
				t.Errorf("NextMultiple(%d, %d) = %d; want %d", test.n, test.step, got, test.want)
			}
		})
	}
}
