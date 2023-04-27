package main

// Interfaces
// An interface is a set of method signatures.
// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.
// Shape interface defines a single method, Area.

// Rectangle and Circle types implement the Shape interface.
type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) increaseRadius() {
	c.Radius++
}

// Function that takes an interface as an argument

func main() {

	// Empty interface
	// An interface with no methods, used to handle values of unknown type.
}
