package ruan_costa

import (
	"errors"
	"math"
)

type Calculadora struct{}

func (c Calculadora) Soma(n1 float64, n2 float64) float64 {
	resultado := n1 + n2
	return resultado
}

func (c Calculadora) Divisao(n1 float64, n2 float64) (float64, error) {
	if n2 == 0 {
		return math.NaN(), errors.New("um número não pode ser dividido por 0")
	}
	resultado := n1 / n2
	return resultado, nil
}

func (c Calculadora) Multiplicacao(n1 float64, n2 float64) float64 {
	resultado := n1 * n2
	return resultado
}

func (c Calculadora) Subtracao(n1 float64, n2 float64) float64 {
	resultado := n1 - n2
	return resultado
}
