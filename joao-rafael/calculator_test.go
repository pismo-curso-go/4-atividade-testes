package main

import (
	"errors"
	"testing"
)

type Test struct {
	num1           int
	num2           int
	ExpectedResult int
	ExpectedError  error
}

func TestAdd(t *testing.T) {
	tests := []Test{
		{3, 5, 8, nil},
		{-2, 4, 2, nil},
		{0, 0, 0, nil},
	}

	for _, test := range tests {
		result := Add(float64(test.num1), float64(test.num2))
		if int(result) != test.ExpectedResult {
			t.Errorf("add(%d, %d) = %d; expected %d", test.num1, test.num2, int(result), test.ExpectedResult)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []Test{
		{10, 4, 6, nil},
		{-3, -3, 0, nil},
		{0, 5, -5, nil},
	}

	for _, test := range tests {
		result := Subtract(float64(test.num1), float64(test.num2))
		if int(result) != test.ExpectedResult {
			t.Errorf("subtract(%d, %d) = %d; expected %d", test.num1, test.num2, int(result), test.ExpectedResult)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []Test{
		{2, 4, 8, nil},
		{0, 100, 0, nil},
		{-3, -3, 9, nil},
	}

	for _, test := range tests {
		result := Multiply(float64(test.num1), float64(test.num2))
		if int(result) != test.ExpectedResult {
			t.Errorf("multiply(%d, %d) = %d; expected %d", test.num1, test.num2, int(result), test.ExpectedResult)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []Test{
		{8, 2, 4, nil},
		{9, 3, 3, nil},
		{5, 0, 0, errors.New("cannot divide by zero")},
	}

	for _, test := range tests {
		result, err := Divide(float64(test.num1), float64(test.num2))

		if test.ExpectedError != nil {
			if err == nil {
				t.Errorf("divide(%d, %d) expected error but got none", test.num1, test.num2)
			}
			continue
		}

		if err != nil {
			t.Errorf("divide(%d, %d) unexpected error: %v", test.num1, test.num2, err)
			continue
		}

		if int(result) != test.ExpectedResult {
			t.Errorf("divide(%d, %d) = %d; expected %d", test.num1, test.num2, int(result), test.ExpectedResult)
		}
	}
}
