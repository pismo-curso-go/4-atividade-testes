package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(2, 3)
	esperado := 5.0

	if resultado != esperado {
		t.Errorf("Soma falhou: esperado %.2f, obtido %.2f", esperado, resultado)
	}
}

func TestSubtracao(t *testing.T) {
	resultado := Subtracao(10, 4)
	esperado := 6.0

	if resultado != esperado {
		t.Errorf("Subtração falhou: esperado %.2f, obtido %.2f", esperado, resultado)
	}
}

func TestMultiplicacao(t *testing.T) {
	resultado := Multiplicacao(-2, 3)
	esperado := -6.0

	if resultado != esperado {
		t.Errorf("Multiplicação falhou: esperado %.2f, obtido %.2f", esperado, resultado)
	}
}

func TestDivisao(t *testing.T) {
	// caso normal
	resultado, err := Divisao(10, 2)
	if err != nil {
		t.Errorf("Erro inesperado na divisão: %v", err)
	}
	if resultado != 5 {
		t.Errorf("Divisão falhou: esperado 5, obtido %.2f", resultado)
	}

	// divisão por zero
	_, err = Divisao(10, 0)
	if err == nil {
		t.Error("Esperado erro ao dividir por zero, mas nenhum erro ocorreu")
	}
}
