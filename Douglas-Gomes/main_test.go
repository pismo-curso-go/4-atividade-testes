package main

import (
	"testing"
)

type casoSimples struct {
	a, b     float64
	esperado float64
}

func testarOperacaoSimples(t *testing.T, nome string, operacao func(float64, float64) float64, casos []casoSimples) {
	t.Run(nome, func(t *testing.T) {
		for _, c := range casos {
			resultado := operacao(c.a, c.b)
			if resultado != c.esperado {
				t.Errorf("%s(%v, %v) = %v; esperado %v", nome, c.a, c.b, resultado, c.esperado)
			}
		}
	})
}

func TestCalculadora(t *testing.T) {
	testarOperacaoSimples(t, "Soma", Soma, []casoSimples{
		{2, 3, 5},
		{-1, -1, -2},
		{0, 5, 5},
	})

	testarOperacaoSimples(t, "Subtrai", Subtrai, []casoSimples{
		{10, 3, 7},
		{-5, -5, 0},
		{0, 4, -4},
	})

	testarOperacaoSimples(t, "Multiplica", Multiplica, []casoSimples{
		{2, 4, 8},
		{-3, 3, -9},
		{0, 5, 0},
	})

	t.Run("Divide", func(t *testing.T) {
		casos := []struct {
			a, b     float64
			esperado float64
			erro     bool
		}{
			{10, 2, 5, false},
			{-6, 3, -2, false},
			{0, 5, 0, false},
			{5, 0, 0, true},
		}

		for _, c := range casos {
			resultado, err := Divide(c.a, c.b)
			if c.erro {
				if err == nil {
					t.Errorf("esperado erro ao dividir %v por %v, mas não houve", c.a, c.b)
				}
			} else {
				if err != nil {
					t.Errorf("não era esperado erro: %v", err)
				}
				if resultado != c.esperado {
					t.Errorf("Divide(%v, %v) = %v; esperado %v", c.a, c.b, resultado, c.esperado)
				}
			}
		}
	})
}
