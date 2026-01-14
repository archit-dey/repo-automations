package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"positive numbers", 10, 20, 30},
		{"negative numbers", -5, -3, -8},
		{"mixed numbers", 10, -5, 5},
		{"zero", 0, 0, 0},
		{"add with zero", 42, 0, 42},
		{"large numbers", 1000000, 2000000, 3000000},
		{"negative and positive", -100, 100, 0},
		{"both negative large", -1000, -2000, -3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"positive numbers", 5, 6, 30},
		{"negative numbers", -5, -3, 15},
		{"mixed numbers", 10, -5, -50},
		{"multiply by zero", 10, 0, 0},
		{"multiply by one", 7, 1, 7},
		{"zero by zero", 0, 0, 0},
		{"large positive numbers", 1000, 2000, 2000000},
		{"negative by one", -42, 1, -42},
		{"multiply negative by positive", -7, 8, -56},
		{"both large negative", -100, -200, 20000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"factorial of 0", 0, 1},
		{"factorial of 1", 1, 1},
		{"factorial of 2", 2, 2},
		{"factorial of 3", 3, 6},
		{"factorial of 4", 4, 24},
		{"factorial of 5", 5, 120},
		{"factorial of 6", 6, 720},
		{"factorial of 7", 7, 5040},
		{"factorial of 10", 10, 3628800},
		{"factorial of 12", 12, 479001600},
		{"negative number", -1, 1},
		{"negative large", -5, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Factorial(tt.input)
			if got != tt.want {
				t.Errorf("Factorial(%d) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestAddCommutativity(t *testing.T) {
	tests := []struct {
		a int
		b int
	}{
		{5, 10},
		{-3, 7},
		{100, 200},
		{0, 42},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result1 := Add(tt.a, tt.b)
			result2 := Add(tt.b, tt.a)
			if result1 != result2 {
				t.Errorf("Add is not commutative: Add(%d, %d) = %d, but Add(%d, %d) = %d",
					tt.a, tt.b, result1, tt.b, tt.a, result2)
			}
		})
	}
}

func TestMultiplyCommutativity(t *testing.T) {
	tests := []struct {
		a int
		b int
	}{
		{5, 10},
		{100, 200},
		{0, 42},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result1 := Multiply(tt.a, tt.b)
			result2 := Multiply(tt.b, tt.a)
			if result1 != result2 {
				t.Errorf("Multiply is not commutative: Multiply(%d, %d) = %d, but Multiply(%d, %d) = %d",
					tt.a, tt.b, result1, tt.b, tt.a, result2)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{"positive division", 10.0, 2.0, 5.0, false},
		{"negative dividend", -10.0, 2.0, -5.0, false},
		{"negative divisor", 10.0, -2.0, -5.0, false},
		{"both negative", -10.0, -2.0, 5.0, false},
		{"division by zero", 10.0, 0.0, 0.0, true},
		{"zero dividend", 0.0, 5.0, 0.0, false},
		{"fractional result", 7.0, 2.0, 3.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Divide(%f, %f) expected error, got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%f, %f) unexpected error: %v", tt.a, tt.b, err)
				}
				if got != tt.want {
					t.Errorf("Divide(%f, %f) = %f; want %f", tt.a, tt.b, got, tt.want)
				}
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{"no arguments", []int{}, 0},
		{"single number", []int{42}, 42},
		{"multiple positive", []int{1, 2, 3, 4, 5}, 15},
		{"with negatives", []int{-5, 10, -3, 8}, 10},
		{"all zeros", []int{0, 0, 0}, 0},
		{"many numbers", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumAll(tt.numbers...)
			if got != tt.want {
				t.Errorf("SumAll(%v) = %d; want %d", tt.numbers, got, tt.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name    string
		numbers []float64
		want    float64
		wantErr bool
	}{
		{"positive numbers", []float64{1, 2, 3, 4, 5}, 3.0, false},
		{"single number", []float64{42}, 42.0, false},
		{"with negatives", []float64{-10, 0, 10}, 0.0, false},
		{"empty slice", []float64{}, 0.0, true},
		{"fractional", []float64{1.5, 2.5, 3.0}, 2.333333, false},
	}

	const tolerance = 0.000001

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Average(tt.numbers)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Average(%v) expected error, got nil", tt.numbers)
				}
			} else {
				if err != nil {
					t.Errorf("Average(%v) unexpected error: %v", tt.numbers, err)
				}
				diff := got - tt.want
				if diff < 0 {
					diff = -diff
				}
				if diff > tolerance {
					t.Errorf("Average(%v) = %f; want %f", tt.numbers, got, tt.want)
				}
			}
		})
	}
}
