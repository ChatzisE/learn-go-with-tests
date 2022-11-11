package main

import "math"

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

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
