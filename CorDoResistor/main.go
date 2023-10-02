package main

import (
	"fmt"
	"strings"
)

var colorMap = map[string]int{
	"black":    0,
	"brown":   1,
	"red": 2,
	"orange":  3,
	"yellow":  4,
	"green":    5,
	"blue":     6,
	"violet":  7,
	"gray":    8,
	"white":   9,
}

func main() {
	resultado := ColorCode("laranja")
	fmt.Println(resultado)
}

// Colors retorna a lista de todas as cores.
func Colors() []string {
	var colors []string
	expectedOrder := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "gray", "white"}

	for _, color := range expectedOrder {
		colors = append(colors, color)
	}
	return colors
}

func ColorCode(color string) int {
	color = strings.ToLower(color)
	return colorMap[color]
}
