package main

import (
	"errors"
	"testing"
)

type Test struct {
	num1 int
	num2 int
	ExpResult int
	ExpError error
}
func Test_Add (t *testing.T) {
	tests := []Test{
		{1, 2, 3, nil},
		{5, 5, 10, nil},
		{-1, -1, -2, nil},
	}

	for _, test := range tests {
		result := Add(test.num1, test.num2)
		if result != test.ExpResult {
			t.Errorf("Add(%d, %d) = %d; want %d", test.num1, test.num2, result, test.ExpResult)
		}
	}
}

func Test_Subtract(t *testing.T) {
	tests := []Test{
		{5, 2, 3, nil},
		{10, 5, 5, nil},
		{-1, -1, 0, nil},
	}

	for _, test := range tests {
		result := Subtract(test.num1, test.num2)
		if result != test.ExpResult {
			t.Errorf("Subtract(%d, %d) = %d; want %d", test.num1, test.num2, result, test.ExpResult)
		}
	}
}

func Test_Multiply(t *testing.T) {
	tests := []Test{
		{2, 3, 6, nil},
		{5, 5, 25, nil},
		{-1, -1, 1, nil},
	}

	for _, test := range tests {
		result := Multiply(test.num1, test.num2)
		if result != test.ExpResult {
			t.Errorf("Multiply(%d, %d) = %d; want %d", test.num1, test.num2, result, test.ExpResult)
		}
	}
}

func Test_Divide(t *testing.T) {
	tests := []Test{
		{6, 2, 3, nil},
		{10, 5, 2, nil},
		{5, 0, 0, errors.New("invalid division by zero")},
	}

	for _, test := range tests {
		result, err := Divide(test.num1, test.num2)
		if err != nil && test.ExpError == nil {
			t.Errorf("Divide(%d, %d) returned an error: %v; want no error", test.num1, test.num2, err)
			continue
		}
		if err == nil && test.ExpError != nil {
			t.Errorf("Divide(%d, %d) = %f; want error %v", test.num1, test.num2, result, test.ExpError)
			continue
		}
		if result != float64(test.ExpResult) {
			t.Errorf("Divide(%d, %d) = %f; want %d", test.num1, test.num2, result, test.ExpResult)
		}
	}
}
