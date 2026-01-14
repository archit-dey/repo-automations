package main

import "testing"

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
