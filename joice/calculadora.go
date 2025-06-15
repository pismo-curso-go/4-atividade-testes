package main

import (
	"errors"
	"fmt"
)

func Soma(a, b float64) float64 {
	return a + b
}

func Subtracao(a, b float64) float64 {
	return a - b
}

func Multiplicacao(a, b float64) float64 {
	return a * b
}

func Divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Não é possiível dividir por zero!")
	}
	return a / b, nil
}

func main() {
	var opcao int
	var a, b float64
	fmt.Println("=== Calculadora Interativa ===")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	fmt.Println("Escolha uma opção: ")
	fmt.Scanln(&opcao)

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&b)

	switch opcao {
	case 1:
		fmt.Printf("Resultado da soma: %.2f\n", Soma(a, b))
	case 2:
		fmt.Printf("Resultado da Subtração: %.2f\n", Subtracao(a, b))
	case 3:
		fmt.Printf("Resultado da Multiplicação: %.2f\n", Multiplicacao(a, b))
	case 4:
		resultado, err := Divisao(a, b)
		if err != nil {
			fmt.Println("Erro na divsão:", err)
		} else {
			fmt.Printf("Divisão: %.2f\n", resultado)
		}
	default:
		fmt.Println("Opção inválida.")
	}

}
