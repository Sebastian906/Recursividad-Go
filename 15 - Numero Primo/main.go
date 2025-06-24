package main

import "fmt"

func esPrimo(n, divisor int) bool {
	if n <= 2 {
		return n == 2
	}
	if n%divisor == 0 {
		return false
	}
	if divisor*divisor > n {
		return true
	}
	return esPrimo(n, divisor+1)
}

func main() {
	numero := 17
	if esPrimo(numero, 2) {
		fmt.Printf("%d es primo\n", numero)
	} else {
		fmt.Printf("%d no es primo\n", numero)
	}
}