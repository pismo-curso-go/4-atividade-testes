package main

import (
	"calculadora"
	"fmt"
)

func main() {

	var c ruan_costa.Calculadora
	fmt.Println(c.Soma(5, 8))
	fmt.Println(c.Divisao(5, 8))
	fmt.Println(c.Subtracao(5, 8))
	fmt.Println(c.Divisao(5, 0))
	fmt.Println(c.Multiplicacao(5, 8))
}
