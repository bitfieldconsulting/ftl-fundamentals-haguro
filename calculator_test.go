package calculator

import (
	"testing"
)

func TestAddSubtractAndMultiply(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		desc string
		fn   func(float64, float64, ...float64) float64
		a    float64
		b    float64
		c    []float64
		want float64
	}{
		{"Adding two positive numbers yields a positive number", Add, 231, 56, []float64{}, 287},
		{"Adding one large and one small number", Add, 9, 123191023123, []float64{}, 123191023132},
		{"Adding two zeros yields zero", Add, 0, 0, []float64{}, 0},
		{"Adding a number with a fraction and a number without one yields a number with fraction", Add, 1.2, 333, []float64{}, 334.2},
		{"Adding a negative number and a positive number", Add, -73, 99, []float64{}, 26},
		{"Adding two negative numbers yields a negative number", Add, -1, -1, []float64{}, -2},
		{"Adding a number to zero yields the same number", Add, -2222, 0, []float64{}, -2222},
		{"Adding five very small numbers", Add, 0.0000005133, 0.00000312, []float64{0.000664, 0.0000532, 0.00001}, 0.0007308333},

		{"Subtracting a positive number from itself yields zero", Subtract, 1, 1, []float64{}, 0},
		{"Subtracting a negative number from itself yields zero", Subtract, -51, -51, []float64{}, 0},
		{"Subtracting multiple positive numbers from a much larger positive number yields a positive number", Subtract, 4323, 12, []float64{11, 39}, 4261},
		{"Subtracting a positive number from a smaller positive number yields a negative number", Subtract, 11.2, 89.8, []float64{}, -78.6},
		{"Subtracting a negative number from a larger negative number yields a positive number", Subtract, -100, -111, []float64{}, 11},
		{"Subtracting multiple positive numbers from zero yields a negative number", Subtract, 0, 6, []float64{6, 15}, -27},
		{"Subtracting a negative number from zero yields a positive number", Subtract, 0, -22, []float64{}, 22},

		{"Multiplying 1 by 1 always yields 1", Multiply, 1, 1, []float64{}, 1},
		{"Multiplying a number by 1 any number of times yields that number", Multiply, 3999191, 1, []float64{1, 1, 1}, 3999191},
		{"Multiplying a number by itself equals that number squared", Multiply, 5091, 5091, []float64{}, 25918281},
		{"Multiplying two negative numbers yields a positive number", Multiply, -10, -10, []float64{}, 100},
		{"Multiplying three negative number yields a negative number", Multiply, -5, -12, []float64{-67}, -4020},
		{"Multiplying a positive number by a negative number yields a negative number", Multiply, 12, -1.5, []float64{}, -18},
		{"Multiplying any number by zero yields zero", Multiply, 43, 0, []float64{}, 0},
	}
	for _, c := range testCases {
		got := c.fn(c.a, c.b, c.c...)
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.desc, c.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		a           float64
		b           float64
		c           []float64
		want        float64
		errExpected bool
	}{
		{"Dividing a number by 1 yields the same number", 10, 1, []float64{}, 10, false},
		{"Dividing a number by 1 multiple times yields the same number", 99, 1, []float64{1, 1}, 99, false},
		{"Dividing a number by zero returns an error", 312, 0, []float64{}, 0, true},
		{"Dividing a number by multiple numbers with zero being one of them returns an error", 50, 1, []float64{5, 0, 200}, 0, true},
		{"Dividing a zero by zero returns an error", 0, 0, []float64{}, 0, true},
		{"Dividing a number by multiple larger numbers yields a fraction", 1, 2, []float64{2, 2}, 0.125, false},
		{"Dividing a number by a large number and a fraction", 100, 0.1, []float64{10}, 100, false},
		{"Dividing a number by a much larger number yields a small fraction", 100, 10000000, []float64{}, 0.00001, false},
	}
	for _, c := range testCases {
		got, err := Divide(c.a, c.b, c.c...)
		errReceived := err != nil
		if c.errExpected != errReceived {
			t.Fatalf("%s: unexpected error status %v", c.name, err)
		}
		if !errReceived && c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		a, want     float64
		errExpected bool
	}{
		{"The square root of 1 is 1", 1, 1, false},
		{"The square root of zero is zero", 0, 0, false},
		{"The square root of a two digit number", 25, 5, false},
		{"The square root of a large number", 1947028, 1395.3594518976106, false},
		{"The square root of a small fraction", 0.00000002, 0.0001414213562373095, false},
		{"Attempting to calculate the square root of a negative number returns an error", -25, 0, true},
	}
	for _, c := range testCases {
		got, err := Sqrt(c.a)
		errReceived := err != nil
		if c.errExpected != errReceived {
			t.Fatalf("%s: unexpected error status %v", c.name, err)
		}
		if !errReceived && c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}
}
func TestEvaluate(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		expression  string
		want        float64
		errExpected bool
	}{
		{"A simple valid addition expression", "1 + 1", 2, false},
		{"A simple valid multiplication expression", "10 * 10.5", 105, false},
		{"A valid expression of subtracting two floats", "3.15 - 12.55", -9.4, false},
		{`Trying to evaluate an unimplemented operation returns an error`, "2 % 5", 0, true},
		{"Trying to evaluate an invalid operation returns an error", "51x5.3", 0, true},
		{"An expression with both operands containing fractions", "43.75 / 3.5", 12.5, false},
		{"An expression containing no spaces", "2/2", 0, true},
		{"An expression with whitepsace around the second operand", "1000/    10   \n	  ", 0, true},
		{`Trying to evaluate an expression with more than two operands returns an error`, "8 * 3 / 9", 0, true},
		{"Trying to evalulate a non-expression text returns an error", "yeah, no", 0, true},
	}
	for _, c := range testCases {
		got, err := Evaluate(c.expression)
		errReceived := err != nil
		if c.errExpected != errReceived {
			t.Fatalf("%s: unexpected error status %v", c.name, err)
		}
		if !errReceived && c.want != got {
			t.Errorf("%s: want %f, got %f", c.name, c.want, got)
		}
	}

}
