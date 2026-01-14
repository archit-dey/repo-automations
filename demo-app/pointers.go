package main

// Swap exchanges the values of two integers using pointers
func Swap(a, b *int) {
	*a, *b = *b, *a
}
