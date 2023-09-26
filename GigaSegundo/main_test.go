package main

import (
	"testing"
	"time"
)

var addCases = []struct {
	descricao string
	entrada   string
	esperado  string
}{
	{
		descricao: "especificação de data apenas",
		entrada:   "2011-04-25",
		esperado:  "2043-01-01T01:46:40",
	},
	{
		descricao: "segundo teste para especificação de data apenas",
		entrada:   "1977-06-13",
		esperado:  "2009-02-19T01:46:40",
	},
	{
		descricao: "terceiro teste para especificação de data apenas",
		entrada:   "1959-07-19",
		esperado:  "1991-03-27T01:46:40",
	},
	{
		descricao: "tempo completo especificado",
		entrada:   "2015-01-24T22:00:00",
		esperado:  "2046-10-02T23:46:40",
	},
	{
		descricao: "tempo completo com rolagem de dia",
		entrada:   "2015-01-24T23:59:59",
		esperado:  "2046-10-03T01:46:39",
	},
}

// Formatos de data usados nos dados de teste
const (
	fmtD  = "2006-01-02"
	fmtDT = "2006-01-02T15:04:05"
)

func TestAddGigasecond(t *testing.T) {
	for _, tc := range addCases {
		t.Run(tc.descricao, func(t *testing.T) {
			entrada := parse(tc.entrada, t)
			esperado := parse(tc.esperado, t)
			obtido := AddGigasecond(entrada)
			if !obtido.Equal(esperado) {
				t.Fatalf("AddGigasecond(%v) = %v, esperado: %v", entrada, obtido, esperado)
			}
		})
	}
}

func parse(s string, t *testing.T) time.Time {
	tt, err := time.Parse(fmtDT, s) // tenta primeiro o formato de data e hora completo
	if err != nil {
		tt, err = time.Parse(fmtD, s) // permite também apenas a data
	}
	if err != nil {
		t.Fatalf("erro na configuração do teste: TestAddGigasecond requer datetime em um dos seguintes formatos:\nformato 1: %q\nformato 2: %q\ngot: %q", fmtD, fmtDT, s)
	}
	return tt
}

func BenchmarkAddGigasecond(b *testing.B) {
	if testing.Short() {
		b.Skip("ignorando teste de benchmark no modo curto.")
	}
	for i := 0; i < b.N; i++ {
		AddGigasecond(time.Time{})
	}
}
