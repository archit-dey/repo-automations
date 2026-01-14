package demo

// Shape interface demonstrates interfaces in Go
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape interface
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter calculates the circumference of a circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// CalculateArea accepts any Shape and returns its area
func CalculateArea(s Shape) float64 {
	return s.Area()
}
