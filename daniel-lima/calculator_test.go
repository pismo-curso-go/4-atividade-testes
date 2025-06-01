package calculator

import (
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		a, b, expected float64
	}{
		{1, 2, 3},
		{-1, -2, -3},
		{0, 0, 0},
		{2.5, 3.5, 6.0},
		{-1, 1, 0},
		{1000000, 1, 1000001},
	}

	for _, tc := range testCases {
		result := Sum(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Sum(%f, %f) = %f; expected %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestSubtraction(t *testing.T) {
	testCases := []struct {
		a, b, expected float64
	}{
		{5, 3, 2},
		{3, 5, -2},
		{0, 0, 0},
		{-1, -1, 0},
		{2.5, 1.5, 1.0},
		{-1, 1, -2},
	}

	for _, tc := range testCases {
		result := Subtraction(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Subtraction(%f, %f) = %f; expected %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestMultiplication(t *testing.T) {
	testCases := []struct {
		a, b, expected float64
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{0, 5, 0},
		{-3, -3, 9},
		{0.5, 4, 2},
		{100, 0.1, 10},
	}

	for _, tc := range testCases {
		result := Multiplication(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Multiplication(%f, %f) = %f; expected %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestDivision(t *testing.T) {
	testCases := []struct {
		a, b, expected float64
		expectError    bool
	}{
		{6, 3, 2, false},
		{1, 2, 0.5, false},
		{-9, 3, -3, false},
		{0, 1, 0, false},
		{-4, -2, 2, false},
		{1, 3, 0.333333, false},
		{1, 0, 0, true},
	}

	for _, tc := range testCases {
		result, err := Division(tc.a, tc.b)

		if tc.expectError {
			if err == nil {
				t.Errorf("Division(%f, %f) should return error but didn't", tc.a, tc.b)
			}
			if err.Error() != "division by zero is not allowed" {
				t.Errorf("Division(%f, %f) returned wrong error message: %v", tc.a, tc.b, err)
			}
		} else {
			if err != nil {
				t.Errorf("Division(%f, %f) returned unexpected error: %v", tc.a, tc.b, err)
			}

			margin := 0.000001
			if (result-tc.expected) > margin || (tc.expected-result) > margin {
				t.Errorf("Division(%f, %f) = %f; expected %f", tc.a, tc.b, result, tc.expected)
			}
		}
	}
}
