package main

import (
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