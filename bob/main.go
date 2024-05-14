package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	result := Hey(("1, 2, 3"))
	fmt.Println(result)
}

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	// Verifica se a string está vazia
	if remark == "" {
		return "Fine. Be that way!"
	}

	// Verifica se a string é um grito
	isYelling := isYelling(remark)

	// Verifica se a string é uma pergunta
	isQuestion := strings.HasSuffix(remark, "?")

	// Respostas baseadas nas condições
	if isYelling && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if isYelling {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	}

	// Se não for uma pergunta, um grito ou uma string vazia, retorna "Whatever."
	return "Whatever."
}

// Função auxiliar para verificar se a string está gritando
func isYelling(remark string) bool {
	hasLetters := false
	for _, char := range remark {
		if unicode.IsLetter(char) {
			hasLetters = true
			if !unicode.IsUpper(char) {
				return false
			}
		}
	}
	return hasLetters
}
