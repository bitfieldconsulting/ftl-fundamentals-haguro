package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, b, want float64
	}{
		{20, 1, 21},
		{1, 1, 2},
		{79, 123123, 123202},
		{0, 0, 0},
		{1.2, 333, 334.2},
	}
	for _, c := range testCases {
		got := calculator.Add(c.a, c.b)
		if c.want != got {
			t.Errorf("want %f, got %f", c.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, b, want float64
	}{
		{1, 1, 0},
		{4323, 12, 4311},
		{11.2, 89.8, -78.6},
		{-100, -111, 11},
		{0, 6, -6},
	}
	for _, c := range testCases {
		got := calculator.Subtract(c.a, c.b)
		if c.want != got {
			t.Errorf("want %f, got %f", c.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, b, want float64
	}{
		{1, 1, 1},
		{100, 100, 10000},
		{-10, -10, 100},
		{12, -1.5, -18},
		{43, 0, 0},
	}
	for _, c := range testCases {
		got := calculator.Multiply(c.a, c.b)
		if c.want != got {
			t.Errorf("want %f, got %f", c.want, got)
		}
	}

}

func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, b, want    float64
		errorExpected bool
	}{
		{1, 1, 1, false},
		{312, 0, 0, true},
		{0, 0, 0, true},
		{1, 2, 0.5, false},
		{100, 10000000, 0.00001, false},
	}
	for _, c := range testCases {
		got, err := calculator.Divide(c.a, c.b)
		if c.errorExpected && err == nil {
			t.Errorf("Expected error, got nil")
		}
		if !c.errorExpected && err != nil {
			t.Errorf("Expected no errors, got \"%v\"", err)
		}
		if c.want != got {
			t.Errorf("want %f, got %f", c.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, want       float64
		errorExpected bool
	}{
		{1, 1, false},
		{0, 0, false},
		{25, 5, false},
		{-25, 0, true},
	}
	for _, c := range testCases {
		got, err := calculator.Sqrt(c.a)
		if c.errorExpected && err == nil {
			t.Errorf("Expected error, got nil")
		}
		if !c.errorExpected && err != nil {
			t.Errorf("Expected no errors, got \"%v\"", err)
		}
		if c.want != got {
			t.Errorf("want %f, got %f", c.want, got)
		}
	}
}
