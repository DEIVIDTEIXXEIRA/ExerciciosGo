package main

import "fmt"

func main() {
	n := 10

	quadradoDaSoma := QuadradoDaSoma(n)
	somaDosQuadrados := SomaDosQuadrados(n)
	diferenca := Diferenca(n)

	fmt.Printf("Quadrado da soma: %d\n", quadradoDaSoma)
    fmt.Printf("Soma dos quadrados: %d\n", somaDosQuadrados)
    fmt.Printf("Diferença: %d\n", diferenca)
}

func QuadradoDaSoma(n int) int {
	soma := n * (n + 1) / 2 // EX: 10+1 = 11 | 11*10 = 110 | 110/2 = 55
	return soma * soma	// 55 * 55 = 3025
}

func SomaDosQuadrados(n int) int {
	soma := 0
    for i := 1; i <= n; i++ {
        soma += i * i
    }
    return soma
}

func Diferenca(n int) int {
	return	QuadradoDaSoma(n) - SomaDosQuadrados(n)
}