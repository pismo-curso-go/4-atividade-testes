package main

import "errors"

func Soma(a, b float64) float64 {
	return a + b
}

func Subtrai(a, b float64) float64 {
	return a - b
}
func Multiplica(a, b float64) float64 {
	return a * b
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("erro: divisão por zero")
	}
	return a / b, nil
}
