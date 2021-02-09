package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, b, want float64
		name       string
	}{
		{20, 1, 21, "Adding two positive numbers yields a positive number"},
		{79, 123191023123, 123191023202, "Adding one large and one small number"},
		{0, 0, 0, "Adding two zeros yields zero"},
		{1.2, 333, 334.2, "Adding a number with a fraction and a number without one yields a number with fraction"},
		{-73, 99, 26, "Adding a negative number and a positive number"},
		{-1, -1, -2, "Adding two negative numbers yields a negative number"},
		{-2222, 0, -2222, "Adding a number to zero yields that number"},
	}
	for _, c := range testCases {
		got := calculator.Add(c.a, c.b)
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, b, want float64
		name       string
	}{
		{1, 1, 0, "Subtracting a positive number from itself yields zero"},
		{-51, -51, 0, "Subtracting a negative number from itself yields zero"},
		{4323, 12, 4311, "Subtracting a positive number from a larger positive number yields a positive number"},
		{11.2, 89.8, -78.6, "Subtracting a positive number from a smaller positive number yields a negative number"},
		{-100, -111, 11, "Subtracting a negative number from a larger negative number yields a positive number"},
		{0, 6, -6, "Subtracting a positive number from zero yields a negative number"},
		{0, -22, 22, "Subtracting a negative number from zero yields a positive number"},
	}
	for _, c := range testCases {
		got := calculator.Subtract(c.a, c.b)
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, b, want float64
		name       string
	}{
		{1, 1, 1, "Multiplying one by one always yields one"},
		{3999191, 1, 3999191, "Multiplying a number by one yields that number"},
		{3, 3, 9, "Multiplying a number by itself equals that number squared"},
		{-10, -10, 100, "Multiplying two negative number yields a positive number"},
		{12, -1.5, -18, "Multiplying a negative number by a positive number yields a negative number"},
		{43, 0, 0, "Multiplying any number by zero yields zero"},
	}
	for _, c := range testCases {
		got := calculator.Multiply(c.a, c.b)
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}

}

func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, b, want    float64
		errorExpected bool
		name          string
	}{
		{10, 1, 10, false, "Dividing a number by 1 yields that number"},
		{312, 0, 0, true, "Dividing a number by zero returns an error"},
		{0, 0, 0, true, "Dividing a zero by zero returns an error"},
		{1, 2, 0.5, false, "Dividing a two numbers may yield a fraction"},
		{100, 10000000, 0.00001, false, "Dividing a number by a much larger number yields a small fraction"},
	}
	for _, c := range testCases {
		got, err := calculator.Divide(c.a, c.b)
		if c.errorExpected && err == nil {
			t.Errorf("%s: expected error, got nil", c.name)
		}
		if !c.errorExpected && err != nil {
			t.Errorf("%s: expected no errors, got \"%v\"", c.name, err)
		}
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a, want       float64
		errorExpected bool
		name          string
	}{
		{1, 1, false, "The square root of one is one"},
		{0, 0, false, "The square root of zero is zero"},
		{25, 5, false, ""},
		{0.0002, 0.01414213562373095, false, "The square root of a small fraction"},
		{-25, 0, true, "Attempting the to get calculate the square root of a negative number returns an error"},
	}
	for _, c := range testCases {
		got, err := calculator.Sqrt(c.a)
		if c.errorExpected && err == nil {
			t.Errorf("%s: expected error, got nil", c.name)
		}
		if !c.errorExpected && err != nil {
			t.Errorf("%s: expected no errors, got \"%v\"", c.name, err)
		}
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}
