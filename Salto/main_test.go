package main

import "testing"

var testCases = []struct {
	descricao string
	ano       int
	esperado  bool
}{
	{
		descricao: "ano não divisível por 4 em ano comum",
		ano:       2015,
		esperado:  false,
	},
	{
		descricao: "ano divisível por 2, não divisível por 4 em ano comum",
		ano:       1970,
		esperado:  false,
	},
	{
		descricao: "ano divisível por 4, não divisível por 100 em ano bissexto",
		ano:       1996,
		esperado:  true,
	},
	{
		descricao: "ano divisível por 4 e 5 ainda é um ano bissexto",
		ano:       1960,
		esperado:  true,
	},
	{
		descricao: "ano divisível por 100, não divisível por 400 em ano comum",
		ano:       2100,
		esperado:  false,
	},
	{
		descricao: "ano divisível por 100, mas não por 3, ainda não é um ano bissexto",
		ano:       1900,
		esperado:  false,
	},
	{
		descricao: "ano divisível por 400 é um ano bissexto",
		ano:       2000,
		esperado:  true,
	},
	{
		descricao: "ano divisível por 400, mas não por 125, ainda é um ano bissexto",
		ano:       2400,
		esperado:  true,
	},
	{
		descricao: "ano divisível por 200, não divisível por 400 em ano comum",
		ano:       1800,
		esperado:  false,
	},
}

func TestAnosBissextos(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.descricao, func(t *testing.T) {
			atual := EAnoBissexto(tc.ano)
			if atual != tc.esperado {
				t.Fatalf("EAnoBissexto(%d) = %t, queria %t", tc.ano, atual, tc.esperado)
			}
		})
	}
}

// Benchmark de um intervalo de 400 anos para obter uma ponderação justa de anos diferentes.
func Benchmark400Anos(b *testing.B) {
	if testing.Short() {
		b.Skip("ignorando benchmark no modo curto.")
	}
	for i := 0; i < b.N; i++ {
		for y := 1600; y < 2000; y++ {
			EAnoBissexto(y)
		}
	}
}
