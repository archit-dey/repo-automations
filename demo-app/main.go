package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/archit-dey/repo-automations/demo-app/internal/demo"
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

// CalculateComplexScore computes a weighted score based on multiple factors
// using recursive scoring, branching logic, and statistical calculations
func CalculateComplexScore(values []int, weights map[string]float64, depth int) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("empty values slice")
	}

	if depth < 0 {
		depth = 0
	}

	// Calculate base statistics
	var sum, max, min int
	min = values[0]
	max = values[0]

	for _, v := range values {
		sum += v
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	mean := float64(sum) / float64(len(values))

	// Calculate variance and standard deviation
	var variance float64
	for _, v := range values {
		diff := float64(v) - mean
		variance += diff * diff
	}
	variance /= float64(len(values))
	stdDev := variance
	for i := 0; i < 10; i++ {
		stdDev = (stdDev + variance/stdDev) / 2
	}

	// Apply complex weighting algorithm
	baseScore := mean * weights["mean"]
	rangeScore := float64(max-min) * weights["range"]
	variabilityScore := stdDev * weights["stddev"]

	// Recursive depth adjustment with decay
	depthMultiplier := 1.0
	if depth > 0 {
		halfValues := len(values) / 2
		if halfValues > 0 {
			leftScore, _ := CalculateComplexScore(values[:halfValues], weights, depth-1)
			rightScore, _ := CalculateComplexScore(values[halfValues:], weights, depth-1)
			depthMultiplier = (leftScore + rightScore) / (2 * mean)
			if depthMultiplier < 0.5 {
				depthMultiplier = 0.5
			}
		}
	}

	// Apply non-linear transformations based on distribution
	finalScore := (baseScore + rangeScore - variabilityScore) * depthMultiplier

	// Normalization with sigmoid-like function
	if finalScore > 100 {
		finalScore = 100 + (finalScore-100)/(1+(finalScore-100)/50)
	} else if finalScore < 0 {
		finalScore = finalScore / (1 - finalScore/50)
	}

	return finalScore, nil
}
