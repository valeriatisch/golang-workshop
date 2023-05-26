// Package mathutil provides utility functions for mathematical operations.
package mathutil

import (
	"math"
	"errors"
)

// Add adds two integers and returns the sum.
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts the second integer from the first and returns the difference.
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two integers and returns the product.
func Multiply(a, b int) int {
	return a * b
}

// Divide divides the first integer by the second and returns the quotient.
// It returns an error if the second integer is zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

// SquareRoot calculates the square root of a float64 number and returns the result.
// It returns an error if the number is negative.
func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Square root of negative number")
	}
	return math.Sqrt(x), nil
}
