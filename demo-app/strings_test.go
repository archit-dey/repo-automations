package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple name", "World", "Hello, World! Welcome to the demo app."},
		{"empty string", "", "Hello, ! Welcome to the demo app."},
		{"long name", "Alice Cooper", "Hello, Alice Cooper! Welcome to the demo app."},
		{"name with numbers", "User123", "Hello, User123! Welcome to the demo app."},
		{"unicode name", "José", "Hello, José! Welcome to the demo app."},
		{"name with special chars", "John-Doe_Jr.", "Hello, John-Doe_Jr.! Welcome to the demo app."},
		{"single letter", "A", "Hello, A! Welcome to the demo app."},
		{"very long name", "Hubert Blaine Wolfeschlegelsteinhausenbergerdorff", "Hello, Hubert Blaine Wolfeschlegelsteinhausenbergerdorff! Welcome to the demo app."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.input)
			if got != tt.want {
				t.Errorf("Greet(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{"simple text", "hello world", map[string]int{"hello": 1, "world": 1}},
		{"repeated words", "hello hello world", map[string]int{"hello": 2, "world": 1}},
		{"case insensitive", "Hello HELLO hello", map[string]int{"hello": 3}},
		{"empty string", "", map[string]int{}},
		{"multiple spaces", "hello  world", map[string]int{"hello": 1, "world": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.text)
			if len(got) != len(tt.want) {
				t.Errorf("CountWords(%q) = %v; want %v", tt.text, got, tt.want)
				return
			}
			for word, count := range tt.want {
				if got[word] != count {
					t.Errorf("CountWords(%q)[%q] = %d; want %d", tt.text, word, got[word], count)
				}
			}
		})
	}
}
