// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"strings"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a float64, b ...float64) float64 {
	for _, bb := range b {
		a += bb
	}
	return a
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a float64, b ...float64) float64 {
	for _, bb := range b {
		a -= bb
	}
	return a
}

// Multiply takes two numbers and returns the result of multiplying them
func Multiply(a float64, b ...float64) float64 {
	for _, bb := range b {
		a *= bb
	}
	return a
}

// Divide takes two numbers and divides the first one by the second one
func Divide(a float64, b ...float64) (float64, error) {
	for _, bb := range b {
		if bb == 0 {
			return 0, errors.New("division by zero")
		}
		a /= bb
	}
	return a, nil
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

//Evaluate evaluates a given string expression and returns the result or an error
func Evaluate(e string) (float64, error) {
	var (
		a, b float64
		op   rune
	)
	r := strings.NewReader(e)
	_, err := fmt.Fscanf(r, "%f %c %f", &a, &op, &b)
	if err != nil {
		return 0, fmt.Errorf("bad expression %q: %v", e, err)
	}
	if r.Len() > 0 {
		return 0, fmt.Errorf("bad expression %q: expression has further content", e)
	}

	switch op {
	case '+':
		return Add(a, b), nil
	case '-':
		return Subtract(a, b), nil
	case '*':
		return Multiply(a, b), nil
	case '/':
		return Divide(a, b)
	default:
		return 0, fmt.Errorf("bad expression %q: invalid operation %q", e, op)
	}
}
