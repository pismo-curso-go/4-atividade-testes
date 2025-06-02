package calculator_test

import (
	"fmt"
	calculator "paulo-valleriote"
	"testing"
)

type TestData struct {
	name         string
	numbers      []int
	expectResult int
	expectError  bool
}

func TestSum(t *testing.T) {
	t.Parallel()

	cases := []TestData{
		{
			name:         "Test if sum is working with min result (1+0)",
			numbers:      []int{1, 0},
			expectResult: 1,
			expectError:  false,
		},
		{
			name:         "Test sum of two algarisms",
			numbers:      []int{1, 1},
			expectResult: 2,
			expectError:  false,
		},
		{
			name:         "Test sum of four algarisms",
			numbers:      []int{1, 2, 3, 4},
			expectResult: 10,
			expectError:  false,
		},
		{
			name:         "Test if sum function throw error with invalid argument",
			numbers:      []int{1},
			expectResult: 0,
			expectError:  true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			t.Parallel()

			res, err := calculator.Sum(tt.numbers...)
			if !tt.expectError && err != nil {
				t.Fatal("error is not expected")
			}

			if res != tt.expectResult {
				t.Fatal("result is not equal as the expected")
			}
		})
	}
}

func TestSub(t *testing.T) {
	t.Parallel()

	cases := []TestData{
		{
			name:         "Test if sub is working with min result (1-0)",
			numbers:      []int{1, 0},
			expectResult: 1,
			expectError:  false,
		},
		{
			name:         "Test sub of two algarisms",
			numbers:      []int{1, 1},
			expectResult: 0,
			expectError:  false,
		},
		{
			name:         "Test sub of two algarisms (positive-negative)",
			numbers:      []int{1, -1},
			expectResult: 2,
			expectError:  false,
		},
		{
			name:         "Test sub of two algarisms (negative-positive)",
			numbers:      []int{-1, 1},
			expectResult: -2,
			expectError:  false,
		},
		{
			name:         "Test sub of positives resulting in a negative subtraction of negatives",
			numbers:      []int{1, 2, -1},
			expectResult: 0,
			expectError:  false,
		},
		{
			name:         "Test sub of positives resulting in a negative subtraction of negative-positive",
			numbers:      []int{1, 2, 1},
			expectResult: -2,
			expectError:  false,
		},
		{
			name:         "Test if sub function throw error with invalid argument",
			numbers:      []int{1},
			expectResult: 0,
			expectError:  true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			t.Parallel()

			res, err := calculator.Sub(tt.numbers...)
			if !tt.expectError && err != nil {
				t.Fatal("error is not expected")
			}

			if res != tt.expectResult {
				fmt.Printf("got result %d\n", res)
				fmt.Printf("expected result %d\n", tt.expectResult)
				t.Fatal("result is not equal as the expected")
			}
		})
	}
}

func TestMulti(t *testing.T) {
	t.Parallel()

	cases := []TestData{
		{
			name:         "Test multiplication by zero",
			numbers:      []int{1, 0},
			expectResult: 0,
			expectError:  false,
		},
		{
			name:         "Test multiplication of one by one",
			numbers:      []int{1, 1},
			expectResult: 1,
			expectError:  false,
		},
		{
			name:         "Test multiplication by one",
			numbers:      []int{2, 1},
			expectResult: 2,
			expectError:  false,
		},
		{
			name:         "Test multiplication by negative one",
			numbers:      []int{2, -1},
			expectResult: -2,
			expectError:  false,
		},

		{
			name:         "Test multiplication by two",
			numbers:      []int{2, 2},
			expectResult: 4,
			expectError:  false,
		},

		{
			name:         "Test multiplication by negative two",
			numbers:      []int{2, -2},
			expectResult: -4,
			expectError:  false,
		},

		{
			name:         "Test multiplication of negative two by two",
			numbers:      []int{-2, 2},
			expectResult: -4,
			expectError:  false,
		},
		{
			name:         "Test if multiplication function throw error with invalid argument",
			numbers:      []int{1},
			expectResult: 0,
			expectError:  true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			t.Parallel()

			res, err := calculator.Multi(tt.numbers...)
			if !tt.expectError && err != nil {
				t.Fatal("error is not expected")
			}

			if res != tt.expectResult {
				fmt.Printf("got result %d\n", res)
				fmt.Printf("expected result %d\n", tt.expectResult)
				t.Fatal("result is not equal as the expected")
			}
		})
	}
}

func TestDiv(t *testing.T) {
	t.Parallel()

	cases := []TestData{
		{
			name:         "Test if division by zero throws error",
			numbers:      []int{1, 0},
			expectResult: 0,
			expectError:  true,
		},
		{
			name:         "Test division by one",
			numbers:      []int{2, 1},
			expectResult: 2,
			expectError:  false,
		},
		{
			name:         "Test division by two",
			numbers:      []int{4, 2},
			expectResult: 2,
			expectError:  false,
		},
		{
			name:         "Test division for negative one",
			numbers:      []int{2, -1},
			expectResult: -2,
			expectError:  false,
		},

		{
			name:         "Test division by two",
			numbers:      []int{2, 2},
			expectResult: 1,
			expectError:  false,
		},

		{
			name:         "Test division by negative two",
			numbers:      []int{2, -2},
			expectResult: -1,
			expectError:  false,
		},
		{
			name:         "Test if division function throw error with invalid argument",
			numbers:      []int{1},
			expectResult: 0,
			expectError:  true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			t.Parallel()

			res, err := calculator.Div(tt.numbers...)
			if !tt.expectError && err != nil {
				t.Fatal("error is not expected")
			}

			if res != tt.expectResult {
				fmt.Printf("got result %d\n", res)
				fmt.Printf("expected result %d\n", tt.expectResult)
				t.Fatal("result is not equal as the expected")
			}
		})
	}
}
