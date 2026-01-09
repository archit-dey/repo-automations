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

func TestGreet(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple name", "World", "Hello, World! Welcome to the demo app."},
		{"empty string", "", "Hello, ! Welcome to the demo app."},
		{"long name", "Alice Cooper", "Hello, Alice Cooper! Welcome to the demo app."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.input)
			if got != tt.want {
				t.Errorf("Greet(%q) = %q; want %q", tt.input, got, tt.want)
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
		{"factorial of 5", 5, 120},
		{"factorial of 10", 10, 3628800},
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

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(100, 200)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}
