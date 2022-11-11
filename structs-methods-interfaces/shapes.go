package main

func Perimeter(rectangle Rectangle) float64 {
	result := 0.0
	result = 2 * (rectangle.Width + rectangle.Height)
	return result
}

func Area(rectangle Rectangle) float64 {
	result := 0.0
	result = rectangle.Width * rectangle.Height
	return result
}

type Rectangle struct {
	Width  float64
	Height float64
}
