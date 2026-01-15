package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/archit-dey/repo-automations/internal/demo"
)

func main() {
	fmt.Println("=== Demo Go Application ===")
	display()

	// Demonstrate math operations
	result := demo.Add(10, 30)
	fmt.Printf("Addition: 10 + 30 = %d\n", result)

	result = demo.Multiply(5, 6)
	fmt.Printf("Multiplication: 5 * 6 = %d\n", result)

	// Demonstrate string operations
	greeting := demo.Greet("World")
	fmt.Println(greeting)

	// Demonstrate factorial calculation
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil && n >= 0 {
			fact := demo.Factorial(n)
			fmt.Printf("Factorial of %d = %d\n", n, fact)
		}
	}
}

func display() {
	fmt.Println("Hello there, this is a demo app")
}
