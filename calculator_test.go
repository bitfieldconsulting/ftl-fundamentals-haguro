package calculator_test

import (
	"calculator"
	"testing"
)

type TestCase struct {
	a, b float64
	want float64
}

var AddTests = []TestCase{
	{20, 1, 21},
	{1, 1, 2},
	{79, 123123, 123202},
	{0, 0, 0},
	{1.2, 333, 334.2},
}

var SubTests = []TestCase{
	{1, 1, 0},
	{4323, 12, 4311},
	{11.2, 89.8, -78.6},
	{-100, -111, 11},
	{0, 6, -6},
}

var MulTests = []TestCase{
	{1, 1, 1},
	{100, 100, 10000},
	{-10, -10, 100},
	{12, -1.5, -18},
	{43, 0, 0},
}

func TestAdd(t *testing.T) {
	t.Parallel()
	for _, c := range AddTests {
		got := calculator.Add(c.a, c.b)
		if c.want != got {
			t.Errorf("want %f, got %f", c.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	for _, c := range SubTests {
		got := calculator.Subtract(c.a, c.b)
		if c.want != got {
			t.Errorf("want %f, got %f", c.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	for _, c := range MulTests {
		got := calculator.Multiply(c.a, c.b)
		if c.want != got {
			t.Errorf("want %f, got %f", c.want, got)
		}
	}

}
