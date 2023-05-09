/*
TODO: Createa a shape area calculator
Define an interface called 'Shape' with a method 'Area() float64'.
Implement the 'Shape' interface for different shapes (e.g., Circle, Rectangle, Triangle).
Create a function 'CalculateTotalArea' that takes a slice of Shapes and returns the sum of their areas.
*/
package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
}

// Circle type
type Circle struct {
	Radius float64
}

// Rectangle type
type Rectangle struct {
	Width  float64
	Height float64
}

// Triangle type
type Triangle struct {
	Base   float64
	Height float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area method for Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

// CalculateTotalArea function
func CalculateTotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	// Create a slice of shapes
	shapes := []Shape{
		Circle{Radius: 10},
		Rectangle{Width: 10, Height: 5},
		Triangle{Base: 12, Height: 6},
	}

	// Calculate the total area of the shapes
	totalArea := CalculateTotalArea(shapes)
	fmt.Println("Total area is:", totalArea)
}
