package main

import "fmt"

func sumaPrimerosN(n int) int {
    if n <= 1 {
        return n
    }
    return n + sumaPrimerosN(n-1)
}

func main() {
    N := 10
    fmt.Printf("La suma de los primeros %d enteros es %d\n", N, sumaPrimerosN(N))
}