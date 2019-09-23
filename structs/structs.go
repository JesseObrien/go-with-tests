package main 

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle is a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Triangle is a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter stores the width and height of something
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculates the Area of something
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
