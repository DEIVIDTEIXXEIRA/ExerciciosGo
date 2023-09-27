package main

import (
	"testing"
)

// Define uma estrutura de teste com casos de teste.
var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "string vazia",
		input:       "",
		expected:    true,
	},
	{
		description: "isograma com apenas caracteres minúsculos",
		input:       "isogram",
		expected:    true,
	},
	{
		description: "palavra com um caractere duplicado",
		input:       "eleven",
		expected:    false,
	},
	{
		description: "palavra com um caractere duplicado do final do alfabeto",
		input:       "zzyzx",
		expected:    false,
	},
	{
		description: "maior isograma relatado em inglês",
		input:       "subdermatoglyphic",
		expected:    true,
	},
	{
		description: "palavra com caractere duplicado em caixa mista",
		input:       "Alphabet",
		expected:    false,
	},
	{
		description: "palavra com caractere duplicado em caixa mista, minúsculas primeiro",
		input:       "alphAbet",
		expected:    false,
	},
	{
		description: "palavra isograma hipotética com hífen",
		input:       "thumbscrew-japingly",
		expected:    true,
	},
	{
		description: "palavra hipotética com caractere duplicado após o hífen",
		input:       "thumbscrew-jappingly",
		expected:    false,
	},
	{
		description: "isograma com hífen duplicado",
		input:       "six-year-old",
		expected:    true,
	},
	{
		description: "nome inventado que é um isograma",
		input:       "Emily Jung Schwartzkopf",
		expected:    true,
	},
	{
		description: "caractere duplicado no meio",
		input:       "accentor",
		expected:    false,
	},
	{
		description: "mesmo primeiro e último caractere",
		input:       "angola",
		expected:    false,
	},
	{
		description: "palavra com caractere duplicado e dois hífens",
		input:       "up-to-date",
		expected:    false,
	},
}

// Função de teste que executa os casos de teste.
func TestPalavrasSaoIsogramas(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := palavrasSaoIsogramas(tc.input); actual != tc.expected {
				t.Fatalf("palavrasSaoIsogramas(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

// Função de benchmark para medir o desempenho.
func BenchmarkPalavrasSaoIsogramas(b *testing.B) {
	if testing.Short() {
		b.Skip("skipando benchmark no modo curto.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			palavrasSaoIsogramas(c.input)
		}
	}
}
