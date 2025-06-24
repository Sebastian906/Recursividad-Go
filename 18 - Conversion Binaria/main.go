package main

import "fmt"

func aBinario(n int) string {
    if n == 0 {
        return ""
    }
    return aBinario(n/2) + fmt.Sprintf("%d", n%2)
}

func main() {
    numero := 28
    binario := aBinario(numero)
    if binario == "" {
        binario = "0"
    }
    fmt.Printf("%d en binario es %s\n", numero, binario)
}