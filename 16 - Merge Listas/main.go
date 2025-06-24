package main

import "fmt"

func merge(L1, L2 []int) []int {
    if len(L1) == 0 {
        return L2
    }
    if len(L2) == 0 {
        return L1
    }
    if L1[0] < L2[0] {
        return append([]int{L1[0]}, merge(L1[1:], L2)...)
    }
    return append([]int{L2[0]}, merge(L1, L2[1:])...)
}

func main() {
    L1 := []int{1, 3, 5, 7}
    L2 := []int{2, 4, 6, 8}
    L3 := merge(L1, L2)
    fmt.Println("Lista fusionada:", L3)
}