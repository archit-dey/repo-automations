package main

import "testing"

func TestRectangle(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 10}

	area := rect.Area()
	if area != 50 {
		t.Errorf("Rectangle{5, 10}.Area() = %f; want 50", area)
	}

	perimeter := rect.Perimeter()
	if perimeter != 30 {
		t.Errorf("Rectangle{5, 10}.Perimeter() = %f; want 30", perimeter)
	}
}

func TestCircle(t *testing.T) {
	circle := Circle{Radius: 5}

	area := circle.Area()
	expected := 3.14159 * 5 * 5
	if area != expected {
		t.Errorf("Circle{5}.Area() = %f; want %f", area, expected)
	}

	perimeter := circle.Perimeter()
	expectedPerimeter := 2 * 3.14159 * 5
	if perimeter != expectedPerimeter {
		t.Errorf("Circle{5}.Perimeter() = %f; want %f", perimeter, expectedPerimeter)
	}
}

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{Width: 4, Height: 5}, 20},
		{"circle", Circle{Radius: 2}, 3.14159 * 2 * 2},
		{"square", Rectangle{Width: 3, Height: 3}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateArea(tt.shape)
			if got != tt.want {
				t.Errorf("CalculateArea(%v) = %f; want %f", tt.shape, got, tt.want)
			}
		})
	}
}
