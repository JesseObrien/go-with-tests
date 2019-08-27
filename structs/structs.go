package structs

// Rectangle stores heights and widths
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter stores the width and height of something
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
