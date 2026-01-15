package demo

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
