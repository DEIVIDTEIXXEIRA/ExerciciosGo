package main

import (
	"errors"
	"fmt"
)

func main() {
	linha1:= "GAGCCTACTAACGGGAT"
	linha2:= "CATCGTAATGACGGCCT"

	distancia, erro := Distance(linha1, linha2)
	if erro != nil {
		return
	}

	fmt.Printf("A Distância de Hamming é: %d", distancia)
}

// Distance retorna a Distância de Hamming
func Distance(a, b string) (int, error) {
	linha1 := a
	linha2 := b

	if len(linha1) != len(linha2) {
		return 0, errors.New("Linhas de tamanhos diferentes")
	}

	distancia := 0 

	for i := 0; i < len(linha1); i++ {
		if linha1[i] != linha2[i] {
			distancia++
		}
	}

	return distancia, nil 
}
