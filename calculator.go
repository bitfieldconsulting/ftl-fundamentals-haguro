// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"strings"
)

// Add takes two (or more) numbers and returns the result of adding them together.
func Add(a, b float64, c ...float64) float64 {
	a += b
	for _, cc := range c {
		a += cc
	}
	return a
}

// Subtract takes two (or more) numbers and returns the result of subtracting them
func Subtract(a, b float64, c ...float64) float64 {
	a -= b
	for _, cc := range c {
		a -= cc
	}
	return a
}

// Multiply takes two (or more) numbers and returns the result of multiplying them
func Multiply(a, b float64, c ...float64) float64 {
	a *= b
	for _, cc := range c {
		a *= cc
	}
	return a
}

// Divide takes two (or more) numbers and divides the them by each other
func Divide(a, b float64, c ...float64) (float64, error) {
	for _, cc := range append([]float64{b}, c...) {
		if cc == 0 {
			return 0, errors.New("division by zero")
		}
		a /= cc
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
	var a, b float64
	var op rune

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
