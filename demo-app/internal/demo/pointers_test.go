package demo

import "testing"

func TestSwap(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
	}{
		{"positive numbers", 5, 10},
		{"negative numbers", -5, -10},
		{"mixed", -5, 10},
		{"zeros", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y := tt.a, tt.b
			Swap(&x, &y)
			if x != tt.b || y != tt.a {
				t.Errorf("After Swap(&%d, &%d), got x=%d, y=%d; want x=%d, y=%d",
					tt.a, tt.b, x, y, tt.b, tt.a)
			}
		})
	}
}
