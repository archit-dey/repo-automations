package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Demo Go Application ===")
	display()

	// Demonstrate math operations
	result := Add(10, 20)
	fmt.Printf("Addition: 10 + 20 = %d\n", result)

	result = Multiply(5, 6)
	fmt.Printf("Multiplication: 5 * 6 = %d\n", result)

	// Demonstrate string operations
	greeting := Greet("World")
	fmt.Println(greeting)

	// Demonstrate factorial calculation
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil && n >= 0 {
			fact := Factorial(n)
			fmt.Printf("Factorial of %d = %d\n", n, fact)
		}
	}
}

func display() {
	fmt.Println("Hello there, this is a demo app")
}

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Greet returns a greeting message
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to the demo app.", name)
}

// Factorial calculates the factorial of n
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Person represents a person with name and age
type Person struct {
	Name string
	Age  int
}

// Introduce returns an introduction message for the person
func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old.", p.Name, p.Age)
}

// IsAdult checks if the person is 18 or older
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// HaveBirthday increments the person's age using a pointer receiver
func (p *Person) HaveBirthday() {
	p.Age++
}

// Divide performs division with error handling
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// SumAll is a variadic function that sums all provided integers
func SumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Average calculates the average of a slice of numbers
func Average(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot calculate average of empty slice")
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers)), nil
}

// CountWords counts the occurrences of each word in a text
func CountWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

// FilterEven filters a slice to return only even numbers
func FilterEven(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

// Shape interface demonstrates interfaces in Go
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape interface
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter calculates the circumference of a circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// CalculateArea accepts any Shape and returns its area
func CalculateArea(s Shape) float64 {
	return s.Area()
}

// Swap exchanges the values of two integers using pointers
func Swap(a, b *int) {
	*a, *b = *b, *a
}

// Contains checks if a slice contains a specific value
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Counter returns a closure that increments and returns a counter
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Multiplier returns a closure that multiplies by a given factor
func Multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
