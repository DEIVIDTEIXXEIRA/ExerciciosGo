package main

import (
	"fmt"
)

func main() {
	numeroInicial := 12

	resultadoConta, _ := CollatzCount(numeroInicial)
	resultadoSequencia := Collatz(numeroInicial)

	fmt.Printf("Sequência de Collatz para %d: %v\n", numeroInicial, resultadoSequencia)
	fmt.Printf("Sequência de Collatz para %d: %v\n", numeroInicial, resultadoConta)
}

// Collatz retorna a sequancia 
func Collatz(n int) []int {
	sequencia := []int{n}

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		sequencia = append(sequencia, n)
	}
	return sequencia
}

// CollatzCount retorna a quantidade de operações
func CollatzCount(n int) (int, error) {
	count := 1 

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count++
		
		if n == 1 {
			break
		}
	}
	return count, nil 
    
}