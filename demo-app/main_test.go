package main

import "testing"

func TestDisplay(t *testing.T) {
	// Test that display() doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("display() panicked: %v", r)
		}
	}()
	display()
}
