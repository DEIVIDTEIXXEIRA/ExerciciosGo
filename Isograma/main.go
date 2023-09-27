package main

import (
    "fmt"
    "strings"
)

func naoContemLetrasRepetidas(s string) bool {
    s = strings.ToLower(s)
    letrasVistas := make(map[rune]bool)

    for _, char := range s {
        if char == ' ' {
            continue
        }

        if letrasVistas[char] {
            return false
        }

        letrasVistas[char] = true
    }

	if s == " " {
        return true
    }

    return true
}

func palavrasSaoIsogramas(frase string) bool {
	if frase == "" {
        return true
    }
	palavras := strings.Fields(frase)

    for _, palavra := range palavras {
        if naoContemLetrasRepetidas(palavra) {
            return true
        }
    }

    return false
}

func main() {
    frase1 := ""
    frase2 := "seis anos"

    resultado1 := palavrasSaoIsogramas(frase1)
    resultado2 := palavrasSaoIsogramas(frase2)

    fmt.Printf("'%s' contém isogramas? %t\n", frase1, resultado1)
    fmt.Printf("'%s' contém isogramas? %t\n", frase2, resultado2)
}