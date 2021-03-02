package calculator_test

import (
	"calculator"
	"testing"
)

func TestAddSubtractAndMultiply(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		desc   string
		fn     func(float64, float64, ...float64) float64
		a      float64
		b      float64
		extras []float64
		want   float64
	}{
		{
			desc:   "Adding two positive numbers yields a positive number",
			fn:     calculator.Add,
			a:      231,
			b:      56,
			extras: []float64{},
			want:   287,
		}, {
			desc:   "Adding one large and one small number",
			fn:     calculator.Add,
			a:      9,
			b:      123191023123,
			extras: []float64{},
			want:   123191023132,
		}, {
			desc:   "Adding two zeros yields zero",
			fn:     calculator.Add,
			a:      0,
			b:      0,
			extras: []float64{},
			want:   0,
		}, {
			desc:   "Adding a number with a fraction and a number without one yields a number with fraction",
			fn:     calculator.Add,
			a:      1.2,
			b:      333,
			extras: []float64{},
			want:   334.2,
		}, {
			desc:   "Adding a negative number and a positive number",
			fn:     calculator.Add,
			a:      -73,
			b:      99,
			extras: []float64{},
			want:   26,
		}, {
			desc:   "Adding two negative numbers yields a negative number",
			fn:     calculator.Add,
			a:      -1,
			b:      -1,
			extras: []float64{},
			want:   -2,
		}, {
			desc:   "Adding a number to zero yields the same number",
			fn:     calculator.Add,
			a:      -2222,
			b:      0,
			extras: []float64{},
			want:   -2222,
		}, {
			desc:   "Adding five very small numbers",
			fn:     calculator.Add,
			a:      0.0000005133,
			b:      0.00000312,
			extras: []float64{0.000664, 0.0000532, 0.00001},
			want:   0.0007308333,
		}, {
			desc:   "Subtracting a positive number from itself yields zero",
			fn:     calculator.Subtract,
			a:      1,
			b:      1,
			extras: []float64{},
			want:   0,
		}, {
			desc:   "Subtracting a negative number from itself yields zero",
			fn:     calculator.Subtract,
			a:      -51,
			b:      -51,
			extras: []float64{},
			want:   0,
		}, {
			desc:   "Subtracting multiple positive numbers from a much larger positive number yields a positive number",
			fn:     calculator.Subtract,
			a:      4323,
			b:      12,
			extras: []float64{11, 39},
			want:   4261,
		}, {
			desc:   "Subtracting a positive number from a smaller positive number yields a negative number",
			fn:     calculator.Subtract,
			a:      11.2,
			b:      89.8,
			extras: []float64{},
			want:   -78.6,
		}, {
			desc:   "Subtracting a negative number from a larger negative number yields a positive number",
			fn:     calculator.Subtract,
			a:      -100,
			b:      -111,
			extras: []float64{},
			want:   11,
		}, {
			desc:   "Subtracting multiple positive numbers from zero yields a negative number",
			fn:     calculator.Subtract,
			a:      0,
			b:      6,
			extras: []float64{6, 15},
			want:   -27,
		}, {
			desc:   "Subtracting a negative number from zero yields a positive number",
			fn:     calculator.Subtract,
			a:      0,
			b:      -22,
			extras: []float64{},
			want:   22,
		}, {
			desc:   "Multiplying 1 by 1 always yields 1",
			fn:     calculator.Multiply,
			a:      1,
			b:      1,
			extras: []float64{},
			want:   1,
		}, {
			desc:   "Multiplying a number by 1 any number of times yields that number",
			fn:     calculator.Multiply,
			a:      3999191,
			b:      1,
			extras: []float64{1, 1, 1},
			want:   3999191,
		}, {
			desc:   "Multiplying a number by itself equals that number squared",
			fn:     calculator.Multiply,
			a:      5091,
			b:      5091,
			extras: []float64{},
			want:   25918281,
		}, {
			desc:   "Multiplying two negative numbers yields a positive number",
			fn:     calculator.Multiply,
			a:      -10,
			b:      -10,
			extras: []float64{},
			want:   100,
		}, {
			desc:   "Multiplying three negative number yields a negative number",
			fn:     calculator.Multiply,
			a:      -5,
			b:      -12,
			extras: []float64{-67},
			want:   -4020,
		}, {
			desc:   "Multiplying a positive number by a negative number yields a negative number",
			fn:     calculator.Multiply,
			a:      12,
			b:      -1.5,
			extras: []float64{},
			want:   -18,
		}, {
			desc:   "Multiplying any number by zero yields zero",
			fn:     calculator.Multiply,
			a:      43,
			b:      0,
			extras: []float64{},
			want:   0,
		},
	}
	for _, c := range testCases {
		got := c.fn(c.a, c.b, c.extras...)
		if c.want != got {
			t.Errorf("%s: want %f, got %f", c.desc, c.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		desc        string
		a           float64
		b           float64
		extras      []float64
		want        float64
		errExpected bool
	}{
		{
			desc:        "Dividing a number by 1 yields the same number",
			a:           10,
			b:           1,
			extras:      []float64{},
			want:        10,
			errExpected: false,
		}, {
			desc:        "Dividing a number by 1 multiple times yields the same number",
			a:           99,
			b:           1,
			extras:      []float64{1, 1},
			want:        99,
			errExpected: false,
		}, {
			desc:        "Dividing a number by zero returns an error",
			a:           312,
			b:           0,
			extras:      []float64{},
			want:        0,
			errExpected: true,
		}, {
			desc:        "Dividing a number by multiple numbers with zero being one of them returns an error",
			a:           50,
			b:           1,
			extras:      []float64{5, 0, 200},
			want:        0,
			errExpected: true,
		}, {
			desc:        "Dividing a zero by zero returns an error",
			a:           0,
			b:           0,
			extras:      []float64{},
			want:        0,
			errExpected: true,
		}, {
			desc:        "Dividing a number by multiple larger numbers yields a fraction",
			a:           1,
			b:           2,
			extras:      []float64{2, 2},
			want:        0.125,
			errExpected: false,
		}, {
			desc:        "Dividing a number by a large number and a fraction",
			a:           100,
			b:           0.1,
			extras:      []float64{10},
			want:        100,
			errExpected: false,
		}, {
			desc:        "Dividing a number by a much larger number yields a small fraction",
			a:           100,
			b:           10000000,
			extras:      []float64{},
			want:        0.00001,
			errExpected: false,
		},
	}
	for _, c := range testCases {
		got, err := calculator.Divide(c.a, c.b, c.extras...)
		errReceived := err != nil
		if c.errExpected != errReceived {
			t.Fatalf("%s: unexpected error status %v", c.desc, err)
		}
		if !errReceived && c.want != got {
			t.Errorf("%s: want %f, got %f", c.desc, c.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		desc        string
		a, want     float64
		errExpected bool
	}{
		{
			desc:        "The square root of 1 is 1",
			a:           1,
			want:        1,
			errExpected: false,
		}, {
			desc:        "The square root of zero is zero",
			a:           0,
			want:        0,
			errExpected: false,
		}, {
			desc:        "The square root of a two digit number",
			a:           25,
			want:        5,
			errExpected: false,
		}, {
			desc:        "The square root of a large number",
			a:           1947028,
			want:        1395.3594518976106,
			errExpected: false,
		}, {
			desc:        "The square root of a small fraction",
			a:           0.00000002,
			want:        0.0001414213562373095,
			errExpected: false,
		}, {
			desc:        "Attempting to calculate the square root of a negative number returns an error",
			a:           -25,
			want:        0,
			errExpected: true,
		},
	}
	for _, c := range testCases {
		got, err := calculator.Sqrt(c.a)
		errReceived := err != nil
		if c.errExpected != errReceived {
			t.Fatalf("%s: unexpected error status %v", c.desc, err)
		}
		if !errReceived && c.want != got {
			t.Errorf("%s: want %f, got %f", c.desc, c.want, got)
		}
	}
}
func TestEvaluate(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		desc        string
		expression  string
		want        float64
		errExpected bool
	}{
		{
			desc:        "A simple valid addition expression",
			expression:  "1 + 1",
			want:        2,
			errExpected: false,
		}, {
			desc:        "A simple valid multiplication expression",
			expression:  "10 * 10.5",
			want:        105,
			errExpected: false,
		}, {
			desc:        "A valid expression of subtracting two floats",
			expression:  "3.15 - 12.55",
			want:        -9.4,
			errExpected: false,
		}, {
			desc:        "Trying to evaluate an unimplemented operation returns an error",
			expression:  "2 % 5",
			want:        0,
			errExpected: true,
		}, {
			desc:        "Trying to evaluate an invalid operation returns an error",
			expression:  "51x5.3",
			want:        0,
			errExpected: true,
		}, {
			desc:        "An expression with both operands containing fractions",
			expression:  "43.75 / 3.5",
			want:        12.5,
			errExpected: false,
		}, {
			desc:        "An expression containing no spaces",
			expression:  "2/2",
			want:        0,
			errExpected: true,
		}, {
			desc: "An expression with whitepsace around the second operand",
			expression: "1000/    10   \n	  ",
			want:        0,
			errExpected: true,
		}, {
			desc:        "Trying to evaluate an expression with more than two operands returns an error",
			expression:  "8 * 3 / 9",
			want:        0,
			errExpected: true,
		}, {
			desc:        "Trying to evalulate a non-expression text returns an error",
			expression:  "yeah, no",
			want:        0,
			errExpected: true,
		},
	}
	for _, c := range testCases {
		got, err := calculator.Evaluate(c.expression)
		errReceived := err != nil
		if c.errExpected != errReceived {
			t.Fatalf("%s: unexpected error status %v", c.desc, err)
		}
		if !errReceived && c.want != got {
			t.Errorf("%s: want %f, got %f", c.desc, c.want, got)
		}
	}

}
