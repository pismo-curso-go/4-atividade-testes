package natalia

import "testing"

func TestCalculadora(t *testing.T) {
	t.Run("Soma", func(soma *testing.T) {
		result := Soma(1, 2)
		if result != 3 {
			t.Errorf("esperando 3, obteve %.2f", result)
		}
	})

	t.Run("Subtração", func(subt *testing.T) {
		result := Subtrai(10, 4)
		if result != 6 {
			t.Errorf("Esperado 6, obteve %.2f", result)
		}
	})

	t.Run("Multiplicação", func(mult *testing.T) {
		result := Multiplica(-2, 3)
		if result != -6 {
			t.Errorf("Esperado -6, obteve %.2f", result)
		}
	})

	t.Run("Divisão válida", func(div *testing.T) {
		result, err := Divide(10, 2)
		if err != nil || result != 5 {
			t.Errorf("Esperado 5 sem erro, obteve %.2f com erro %v", result, err)
		}
	})

	t.Run("Divisão por zero", func(div *testing.T) {
		_, err := Divide(5, 0)
		if err == nil {
			t.Error("Esperado erro ao dividir por zero, mas não ocorreu")
		}
	})
}
