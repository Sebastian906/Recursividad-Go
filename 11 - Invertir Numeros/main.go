package main

import (
	"fmt"
	"strconv"
)

func invertirNumero(n int, invertido string) int {
	if n == 0 {
		result, _ := strconv.Atoi(invertido)
		return result
	}
	ultimoDigito := n % 10
	return invertirNumero(n/10, invertido + fmt.Sprintf("%d", ultimoDigito))
}

func main() {
	numero := 12345
	fmt.Println("Original: ", numero)
	invertido := invertirNumero(numero, "")
	fmt.Println("Invertido: ", invertido)
}