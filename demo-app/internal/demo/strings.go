package demo

import (
	"fmt"
	"strings"
)

// Greet returns a greeting message
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to the demo app.", name)
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
