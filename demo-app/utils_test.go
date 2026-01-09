package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"simple palindrome", "racecar", true},
		{"palindrome with spaces", "A man a plan a canal Panama", true},
		{"not a palindrome", "hello", false},
		{"single character", "a", true},
		{"empty string", "", true},
		{"numeric palindrome", "12321", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple string", "hello", "olleh"},
		{"palindrome", "racecar", "racecar"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"with spaces", "hello world", "dlrow olleh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.want {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"even positive", 4, true},
		{"odd positive", 7, false},
		{"zero", 0, true},
		{"even negative", -4, true},
		{"odd negative", -7, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEven(tt.input)
			if got != tt.want {
				t.Errorf("IsEven(%d) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"first larger", 10, 5, 10},
		{"second larger", 5, 10, 10},
		{"equal", 5, 5, 5},
		{"negative numbers", -5, -10, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"first smaller", 5, 10, 5},
		{"second smaller", 10, 5, 5},
		{"equal", 5, 5, 5},
		{"negative numbers", -5, -10, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Min(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Min(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"positive numbers", []int{1, 2, 3, 4, 5}, 15},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed numbers", []int{-5, 10, -3, 8}, 10},
		{"empty slice", []int{}, 0},
		{"single number", []int{42}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.input)
			if got != tt.want {
				t.Errorf("Sum(%v) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man a plan a canal Panama")
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("hello world this is a test string")
	}
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
