package main

import (
	"fmt"
	"testing"
)

// Defina uma estrutura de teste com casos de teste.
var tests = []struct{ input, QuadradoDaSoma, SomaDosQuadrados int }{
	{input: 5, QuadradoDaSoma: 225, SomaDosQuadrados: 55},
	{input: 10, QuadradoDaSoma: 3025, SomaDosQuadrados: 385},
	{input: 100, QuadradoDaSoma: 25502500, SomaDosQuadrados: 338350},
}

// TestQuadradoDaSoma testa a função QuadradoDaSoma.
func TestQuadradoDaSoma(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Quadrado da soma de 1 a %d", test.input), func(t *testing.T) {
			// Verifica se o resultado da função QuadradoDaSoma é igual ao valor esperado.
			if got := QuadradoDaSoma(test.input); got != test.QuadradoDaSoma {
				t.Fatalf("QuadradoDaSoma(%d) = %d, queria: %d", test.input, got, test.QuadradoDaSoma)
			}
		})
	}
}

// TestSomaDosQuadrados testa a função SomaDosQuadrados.
func TestSomaDosQuadrados(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Soma dos quadrados de 1 a %d", test.input), func(t *testing.T) {
			// Verifica se o resultado da função SomaDosQuadrados é igual ao valor esperado.
			if got := SomaDosQuadrados(test.input); got != test.SomaDosQuadrados {
				t.Fatalf("SomaDosQuadrados(%d) = %d, queria: %d", test.input, got, test.SomaDosQuadrados)
			}
		})
	}
}

// TestDiferenca testa a função Diferenca.
func TestDiferenca(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Diferença entre Quadrado da Soma e Soma dos Quadrados de %d", test.input), func(t *testing.T) {
			// Calcula o valor esperado da diferença.
			want := test.QuadradoDaSoma - test.SomaDosQuadrados

			// Verifica se o resultado da função Diferenca é igual ao valor esperado.
			if got := Diferenca(test.input); got != want {
				t.Fatalf("Diferenca(%d) = %d, queria: %d", test.input, got, want)
			}
		})
	}
}

// Funções de benchmark para medir o desempenho em um único número (100, do problema original PE)
// para evitar a sobrecarga de iterar sobre testes.
func BenchmarkQuadradoDaSoma(b *testing.B) {
	if testing.Short() {
		b.Skip("pulando benchmark no modo curto.")
	}
	for i := 0; i < b.N; i++ {
		QuadradoDaSoma(100)
	}
}

func BenchmarkSomaDosQuadrados(b *testing.B) {
	if testing.Short() {
		b.Skip("pulando benchmark no modo curto.")
	}
	for i := 0; i < b.N; i++ {
		SomaDosQuadrados(100)
	}
}

func BenchmarkDiferenca(b *testing.B) {
	if testing.Short() {
		b.Skip("pulando benchmark no modo curto.")
	}
	for i := 0; i < b.N; i++ {
		Diferenca(100)
	}
}
