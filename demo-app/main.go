package main

import (
	"fmt"
	"os"
	"strconv"
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
