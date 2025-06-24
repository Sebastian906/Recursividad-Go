package main

import "fmt"

func existeElemento(L []int, x int) string {
    if len(L) == 0 {
        return "No existe"
    }
    if L[0] == x {
        return "Existe"
    }
    return existeElemento(L[1:], x)
}

func main() {
    lista := []int{5, 3, 8, 1, 9, 2}
    elemento := 8
    fmt.Printf("El elemento %d %s en la lista\n", elemento, existeElemento(lista, elemento))
}