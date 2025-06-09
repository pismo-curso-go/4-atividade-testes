package ruan_costa

import (
	"math"
	"testing"
)

func TestSoma(t *testing.T) {
	calc := Calculadora{}
	tests := []struct {
		name     string
		n1, n2   float64
		expected float64
	}{
		{"Soma simples", 3, 7, 10},
		{"Soma com zero", 5, 0, 5},
		{"Soma negativa", -3, -2, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Soma(tt.n1, tt.n2)
			if result != tt.expected {
				t.Errorf("esperado %v, obtido %v", tt.expected, result)
			}
		})
	}
}

func TestSubtracao(t *testing.T) {
	calc := Calculadora{}
	tests := []struct {
		name     string
		n1, n2   float64
		expected float64
	}{
		{"Subtração simples", 10, 3, 7},
		{"Subtração com negativos", -4, -2, -2},
		{"Subtração que dá zero", 5, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtracao(tt.n1, tt.n2)
			if result != tt.expected {
				t.Errorf("esperado %v, obtido %v", tt.expected, result)
			}
		})
	}
}

func TestMultiplicacao(t *testing.T) {
	calc := Calculadora{}
	tests := []struct {
		name     string
		n1, n2   float64
		expected float64
	}{
		{"Multiplicação simples", 2, 3, 6},
		{"Multiplicação com zero", 7, 0, 0},
		{"Multiplicação com negativos", -2, 4, -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiplicacao(tt.n1, tt.n2)
			if result != tt.expected {
				t.Errorf("esperado %v, obtido %v", tt.expected, result)
			}
		})
	}
}

func TestDivisao(t *testing.T) {
	calc := Calculadora{}
	tests := []struct {
		name          string
		n1, n2        float64
		expected      float64
		expectingErro bool
	}{
		{"Divisão simples", 10, 2, 5, false},
		{"Divisão com zero", 10, 0, math.NaN(), true},
		{"Divisão negativa", -8, 2, -4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divisao(tt.n1, tt.n2)
			if tt.expectingErro {
				if err == nil {
					t.Errorf("esperava erro, mas não recebeu nenhum")
				}
			} else {
				if err != nil {
					t.Errorf("não esperava erro, mas recebeu: %v", err)
				}
				if result != tt.expected {
					t.Errorf("esperado %v, obtido %v", tt.expected, result)
				}
			}
		})
	}
}
