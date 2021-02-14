// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and divides the first one by the second one
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Sqrt returns the square root of a positive number, an error otherwise.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of a negative number")
	}
	x, y := a/2, 0.0
	for i := 0; i < 5000; i++ {
		if x == 0 {
			break
		}
		y = 0.5 * (x + (a / x))
		x = y
	}

	return y, nil
}
