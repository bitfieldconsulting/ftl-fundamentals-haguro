package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a    float64
		b    []float64
		want float64
		name string
	}{
		{20, []float64{1}, 21, "Adding two positive numbers yields a positive number"},
		{231, []float64{56}, 287, "Adding three positive numbers yields a positive number"},
		{79, []float64{123191023123}, 123191023202, "Adding one large and one small number"},
		{0, []float64{0}, 0, "Adding two zeros yields zero"},
		{1.2, []float64{333}, 334.2, "Adding a number with a fraction and a number without one yields a number with fraction"},
		{-73, []float64{99}, 26, "Adding a negative number and a positive number"},
		{-1, []float64{-1}, -2, "Adding two negative numbers yields a negative number"},
		{-2222, []float64{0}, -2222, "Adding a number to zero yields that number"},
		{0.0000005133, []float64{0.00000312, 0.000664, 0.0000532, 0.00001}, 0.0007308333, "Adding five very small numbers"},
	}
	for _, c := range testCases {
		got := Add(c.a, c.b...)
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a    float64
		b    []float64
		want float64
		name string
	}{
		{1, []float64{1}, 0, "Subtracting a positive number from itself yields zero"},
		{-51, []float64{-51}, 0, "Subtracting a negative number from itself yields zero"},
		{4323, []float64{12, 11, 39}, 4261, "Subtracting 3 positive number from a much larger positive number yields a positive number"},
		{11.2, []float64{89.8}, -78.6, "Subtracting a positive number from a smaller positive number yields a negative number"},
		{-100, []float64{-111}, 11, "Subtracting a negative number from a larger negative number yields a positive number"},
		{0, []float64{6, 6, 15}, -27, "Subtracting multiple positive number from zero yields a negative number"},
		{0, []float64{-22}, 22, "Subtracting a negative number from zero yields a positive number"},
	}
	for _, c := range testCases {
		got := Subtract(c.a, c.b...)
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a    float64
		b    []float64
		want float64
		name string
	}{
		{1, []float64{1}, 1, "Multiplying one by one always yields one"},
		{3999191, []float64{1, 1, 1}, 3999191, "Multiplying a number by one any number of times yields that number"},
		{3, []float64{3}, 9, "Multiplying a number by itself equals that number squared"},
		{-10, []float64{-10}, 100, "Multiplying two negative number yields a positive number"},
		{-5, []float64{-12, -67}, -4020, "Multiplying three negative number yields a negative number"},
		{12, []float64{-1.5}, -18, "Multiplying a negative number by a positive number yields a negative number"},
		{43, []float64{0}, 0, "Multiplying any number by zero yields zero"},
	}
	for _, c := range testCases {
		got := Multiply(c.a, c.b...)
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		a             float64
		b             []float64
		want          float64
		errorExpected bool
		name          string
	}{
		{10, []float64{1}, 10, false, "Dividing a number by 1 yields that number"},
		{99, []float64{1, 1, 1}, 99, false, "Dividing a number by one multiple times yields that same number"},
		{312, []float64{0}, 0, true, "Dividing a number by zero returns an error"},
		{50, []float64{1, 5, 0, 200}, 0, true, "Dividing a number by multiple numbers with zero as one of them returns an error"},
		{0, []float64{0}, 0, true, "Dividing a zero by zero returns an error"},
		{1, []float64{2, 2, 2}, 0.125, false, "Dividing a number by multiple larger numbers yields a fraction"},
		{100, []float64{0.1, 10}, 100, false, "Dividing a number by a large number and a fraction"},
		{100, []float64{10000000}, 0.00001, false, "Dividing a number by a much larger number yields a small fraction"},
		{100, []float64{10000000}, 0.00001, false, "Dividing a number by a much larger number yields a small fraction"},
	}
	for _, c := range testCases {
		got, err := Divide(c.a, c.b...)
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
		{25, 5, false, "The square root of a two digit number"},
		{1947028, math.Sqrt(1947028), false, "The square root of a large number"},
		{0.00000002, math.Sqrt(0.00000002), false, "The square root of a small fraction"},
		{-25, 0, true, "Attempting the to calculate the square root of a negative number returns an error"},
	}
	for _, c := range testCases {
		got, err := Sqrt(c.a)
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
func TestEvaluate(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		expression    string
		want          float64
		errorExpected bool
		name          string
	}{
		{"1 + 1", 2, false, "A simple expression containing only signature"},
		{"2 % 5", 0, true, `An "invalid" operation (unimplemented operation)`},
		{"51x5.3", 0, true, "An invalid operation"},
		{"43.75 / 3.5", 12.5, false, "An expression with both operands containing fractions"},
		{"2/2", 1, false, "An expression containing no spaces"},
		{"  3- 1.2", 1.8, false, "An expression containing trailing whitespaces"},
		{"8 * 3 / 9", 0, true, `An "invalid" expression (more than two operands)`},
		{"yeah, no", 0, true, "An invalid expression"},
		{"1000/    10   \n	  ", 100, false, "An expression with whitepsace around the second operand"},
	}
	for _, c := range testCases {
		got, err := Evaluate(c.expression)
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
