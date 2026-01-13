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

func TestGreet(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple name", "World", "Hello, World! Welcome to the demo app."},
		{"empty string", "", "Hello, ! Welcome to the demo app."},
		{"long name", "Alice Cooper", "Hello, Alice Cooper! Welcome to the demo app."},
		{"name with numbers", "User123", "Hello, User123! Welcome to the demo app."},
		{"unicode name", "José", "Hello, José! Welcome to the demo app."},
		{"name with special chars", "John-Doe_Jr.", "Hello, John-Doe_Jr.! Welcome to the demo app."},
		{"single letter", "A", "Hello, A! Welcome to the demo app."},
		{"very long name", "Hubert Blaine Wolfeschlegelsteinhausenbergerdorff", "Hello, Hubert Blaine Wolfeschlegelsteinhausenbergerdorff! Welcome to the demo app."},
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

func TestDisplay(t *testing.T) {
	// Test that display() doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("display() panicked: %v", r)
		}
	}()
	display()
}

// Test Add with table-driven subtests for commutativity
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

// Test Multiply with table-driven subtests for commutativity
func TestMultiplyCommutativity(t *testing.T) {
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
			result1 := Multiply(tt.a, tt.b)
			result2 := Multiply(tt.b, tt.a)
			if result1 != result2 {
				t.Errorf("Multiply is not commutative: Multiply(%d, %d) = %d, but Multiply(%d, %d) = %d",
					tt.a, tt.b, result1, tt.b, tt.a, result2)
			}
		})
	}
}

// Test Person struct and methods
func TestPersonIntroduce(t *testing.T) {
	tests := []struct {
		name   string
		person Person
		want   string
	}{
		{"adult", Person{"Alice", 30}, "Hi, I'm Alice and I'm 30 years old."},
		{"child", Person{"Bob", 10}, "Hi, I'm Bob and I'm 10 years old."},
		{"zero age", Person{"Charlie", 0}, "Hi, I'm Charlie and I'm 0 years old."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.Introduce()
			if got != tt.want {
				t.Errorf("Introduce() = %q; want %q", got, tt.want)
			}
		})
	}
}

func TestPersonIsAdult(t *testing.T) {
	tests := []struct {
		name   string
		person Person
		want   bool
	}{
		{"adult", Person{"Alice", 30}, true},
		{"exactly 18", Person{"Bob", 18}, true},
		{"minor", Person{"Charlie", 17}, false},
		{"child", Person{"David", 5}, false},
		{"zero age", Person{"Eve", 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.IsAdult()
			if got != tt.want {
				t.Errorf("IsAdult() = %v; want %v", got, tt.want)
			}
		})
	}
}

func TestPersonHaveBirthday(t *testing.T) {
	person := Person{"Alice", 25}
	person.HaveBirthday()
	if person.Age != 26 {
		t.Errorf("After HaveBirthday(), Age = %d; want 26", person.Age)
	}

	person.HaveBirthday()
	person.HaveBirthday()
	if person.Age != 28 {
		t.Errorf("After 3 birthdays, Age = %d; want 28", person.Age)
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

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{"simple text", "hello world", map[string]int{"hello": 1, "world": 1}},
		{"repeated words", "hello hello world", map[string]int{"hello": 2, "world": 1}},
		{"case insensitive", "Hello HELLO hello", map[string]int{"hello": 3}},
		{"empty string", "", map[string]int{}},
		{"multiple spaces", "hello  world", map[string]int{"hello": 1, "world": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.text)
			if len(got) != len(tt.want) {
				t.Errorf("CountWords(%q) = %v; want %v", tt.text, got, tt.want)
				return
			}
			for word, count := range tt.want {
				if got[word] != count {
					t.Errorf("CountWords(%q)[%q] = %d; want %d", tt.text, word, got[word], count)
				}
			}
		})
	}
}

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    []int
	}{
		{"mixed numbers", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"all even", []int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
		{"all odd", []int{1, 3, 5, 7}, []int{}},
		{"empty slice", []int{}, []int{}},
		{"with negatives", []int{-4, -3, -2, -1, 0, 1, 2}, []int{-4, -2, 0, 2}},
		{"only zero", []int{0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterEven(tt.numbers)
			if len(got) != len(tt.want) {
				t.Errorf("FilterEven(%v) = %v; want %v", tt.numbers, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("FilterEven(%v) = %v; want %v", tt.numbers, got, tt.want)
					break
				}
			}
		})
	}
}

func TestRectangle(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 10}

	area := rect.Area()
	if area != 50 {
		t.Errorf("Rectangle{5, 10}.Area() = %f; want 50", area)
	}

	perimeter := rect.Perimeter()
	if perimeter != 30 {
		t.Errorf("Rectangle{5, 10}.Perimeter() = %f; want 30", perimeter)
	}
}

func TestCircle(t *testing.T) {
	circle := Circle{Radius: 5}

	area := circle.Area()
	expected := 3.14159 * 5 * 5
	if area != expected {
		t.Errorf("Circle{5}.Area() = %f; want %f", area, expected)
	}

	perimeter := circle.Perimeter()
	expectedPerimeter := 2 * 3.14159 * 5
	if perimeter != expectedPerimeter {
		t.Errorf("Circle{5}.Perimeter() = %f; want %f", perimeter, expectedPerimeter)
	}
}

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{Width: 4, Height: 5}, 20},
		{"circle", Circle{Radius: 2}, 3.14159 * 2 * 2},
		{"square", Rectangle{Width: 3, Height: 3}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateArea(tt.shape)
			if got != tt.want {
				t.Errorf("CalculateArea(%v) = %f; want %f", tt.shape, got, tt.want)
			}
		})
	}
}

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

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		item  string
		want  bool
	}{
		{"found at start", []string{"apple", "banana", "cherry"}, "apple", true},
		{"found in middle", []string{"apple", "banana", "cherry"}, "banana", true},
		{"found at end", []string{"apple", "banana", "cherry"}, "cherry", true},
		{"not found", []string{"apple", "banana", "cherry"}, "grape", false},
		{"empty slice", []string{}, "apple", false},
		{"case sensitive", []string{"Apple", "Banana"}, "apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.slice, tt.item)
			if got != tt.want {
				t.Errorf("Contains(%v, %q) = %v; want %v", tt.slice, tt.item, got, tt.want)
			}
		})
	}
}
