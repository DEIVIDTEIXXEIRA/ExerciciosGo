package main

import "fmt"

func main() {
	ano := 2000
	resultado := EAnoBissexto(ano)
	fmt.Printf("O ano %d é bisexto?\n R: %t", ano, resultado)
}

func EAnoBissexto(year int) bool {
	switch {
	case year%4 == 0 && year%100 == 0 && year%400 == 0:
		return true
	case year%4 == 0 && year%100 != 0:
		return true
	case year%4 == 0 && year%100 == 0:
		return false
	}
	return false
}
