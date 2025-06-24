package main

import "fmt"

func contarDigitosPares(n int) int {
    if n == 0 {
        return 0
    }
    digito := n % 10
    if digito%2 == 0 {
        return 1 + contarDigitosPares(n/10)
    }
    return contarDigitosPares(n/10)
}

func main() {
    numero := 23025
    fmt.Printf("El número %d tiene %d dígitos pares\n", numero, contarDigitosPares(numero))
}