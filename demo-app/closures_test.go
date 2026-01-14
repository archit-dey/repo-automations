package main

import "testing"

func TestCounter(t *testing.T) {
	counter := Counter()

	// Test sequential calls
	first := counter()
	if first != 1 {
		t.Errorf("First call to counter() = %d; want 1", first)
	}

	second := counter()
	if second != 2 {
		t.Errorf("Second call to counter() = %d; want 2", second)
	}

	third := counter()
	if third != 3 {
		t.Errorf("Third call to counter() = %d; want 3", third)
	}

	// Test that multiple counters are independent
	counter2 := Counter()
	first2 := counter2()
	if first2 != 1 {
		t.Errorf("First call to counter2() = %d; want 1", first2)
	}

	// Original counter should still be at 3
	fourth := counter()
	if fourth != 4 {
		t.Errorf("Fourth call to original counter() = %d; want 4", fourth)
	}
}

func TestMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{"multiply by 2", 2, 5, 10},
		{"multiply by 3", 3, 4, 12},
		{"multiply by 10", 10, 7, 70},
		{"multiply by 0", 0, 100, 0},
		{"multiply by 1", 1, 42, 42},
		{"multiply by negative", -2, 5, -10},
		{"multiply negative by negative", -3, -4, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiply := Multiplier(tt.factor)
			got := multiply(tt.input)
			if got != tt.want {
				t.Errorf("Multiplier(%d)(%d) = %d; want %d", tt.factor, tt.input, got, tt.want)
			}
		})
	}
}

func TestMultiplierIndependence(t *testing.T) {
	// Test that multiple multipliers are independent
	double := Multiplier(2)
	triple := Multiplier(3)

	if double(5) != 10 {
		t.Errorf("double(5) = %d; want 10", double(5))
	}

	if triple(5) != 15 {
		t.Errorf("triple(5) = %d; want 15", triple(5))
	}

	// Test that they maintain their own state
	if double(10) != 20 {
		t.Errorf("double(10) = %d; want 20", double(10))
	}

	if triple(10) != 30 {
		t.Errorf("triple(10) = %d; want 30", triple(10))
	}
}
