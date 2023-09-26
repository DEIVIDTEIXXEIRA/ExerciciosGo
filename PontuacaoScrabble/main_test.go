package main

import "testing"

type testeDeScrabble struct {
	descricao string
	entrada   string
	esperado  int
}

var testesDeScoreDoScrabble = []testeDeScrabble{
	{
		descricao: "letra minúscula",
		entrada:   "a",
		esperado:  1,
	},
	{
		descricao: "letra maiúscula",
		entrada:   "A",
		esperado:  1,
	},
	{
		descricao: "letra valiosa",
		entrada:   "f",
		esperado:  4,
	},
	{
		descricao: "palavra curta",
		entrada:   "at",
		esperado:  2,
	},
	{
		descricao: "palavra curta e valiosa",
		entrada:   "zoo",
		esperado:  12,
	},
	{
		descricao: "palavra média",
		entrada:   "street",
		esperado:  6,
	},
	{
		descricao: "palavra média e valiosa",
		entrada:   "quirky",
		esperado:  22,
	},
	{
		descricao: "palavra longa com mistura de maiúsculas",
		entrada:   "OxyphenButazone",
		esperado:  41,
	},
	{
		descricao: "palavra semelhante ao inglês",
		entrada:   "pinata",
		esperado:  8,
	},
	{
		descricao: "entrada vazia",
		entrada:   "",
		esperado:  0,
	},
	{
		descricao: "alfabeto completo disponível",
		entrada:   "abcdefghijklmnopqrstuvwxyz",
		esperado:  87,
	},
}

func TestScore(t *testing.T) {
	for _, tc := range testesDeScoreDoScrabble {
		t.Run(tc.descricao, func(t *testing.T) {
			if atual := Score(tc.entrada); atual != tc.esperado {
				t.Errorf("Score(%q) = %d, queria:%d", tc.entrada, atual, tc.esperado)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("ignorando benchmark no modo curto.")
	}
	for i := 0; i < b.N; i++ {
		for _, teste := range testesDeScoreDoScrabble {
			Score(teste.entrada)
		}
	}
}
