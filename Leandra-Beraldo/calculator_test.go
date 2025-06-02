package main

import (
	"errors"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{-2, -3, -5},
	}
	for _, tt := range tests {
		result := sum(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("sum(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{5, 3, 2},
		{0, 0, 0},
		{-1, -1, 0},
		{2, 5, -3},
	}
	for _, tt := range tests {
		result := subtract(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("subtract(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 3, -6},
		{-2, -3, 6},
	}
	for _, tt := range tests {
		result := multiply(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("multiply(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b        float64
		expected    float64
		expectedErr error
	}{
		{6, 3, 2, nil},
		{5, 2, 2.5, nil},
		{0, 1, 0, nil},
		{1, 0, 0, errors.New("Division by zero is not possible.")},
	}
	for _, tt := range tests {
		result, err := divide(tt.a, tt.b)
		if tt.b == 0 {
			if err == nil {
				t.Errorf("divide(%v, %v) expected error, got nil", tt.a, tt.b)
			}
		} else {
			if err != nil {
				t.Errorf("divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
			}
			if result != tt.expected {
				t.Errorf("divide(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		}
	}
}
