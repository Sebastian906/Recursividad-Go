package main

import "fmt"

func sonIguales(L1, L2 []int) bool {
    if len(L1) != len(L2) {
        return false
    }
    if len(L1) == 0 {
        return true
    }
    if L1[0] != L2[0] {
        return false
    }
    return sonIguales(L1[1:], L2[1:])
}

func main() {
    L1 := []int{1, 2, 3, 4}
    L2 := []int{1, 2, 3, 4}
    if sonIguales(L1, L2) {
        fmt.Println("Las listas son iguales")
    } else {
        fmt.Println("Las listas son diferentes")
    }
}