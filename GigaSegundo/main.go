package main

import (
	"fmt"
	"time"
)

// AddGigasecond adiciona um gigasegundo (1.000.000.000 de segundos) a uma data e hora fornecidas.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1e9) * time.Second
	return t.Add(gigasecond)
}

func main() {
	dataString := "23-09-2023"

	data, err := time.Parse("02-01-2006", dataString)
	if err != nil {
		fmt.Println("Erro ao analisar a data:", err)
		return
	}

	resultado := AddGigasecond(data)
	fmt.Printf("Você completará um gigasegundo em: %02d-%02d-%04d", resultado.Day(), resultado.Month(), resultado.Year())
}
